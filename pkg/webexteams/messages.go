package webexteams

import (
	"bytes"
	"encoding/json"
	"net/http"
	"twitter-man/pkg/adaptivecards"
)

// SendAdaptiveCardInput contains fields to customize the Webex API call
type SendAdaptiveCardInput struct {
	ToPersonEmail string       `json:"toPersonEmail"`
	Markdown      string       `json:"markdown"`
	Attachements  []Attachment `json:"attachments"`
}

// Attachment contains the actual adaptive card
type Attachment struct {
	ContentType string             `json:"contentType"`
	Content     adaptivecards.Card `json:"content"`
}

// SendAdaptiveCard will send the adaptive card through the Webex Teams API
func (v *Client) SendAdaptiveCard(input *SendAdaptiveCardInput) (*http.Response, error) {
	url := "https://webexapis.com/v1/messages"

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+v.Token)
	client := http.DefaultClient
	return client.Do(req)
}
