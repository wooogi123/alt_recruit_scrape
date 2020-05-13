package libs

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

func InitDb() {
	db, err := sql.Open("sqlite3", "file::recruits?mode=memory?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec("create table if not exists MMA (href string, title string, start_at timestamp, end_at timestamp)")
}
