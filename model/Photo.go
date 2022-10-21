package model

type Photo struct {
	ID        int32  `gorm:"unique;primaryKey" json:"photo_id"`
	Title     string `gorm:"column:title;not null;type:varchar(255)" json:"title"`
	Caption   string `gorm:"column:caption;type:varchar(200)" json:"caption"`
	Photo_url string `gorm:"column:photo_url;not null;type:varchar(200)" json:"photo_url"`
	CreatedAt string `gorm:"column:created_at;type:varchar(25)" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(25)" json:"updated_at"`
	UserID    int32  `gorm:"column:user_id" json:"user_id"`
	User      *User  `gorm:"references:id;foreignkey:user_id" json:"User"` // use Refer as association foreign key
}

type ResponsePostPhoto struct {
	ID        int32  `gorm:"column:id" json:"id"`
	Title     string `gorm:"column:title" json:"title"`
	Caption   string `gorm:"column:caption" json:"caption"`
	Photo_url string `gorm:"column:photo_url" json:"photo_url"`
	UserID    int    `gorm:"column:user_id" json:"user_id"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
}

func (ResponsePostPhoto) TableName() string {
	return "photos"
}
