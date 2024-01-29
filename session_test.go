package session

import (
	"github.com/zj-kenzhou/go-session/config"
	"github.com/zj-kenzhou/go-session/model"
	"testing"
)

func testConfig() {
	config.SetConfig(config.SessionConfig{
		Timeout:          604800,
		ActivityTimeout:  7200,
		KeyPrefix:        "test",
		Host:             []string{"127.0.0.1:6379"},
		Username:         "root",
		Password:         "",
		SentinelUsername: "",
		SentinelPassword: "",
		Db:               2,
		MasterName:       "",
		ClientName:       "",
		PoolSize:         100,
		PoolTimeout:      10,
		MinIdleConns:     1,
		MaxIdleConns:     10,
		ConnMaxIdleTime:  60,
		ConnMaxLifetime:  480,
	})
}

func TestLogin(t *testing.T) {
	testConfig()
	token, err := Login(model.LoginModel{
		LoginId:   "1234",
		Ip:        "127.0.0.1",
		Device:    "pc",
		UserAgent: "chrome",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestLogout(t *testing.T) {
	testConfig()
	err := Logout("5ePlQ06z_EQnXu71xI76e")
	if err != nil {
		t.Error(err)
	}
}

func TestLogoutByLoginId(t *testing.T) {
	testConfig()
	err := LogoutByLoginId("123")
	if err != nil {
		t.Error(err)
	}
}

func TestCheckLogin(t *testing.T) {
	testConfig()
	loginId, err := CheckLogin("5ePlQ06z_EQnXu71xI76e")
	if err != nil {
		t.Error(err)
	}
	t.Log(loginId)
}

func TestSetSessionValue(t *testing.T) {
	testConfig()
	err := SetSessionValue("5ePlQ06z_EQnXu71xI76e", "test", 2222)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSessionValue(t *testing.T) {
	testConfig()
	value, err := GetSessionValue("5ePlQ06z_EQnXu71xI76e", "test")
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}

func TestRenewTimeout(t *testing.T) {
	testConfig()
	err := RenewTimeout("1234", "5ePlQ06z_EQnXu71xI76e")
	if err != nil {
		t.Error(err)
	}

}
