package webexteams_test

import (
	"fmt"
	"os"
	"testing"
	"twitter-man/pkg/adaptivecards"
	"twitter-man/pkg/webexteams"
)

func TestSendAdaptiveCards(t *testing.T) {
	token := os.Getenv("WEBEXTEAMS_TOKEN")
	client := &webexteams.Client{Token: token}
	email := ""
	tweet, _ := adaptivecards.Tweet(&adaptivecards.TweetOptions{
		Title:               "Today is an incredibly proud day",
		CreatorProfileImage: "https://pbs.twimg.com/profile_images/595243565900374016/6OoNrLqS_400x400.jpg",
		CreatorName:         "Chuck Robbins",
		Date:                "October 2, 2019",
		Description:         "Today is an incredibly proud day for @Cisco - being named #1 best place to work in the world by Great Places to Work!  I am honored to lead this team and proud of all the impact we are having around the world! #WorldsBestWorkplaces",
		ViewURL:             "https://twitter.com/ChuckRobbins/status/1179368668569964547",
	})
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

		resp, err := client.SendAdaptiveCard(&webexteams.SendAdaptiveCardInput{
			ToPersonEmail: email,
			Markdown:      "Hi, just sending you a tweet from Chuck!",
			Attachements: []webexteams.Attachment{
				webexteams.Attachment{
					ContentType: "application/vnd.microsoft.card.adaptive",
					Content:     *tweet,
				},
			},
		})

		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatal(err)
		}
	}
}
