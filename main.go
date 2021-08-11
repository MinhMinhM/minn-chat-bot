package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)
//
//type LineMessage struct {
//	Destination string 		`json:"destination"`
//	Events      []struct {
//		ReplyToken string 	`json:"replyToken"`
//		Type       string	`json:"type"`
//		Timestamp  int64  	`json:"timestamp"`
//		Source     struct {
//			Type   string 	`json:"type"`
//			UserID string 	`json:"userId"`
//		}`json:"source"`
//		Message struct {
//			ID   string 	`json:"id"`
//			Type string 	`json:"type"`
//			Text string 	`json:"text"`
//		} `json:"message"`
//	} `json:"events"`
//}
//
//
//type ReplyMessage struct {
//	ReplyToken 	string `json:"replyToken"`
//	Messages   	[]Text `json:"messages"`
//}
//
//type Text struct {
//	Type 		string `json:"type"`
//	Text 		string `json:"text"`
//}
//
//type ProFile struct {
//	UserID        string `json:"userId"`
//	DisplayName   string `json:"displayName"`
//	PictureURL    string `json:"pictureUrl"`
//	StatusMessage string `json:"statusMessage"`
//}

//var ChannelToken = "L+cQvuHznzfIrgGNmhGpq727IvYZRnOab3IxPetVdPBuVgsYmFhc2TbTwBqaovgC6wCr6/zWjOSgsGccVTzHiiidvkD7dJc3GkdSmlZf8YSpi7Qf/kWTVAiA9h28bfvP+IpdSRGPT4CP8qzMK5CbqAdB04t89/1O/w1cDnyilFU="
//var ChannelToken = "Ys4yNEZl1H9WfRiyV9Db93+wR94d9Pk2q4AskWSEQB5bUyZCQNUngovXtk4qFd4f7HfwyZW0TAzvNisTqKa0QF54poGcsNXtDpNJQcFY+5Zc9mfy9lI+6MeRxU60aeOiSceoLlWjsi+FPc8ZEYxiVQdB04t89/1O/w1cDnyilFU="

func main() {


	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.POST("/webhook", WebHook)
	e.Logger.Fatal(e.Start(":1323"))

}



//
//func main() {
//	e := echo.New()
//	e.GET("/", func(c echo.Context) error {
//		return c.String(http.StatusOK, "ok")
//	})
//
//	e.POST("/webhook", func(c echo.Context) error {
//
//		Line := new(LineMessage)
//		fmt.Printf("%+v", Line.Events)
//		fmt.Println(len(Line.Events))
//		fmt.Println("//////////////////////////////")
//		//fmt.Println(Line.Events[0].ReplyToken)
//		if err := c.Bind(Line); err != nil {
//			log.Println("err")
//			return c.String(http.StatusOK, "error")
//		}
//
//		fullname := getProfile(Line.Events[0].Source.UserID)
//		text := Text{
//			Type : "text",
//			Text : "ข้อความเข้ามา : " + Line.Events[0].Message.Text  + " ยินดีต้อนรับ : " + fullname,
//		}
//
//		incomingText:=Line.Events[0].Message.Text
//		fmt.Println(incomingText[13:])
//		fmt.Println("////////////////")
//		if strings.Contains(incomingText, "ซื้อ"){
//			text = Text{
//				Type : "text",
//				Text : "Minh Golang Bot:ขอบคุณ คุณ " + fullname  + " ที่สอบถาม ขณะนี้ " + incomingText[13:]+ " หมด",
//			}
//		}
//
//
//		message := ReplyMessage{
//			ReplyToken : Line.Events[0].ReplyToken ,
//			Messages : []Text{
//				text,
//			},
//		}
//
//		fmt.Println(message)
//		replyMessageLine(message)
//
//		log.Println("%% message success")
//		return c.String(http.StatusOK, "ok")
//
//	})
//
//	e.Logger.Fatal(e.Start(":1323"))
//}

//
//func replyMessageLine(Message ReplyMessage) error {
//	value, _ := json.Marshal(Message)
//
//	url := "https://api.line.me/v2/bot/message/reply"
//
//	var jsonStr = []byte(value)
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Add("Authorization", "Bearer "+ChannelToken)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil
//	}
//	defer resp.Body.Close()
//
//	log.Println("response Status:", resp.Status)
//	log.Println("response Headers:", resp.Header)
//	body, _ := ioutil.ReadAll(resp.Body)
//	log.Println("response Body:", string(body))
//
//	return err
//}
//
//
//func getProfile(userId string) string {
//
//	url := "https://api.line.me/v2/bot/profile/" + userId
//
//	req, _ := http.NewRequest("GET", url, nil)
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Add("Authorization", "Bearer "+ChannelToken)
//
//	res, _ := http.DefaultClient.Do(req)
//
//	defer res.Body.Close()
//	body, _ := ioutil.ReadAll(res.Body)
//
//	var profile ProFile
//	if err := json.Unmarshal(body, &profile); err != nil {
//		log.Println("%% err \n")
//	}
//	log.Println(profile.DisplayName)
//	return profile.DisplayName
//
//}
