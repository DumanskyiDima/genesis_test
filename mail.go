package main

import (
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func mail() {
	// SMTP server configuration
	smtpHost := os.Getenv("SMTP_SERVER")
	// smtpHost := "smtp.gmail.com"
	smtpPort := os.Getenv("587")

	sender := "dt.dumanskyi@gmail.com"
	password := "kesfhtifzgwnksgm"

	// Recipient email address
	to := []string{"dumanskiy.dima@gmail.com"}

	// Email content
	subject := "Test Email"
	body := "This is a test email sent from Go."

	// SMTP authentication configuration
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	// Compose email message
	message := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	address := smtpHost + ":" + strconv.Itoa(smtpPort)
	log.Println(address)

	// Send email
	err := smtp.SendMail(address, auth, sender, to, message)
	if err != nil {
		log.Fatal("Error sending email:", err)
	}
	log.Println("Email sent successfully")
}
