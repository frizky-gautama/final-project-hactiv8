package param

var UpdateUserPayload struct {
	Email    string `binding:"required,email" json:"email"`
	Username string `binding:"required" json:"username"`
}
