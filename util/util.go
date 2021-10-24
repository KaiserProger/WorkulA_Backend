package util

import (
	"WorkulA/db"
	"WorkulA/models/user"
	"crypto/sha256"
)

func GenerateUserID() int64 {
	var c int64 = 0
	db.Connection.Model(&user.User{}).Count(&c)
	return c + 1
}
func GenerateHash(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}
