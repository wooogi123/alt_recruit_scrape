package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type mmaRecruit struct {
	href    string
	title   string
	startAt time.Time
	endAt   time.Time
}

func main() {
	db, _ := sql.Open("sqlite3", "db/recruits.db")
	db.Exec("create table if not exists test (href string, title string, start_at timestamp, end_at timestamp)")
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into test (href, title, start_at, end_at) values (?, ?, ?, ?)")
	_, err := stmt.Exec("test", "etst", time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
