package entity

type User struct {
	BaseEntity
	UserName string `gorm:"not null;type:varchar(255)"`
	Password string `gorm:"not null;type:nvarchar(500)"`
}
