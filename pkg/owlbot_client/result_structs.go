package owlbot_client

// MIT: https://github.com/mrbentarikau/pagst/blob/master/stdcommands/owlbot/owlbot.go

type Result struct {
	Word          string             `json:"word"`
	Definitions   []ResultDefinition `json:"definitions"`
	Pronunciation string             `json:"pronunciation"`
}

type ResultDefinition struct {
	Type       string `json:"type"`
	Definition string `json:"definition"`
	Example    string `json:"example"`
	ImageURL   string `json:"image_url"`
	Emoji      string `json:"emoji"`
}
