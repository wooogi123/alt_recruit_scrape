package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InsertDb(db *sql.DB, recruits []Recruit) {
	stmt, err := db.Prepare("insert into MMA (href, title, start_at, end_at) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, recruit := range recruits {
		_, err = stmt.Exec(recruit.Href, recruit.Title, recruit.StartAt, recruit.EndAt)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func SelectDb(db *sql.DB) (recruits []Recruit) {
	rows, err := db.Query("select * from MMA order by start_at desc")
	if err != nil {
		log.Fatal(err)
	}

	var recruit Recruit

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&recruit.Href, &recruit.Title, &recruit.StartAt, &recruit.EndAt)
		if err != nil {
			log.Fatal(err)
		}
		recruits = append(recruits, recruit)
	}
	return
}

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec("create table if not exists MMA (href string, title string, start_at timestamp, end_at timestamp)")
	InsertDb(db, MMAScrape())
	recruits := SelectDb(db)
	Send(recruits)
}
