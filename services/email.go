package services

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/DumanskyiDima/genesis_test/database"
	"github.com/DumanskyiDima/genesis_test/models"
)

func GenerateEmailContent(email string) ([]byte, error) {
	rate, err := FetchExchangeRate()
	if err != nil {
		log.Println("Failed to fetch exchange rate: ", err)
		return nil, err
	}

	subject := "Update on USD/UAH exchange rate"
	body := "Hello, Current USD/UAH exchange rate is: " + fmt.Sprintf("%f", rate["rate"])
	message := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	return message, nil
}

func SendEmail(ctx context.Context, user models.User) error {
	smtpHost := os.Getenv("SMTP_SERVER")
	smtpPort := 587

	sender := os.Getenv("SENDER_EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{user.Email}

	message, err := GenerateEmailContent(to[0])
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", sender, password, smtpHost)
	address := smtpHost + ":" + strconv.Itoa(smtpPort)

	sendChan := make(chan error, 1)
	go func() {
		log.Printf("Sending email to %s\n", to[0])
		sendChan <- smtp.SendMail(address, auth, sender, to, message)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-sendChan:
		if err != nil {
			return err
		}
	}

	log.Println("Email sent successfully")
	return nil
}

func SendEmailBatch(users []models.User, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, user := range users {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := SendEmail(ctx, user); err != nil {
			log.Printf("Error sending email to %s: %v\n", user.Email, err)
		}
	}
}

func SendEmails() {
	users := database.GetSubscribedUsers()
	fmt.Println("Emails will be sent to", len(users), "users")

	const batchSize = 5
	var wg sync.WaitGroup

	for i := 0; i < len(users); i += batchSize {
		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}

		wg.Add(1)
		go SendEmailBatch(users[i:end], &wg)
	}

	wg.Wait()
}
