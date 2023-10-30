/*
Copyright 2023 Peter Ydd <peterydd@outlook.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
