package dto

//USUARIOS

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Account  string `json:"account_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersDto []UserDto
