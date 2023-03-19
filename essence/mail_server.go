package main

import (
	"log"
	"net/smtp"
	"os"
	"os/exec"

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
	exec.Command("source .env")
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
	if err := smtp.SendMail(smtpHost, smtpAuth, cfg.from, []string{cfg.to}, []byte(msg)); err != nil {
		log.Fatal(err)
	}
	// conn, err := tls.Dial("tcp", "smtp.gmail.com:465", &tls.Config{})
}
