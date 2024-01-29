package session

import (
	"github.com/zj-kenzhou/go-session/internal/service"
	"github.com/zj-kenzhou/go-session/model"
)

func Login(model model.LoginModel) (string, error) {
	token, err := service.GetSessionService().Login(model)
	if err != nil {
		return "", err
	}
	doLogin(model.LoginId, token, model)
	return token, err
}
func Logout(token string) error {
	err := service.GetSessionService().Logout(token)
	if err == nil {
		doLogout(token)
	}
	return err
}
func LogoutByLoginId(loginId any) error {
	err := service.GetSessionService().LogoutByLoginId(loginId)
	if err == nil {
		doLogoutByLoginId(loginId)
	}
	return err
}
func CheckLogin(tokenValue string) (string, error) {
	return service.GetSessionService().CheckLogin(tokenValue)
}
func GetSessionValue(tokenValue string, key string) (any, error) {
	return service.GetSessionService().GetSessionValue(tokenValue, key)
}
func SetSessionValue(tokenValue string, key string, value any) error {
	return service.GetSessionService().SetSessionValue(tokenValue, key, value)
}
func RenewTimeout(loginId any, token string) error {
	return service.GetSessionService().RenewTimeout(loginId, token)
}
