package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB

// Init init db
func Init() {
	var err error
	
	db, err = gorm.Open(viper.GetString("db.type"),
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.addr"),
			viper.GetString("db.dbname")))
	
	if err != nil {
		log.Fatalf("database connection failed. Database name: %s", viper.GetString("db.dbname"))
	}
	
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	db.AutoMigrate(&TodoModel{})
}

// Close close db
func Close() {
	defer db.Close()
}
