package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func New() *Database {
	return &Database{}
}

func (d *Database) Connect(dsn string, engine string) {

	var err error

	if engine == "sqlserver" {
		d.Conn, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	} else if engine == "mysql" {
		d.Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		log.Fatalf("invalid database engine [%s]", engine)
	}

	if err != nil {
		log.Fatalf("error connecting to the database. %s", err.Error())
	}

	log.Println("Connected to the database")
}
