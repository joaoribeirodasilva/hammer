package database

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func New() *Database {
	return &Database{}
}

func (d *Database) Connect(dsn string) {

	var err error
	d.Conn, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to the database. %s", err.Error())
	}

	log.Println("Connected to the database")
}
