package twitter_test

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"twitter-man/pkg/twitter"

	gotwitter "github.com/dghubble/go-twitter/twitter"
)

func TestClient(t *testing.T) {

	client := &twitter.Client{
		ConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		AccessToken:    os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("TWITTER_ACCESS_SECRET"),
	}

	demux := gotwitter.NewSwitchDemux()
	demux.Tweet = func(tweet *gotwitter.Tweet) {
		log.Println("tweet!")
		log.Println(tweet.CreatedAt)
		log.Println(tweet.Text)
		log.Println(tweet.User.ProfileImageURLHttps)
		log.Println(tweet.User.Name)
		log.Println(tweet.FullText)

		url := fmt.Sprintf("https://twitter.com/%s/status/%s", tweet.User.IDStr, tweet.IDStr)
		log.Println(url)
	}

	log.Println("Starting Stream...")
	end := make(chan bool, 1)
	err := client.Stream([]string{"cat"}, &demux, end)
	if err != nil {
		t.Fatal(err)
	}

	// // Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	end <- true
	// }()
	// // end <- true
	log.Println(<-ch)
}
