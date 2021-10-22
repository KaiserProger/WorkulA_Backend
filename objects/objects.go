package objects

type UnifiedQueryObject struct {
	Session *Session               `json:"session"`
	Params  map[string]interface{} `json:"params"`
}
type Session struct {
	UserId     int    `gorm:"primary_key" json:"user_id,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
}
type User struct {
	UserId   int    `gorm:"primary_key" json:"user_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
type Todo struct {
}
