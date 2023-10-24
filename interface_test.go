package database

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestMysql(t *testing.T) {
	var dataBase Database = &Mysql{
		DriverName:     "mysql",
		DataSourceName: "test:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
	}

	db, err := dataBase.Open()
	if err != nil {
		spew.Dump(err)
	}
	t.Log(db)

	result, err := db.Exec("drop table if exists test")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(result)

	result, err = db.Exec("create table if not exists test (id int)")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(result)

	result, err = db.Exec("insert into test (id) values (1)")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(result)

	_, err = db.Query("select id from test")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(err)

	err = db.Close()
	if err != nil {
		spew.Dump(err)
	}
	t.Log(err)
}

func TestOracle(t *testing.T) {
	var dataBase Database = &Oracle{
		DriverName:     "oracle",
		DataSourceName: "oracle://test:password@127.0.0.1:1521/FREEPDB1",
	}

	db, err := dataBase.Open()
	if err != nil {
		spew.Dump(err)
	}
	t.Log(db)

	result, err := db.Exec("drop table test purge")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(result)

	result, err = db.Exec("create table test (id int)")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(result)

	result, err = db.Exec("insert into test (id) values (1)")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(result)

	_, err = db.Query("select id from test")
	if err != nil {
		spew.Dump(err)
	}
	t.Log(err)

	err = db.Close()
	if err != nil {
		spew.Dump(err)
	}
	t.Log(err)
}
