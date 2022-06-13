package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	lambda.Start(garbageAutomation)
}

func garbageAutomation() {
	garbagePickups, err := readFile("garbage.csv")

	if err != nil {
		log.Fatal("There was an error reading CSV file")
	}

	currentDate := time.Now()
	for _, p := range garbagePickups {
		date, err := time.Parse("2006-01-02", p[1])

		if err != nil {
			log.Fatal("There was an issue parsing CSV date into object")
		}

		if currentDate.Day()+1 == date.Day() && currentDate.Month() == date.Month() {
			sendSMS(p[0])
		}
	}
}

func readFile(csvFile string) ([][]string, error) {
	f, err := os.Open(csvFile)

	if err != nil {
		log.Fatal("There was an error reading CSV file")
	}

	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal("there was an error reading first line in CSV file")
		return [][]string{}, err
	}

	garbagePickups, err := r.ReadAll()

	if err != nil {
		log.Fatal("There was an error reading all CSV values")
		return [][]string{}, err
	}

	return garbagePickups, nil
}

func sendSMS(msg string) {

	env := os.Getenv("APP_ENV")

	if env == "" || env == "development" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("There was an error loading .env file")
		}
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("PHONE_NUMBER"))
	params.SetBody(msg + " kanta")
	params.SetMessagingServiceSid(os.Getenv("SERVICE_ID"))
	resp, err := client.Api.CreateMessage(params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
