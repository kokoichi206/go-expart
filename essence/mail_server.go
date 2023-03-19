package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/jhillyerd/enmime"
	"github.com/joho/godotenv"
)

type mail struct {
	from string
	user string
	pass string
	to   string
}

func env() mail {
	return mail{
		from: os.Getenv("SMTP_USER"),
		user: os.Getenv("SMTP_USER"),
		pass: os.Getenv("SMTP_PASS"),
		to:   os.Getenv("MAIL_TO"),
	}
}

func sendMail() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	smtpHost := "smtp.gmail.com:587"
	cfg := env()
	smtpAuth := smtp.PlainAuth(
		"",
		cfg.from,
		cfg.pass,
		"smtp.gmail.com",
	)
	msg := "こんちゃ"

	sender := enmime.NewSMTP(smtpHost, smtpAuth)

	master := enmime.Builder().
		From("kokoichi206", cfg.from).
		Subject("THE email title").
		Text([]byte(msg)).
		HTML([]byte("<p>本文<br /><br />だよ</p>")).
		AddFileAttachment("./static/logo.png")

	if err := master.To(cfg.to, cfg.to).Send(sender); err != nil {
		log.Fatal(err)
	}
	// if err := smtp.SendMail(smtpHost, smtpAuth, cfg.from, []string{cfg.to}, []byte(msg)); err != nil {
	// 	log.Fatal(err)
	// }
}
