package main

import (
	db "./db"
	scrape "./scrape"
	"fmt"
)

func main() {
	db.InitDb()
	fmt.Println(scrape.MMAScrape())
}
