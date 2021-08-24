package chat

type Chat struct {
	Id   string `gorm:"primary_key" json:"chat_id"`
	Name string `gorm:"name"`
}

func NewChat() {

}
