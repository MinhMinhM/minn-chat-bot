package test


type Form1Request struct {
	DateSubmitted string `json:"date_submitted"`
	UniqueId int `json:"unique_id"`
	AgentEmail string `json:"agent_email"`
	RegisService string `json:"regis_service"`
	File1TS string `json:"file1_ts"`
	File2TS string `json:"file2_ts"`
	TelephoneNo string `json:"telephone_no"`
	SecondTelephoneNo string `json:"second_telephone_no"`
	ShippingAddress string `json:"shipping_address"`
	IsTax string `json:"is_tax"`
	NameSurname string `json:"name_surname"`
	NationalID string `json:"national_id"`
	TaxAddress string `json:"tax_address"`
	TaxName string `json:"tax_name"`
	File1Ekyc string `json:"file1_ekyc"`
	File2Ekyc string `json:"file2_ekyc"`
	Signature string `json:"signature"`
}

func (u *Form1Request)TableName() string{
	return "form1"
}