package adaptivecards

// Card is an incomplete representation of an adaptivecard. https://adaptivecards.io/explorer/
type Card struct {
	Type    string   `json:"type"`
	Body    []Body   `json:"body"`
	Actions []Action `json:"actions"`
	Schema  string   `json:"$schema"`
	Version string   `json:"version"`
}

// Body is the card elements to show in the primary card region.
type Body struct {
	Type    string   `json:"type"`
	Size    string   `json:"size,omitempty"`
	Weight  string   `json:"weight,omitempty"`
	Text    string   `json:"text,omitempty"`
	Columns []Column `json:"columns,omitempty"`
	Wrap    bool     `json:"wrap,omitempty"`
}

// Action is the action to show in the card’s action bar.
type Action struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Column is a member of the ColumnSet
type Column struct {
	Type  string `json:"type"`
	Items []Item `json:"items"`
	Width string `json:"width"`
}

// Item is a member of a Column
type Item struct {
	Type     string `json:"type,omitempty"`
	Style    string `json:"style,omitempty"`
	URL      string `json:"url,omitempty"`
	Size     string `json:"size,omitempty"`
	Weight   string `json:"weight,omitempty"`
	Text     string `json:"text,omitempty"`
	Wrap     bool   `json:"wrap,omitempty"`
	IsSubtle bool   `json:"isSubtle,omitempty"`
	Spacing  string `json:"spacing,omitempty"`
}
