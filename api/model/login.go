package model

type UserLogin struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
