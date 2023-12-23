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
	ID uint		`gorm:"primaryKey"`
	Imo string
	Name string
	Date time.Time
}

type Departures struct {
	ID uint		`gorm:"primaryKey"`
	Imo string
	Name string
	Destination string
	Date time.Time
}

type BunkeringShips struct {
	Imo string	`gorm:"primaryKey"`
	Name string
	Available bool
}

type Tugs struct {
	Imo string	`gorm:"primaryKey"`
	Name string
	Available bool
}

func (ShipsInPort) TableName() string {
	return "ships_in_port"
  }

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(&ShipsInPort{}, &Arrivals{}, &Departures{}, &BunkeringShips{}, &Tugs{})
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

func SetUpTugs(db *gorm.DB, tugs []Tugs) {
	for _, tug := range tugs {
		db.Create(&Tugs{Imo: tug.Imo, Name: tug.Name, Available: tug.Available})
	}
}

func InsertNewArrival(db *gorm.DB, imo string, name string) int {
	ship := ShipsInPort{}
	result := db.Find(&ship, imo).Limit(1)
	if result.RowsAffected != 0 {
		return -1	// the ship is already in this port
	}
	db.Create(&Arrivals{Imo: imo, Name: name, Date: time.Now()})
	db.Create(&ShipsInPort{Imo: imo, Name: name})
	return 0
}

func InsertNewDeparture(db *gorm.DB, imo string, name string, destination string) int {
	ship := ShipsInPort{}
	result := db.Find(&ship, imo).Limit(1)
	if result.RowsAffected == 0 {
		return -1	// the ship is not in this port
	}
	db.Create(&Departures{Imo: imo, Name: name, Destination: destination, Date: time.Now()})
	db.Delete(&ShipsInPort{}, imo)
	return 0
}

func Bunkering(db *gorm.DB, imo string) (int, *BunkeringShips) {
	ship := ShipsInPort{}

	result := db.Find(&ship, imo).Limit(1)
	if result.RowsAffected == 0 {
		return -1, nil	// the ship is not in this port
	}
	tanker := BunkeringShips{}
	result = db.Where("available = ?", true).Find(&tanker).Limit(1)
	if result.RowsAffected == 0 {
		return 0, nil	// all the bunkering ships are unavailable
	}
	db.Model(&tanker).Select("available").Updates(BunkeringShips{Available: false})
	return 1, &tanker
}

func BunkeringEnd(db *gorm.DB, tankerImo string) {
	tanker := BunkeringShips{Imo: tankerImo}
	db.Model(&tanker).Select("available").Updates(BunkeringShips{Available: true})
}

func AcquireTugs(db *gorm.DB, imo string, requestType string, tugsNumber int) (int, []Tugs) {
	if requestType == "departure" {		// the ship must be in this port
		ship := ShipsInPort{}
		result := db.Find(&ship, imo).Limit(1)
		if result.RowsAffected == 0 {
			return -1, nil	// the ship is not in this port
		}
	}

	tugs := []Tugs{}
	result := db.Limit(tugsNumber).Where("available = ?", true).Find(&tugs)
	if result.RowsAffected < int64(tugsNumber) {
		return 0, nil	// not enough available tugs
	}
	db.Model(&tugs).Select("available").Updates(Tugs{Available: false})
	return 1, tugs
}

func ReleaseTugs(db *gorm.DB, imoList []string) {
	tug := Tugs{}
	for _, imo := range imoList {
		tug.Imo = imo
		db.Model(&tug).Select("available").Updates(Tugs{Available: true})
	}
}