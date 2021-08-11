package test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
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

func (h *handler)WebHook(c echo.Context) error {
	//Initiate Line channel
	ChannelToken:="Ys4yNEZl1H9WfRiyV9Db93+wR94d9Pk2q4AskWSEQB5bUyZCQNUngovXtk4qFd4f7HfwyZW0TAzvNisTqKa0QF54poGcsNXtDpNJQcFY+5Zc9mfy9lI+6MeRxU60aeOiSceoLlWjsi+FPc8ZEYxiVQdB04t89/1O/w1cDnyilFU="
	ChannelSecret:="6c2154dbfbd911af13de10982e2db6cf"
	bot, err := linebot.New(ChannelSecret, ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	//Get Line event
	Line := new(LineMessage)
	if err := c.Bind(Line); err != nil {
		log.Println("err")
		return c.String(http.StatusOK, "error")
	}

	fullname := h.sv.getProfile(Line.Events[0].Source.UserID,ChannelToken)
	fmt.Println("///////////////////")
	fmt.Println(fullname)
	fmt.Println(Line.Events[0].Message)
	fmt.Println("///////////////////")


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

		flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)
		bot.PushMessage(Line.Events[0].Source.UserID, msg1).Do()
		bot.ReplyMessage(Line.Events[0].ReplyToken, flexMessage).Do()

		return c.NoContent(http.StatusOK)
	}


	if Line.Events[0].Message.Type=="text" {
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
		} else{
			return c.NoContent(http.StatusOK)
		}

		message := ReplyMessage{
			ReplyToken: Line.Events[0].ReplyToken,
			Messages:   []Text{text},
		}

		//Sent message using api instead of SDK(just as an example)
		h.sv.replyMessageLine(message,ChannelToken)
		return c.NoContent(http.StatusOK)
	}

	//log.Println("%% message success")
	return c.NoContent(http.StatusOK)
}









