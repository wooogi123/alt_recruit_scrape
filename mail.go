package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"time"
)

func Parse(recruits []Recruit) (body string) {
	fm := template.FuncMap{
		"formatDate": func(t time.Time) string {
			return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
		},
	}
	t := template.Must(template.New("template.html").Funcs(fm).ParseFiles("template.html"))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, recruits)
	if err != nil {
		log.Fatal(err)
	}
	body = buf.String()
	return
}

func Send(recruits []Recruit) {
	from := os.Getenv("EMAIL")
	pass := os.Getenv("PASSWORD")
	to := os.Getenv("EMAIL")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: 산업지원 병역일터 공고 목록\n"
	body := Parse(recruits)
	msg := subject + mime + "\n" + body

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
