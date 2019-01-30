package processtweets

import (
	"testing"

	"github.com/sleepypioneer/automated-journaling-with-GO/retrievetweets"
	"github.com/stretchr/testify/assert"
)

var (
	mockTweet = retrievetweets.Tweet{
		Date: "02-02-2019",
		Text: "#R2D97 #100DaysOfCode started on phase two of the #PoweredByBertelsmann #UdacityDataScholars Nano Degree. Excited t… https://t.co/rfLGHfJaWN",
		ID:   "112",
	}
)

func TestRemoveNonPlainText(t *testing.T) {
	expected := "started on phase two of the   Nano Degree. Excited t…"
	assert.Equal(t, expected, removeNonPlainText(mockTweet.Text))
}

func TestSeparateTweetsText(t *testing.T) {
	expected := []tweetText{
		tweetText{
			ID:       "112",
			Date:     "02-02-2019",
			Round:    "",
			Day:      "",
			Text:     "started on phase two of the   Nano Degree. Excited t…",
			Hashtags: []string{"#R2D97", "#100DaysOfCode", "#PoweredByBertelsmann", "#UdacityDataScholars"},
			// Mentions: []string{},
			Links: []string{"https://t.co/rfLGHfJaWN"},
		},
	}
	SeparateTweetsText(mockTweet)
	assert.Equal(t, expected, SplitTweets)
}
