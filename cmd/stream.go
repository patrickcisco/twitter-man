/*
Copyright © 2020 patrickcisco

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"twitter-man/pkg/adaptivecards"
	"twitter-man/pkg/twitter"
	"twitter-man/pkg/webexteams"

	gotwitter "github.com/dghubble/go-twitter/twitter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Begin streaming twitter tweets",
	Long:  `Stream twitter tweets from preconfigured handlers and filters`,
	Run: func(cmd *cobra.Command, args []string) {

		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		// Output to stdout instead of the default stderr
		// Can be any io.Writer
		log.SetOutput(os.Stdout)

		// Parse the log level setting (passed in as a cli flag)
		level, err := log.ParseLevel(LogLevel)
		if err != nil {
			panic(err)
		}
		// Set the log level for our logger
		log.SetLevel(level)

		// extract the to_person_email flag
		email, err := cmd.Flags().GetString("to_person_email")
		if err != nil || email == "" {
			log.Error("missing webex email")
			return
		}
		// extract the "tags" or "queries" to stream from
		tags, err := cmd.Flags().GetStringSlice("tag")
		if err != nil {
			log.Error("missing tags")
			return
		}

		log.Info("creating twitter client")
		twitterClient := &twitter.Client{
			ConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
			ConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
			AccessToken:    os.Getenv("TWITTER_ACCESS_TOKEN"),
			AccessSecret:   os.Getenv("TWITTER_ACCESS_SECRET"),
		}

		log.Info("creating webexteams client")
		webexteamsClient := &webexteams.Client{Token: os.Getenv("WEBEXTEAMS_TOKEN")}

		log.Info("creating twitter switch demux")

		demux := gotwitter.NewSwitchDemux()

		// this function will be called every time a new tweet occurs on the stream
		demux.Tweet = func(tweet *gotwitter.Tweet) {
			log.Debug("new tweet received")

			// handle the tweet concurrently
			// we don't want it blocking us from processing additional tweets in the stream
			go func(tweet gotwitter.Tweet) {
				// create an adaptive card from the new tweet
				log.Info("created adaptivecard tweet")
				adaptivecardTweet, _ := adaptivecards.Tweet(&adaptivecards.TweetOptions{
					Title:               "NEW TWEET",
					CreatorProfileImage: tweet.User.ProfileImageURLHttps,
					CreatorName:         tweet.User.Name,
					Date:                tweet.CreatedAt,
					Description:         tweet.Text,
					ViewURL:             fmt.Sprintf("https://twitter.com/%s/status/%s", tweet.User.IDStr, tweet.IDStr),
				})
				// create our Webex Teams input
				log.Info("sending adaptive card to webexteams")
				adaptivecardInput := &webexteams.SendAdaptiveCardInput{
					ToPersonEmail: email,
					Markdown:      "Sending you another tweet!",
					Attachements: []webexteams.Attachment{
						webexteams.Attachment{
							ContentType: "application/vnd.microsoft.card.adaptive",
							Content:     *adaptivecardTweet,
						},
					},
				}
				// if the debug log level is set, this will log the JSON payload being sent to Webex Teams
				debugJSON(adaptivecardInput)
				// make the Webex Teams API request
				res, err := webexteamsClient.SendAdaptiveCard(adaptivecardInput)
				if err != nil {
					log.Error("error sending adaptive tweet", err.Error())
				} else if res.StatusCode != 200 {
					log.Error("error sending adaptive tweet", "unexpected status code", res.StatusCode)
				} else {
					log.Info("success sending adaptive card")
				}
			}(*tweet)
		}
		// start our Twitter stream
		log.Info("starting stream")
		// make a bool channel that we can use to programmatically stop our client from streaming by calling
		// end <- true
		end := make(chan bool, 1)
		err = twitterClient.Stream(tags, &demux, end)
		if err != nil {
			log.Error("error starting stream", err.Error())
			return
		}
		// Wait for SIGINT and SIGTERM (HIT CTRL-C)
		// CTRL-C will cancel the stream and exit out of the progam
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		// Using <-ch is important here.
		// This tells our program to wait for SIGINT and SIGTERM (HIT CTRL-C)
		// The Twitter client will continue to stream until this condition is met.
		log.Println(<-ch)
	},
}

func debugJSON(item interface{}) {
	data, err := json.Marshal(item)
	if err != nil {
		log.Debug("unable to convert item to JSON")
	}
	log.Debug(string(data))
}

func init() {
	rootCmd.AddCommand(streamCmd)
	streamCmd.Flags().StringSliceP("tag", "t", []string{}, "Hashtags to stream")
	streamCmd.Flags().StringP("to_person_email", "e", "", "Webex Teams email")
}
