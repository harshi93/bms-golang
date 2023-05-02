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

// help initialize connection to db
func Connect() {
	dbconn, err := gorm.Open("mysql", "harshi93:harsh120522/gobookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = dbconn
}

func GetDB() *gorm.DB {
	return db
}
