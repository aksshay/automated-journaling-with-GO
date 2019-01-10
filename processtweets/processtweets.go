package processtweets

import (
	regexp "regexp"

	"github.com/sleepypioneer/automated-journaling-with-GO/retrievetweets"
)

// tweetText is a tweets components split up
type tweetText struct {
	ID       string
	Date     string
	Round    string
	Day      string
	Text     string
	Hashtags []string
	Mentions []string
	Links    []string
}

const (
	mockTweet = "#R2D97 #100DaysOfCode started on phase two of the #PoweredByBertelsmann #UdacityDataScholars Nano Degree. Excited t… https://t.co/rfLGHfJaWN"
)

var (
	regRound    = regexp.MustCompile(`#R(\S+)`)
	regDay      = regexp.MustCompile(`#R2D(\S+)`)
	regHashtags = regexp.MustCompile(`#(\S+)`)
	regHTTP     = regexp.MustCompile(`http(\S+)`)
	// regSymbols = regexp.MustCompile(`@(\S+)`)
	regMentions = regexp.MustCompile(`@(\S+)`)
	// SplitTweets is the retrieved tweets split out
	SplitTweets []tweetText
)

func removeNonPlainText(s string) string {
	regHTTP.ReplaceAllString(s, "")
	regRound.ReplaceAllString(s, "")
	regHashtags.ReplaceAllString(s, "")
	regMentions.ReplaceAllString(s, "")
	// regSymbols.ReplaceAllString(s, "")
	return s
}

func extractAsset(s string, re *regexp.Regexp) []string {
	return re.FindAllString(s, -1)
}

// SeparateTweetsText separates tweets out
func SeparateTweetsText(inputTweet retrievetweets.Tweet) {
	splitTweet := tweetText{
		ID:       inputTweet.ID,
		Date:     inputTweet.Date,
		Round:    "",
		Day:      "",
		Text:     removeNonPlainText(inputTweet.Text),
		Hashtags: extractAsset(inputTweet.Text, regHashtags),
		Mentions: extractAsset(inputTweet.Text, regMentions),
		Links:    extractAsset(inputTweet.Text, regHTTP),
	}

	SplitTweets = append(SplitTweets, splitTweet)
}
