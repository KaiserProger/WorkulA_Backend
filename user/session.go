package user

import (
	"log"
	"workula/objects"
	"workula/util"

	"github.com/labstack/echo/v4"
)

type Session objects.Session

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
	db.Model(&Session{}).Where("user_id = ?", user_id).FirstOrCreate(session)
	return session
}
func DecodeSession(c echo.Context) *Session {
	session := &Session{}
	err := c.Bind(session)
	log.Printf("%v", session)
	util.CheckErrors("DecodeUser", err)
	return session
}
func DecodeUnifiedQueryObject(c echo.Context) *objects.UnifiedQueryObject {
	uqo := &objects.UnifiedQueryObject{}
	err := c.Bind(uqo)
	util.CheckErrors("DecodeUQO", err)
	return uqo
}
