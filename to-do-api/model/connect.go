package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var con *sql.DB

func Connect() *sql.DB {

	os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	
	_, err2 := con.Exec("CREATE TABLE IF NOT EXISTS TODO( name varchar(50), todo varchar(100));")
	if err2 != nil {
		log.Fatal(err2)
	}

	return db
}