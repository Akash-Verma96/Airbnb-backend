package dto

type CreateUser struct {
	Name string
	Email string
	Password string
}

type LoginUser struct {
	Email string
	Password string
}

type DeleteRequest struct {
	Id int64
}