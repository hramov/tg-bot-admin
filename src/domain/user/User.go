package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	GeoLabel string `json:"geo_label"`
	ChatId   string `json:"chat_id"`
	Password string `json:"password"`
}
