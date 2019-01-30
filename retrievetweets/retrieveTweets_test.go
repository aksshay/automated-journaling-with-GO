package retrievetweets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockTweets = []Tweet{
	{
		"12-12-2018",
		"#100DaysOfCode Tweet one",
		"1",
	},
	{
		"12-12-2018",
		"#R2D2 Tweet two",
		"2",
	},
	{
		"12-12-2018",
		"Tweet three",
		"3",
	},
}

func TestFilterTweets(t *testing.T) {
	expected := struct {
		tweets     []Tweet
		maxIDQuery string
	}{
		[]Tweet{{Date: "12-12-2018", Text: "#R2D2 Tweet two", ID: "2"}},
		"&max_id=3",
	}
	filteredTweets, maxIDQuery := filterTweets(mockTweets, "0")

	assert.Equal(t, expected.tweets, filteredTweets)
	assert.Equal(t, expected.maxIDQuery, maxIDQuery)
}
