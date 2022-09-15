package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("----------------------")

	}
}

func main() {
	os.Setenv("SlACK_BOT_TOKEN", "<redacted>")
	os.Setenv("SLACK_APP_TOKEN", "<redacted>")

	bot := slacker.NewClient(os.Getenv("SlACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		//Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, responce slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("ERROR In year conversion")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			responce.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
