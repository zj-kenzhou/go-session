package domain

import "time"

type SessionEntity struct {
	LoginId   string         `json:"loginId"`
	Ip        string         `json:"ip"`
	Token     string         `json:"token"`
	LoginTime time.Time      `json:"loginTime"`
	Device    string         `json:"device"`
	UserAgent string         `json:"userAgent"`
	Data      map[string]any `json:"data"`
}
