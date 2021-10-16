package test

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




//type Form1Json struct {
//	DateSubmitted time.Time `json:"date_submitted"`
//	UniqueId int `json:"UniqueID"`
//	AgentEmail string `json:"Agent Email"`
//	RegisService string `json:"บริการที่สมัคร"`
//	File1TS string `json:"File1 True store"`
//	File2TS string `json:"File2 True store"`
//	TelephoneNo string `json:"บอร์ที่ใช้ในการสมัครบริการ"`
//	SecondTelephoneNo string `json:"เบอร์ติดต่อสำรอง"`
//	ShippingAddress string `json:"ที่อยู่ในการจัดส่ง"`
//	IsTax string `json:"ใบกำกับภาษี"`
//	NameSurname string `json:"ชื่อ-นามสกุล"`
//	NationalID string `json:"เลขบัตรประชาชน"`
//	TaxAddress string `json:"ที่อยู่สำหรับใบกำกับภาษี"`
//	TaxName string `json:"ชื่อ-นามสกุล"`
//	TaxNationalID string `json:"เลขบัตรประชาชน"`
//	File1Ekyc string `json:"File1 ekyc"`
//	File2Ekyc string `json:"File2 ekyc"`
//	Signature string `json:"Signature"`
//}

//type Form1Request struct {
//	UniqueId int `json:"unique_id"`
//	AgentEmail string `json:"agent_email"`
//	RegisService string `json:"regis_service"`
//	File1TS string `json:"file1_ts"`
//	File2TS string `json:"file2_ts"`
//	TelephoneNo string `json:"telephone_no"`
//	SecondTelephoneNo string `json:"second_telephone_no"`
//	ShippingAddress string `json:"shipping_address"`
//	IsTax string `json:"is_tax"`
//	NameSurname string `json:"name_surname"`
//	NationalID string `json:"national_id"`
//	TaxAddress string `json:"tax_address"`
//	TaxName string `json:"tax_name"`
//	TaxNationalID string `json:"tax_national_id"`
//	File1Ekyc string `json:"file1_ekyc"`
//	File2Ekyc string `json:"file2_ekyc"`
//	Signature string `json:"signature"`
//}
//
//func (u *Form1Request)TableName() string{
//	return "form1"
//}