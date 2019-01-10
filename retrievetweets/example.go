package retrievetweets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/dghubble/oauth1"
)

func handleGetTweets(w http.ResponseWriter, r *http.Request) {
	var maxIDQuery string
	var tweets []Tweet
	userID := os.Getenv("USERID")

	// httpClient will automatically authorize http.Requests
	httpClient := Config.Client(oauth1.NoContext, Token)

Outer:
	for i := 0; i < pages; i++ {
		// Twitter API request
		// userID is the Twitter handle
		// maxIDQuery is the last result on each page, so the API knows what the next page is
		path := fmt.Sprintf("https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=%v&include_rts=false&count=200%v", userID, maxIDQuery)
		if strings.Contains(path, "favicon.ico") { // API returns this so skip it, not needed
			break
		}

		resp, err := httpClient.Get(path)
		if err != nil {
			respondWithError(err, w)
			break
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			respondWithError(err, w)
			break
		}

		var gotTweets []Tweet
		err = json.Unmarshal(body, &gotTweets)
		if err != nil {
			respondWithError(err, w)
			break
		}

		// range through Tweets to clear out unneeded info
		for i, t := range gotTweets {

			if i == len(gotTweets)-1 {
				// this is the logic to tell Twitter API where the pages are
				if maxIDQuery == fmt.Sprintf("&max_id=%v", t.ID) {
					break Outer
				}
				maxIDQuery = fmt.Sprintf("&max_id=%v", t.ID)
			}

			// remove @ mentions and links from returned Tweets
			regAt := regexp.MustCompile(`@(\S+)`)
			t.Text = regAt.ReplaceAllString(t.Text, "")
			regHTTP := regexp.MustCompile(`http(\S+)`)
			t.Text = regHTTP.ReplaceAllString(t.Text, "")
			tweets = append(tweets, t)
		}
	}

	var result []string

	for _, t := range tweets {
		result = append(result, t.Text)
	}

	stringResult := strings.Join(result, "\n")

	w.WriteHeader(200)
	w.Write([]byte(stringResult))
}

// https://medium.com/@mycoralhealth/build-your-own-blockchain-twitter-recorder-in-go-4fa504e912c3
