package dtos

type CreateUseDTO struct {
	Name  string   `json:"name"`
	Login LoginDTO `json:"login"`
}
