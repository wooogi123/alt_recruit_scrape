package mail

import (
	scrape "../scrape"
	"log"
	"net/smtp"
	"os"
)

func Send(recruits []scrape.Recruit) {
	from := os.Getenv("FROM")
	pass := os.Getenv("PASS")
	to := os.Getenv("TO")
	msg := "From: " + from + "\n"
	msg += "To: " + to + "\n"
	msg += "Subject: 산업지원 병역일터 공고 목록\n\n"

	for _, recruit := range recruits {
		msg += recruit.Title + "\n"
	}

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}
