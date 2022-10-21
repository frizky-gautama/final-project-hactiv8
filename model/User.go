package model

type User struct {
	ID        int32  `gorm:"unique;primaryKey" json:"user_id"`
	Username  string `gorm:"unique;column:username;not null;type:varchar(200)" json:"username"`
	Email     string `gorm:"unique;column:email;not null;type:varchar(200)" json:"email"`
	Password  string `gorm:"column:password;not null;type:varchar(100)" json:"password"`
	Age       int    `gorm:"column:age;not null;type:int(10)" json:"age"`
	CreatedAt string `gorm:"column:created_at;type:varchar(25)" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(25)" json:"updated_at"`
}

type RegResponse struct {
	ID       int32  `gorm:"column:id" json:"user_id"`
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Age      int    `gorm:"column:age" json:"age"`
}

type UpdateResponse struct {
	ID        int32  `gorm:"column:id" json:"user_id"`
	Username  string `gorm:"column:username" json:"username"`
	Email     string `gorm:"column:email" json:"email"`
	Age       int    `gorm:"column:age" json:"age"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func (RegResponse) TableName() string {
	return "users"
}

func (UpdateResponse) TableName() string {
	return "users"
}
