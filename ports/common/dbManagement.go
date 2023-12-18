package common

import (
	"log"
	"time"
	"gorm.io/gorm"
)

type ShipsInPort struct {
	Imo string	`gorm:"primaryKey"`
	Name string
}

type Tabler interface {
	TableName() string
}

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

type BunkeringShips struct {
	Imo string	`gorm:"primaryKey"`
	Name string
	Available bool
}

func (ShipsInPort) TableName() string {
	return "ships_in_port"
  }

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(&ShipsInPort{}, &Arrivals{}, &Departures{}, &BunkeringShips{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func SetUpBunkeringShips(db *gorm.DB, ships []BunkeringShips) {
	for _, tanker := range ships {
		db.Create(&BunkeringShips{Imo: tanker.Imo, Name: tanker.Name, Available: tanker.Available})
	}
}

func InsertNewArrival(db *gorm.DB, imo string, name string) {
	db.Create(&Arrivals{Imo: imo, Name: name, Date: time.Now()})
	db.Create(&ShipsInPort{Imo: imo, Name: name})
}

func InsertNewDeparture(db *gorm.DB, imo string, name string, destination string) {
	db.Create(&Departures{Imo: imo, Name: name, Destination: destination, Date: time.Now()})
	db.Delete(&ShipsInPort{}, imo)
}

//NOT TESTED
func Bunkering(db *gorm.DB, imo string) (int, *BunkeringShips) {
	ship := ShipsInPort{}

	result := db.First(&ship, imo)
	if result.RowsAffected == 0 {
		return -1, nil	// the ship is not in this port
	}
	tanker := BunkeringShips{}
	result = db.Where("available = ?", true).First(&tanker)
	if result.RowsAffected == 0 {
		return 0, nil	// all the bunkering ships are unavailable
	}
	// UPDATE AVAILABLE = FALSE ON THE DB
	return 1, &tanker
}