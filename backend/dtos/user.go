package dtos

type UserDTO struct {
	Name  string   `json:"name"`
	Login LoginDTO `json:"login"`
}
