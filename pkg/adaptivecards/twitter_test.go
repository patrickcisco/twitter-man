package adaptivecards_test

import (
	"testing"
	"twitter-man/pkg/adaptivecards"
)

func TestTwitterAdaptiveCard(t *testing.T) {
	options := &adaptivecards.TweetOptions{
		Title:               "Just A Test",
		CreatorProfileImage: "https://google.com",
		CreatorName:         "Twitter-Man",
		Date:                "Today",
		Description:         "Just a description",
		ViewURL:             "https://twitter.com",
	}
	data, _ := adaptivecards.Tweet(options)
	t.Log(data)
}
