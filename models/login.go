package models

type Login struct {
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

type Register struct {
	Email           string `json:"email" xml:"email" form:"email"`
	Password        string `json:"password" xml:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" xml:"confirm_password" form:"confirm_password"`
}
