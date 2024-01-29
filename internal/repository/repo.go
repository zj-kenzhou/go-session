package repository

import "github.com/zj-kenzhou/go-session/internal/domain"

type SessionRepo interface {
	GetLoginId(token string) (string, error)
	SetToken(loginId any, token string) error
	RemoveToken(token string) error
	GetToken(loginId any) ([]string, error)
	GetSession(loginId any, token string) (domain.SessionEntity, error)
	SetSession(entity domain.SessionEntity) error
	RemoveSession(loginId any, token string) error
	RenewTimeout(loginId any, token string) error
}
