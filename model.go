package main

type LineMessage struct {
	Destination string 		`json:"destination"`
	Events      []struct {
		ReplyToken string 	`json:"replyToken"`
		Type       string	`json:"type"`
		Timestamp  int64  	`json:"timestamp"`
		Source     struct {
			Type   string 	`json:"type"`
			UserID string 	`json:"userId"`
		}`json:"source"`
		Message struct {
			ID   string 	`json:"id"`
			Type string 	`json:"type"`
			Text string 	`json:"text"`
		} `json:"message"`
	} `json:"events"`
}


type ReplyMessage struct {
	ReplyToken 	string `json:"replyToken"`
	Messages   	[]Text `json:"messages"`
}

type Text struct {
	Type 		string `json:"type"`
	Text 		string `json:"text"`
}

type Texts []Text

type ProFile struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}


type T struct {
	Type string `json:"type"`
	Hero struct {
		Type        string `json:"type"`
		Url         string `json:"url"`
		Size        string `json:"size"`
		AspectRatio string `json:"aspectRatio"`
		AspectMode  string `json:"aspectMode"`
		Action      struct {
			Type string `json:"type"`
			Uri  string `json:"uri"`
		} `json:"action"`
	} `json:"hero"`
	Body struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Contents []struct {
			Type     string `json:"type"`
			Text     string `json:"text,omitempty"`
			Weight   string `json:"weight,omitempty"`
			Size     string `json:"size,omitempty"`
			Layout   string `json:"layout,omitempty"`
			Margin   string `json:"margin,omitempty"`
			Contents []struct {
				Type     string `json:"type"`
				Size     string `json:"size,omitempty"`
				Url      string `json:"url,omitempty"`
				Text     string `json:"text,omitempty"`
				Color    string `json:"color,omitempty"`
				Margin   string `json:"margin,omitempty"`
				Flex     int    `json:"flex,omitempty"`
				Layout   string `json:"layout,omitempty"`
				Spacing  string `json:"spacing,omitempty"`
				Contents []struct {
					Type  string `json:"type"`
					Text  string `json:"text"`
					Color string `json:"color"`
					Size  string `json:"size"`
					Flex  int    `json:"flex"`
					Wrap  bool   `json:"wrap,omitempty"`
				} `json:"contents,omitempty"`
			} `json:"contents,omitempty"`
			Spacing string `json:"spacing,omitempty"`
		} `json:"contents"`
	} `json:"body"`
	Footer struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Spacing  string `json:"spacing"`
		Contents []struct {
			Type   string `json:"type"`
			Style  string `json:"style,omitempty"`
			Height string `json:"height,omitempty"`
			Action struct {
				Type  string `json:"type"`
				Label string `json:"label"`
				Uri   string `json:"uri"`
			} `json:"action,omitempty"`
			Size string `json:"size,omitempty"`
		} `json:"contents"`
		Flex int `json:"flex"`
	} `json:"footer"`

}
type T2 struct {
	Type string `json:"type"`
	Body struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Contents []ContentFlex `json:"contents"`
	} `json:"body"`
}


type ContentFlex struct {
	Type string `json:"type"`
	Text string `json:"text"`
}