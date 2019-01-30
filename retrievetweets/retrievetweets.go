package retrievetweets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/dghubble/oauth1"
)

// Tweet is a collection of important info in each Tweet
type Tweet struct {
	Date string `json:"created_at"`
	Text string `json:"text"`
	ID   string `json:"id_str"`
}

// RetrievedTweets - returned & filtered Tweets
type RetrievedTweets struct {
	DateRange string  `json:"date_range"`
	Tweets    []Tweet `json:"tweets"`
}

// ReturnedTweets - instance of returnedTweets to be written out to JSON
var ReturnedTweets RetrievedTweets

// Twitter API page limit
const pages = 18

// Config for Twitter API
var Config *oauth1.Config

// Token for Twitter API
var Token *oauth1.Token

func respondWithError(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

func filterTweets(retrievedTweets []Tweet, maxIDQuery string) ([]Tweet, string) {
	var filteredTweets []Tweet
	for i, t := range retrievedTweets {
		if i == len(retrievedTweets)-1 {
			// this is the logic to tell Twitter API where the pages are
			if maxIDQuery == fmt.Sprintf("&max_id=%v", t.ID) {
				return filteredTweets, "end"
			}
		}
		// find Tweets with #100DaysOfCode and #R2D (round 2) hashtag
		// regOne := regexp.MustCompile(`#100DaysOfCode`)
		regTwo := regexp.MustCompile(`#R2D`)
		if /*regOne.MatchString(t.Text) && */ regTwo.MatchString(t.Text) {
			filteredTweets = append(filteredTweets, t)
		}

		// remove @ mentions and links from returned Tweets
		// regAt := regexp.MustCompile(`@(\S+)`)
		// t.Text = regAt.ReplaceAllString(t.Text, "")
		// regHTTP := regexp.MustCompile(`http(\S+)`)
		// t.Text = regHTTP.ReplaceAllString(t.Text, "")
		// tweets = append(tweets, t)

		maxIDQuery = fmt.Sprintf("&max_id=%v", t.ID)
	}
	return filteredTweets, maxIDQuery
}

// Retrieve retrieves tweets
func Retrieve( /*w http.ResponseWriter, r *http.Request*/ ) {
	userID := os.Getenv("USERID")
	// httpClient will automatically authorize http.Requests
	httpClient := Config.Client(oauth1.NoContext, Token)
	var maxIDQuery string

	for i := 0; i < pages; i++ {
		// Twitter API request
		// userID is the Twitter handle
		// maxIDQuery is the last result on each page, so the API knows what the next page is
		path := fmt.Sprintf("https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=%v&include_rts=false&count=200%v", userID, maxIDQuery)

		resp, err := httpClient.Get(path)
		if err != nil {
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		}

		var retrievedTweets []Tweet
		err = json.Unmarshal(body, &retrievedTweets)
		if err != nil {
			break
		}
		var filteredTweets []Tweet
		filteredTweets, maxIDQuery = filterTweets(retrievedTweets, maxIDQuery)

		ReturnedTweets.Tweets = append(ReturnedTweets.Tweets, filteredTweets...)
		// this is the logic to tell Twitter API where the pages are
		if maxIDQuery == "end" {
			break
		}
	}
	returnedTweetsJSON, _ := json.Marshal(ReturnedTweets)
	err := ioutil.WriteFile("returnedTweets.json", returnedTweetsJSON, 0644)
	if err != nil {

	}
}
