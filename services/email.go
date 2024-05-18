package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	. "github.com/DumanskyiDima/genesis_test/database"

	. "github.com/DumanskyiDima/genesis_test/models"
	// . "github.com/DumanskyiDima/genesis_test/views"
)

func GenerateEmailContent(email string) []byte {
	// todo add currency rate
	subject, body := "Update on exchange rate", "Hello, the exchange rate has been updated."
	message := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	return message
}

func SendEmail(user User) error {
	smtpHost := os.Getenv("SMTP_SERVER")
	smtpPort := 573

	sender := os.Getenv("SENDER_EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{user.Email}
	message := GenerateEmailContent(to[0])

	auth := smtp.PlainAuth("", sender, password, smtpHost)
	address := smtpHost + ":" + strconv.Itoa(smtpPort)

	err := smtp.SendMail(address, auth, sender, to, message)
	if err != nil {
		return err
	}
	log.Println("Email sent successfully")
	return nil
}

func SendEmailBatch(users []User) {
	for _, user := range users {
		if err := SendEmail(user); err != nil {
			log.Printf("Error sending email to %s: %v\n", user.Email, err)
		}
	}
}

func SendEmails() {
	users := GetSubscribedUsers()
	fmt.Println("Emails will be sent to", len(users), "users")

	batchSize := 5
	for i := 0; i < len(users); i += batchSize {
		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}
		SendEmailBatch(users[i:end])
	}
}
