package user

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateDto struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	GeoLabel string `json:"geo_label"`
	ChatId   string `json:"chat_id"`
	Password string `json:"password"`
}

type UpdateDto struct {
	Id int `json:"id"`
	CreateDto
}
