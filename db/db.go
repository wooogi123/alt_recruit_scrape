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

func InsertDb(recruits []mmaRecruit) {
	db, err := sql.Open("sqlite3", "recruits.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into MMA (href, title, start_at, end_at) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, recruit := range recruits {
		_, err = stmt.Exec(recruit.href, recruit.title, recruit.startAt, recruit.endAt)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InitDb() {
	db, err := sql.Open("sqlite3", "recruits.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec("create table if not exists MMA (href string, title string, start_at timestamp, end_at timestamp)")
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
