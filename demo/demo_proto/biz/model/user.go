package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(128) not null;unique_index"`
	Password string `gorm:"type:varchar(64) not null"`
}

func (User) TableName() string {
	return "user"
}
