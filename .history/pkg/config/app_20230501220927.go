/*
	this file is intended to return db object

to be consumed by other files that want to
interact with db
*/
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)
