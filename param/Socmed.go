package param

var InsertSocmed struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

var UpdateSocmed struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type ResponseInsertSocmed struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int32  `json:"user_id"`
	CreatedAt      string `son:"created_at"`
}

type ResponseSocmed struct {
	ID             int32              `json:"id"`
	Name           string             `json:"name"`
	SocialMediaUrl string             `json:"social_media_url"`
	UserID         int32              `json:"user_id"`
	CreatedAt      string             `son:"created_at"`
	UpdatedAt      string             `json:"updated_at"`
	User           ResponseUserSocmed `json:"User"`
}

type ResponseUserSocmed struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseUpdateSocmed struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int32  `json:"user_id"`
	UpdatedAt      string `json:"updated_at"`
}
