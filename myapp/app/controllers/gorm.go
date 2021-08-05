package controllers

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"myapp/app/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}
	db.DB()
	db.AutoMigrate(&models.Video{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Pornstar{})
	db.AutoMigrate(&models.Casting{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.Tagging{})
	
	DB=db
}
