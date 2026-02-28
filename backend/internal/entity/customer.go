package entity

type Customer struct {
	BaseEntity
	Name    string `gorm:"not null;type:varchar(255)"`
	Surname string `gorm:"default:null;type:varchar(255)"`
	Email   string `gorm:"default:null;type:varchar(255)"`
	Phone   string `gorm:"default:null;type:varchar(255)"`
	Address string `gorm:"default:null;type:varchar(255)"`
	Image   string `gorm:"default:null;type:varchar(255)"`
}
