package session

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Session struct {
	UserId     int
	SessionKey string
	CreatedAt  int64
}

func New(user_id int, password string) *Session {
	created_at := time.Now().Unix()
	hash := sha256.Sum256([]byte(password + fmt.Sprint(created_at)))
	return &Session{
		UserId:     user_id,
		SessionKey: string(hash[:]),
		CreatedAt:  created_at,
	}
}
