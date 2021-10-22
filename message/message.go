package message

import (
	"encoding/json"
	"workula/util"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var db *gorm.DB

type Message struct {
	MessageId int    `gorm:"primary_key" json:"message_id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	ChatId    int    `json:"chat_id,omitempty"`
	Text      string `json:"text,omitempty"`
}

var ConnectionsStorage map[int]*websocket.Conn = make(map[int]*websocket.Conn)

func Broadcast(message *Message) {
	for _, i := range ConnectionsStorage {
		i.WriteJSON(message)
	}
}
func OpenWS(conn *websocket.Conn, user_id int) {
	ConnectionsStorage[user_id] = conn
}
func CloseWS(user_id int) {
	delete(ConnectionsStorage, user_id)
}
func Init(ddb *gorm.DB) {
	db = ddb
}
func NewMessage(message_id int, user_id int, text string) *Message {
	return &Message{
		MessageId: message_id,
		UserId:    user_id,
		Text:      text,
	}
}
func AppendMessage(message *Message) {
	message.MessageId = GetMessagesCount() + 1
	db.Model(&Message{}).Create(message)
	Broadcast(message)
}
func DecodeMessage(source []byte) *Message {
	message := &Message{}
	err := json.Unmarshal(source, message)
	util.CheckErrors("DecodeMessage", err)
	return message
}
func GetMessagesCount() int {
	message := &Message{}
	db.Model(&Message{}).Last(message)
	return message.MessageId
}
func GetLastPages(count int) []*Message {
	message_count := count * 10
	rows, err := db.Model(&Message{}).Order("message_id DESC").Limit(message_count).Rows()
	util.CheckErrors("GetLastPages", err)
	messages := make([]*Message, message_count)
	incrementer := 0
	for rows.Next() {
		messages[incrementer] = &Message{}
		err = rows.Scan(messages[incrementer])
		util.CheckErrors("GetLastPages CheckingRows", err)
		incrementer += 1
	}
	return messages
}
