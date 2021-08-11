package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
	//authUser model.AuthUser
	sv *testService

}

// NewHandler return +Handler
func NewHandler() *handler {
	// Inject service, store here!!

	store := NewTestStore()
	service := NewTestService(store)
	return &handler{
		sv: service,
	}
}

func (h *handler) webhaook(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}

func WebHook(c echo.Context) error {

	bot, err := linebot.New("6c2154dbfbd911af13de10982e2db6cf",
		"Ys4yNEZl1H9WfRiyV9Db93+wR94d9Pk2q4AskWSEQB5bUyZCQNUngovXtk4qFd4f7HfwyZW0TAzvNisTqKa0QF54poGcsNXtDpNJQcFY+5Zc9mfy9lI+6MeRxU60aeOiSceoLlWjsi+FPc8ZEYxiVQdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}

	Line := new(LineMessage)

	//fmt.Println(Line.Events[0].ReplyToken)
	if err := c.Bind(Line); err != nil {
		log.Println("err")
		return c.String(http.StatusOK, "error")
	}

	fmt.Println(Line.Events[0].Message)

	if Line.Events[0].Message.Type=="location"{
		//text := Text{}
		//text = Text{
		//	Type: "text",
		//	Text: "Minh Golang Bot: ที่อยู่ของท่าน อยู่นอกเขตให้บริการ",
		//}
		//
		//message := ReplyMessage{
		//	ReplyToken: Line.Events[0].ReplyToken,
		//	Messages:   []Text{text},
		//}
		//replyMessageLine(message)

		// Unmarshal JSON
		flexContainer, err := linebot.UnmarshalFlexMessageJSON(flexs)
		if err != nil {
			log.Println(err)
		}
		// New Flex Message
		msg1:=linebot.NewTextMessage("พื้นที่ของท่าน อยู่ในเขตให้บริการ")
		//message:=Texts{}
		//message
		flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)
		//fmt.Printf("%+v\n", flexMessage)
		//fmt.Println("////////////////")
		//fmt.Printf("%+v\n", msg1)
		// Reply Message
		//bot.ReplyMessage(Line.Events[0].ReplyToken, linebot.NewTextMessage("พื้นที่ของท่าน อยู่ในเขตให้บริการ")).Do()
		bot.PushMessage(Line.Events[0].Source.UserID, msg1).Do()
		bot.ReplyMessage(Line.Events[0].ReplyToken, flexMessage).Do()


		return c.NoContent(http.StatusOK)
	}


	if Line.Events[0].Message.Type=="text" {
		fullname := getProfile(Line.Events[0].Source.UserID)
		incomingText := Line.Events[0].Message.Text
		text := Text{}
		//if strings.Contains(incomingText, "ซื้อ"){
		if len(incomingText)>=12 && incomingText[0:12] == "ซื้อ"{
			text = Text{
				Type: "text",
				Text: "ขอบคุณ คุณ " + fullname + " ที่สนใจ" + incomingText[13:] + " อีกซักครู่ พนักงานประจำสาขาจะมาตอบในไม่ช้า",
			}
		} else if incomingText=="ติดต่อพนักงาน"{
			text = Text{
				Type : "text",
				Text : "ขอบคุณ คุณ "+fullname+ " ที่ส่งข้อความถึงเรา พนักงานประจำสาขาจะมาตอบในไม่ช้า",
			}
			//return c.NoContent(http.StatusOK)
		} else{
			return c.NoContent(http.StatusOK)
		}

		message := ReplyMessage{
			ReplyToken: Line.Events[0].ReplyToken,
			Messages:   []Text{text},
		}
		replyMessageLine(message)
		return c.NoContent(http.StatusOK)
	}

	log.Println("%% message success")
	return c.NoContent(http.StatusOK)
}



func replyMessageLine(Message ReplyMessage) error {
	value, _ := json.Marshal(Message)

	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return err
}


func getProfile(userId string) string {

	url := "https://api.line.me/v2/bot/profile/" + userId

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var profile ProFile
	if err := json.Unmarshal(body, &profile); err != nil {
		log.Println("%% err \n")
	}
	log.Println(profile.DisplayName)
	return profile.DisplayName

}


