package user

import (
	"log"
	"workula/util"

	"github.com/labstack/echo/v4"
)

type Session struct {
	UserId     int    `gorm:"primary_key" json:"user_id,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
}

func NewSession(user_id int, session_key [32]byte, mstime int64) *Session {
	session := &Session{
		UserId:     user_id,
		SessionKey: string(session_key[:]),
		CreatedAt:  mstime,
	}
	db.Model(&Session{}).Create(session)
	return session
}
func GetSessionByID(user_id int) *Session {
	session := &Session{}
	db.Model(&Session{}).Where("user_id = ?", user_id).Take(session)
	return session
}
func DecodeSession(c echo.Context) *Session {
	session := &Session{}
	err := c.Bind(session)
	log.Printf("%v", session)
	util.CheckErrors("DecodeUser", err)
	return session
}
