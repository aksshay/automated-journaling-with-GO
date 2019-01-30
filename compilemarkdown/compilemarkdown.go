package compilemarkdown

import (
	"io/ioutil"

	"github.com/sleepypioneer/automated-journaling-with-GO/processtweets"
)

func hashtags(hashtags []string) string {
	var s string
	s = s + "### Hashtags: "
	for _, h := range hashtags {
		s = s + h + ", "
	}
	s = s + "\n"
	return s
}

func mentions(mentions []string) string {
	var s string
	s = s + "### Mentions: "
	for _, m := range mentions {
		s = s + m + ", "
	}
	s = s + "\n"
	return s
}

func links(links []string) string {
	var s string
	s = s + "### Links: "
	for _, l := range links {
		s = s + l + ", "
	}
	s = s + "\n"
	return s
}

func tweetLink(ID string) string {
	var s string
	s = s + "Link to Tweet: https://twitter.com/" + ID + "\n\n"
	return s
}

func formatTweet(t processtweets.TweetText) string {
	var s string
	s = s + "link\n"
	s = s + header("Round: "+t.Round+" Day: "+t.Day+"\n")
	s = s + "Date " + t.Date + "\n"
	s = s + "### Todays progress: " + t.Text + "\n\n"
	if len(t.Hashtags) > 0 {
		s = s + hashtags(t.Hashtags)
	}
	if len(t.Mentions) > 0 {
		s = s + mentions(t.Mentions)
	}
	if len(t.Links) > 0 {
		s = s + links(t.Links)
	}
	s = s + tweetLink(t.ID)
	s = s + toTop()
	s = s + lineBreak()
	return s
}

// WriteToMarkdown writes out to a markdown file the constructed data
func WriteToMarkdown() {
	var mdString string
	for _, tweet := range processtweets.SplitTweets {
		mdString = mdString + formatTweet(tweet)

	}
	mdByte := []byte(mdString)
	err := ioutil.WriteFile("log.md", mdByte, 0644)
	if err != nil {
		panic(err)
	}
}
