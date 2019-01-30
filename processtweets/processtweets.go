package processtweets

import (
	regexp "regexp"
	"strings"

	"github.com/sleepypioneer/automated-journaling-with-GO/retrievetweets"
)

// TweetText is a tweets components split up
type TweetText struct {
	ID       string
	Date     string
	Round    string
	Day      string
	Text     string
	Hashtags []string
	Mentions []string
	Links    []string
}

var (
	regRound    = regexp.MustCompile(`#R(\S+)`)
	regDay      = regexp.MustCompile(`#R2D(\S+)`)
	regHashtags = regexp.MustCompile(`#(\S+)`)
	regHTTP     = regexp.MustCompile(`http(\S+)`)
	// regSymbols = regexp.MustCompile(`@(\S+)`)
	regMentions = regexp.MustCompile(`@(\S+)`)
	// SplitTweets is the retrieved tweets split out
	SplitTweets []TweetText
)

func removeNonPlainText(s string) string {
	s = regHTTP.ReplaceAllString(s, "")
	s = regRound.ReplaceAllString(s, "")
	s = regHashtags.ReplaceAllString(s, "")
	s = regMentions.ReplaceAllString(s, "")
	// regSymbols.ReplaceAllString(s, "")
	return strings.Trim(s, " ")
}

func extractAsset(s string, re *regexp.Regexp) []string {
	return re.FindAllString(s, -1)
}

// SeparateTweetsText separates tweets out
func SeparateTweetsText(inputTweet retrievetweets.Tweet) {
	splitTweet := TweetText{
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
