package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	err := godotenv.Load() // Loading the environment variables from the .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slackToken := os.Getenv("SLACK_API_KEY") // Getting the sack api key from the .env file
	if slackToken == "" {
		log.Fatal("Slack API token not found in .env file")
	}

	//initializing the slack client
	api := slack.New(slackToken)

	//Reminder message or motivation message
	channel := "#project"
	message := "💡 *Hello Fellow Devs!* 💡\n\n" +
		"🚀 Keep up the great work, you're making excellent progress! 💻\n" +
		"💧 Remember to take a short break and stay hydrated as you code. Your well-being is important! 😊\n" +
		"📌 Quick Reminder: Please aim to complete this task by the end of today, and don’t forget to push your latest code updates to the repository so we can focus on planning and implementing new features tomorrow."

	err = sendSlackMessage(api, channel, message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Println("Reminder sent successfully!")
}

// function to send the message to the slack channel
func sendSlackMessage(api *slack.Client, channel, message string) error {
	_, _, err := api.PostMessage(channel, slack.MsgOptionText(message, false))
	return err
}
