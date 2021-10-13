package core

import (
	"gorm.io/gorm"
)

type BaseContext struct {
	Mysql *gorm.DB
}