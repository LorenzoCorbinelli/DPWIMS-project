package common

import (
	"log"
	"time"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
)

type Arrival struct {
	Imo string	`gorm:"primaryKey"`
	Name string
	Date time.Time
}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(&Arrival{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func InsertNewArrival(db *gorm.DB, imo string, name string) {
	db.Create(&Arrival{Imo: imo, Name: name, Date: time.Now()})
}