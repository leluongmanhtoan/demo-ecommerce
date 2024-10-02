package service

import "context"

type (
	IAuth interface {
	}
	Auth struct {
		UIBaseUrl string
	}
)

var AuthService IAuth

func NewAuthService(uiBaseUrl string) IAuth {
	return &Auth{
		UIBaseUrl: uiBaseUrl,
	}
}

func (auth *Auth) Login(ctx context.Context)
