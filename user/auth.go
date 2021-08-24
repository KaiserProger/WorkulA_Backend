package user

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"workula/util"

	"github.com/labstack/echo/v4"
)

type InvalidPasswordError struct {
}

func (err InvalidPasswordError) Error() string {
	return "Invalid password!"
}

func SignIn(email string, password string, mstime int64) ([32]byte, bool) {
	if GetUserByEmail(email).Password != password {
		return [32]byte{}, false
	}
	key := sha256.Sum256([]byte(password + strconv.FormatInt(mstime, 10)))
	return key, true
}
func VerifyKey(user_id int, key string) bool {
	session := GetSessionByID(user_id)
	k := sha256.Sum256([]byte(GetUserByID(user_id).Password + strconv.FormatInt(session.CreatedAt, 10)))
	return string(k[:]) == key
}
func SignInHandler(c echo.Context, _user *User) error {
	_user = GetUserByEmail(_user.Email)
	log.Println("Sign in: Got user from DB")
	mstime := time.Now().Unix()
	key, is := SignIn(_user.Email, _user.Password, mstime)
	log.Println("Sign in: Key generated")
	if !is {
		return c.JSON(http.StatusForbidden, &Session{})
	}
	session := NewSession(_user.UserId, key, mstime)
	log.Println("Sign in: Session created")
	return c.JSON(http.StatusOK, session)
}
func DecodeUser(c echo.Context) *User {
	_user := &User{}
	err := json.NewDecoder(c.Request().Body).Decode(_user)
	log.Printf("%v", _user)
	util.CheckErrors("DecodeUser", err)
	return _user
}
