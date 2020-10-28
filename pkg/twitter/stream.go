package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	gotwitter "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Client contains fields to access the twitter API
type Client struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// Stream begins a Twitter Filter Stream based on a slice of queries
// demux allows you to configure what to happen when tweets, dms, etc. are found in the stream
// To cancel the stream, send a boolean value to the end channel.
func (v *Client) Stream(query []string, demux *gotwitter.SwitchDemux, end <-chan bool) error {
	config := oauth1.NewConfig(v.ConsumerKey, v.ConsumerSecret)
	token := oauth1.NewToken(v.AccessToken, v.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	filterParams := &twitter.StreamFilterParams{
		Track:         query,
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		return err
	}

	go demux.HandleChan(stream.Messages)
	go func(s *gotwitter.Stream) {
		defer stream.Stop()
		<-end
	}(stream)
	return nil
}
