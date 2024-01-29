package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zj-kenzhou/go-session/config"
	"github.com/zj-kenzhou/go-session/internal/domain"
	"github.com/zj-kenzhou/go-session/internal/repository"
	"time"
)

const dataType = "login"

const sessionKey = "session"

func getKeyPrefix() string {
	if config.GetConfig().KeyPrefix != "" {
		return config.GetConfig().KeyPrefix + ":" + dataType + ":"
	}
	return dataType + ":"
}

func getTimeOut() time.Duration {
	return time.Duration(config.GetConfig().ActivityTimeout) * time.Second
}

type redisRepository struct {
	client redis.UniversalClient
}

func (r *redisRepository) GetLoginId(token string) (string, error) {
	value, err := r.client.Get(context.Background(), getKeyPrefix()+token).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}

func (r *redisRepository) SetToken(loginId any, token string) error {
	return r.client.Set(context.Background(), getKeyPrefix()+token, loginId, getTimeOut()).Err()
}

func (r *redisRepository) RemoveToken(token string) error {
	return r.client.Del(context.Background(), getKeyPrefix()+token).Err()
}

func (r *redisRepository) GetToken(loginId any) ([]string, error) {
	return r.client.Keys(context.Background(), fmt.Sprintf("%s%s:%v:*", getKeyPrefix(), sessionKey, loginId)).Result()
}

func (r *redisRepository) RemoveSession(loginId any, token string) error {
	return r.client.Del(context.Background(), fmt.Sprintf("%s%s:%v:%s", getKeyPrefix(), sessionKey, loginId, token)).Err()
}

func (r *redisRepository) GetSession(loginId any, token string) (domain.SessionEntity, error) {
	resString, err := r.client.Get(context.Background(), fmt.Sprintf("%s%s:%v:%s", getKeyPrefix(), sessionKey, loginId, token)).Result()
	res := domain.SessionEntity{}
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(resString), &res)
	return res, err
}

func (r *redisRepository) SetSession(entity domain.SessionEntity) error {
	key := fmt.Sprintf("%s%s:%v:%s", getKeyPrefix(), sessionKey, entity.LoginId, entity.Token)
	bytes, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, string(bytes), getTimeOut()).Err()
}

func (r *redisRepository) RenewTimeout(loginId any, token string) error {
	// 先判断是否超过最长登录时间
	session, err := r.GetSession(loginId, token)
	if err != nil {
		return err
	}
	if time.Now().Sub(session.LoginTime) > time.Duration(config.GetConfig().Timeout)*time.Second {
		return nil
	}
	key := fmt.Sprintf("%s%s:%v:%s", getKeyPrefix(), sessionKey, loginId, token)
	err = r.client.Expire(context.Background(), getKeyPrefix()+token, getTimeOut()).Err()
	if err != nil {
		return err
	}
	return r.client.Expire(context.Background(), key, getTimeOut()).Err()
}

func NewRedisRepo() repository.SessionRepo {
	sessionConfig := config.GetConfig()
	return &redisRepository{
		client: redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:            sessionConfig.Host,
			ClientName:       sessionConfig.ClientName,
			DB:               sessionConfig.Db,
			Username:         sessionConfig.Username,
			Password:         sessionConfig.Password,
			SentinelUsername: sessionConfig.SentinelUsername,
			SentinelPassword: sessionConfig.SentinelPassword,
			PoolSize:         sessionConfig.PoolSize,
			PoolTimeout:      time.Duration(sessionConfig.PoolTimeout) * time.Second,
			MinIdleConns:     sessionConfig.MinIdleConns,
			MaxIdleConns:     sessionConfig.MaxIdleConns,
			ConnMaxIdleTime:  time.Duration(sessionConfig.ConnMaxIdleTime) * time.Minute,
			ConnMaxLifetime:  time.Duration(sessionConfig.ConnMaxLifetime) * time.Minute,
			MasterName:       sessionConfig.MasterName,
		}),
	}
}
