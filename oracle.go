package database

import (
	"database/sql"
	_ "github.com/sijms/go-ora/v2"
	"log"
	"time"
)

type Oracle struct {
	DriverName     string
	DataSourceName string
}

func (o *Oracle) Open() (*sql.DB, error) {
	db, err := sql.Open(o.DriverName, o.DataSourceName)
	if err != nil {
		log.Fatal("unable to use oracle data source name", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("unable to ping oracle data source", err)
	}
	log.Println("ping oracle data source success")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}
