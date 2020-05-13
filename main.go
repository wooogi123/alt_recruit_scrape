package main

import (
	db "./db"
	scrape "./scrape"
)

func main() {
	db.InitDb()
	db.InsertDb(scrape.MMAScrape())
}
