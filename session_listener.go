package session

import (
	"github.com/zj-kenzhou/go-session/model"
)

var listenerList []Listener

type Listener interface {
	// DoLogin 每次登录时触发
	DoLogin(loginId any, tokenValue string, loginModel model.LoginModel)

	// DoLogout 每次注销时触发
	DoLogout(tokenValue string)

	// DoLogoutByLoginId 每次注销时触发
	DoLogoutByLoginId(loginId any)

	//DoRenewTimeout 每次Token续期时触发
	DoRenewTimeout(tokenValue string, loginId any, timout int64)
}

func RegisterListener(listener Listener) {
	if listener != nil {
		listenerList = append(listenerList, listener)
	}
}

// DoLogin 每次登录时触发
func doLogin(loginId any, tokenValue string, loginModel model.LoginModel) {
	for _, listener := range listenerList {
		listener.DoLogin(loginId, tokenValue, loginModel)
	}
}

// DoLogout 每次注销时触发
func doLogout(tokenValue string) {
	for _, listener := range listenerList {
		listener.DoLogout(tokenValue)
	}
}

func doLogoutByLoginId(tokenValue any) {
	for _, listener := range listenerList {
		listener.DoLogoutByLoginId(tokenValue)
	}
}

// doRenewTimeout 每次续期
func doRenewTimeout(loginType string, loginId any, timeout int64) {
	for _, listener := range listenerList {
		listener.DoRenewTimeout(loginType, loginId, timeout)
	}
}
