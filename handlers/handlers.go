package handlers

import (
	"log"
	"net/http"
	"workula/message"
	"workula/objects"
	"workula/user"
	"workula/util"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

func CheckSession(c echo.Context) (*user.Session, error) {
	session := user.DecodeSession(c)
	if !user.VerifyKey(session.UserId, session.SessionKey) {
		return nil, c.JSON(http.StatusForbidden, &message.Message{
			Text: "You are not authorized!",
		})
	}
	return session, nil
}
func CheckSessionObject(session *objects.Session) (*objects.Session, error) {
	if !user.VerifyKey(session.UserId, session.SessionKey) {
		return nil, echo.ErrBadRequest
	}
	return session, nil
}
func SearchUser(c echo.Context) error {
	uqo := user.DecodeUnifiedQueryObject(c)
	_, err := CheckSessionObject(uqo.Session)
	util.CheckErrors("SearchUser", err)
	name, ok := uqo.Params["Name"].(string)
	if !ok {
		log.Printf("SearchUser: Can't convert name parameter to string!")
		return echo.ErrBadRequest
	}
	users := user.FindUserByName(name)
	return c.JSON(http.StatusOK, users)
}

func SignIn(c echo.Context) error {
	_user := user.DecodeUser(c)
	return user.SignInHandler(c, _user)
}
func SignUp(c echo.Context) error {
	_user := user.DecodeUser(c)
	log.Println("User decoded")
	uuser := user.NewUser(_user.Name, _user.Email, _user.Password)
	log.Println("User created")
	user.AppendUserToDB(uuser)
	log.Println("User appended to DB")
	return user.SignInHandler(c, _user)
}
func Connect(c echo.Context) error {
	session, err := CheckSession(c)
	util.CheckErrors("Connect CheckSession", err)
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	message.OpenWS(conn, session.UserId)
	util.CheckErrors("Connect OpenWebSocket", err)
	defer message.CloseWS(session.UserId)
	defer conn.Close()
	conn.WriteJSON(message.GetLastPages(2))
cycle:
	for {
		decoded_msg := &message.Message{}
		err = conn.ReadJSON(decoded_msg)
		util.CheckErrors("ConnectLoop", err)
		switch decoded_msg.Text {
		case "close":
			break cycle
		case "load":
			conn.WriteJSON(message.GetLastPages(1))
		default:
			message.AppendMessage(decoded_msg)
		}
	}
	return nil
}
