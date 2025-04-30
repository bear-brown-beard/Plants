package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `gorm:"column:first_name;type:varchar(100);not null" json:"first_name" binding:"required"`
	LastName  string  `gorm:"column:last_name;type:varchar(100);not null" json:"last_name" binding:"required"`
	Email     string  `gorm:"column:email;type:varchar(100);unique" json:"email" binding:"required,email"`
	Password  string  `gorm:"column:password;type:varchar(100);not null" json:"password" binding:"required,min=6"`
	City      string  `gorm:"column:city;type:varchar(100);" json:"city"`
	Plants    []Plant `gorm:"foreignKey:UserID" json:"plants"`
}
