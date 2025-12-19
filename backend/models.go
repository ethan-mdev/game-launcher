package backend

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type UserProfile struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image"`
	CreatedAt    string `json:"created_at"`
}

type PatchingService struct {
	app         *App
	patchingURL string
	gameDir     string
}

type Manifest struct {
	Version string     `json:"version"`
	Files   []FileHash `json:"files"`
}

type FileHash struct {
	FileName  string `json:"fileName"`
	Directory string `json:"directory"`
	Hash      string `json:"hash"`
}

type UpdateCheckResult struct {
	NeedsUpdate    bool       `json:"needsUpdate"`
	CurrentVersion string     `json:"currentVersion"`
	ServerVersion  string     `json:"serverVersion"`
	FilesToUpdate  []FileHash `json:"filesToUpdate"`
}

type VerifyResult struct {
	Valid      bool     `json:"valid"`
	Mismatches []string `json:"mismatches"`
	Missing    []string `json:"missing"`
}
