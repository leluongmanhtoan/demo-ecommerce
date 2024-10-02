package model

type LoginForm struct {
	Username string `json:"username" binding:"required,min=2,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type (
	LoginResponse struct {
		*AuthUserData
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		ExpiredIn    int    `json:"expired_in"`
		Fullname     string `json:"fullname"`
	}

	AuthUserData struct {
		UserId   string   `json:"user_id"`
		UserName string   `json:"username"`
		Level    string   `json:"level"`
		Scopes   []string `json:"scopes"`
	}
)
