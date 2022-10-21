package param

var LoginReq struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
}

var RegisReq struct {
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
	Username string `binding:"required" json:"username"`
	Age      int    `binding:"required,gt=8" json:"age"`
}
