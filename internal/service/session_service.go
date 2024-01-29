package service

import (
	"github.com/zj-kenzhou/go-session/model"
)

type SessionService interface {
	Login(model model.LoginModel) (string, error)
	Logout(token string) error
	LogoutByLoginId(loginId any) error
	CheckLogin(tokenValue string) (string, error)
	GetSessionValue(tokenValue string, key string) (any, error)
	SetSessionValue(tokenValue string, key string, value any) error
	RenewTimeout(loginId any, token string) error
}
