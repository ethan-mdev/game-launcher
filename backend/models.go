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
