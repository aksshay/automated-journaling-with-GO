package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"

	"github.com/sleepypioneer/automated-journaling-with-GO/compilemarkdown"
	"github.com/sleepypioneer/automated-journaling-with-GO/processtweets"
	"github.com/sleepypioneer/automated-journaling-with-GO/retrievetweets"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Load API credentials
	retrievetweets.Config = oauth1.NewConfig(os.Getenv("APIKEY"), os.Getenv("APISECRET"))
	retrievetweets.Token = oauth1.NewToken(os.Getenv("TOKEN"), os.Getenv("TOKENSECRET"))
	retrievetweets.Retrieve()
	for _, t := range retrievetweets.ReturnedTweets.Tweets {
		processtweets.SeparateTweetsText(t)
	}
	fmt.Println("Tweets processed")
	compilemarkdown.WriteToMarkdown()
	fmt.Println("finished")
}
