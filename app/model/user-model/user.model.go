package usermodel

type UserModel struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
}
