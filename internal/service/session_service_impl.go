package service

import (
	"github.com/matoous/go-nanoid/v2"
	"github.com/zj-kenzhou/go-session/internal/domain"
	"github.com/zj-kenzhou/go-session/internal/infra"
	"github.com/zj-kenzhou/go-session/internal/repository"
	"github.com/zj-kenzhou/go-session/model"
	"strings"
	"sync"
	"time"
)

var _sessionService SessionService
var _once sync.Once

type sessionService struct {
	repo repository.SessionRepo
}

func (s *sessionService) Login(model model.LoginModel) (string, error) {
	token, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	err = s.repo.SetToken(model.LoginId, token)
	if err != nil {
		return "", err
	}
	return token, s.repo.SetSession(domain.SessionEntity{
		LoginId:   model.LoginId,
		Ip:        model.Ip,
		Token:     token,
		LoginTime: time.Now(),
		Device:    model.Device,
		UserAgent: model.UserAgent,
		Data:      map[string]any{},
	})
}

func (s *sessionService) Logout(token string) error {
	loginId, err := s.repo.GetLoginId(token)
	if err != nil {
		return err
	}
	if loginId != "" {
		err = s.repo.RemoveToken(token)
		if err != nil {
			return err
		}
		return s.repo.RemoveSession(loginId, token)
	}
	return nil
}

func (s *sessionService) LogoutByLoginId(loginId any) error {
	tokens, err := s.repo.GetToken(loginId)
	if err != nil {
		return err
	}
	for _, key := range tokens {
		keyList := strings.Split(key, ":")
		token := keyList[len(keyList)-1]
		err = s.repo.RemoveToken(token)
		if err != nil {
			return err
		}
		err = s.repo.RemoveSession(loginId, token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *sessionService) CheckLogin(tokenValue string) (string, error) {
	return s.repo.GetLoginId(tokenValue)
}

func (s *sessionService) GetSessionValue(tokenValue string, key string) (any, error) {
	loginId, err := s.repo.GetLoginId(tokenValue)
	if err != nil {
		return nil, err
	}
	session, err := s.repo.GetSession(loginId, tokenValue)
	if err != nil {
		return nil, err
	}
	return session.Data[key], nil
}

func (s *sessionService) SetSessionValue(tokenValue string, key string, value any) error {
	loginId, err := s.repo.GetLoginId(tokenValue)
	if err != nil {
		return err
	}
	session, err := s.repo.GetSession(loginId, tokenValue)
	if err != nil {
		return err
	}
	session.Data[key] = value
	return s.repo.SetSession(session)
}

func (s *sessionService) RenewTimeout(loginId any, token string) error {
	return s.repo.RenewTimeout(loginId, token)
}

func GetSessionService() SessionService {
	_once.Do(func() {
		_sessionService = &sessionService{
			repo: infra.NewRedisRepo(),
		}
	})
	return _sessionService
}
