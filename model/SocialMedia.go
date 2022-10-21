package model

type SocialMedia struct {
	ID             int32  `gorm:"unique;primaryKey" json:"socialmedia_id"`
	Name           string `gorm:"unique;column:name;not null;type:varchar(200)" json:"name"`
	SocialMediaUrl string `gorm:"unique;column:social_media_url;not null;type:varchar(200)" json:"social_media_url"`
	UserID         int32  `gorm:"column:user_id" json:"user_id"`
	CreatedAt      string `gorm:"column:created_at;type:varchar(25)" json:"created_at"`
	UpdatedAt      string `gorm:"column:updated_at;type:varchar(25)" json:"updated_at"`
	User           *User  `gorm:"references:id;foreignkey:user_id" json:"User"` // use Refer as association foreign key
}
