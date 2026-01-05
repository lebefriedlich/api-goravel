package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name     string
	Email    string `gorm:"size:150;uniqueIndex;not null"`
	Password string
	Role     string `gorm:"size:20;not null;default:user"`
}

func (u *User) GetKey() any {
	return u.ID
}
