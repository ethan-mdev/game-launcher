package backend

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewPatchingService(app *App, gameDir, patchingURL string) *PatchingService {
	return &PatchingService{
		app:         app,
		patchingURL: patchingURL,
		gameDir:     gameDir,
	}
}

func (p *PatchingService) DownloadUpdates(files []FileHash, accessToken string) error {
	serverManifest, err := p.fetchManifest(accessToken)
	if err != nil {
		return err
	}

	totalFiles := len(files)
	for i, file := range files {
		runtime.EventsEmit(p.app.ctx, "patch:progress", map[string]interface{}{
			"current": i + 1,
			"total":   totalFiles,
			"file":    file.FileName,
			"status":  "downloading",
		})

		if err := p.downloadFile(file, accessToken); err != nil {
			return err
		}

		if err := p.verifyFile(file); err != nil {
			return err
		}

		progress := float64(i+1) / float64(totalFiles) * 100
		runtime.EventsEmit(p.app.ctx, "patch:file-complete", map[string]interface{}{
			"file":     file.FileName,
			"progress": progress,
		})
	}

	if err := p.updateLocalManifest(serverManifest.Version, files); err != nil {
		return err
	}

	runtime.EventsEmit(p.app.ctx, "patch:complete", nil)
	return nil
}

func (p *PatchingService) CheckForUpdates(accessToken string) (*UpdateCheckResult, error) {
	localVersion := p.getLocalVersion()

	req, err := http.NewRequest("GET", p.patchingURL+"/manifest", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	if localVersion != "unknown" {
		req.Header.Set("If-None-Match", localVersion)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return &UpdateCheckResult{
			NeedsUpdate:    false,
			CurrentVersion: localVersion,
			ServerVersion:  localVersion,
			FilesToUpdate:  []FileHash{},
		}, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to check for updates: status code %d", resp.StatusCode)
	}

	var manifest Manifest
	err = json.NewDecoder(resp.Body).Decode(&manifest)
	if err != nil {
		return nil, err
	}

	serverVersion := manifest.Version

	localFiles, err := p.hashLocalFiles()
	if err != nil {
		return nil, err
	}

	toDownload := make([]FileHash, 0)
	for _, serverFile := range manifest.Files {
		localHash, exists := localFiles[serverFile.FileName]
		if !exists || localHash != serverFile.Hash {
			toDownload = append(toDownload, serverFile)
		}
	}

	// Save manifest version even if no files need updating
	if len(toDownload) == 0 && localVersion != serverVersion {
		p.updateLocalManifest(serverVersion, []FileHash{})
	}

	return &UpdateCheckResult{
		NeedsUpdate:    len(toDownload) > 0,
		CurrentVersion: localVersion,
		ServerVersion:  serverVersion,
		FilesToUpdate:  toDownload,
	}, nil

}

func (p *PatchingService) RepairGameFiles(accessToken string) error {
	localFiles, err := p.hashLocalFiles()
	if err != nil {
		return err
	}

	verifyResult, err := p.verifyWithServer(localFiles, accessToken)
	if err != nil {
		return err
	}

	if verifyResult.Valid {
		runtime.EventsEmit(p.app.ctx, "patch:complete", nil)
		return nil
	}

	toRepair := make(map[string]bool)
	for _, file := range verifyResult.Mismatches {
		toRepair[file] = true
	}
	for _, file := range verifyResult.Missing {
		toRepair[file] = true
	}

	serverManifest, err := p.fetchManifest(accessToken)
	if err != nil {
		return fmt.Errorf("failed to fetch manifest: %w", err)
	}

	filesToDownload := []FileHash{}
	for _, serverFile := range serverManifest.Files {
		if toRepair[serverFile.FileName] {
			filesToDownload = append(filesToDownload, serverFile)
		}
	}

	return p.DownloadUpdates(filesToDownload, accessToken)
}

func (p *PatchingService) fetchManifest(accessToken string) (*Manifest, error) {
	req, err := http.NewRequest("GET", p.patchingURL+"/manifest", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch manifest: status code %d", resp.StatusCode)
	}

	var manifest Manifest
	err = json.NewDecoder(resp.Body).Decode(&manifest)
	if err != nil {
		return nil, err
	}

	return &manifest, nil
}

func (p *PatchingService) updateLocalManifest(version string, updatedFiles []FileHash) error {
	manifestPath := filepath.Join(p.gameDir, "manifest.json")
	var manifest Manifest

	data, err := os.ReadFile(manifestPath)
	if err == nil {
		json.Unmarshal(data, &manifest)
	}

	manifest.Version = version
	for _, newFile := range updatedFiles {
		found := false
		for i, localFile := range manifest.Files {
			if localFile.FileName == newFile.FileName {
				manifest.Files[i] = newFile
				found = true
				break
			}
		}
		if !found {
			manifest.Files = append(manifest.Files, newFile)
		}
	}
	data, err = json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(manifestPath, data, 0644)
}

func (p *PatchingService) verifyWithServer(localFiles map[string]string, accessToken string) (*VerifyResult, error) {
	jsonData, err := json.Marshal(localFiles)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", p.patchingURL+"/verify", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify files: status code %d", resp.StatusCode)
	}

	var result VerifyResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *PatchingService) downloadFile(file FileHash, accessToken string) error {
	url := fmt.Sprintf("%s/files/%s", p.patchingURL, file.FileName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s, status code: %d", file.FileName, resp.StatusCode)
	}

	localPath := filepath.Join(p.gameDir, file.Directory, file.FileName)
	dir := filepath.Dir(localPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write to temp file first
	tempPath := localPath + ".tmp"
	out, err := os.Create(tempPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	out.Close()

	if err != nil {
		os.Remove(tempPath) // Clean up on error
		return err
	}

	// Atomic rename
	return os.Rename(tempPath, localPath)
}

func (p *PatchingService) verifyFile(file FileHash) error {
	localPath := filepath.Join(p.gameDir, file.Directory, file.FileName)
	hash, err := computeFileHash(localPath)
	if err != nil {
		return err
	}

	if hash != file.Hash {
		return fmt.Errorf("hash mismatch for file: %s", file.FileName)
	}

	return nil
}

func (p *PatchingService) hashLocalFiles() (map[string]string, error) {
	fileHashes := make(map[string]string)
	err := filepath.Walk(p.gameDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || info.Name() == "manifest.json" {
			return nil
		}

		hash, err := computeFileHash(path)
		if err != nil {
			return err
		}

		// Use just the filename, not the full path
		fileHashes[info.Name()] = hash
		return nil
	})

	return fileHashes, err
}

func (p *PatchingService) getLocalVersion() string {
	manifestPath := filepath.Join(p.gameDir, "manifest.json")
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return "unknown"
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return "unknown"
	}
	return manifest.Version
}

func computeFileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:]), nil
}
