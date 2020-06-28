package main

type UserDetail struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []UserDetail
}
