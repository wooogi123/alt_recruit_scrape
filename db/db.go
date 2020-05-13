package libs

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
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

func InitDb() {
	db, err := sql.Open("sqlite3", "recruits.db")
	if err != nil {
		log.Fatal(err)
	}
	mmaInit(db)
}

func init() {
	dbname := "recruits.db"
	_, err := os.Stat(dbname)
	if os.IsNotExist(err) {
		file, err := os.Create(dbname)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}
