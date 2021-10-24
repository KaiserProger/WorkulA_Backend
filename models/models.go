package models

import (
	"WorkulA/models/session"
	"WorkulA/models/user"
)

var Model = map[string]interface{}{
	"User":    &user.User{},
	"Session": &session.Session{},
}
