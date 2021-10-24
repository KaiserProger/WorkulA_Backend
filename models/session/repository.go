package session

import (
	"errors"
)

type SessionRepository struct {
	Cache []*Session
}

var Repository = &SessionRepository{}
var ErrInvalidKey = errors.New("invalid key")
var ErrNoSessionFound = errors.New("no session found")

func (s *SessionRepository) Init() {
	Repository.Cache = make([]*Session, 0)
}
func (s *SessionRepository) Insert(session *Session) {
	s.Cache = append(s.Cache, session)
}
func (s *SessionRepository) Erase(session *Session) error {
	for i, v := range s.Cache {
		if v.UserId == session.UserId {
			x := len(s.Cache) - 1
			s.Cache[i] = s.Cache[x]
			s.Cache = s.Cache[:x]
			return nil
		}
	}
	return ErrNoSessionFound
}
func (s *SessionRepository) Check(session *Session) error {
	for _, v := range s.Cache {
		if v.UserId == session.UserId {
			if v.SessionKey == session.SessionKey {
				return nil
			} else {
				return ErrInvalidKey
			}
		}
	}
	return ErrNoSessionFound
}
func (s *SessionRepository) GetByUserID(user_id int) *Session {
	for _, v := range s.Cache {
		if v.UserId == user_id {
			return v
		}
	}
	return nil
}
