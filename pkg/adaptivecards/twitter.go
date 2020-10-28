package adaptivecards

// TweetOptions options that can be passed to configure the adaptivecards
type TweetOptions struct {
	Title               string
	CreatorProfileImage string
	CreatorName         string
	Date                string
	Description         string
	ViewURL             string
}

// Tweet creates a Card given the available TweetOptions
func Tweet(options *TweetOptions) (*Card, error) {
	card := &Card{
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Version: "1.2",
		Type:    "AdaptiveCard",
		Actions: []Action{
			Action{
				Type:  "Action.OpenUrl",
				Title: "View",
				URL:   options.ViewURL,
			},
		},
		Body: []Body{
			Body{
				Type:   "TextBlock",
				Size:   "Medium",
				Weight: "Bolder",
				Text:   options.Title,
			},
			Body{
				Type: "ColumnSet",
				Columns: []Column{
					Column{
						Type:  "Column",
						Width: "auto",
						Items: []Item{
							Item{
								Type:  "Image",
								Style: "Person",
								URL:   options.CreatorProfileImage,
								Size:  "Small",
							},
						},
					},
					Column{
						Type:  "Column",
						Width: "stretch",
						Items: []Item{
							Item{
								Type:   "TextBlock",
								Weight: "Bolder",
								Text:   options.CreatorName,
								Wrap:   true,
							},

							Item{
								Type:     "TextBlock",
								Spacing:  "None",
								Text:     "Created " + options.Date,
								IsSubtle: true,
								Wrap:     true,
							},
						},
					},
				},
			},
			Body{
				Type: "TextBlock",
				Wrap: true,
				Text: options.Description,
			},
		},
	}
	return card, nil
}
