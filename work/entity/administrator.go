package entity

import "time"

type AdministratorsInfo struct {
	ID     uint64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	LastLoginDate time.Time `gorm:"column:last_login_date" json:"last_login_date"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
	Status uint `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" format:"YYYY-mm-dd"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
func (AdministratorsInfo) TableName() string {
	return "administrators"
}

type Administrator struct {
	Name string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	LastLoginDate time.Time `gorm:"column:last_login_date" json:"last_login_date"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
	Status uint `gorm:"column:status" json:"status"`
}
