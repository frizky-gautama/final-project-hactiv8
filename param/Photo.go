package param

var InsertPhoto struct {
	Title     string `binding:"required" json:"title"`
	Caption   string `binding:"required" json:"caption"`
	Photo_url string `binding:"required" json:"photo_url"`
}

type ResponsePhoto struct {
	ID        int32             `json:"id"`
	Title     string            `json:"title"`
	Caption   string            `json:"caption"`
	Photo_url string            `json:"photo_url"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	UserID    int32             `json:"user_id"`
	User      ResponseUserPhoto `json:"User"`
}

type ResponseUserPhoto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

var UpdatePhoto struct {
	Title     string `binding:"required" json:"title"`
	Caption   string `binding:"required" json:"caption"`
	Photo_url string `binding:"required" json:"photo_url"`
}

type ResponseUpdatePhoto struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	UserID    int32  `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
