package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/sleepypioneer/automated-journaling-with-GO/processtweets"
	"github.com/sleepypioneer/automated-journaling-with-GO/retrievetweets"
)

// makeMuxRouter defines and creates routes
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", printToScreen).Methods("GET")
	return muxRouter
}

func printToScreen(w http.ResponseWriter, r *http.Request) {
	var result []string

	// only for printing results
	for _, t := range processtweets.SplitTweets {
		result = append(result, t.ID)
		result = append(result, t.Date)
		result = append(result, t.Text)
		result = append(result, "------------------------------------------------------------------------------------")
	}

	stringResult := strings.Join(result, "\n")

	w.WriteHeader(200)
	w.Write([]byte(stringResult))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Load API credentials
	retrievetweets.Config = oauth1.NewConfig(os.Getenv("APIKEY"), os.Getenv("APISECRET"))
	retrievetweets.Token = oauth1.NewToken(os.Getenv("TOKEN"), os.Getenv("TOKENSECRET"))
	retrievetweets.Retrieve()
	for _, t := range retrievetweets.Tweets {
		processtweets.SeparateTweetsText(t)
	}

	s := &http.Server{
		Addr:           os.Getenv("PORT"),
		Handler:        makeMuxRouter(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
