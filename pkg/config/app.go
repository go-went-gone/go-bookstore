package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//dsn := "admin:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	//openedDB, err := gorm.Open("mysql", "admin:password@tcp(localhost:9010)/gorm?charset=utf8&parseTime=True&loc=Local")
	//openedDB, err := gorm.Open("mysql", dsn)
	// openedDB, err := gorm.Open("mysql", "admin:password@/simplerest") // works to push but doesnt save string properly
	openedDB, err := gorm.Open("mysql", "root:roottoor@/zmhdb?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	db = openedDB
}

func GetDB() *gorm.DB {
	return db
}
