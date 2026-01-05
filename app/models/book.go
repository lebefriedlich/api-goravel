package models

import (
	"github.com/goravel/framework/database/orm"
)

type Book struct {
	orm.Model
	Title         string
	Author        string
	PublishedYear int
	Stock         int
}
