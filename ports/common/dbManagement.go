package common

import (
	"log"
	"time"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
)

type Arrivals struct {
	Imo string	`gorm:"primaryKey"`
	Name string
	Date time.Time
}

type Departures struct {
	Imo string	`gorm:"primaryKey"`
	Name string
	Destination string
	Date time.Time
}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(&Arrivals{}, &Departures{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func InsertNewArrival(db *gorm.DB, imo string, name string) {
	db.Create(&Arrivals{Imo: imo, Name: name, Date: time.Now()})
}

func InsertNewDeparture(db *gorm.DB, imo string, name string, destination string) {
	db.Create(&Departures{Imo: imo, Name: name, Destination: destination, Date: time.Now()})
}