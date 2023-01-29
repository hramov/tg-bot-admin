package user

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateDto struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email" validate:"required, email"`
	GeoLabel string `json:"geo_label"`
	ChatId   string `json:"chat_id"`
	Password string `json:"password" validate:"required"`
}

type UpdateDto struct {
	Id int `json:"id"`
	CreateDto
}

type LoginResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
