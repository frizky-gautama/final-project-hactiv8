package param

var InsertComment struct {
	Message  string `binding:"required" json:"message"`
	Photo_id int32  `binding:"required" json:"photo_id"`
}

var UpdateComment struct {
	Message string `binding:"required" json:"message"`
}

type ResponseInsertComment struct {
	ID        int32  `json:"id"`
	Message   string `json:"message"`
	PhotoID   int32  `json:"photo_id"`
	UserID    int32  `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type ResponseComment struct {
	ID        int32                `json:"id"`
	Message   string               `json:"message"`
	PhotoID   int32                `json:"photo_id"`
	UserID    int32                `json:"user_id"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
	User      ResponseUserComment  `json:"User"`
	Photo     ResponsePhotoComment `json:"Photo"`
}

type ResponseUserComment struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponsePhotoComment struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	UserID    int32  `json:"user_id"`
}

type ResponseUpdateComment struct {
	ID        int32  `json:"id"`
	Message   string `json:"message"`
	PhotoID   int32  `json:"photo_id"`
	UserID    int32  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
