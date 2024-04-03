package main

import (
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Notifier interface {
	Notify(msg string) error
}

type EmailNotifier struct {
	fromEmail string
	password  string
	smtpHost  string
	smtpPort  string
}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{
		fromEmail: os.Getenv("FROM_EMAIL"),
		password:  os.Getenv("APP_PASSWORD"),
		smtpHost:  "smtp.gmail.com",
		smtpPort:  "587",
	}
}

func (n *EmailNotifier) Notify(msg string) error {
	message := []byte(fmt.Sprintf("Subject: smeti \n\n %s\n", msg))
	receivers := strings.Split(os.Getenv("EMAIL_RECEIVERS"), ",")

	auth := smtp.PlainAuth("", n.fromEmail, n.password, n.smtpHost)
	addr := fmt.Sprintf("%s:%s", n.smtpHost, n.smtpPort)

	err := smtp.SendMail(addr, auth, n.fromEmail, receivers, message)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(garbageAutomation)
}

func garbageAutomation() {
	garbagePickups, err := readFile("garbage.csv")
	if err != nil {
		fmt.Printf("there was an error reading file: %v", err)
	}

	todayGarbagePickups := getTodayGarbagePickups(garbagePickups)
	// only send email notifications for now
	notifier := NewEmailNotifier()

	if len(todayGarbagePickups) > 0 {
		msg := strings.Join(todayGarbagePickups, ", ")
		err := notifier.Notify(msg)
		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Println("Notification sent successfully")
	}
}

func getTodayGarbagePickups(garbagePickups [][]string) []string {
	todayGarbagePickups := make([]string, 0, 5)
	currentDate := time.Now()

	for _, p := range garbagePickups {
		date, err := time.Parse("2006-01-02", p[1])

		if err != nil {
			log.Fatal("There was an issue parsing CSV date into object")
		}

		if currentDate.Day()+1 == date.Day() && currentDate.Month() == date.Month() {
			todayGarbagePickups = append(todayGarbagePickups, p[0])
		}
	}

	return todayGarbagePickups
}

func readFile(name string) ([][]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	garbagePickups, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return garbagePickups, nil
}
