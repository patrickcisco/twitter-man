# twitter-man

## Overview

A Go cli tool that streams tweets from Twitter to Webex Teams via adaptivecards. 

## Requirements
- A Webex Teams Bot Access Token (https://developer.webex.com)
- Twitter Developer Account + Project with Read Only permissions (https://developer.twitter.com) 

## Environment Variables
Secrets are passed into the cli application through environment variables. The following environment variables are required to run this application.
 
```bash
# Twitter environment variables
TWITTER_CONSUMER_KEY
TWITTER_CONSUMER_SECRET
TWITTER_ACCESS_TOKEN
TWITTER_ACCESS_SECRET

# Webex Teams environment variables
WEBEXTEAMS_TOKEN
```



### Examples
```
twitter-man --help
A cli application that streams tweets from Twitter

Usage:
  twitter-man [command]

Available Commands:
  help        Help about any command
  stream      Begin streaming twitter tweets

Flags:
      --config string   config file (default is $HOME/.twitter-man.yaml)
  -h, --help            help for twitter-man
      --log string      log level ( (default "info")

Use "twitter-man [command] --help" for more information about a command.
```

```
twitter-man stream --tag "#ciscolive" --tag "#ciscodevnet" --to_person_email "nobody@cisco.com"
```

![example usage](./assets/example-usage.gif)