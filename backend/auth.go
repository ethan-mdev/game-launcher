package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(username, password string) (*AuthResponse, error) {
	req := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(AuthBaseURL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid username or password")
	}
	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid request sent")
	}
	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("internal server error")
	}

	var respData AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	return &respData, nil
}

func (a *AuthService) Logout(refreshToken string) error {
	body, _ := json.Marshal(map[string]string{
		"refresh_token": refreshToken,
	})

	resp, err := http.Post(AuthBaseURL+"/logout", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to logout")
	}

	return nil
}

func (a *AuthService) Refresh(refreshToken string) (*AuthResponse, error) {
	body, _ := json.Marshal(map[string]string{
		"refresh_token": refreshToken,
	})

	resp, err := http.Post(AuthBaseURL+"/refresh", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var auth AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

func (a *AuthService) Register(username, email, password string) (*AuthResponse, error) {
	req := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Username: username,
		Email:    email,
		Password: password,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(AuthBaseURL+"/register", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid request sent")
	}
	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("internal server error")
	}

	var respData AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	return &respData, nil
}

func (a *AuthService) GetProfile(userID string) (*UserProfile, error) {
	resp, err := http.Get(fmt.Sprintf("%s/profile/%s", AuthBaseURL, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get profile")
	}

	var profile UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func (a *AuthService) UpdateProfileImage(accessToken, profileImage string) error {
	body, _ := json.Marshal(map[string]string{
		"profile_image": profileImage,
	})

	req, err := http.NewRequest("PUT", AuthBaseURL+"/profile", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update profile image")
	}

	return nil
}

func (a *AuthService) GetGameCredentials(accessToken string) (*GameCredentials, error) {
	req, err := http.NewRequest("GET", AuthBaseURL+"/game/credentials", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the "not linked" case - return nil without error so frontend can handle it
	if resp.StatusCode == http.StatusForbidden {
		var errResp GameCredentialsError
		json.NewDecoder(resp.Body).Decode(&errResp)
		return nil, fmt.Errorf(errResp.Error)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch game credentials")
	}

	var creds GameCredentials
	if err := json.NewDecoder(resp.Body).Decode(&creds); err != nil {
		return nil, err
	}

	return &creds, nil
}

func (a *AuthService) VerifyGameAccount(accessToken string) error {
	req, err := http.NewRequest("POST", AuthBaseURL+"/game/verify", nil)
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
		var errResp map[string]string
		json.NewDecoder(resp.Body).Decode(&errResp)
		if msg, ok := errResp["message"]; ok {
			return fmt.Errorf(msg)
		}
		return fmt.Errorf("failed to verify account")
	}

	return nil
}
