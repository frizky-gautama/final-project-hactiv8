package model

type Comment struct {
	ID        int32  `gorm:"unique;primaryKey" json:"comment_id"`
	UserID    int32  `gorm:"column:user_id;not null" json:"user_id"`
	PhotoID   int32  `gorm:"column:photo_id" json:"photo_id"`
	Message   string `gorm:"column:message;not null;type:varchar(200)" json:"message"`
	CreatedAt string `gorm:"column:created_at;type:varchar(25)" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(25)" json:"updated_at"`
	User      *User  `gorm:"references:id;foreignkey:user_id" json:"User"`   // use Refer as association foreign key
	Photo     *Photo `gorm:"references:id;foreignkey:photo_id" json:"Photo"` // use Refer as association foreign key
}
