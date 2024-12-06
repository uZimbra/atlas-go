package dtos

type UserType string

const (
	Parent   UserType = "Parent"
	Children UserType = "Children"
)

type CreateUserInputDto struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	UserType UserType `json:"user_type"`
}
