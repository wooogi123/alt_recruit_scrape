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

func mmaInit(db *sql.DB) {
	db.Exec("create table if not exists MMA (href string, title string, start_at timestamp, end_at timestamp)")
}

func init() {
	db, err := sql.Open("sqlite3", "db/recruits.db")
	if err != nil {
		log.Fatal(err)
	}
	mmaInit(db)
}
