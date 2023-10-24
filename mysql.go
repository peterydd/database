package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Mysql struct {
	DriverName     string
	DataSourceName string
}

func (m *Mysql) Open() (*sql.DB, error) {
	db, err := sql.Open(m.DriverName, m.DataSourceName)
	if err != nil {
		log.Fatal("unable to use mysql data source name", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("unable to ping mysql data source", err)
	}
	log.Println("ping mysql data source success")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}
