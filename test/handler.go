package test

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"minh-chat-bot/core"
	"net/http"
	"os"
	"strconv"
)

type handler struct {
	sv *testService

}

// NewHandler return +Handler
func NewHandler(ctx *core.BaseContext) *handler {
	// Inject service here!!
	store := NewTestStore(ctx)
	service := NewTestService(ctx,store)
	return &handler{
		sv: service,
	}
}

//func (h *handler) getChannelToken(shopId string) echo.MiddlewareFunc {
//		ChannelToken:=os.Getenv("ChannelToken")
//		ChannelSecret:=os.Getenv("ChannelSecret")
//		return nil
//	}

func (h *handler)WebHook(c echo.Context) error {

	//Initiate Line channel
	route:=c.Path()
	ChannelToken:=os.Getenv("ChannelToken"+route[9:])
	ChannelSecret:=os.Getenv("ChannelSecret"+route[9:])
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
		// Unmarshal JSON
		flexContainer, err := linebot.UnmarshalFlexMessageJSON(flexs)
		if err != nil {
			log.Println(err)
		}
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



	sentry.WithScope(func(scope *sentry.Scope) {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelWarning)
			scope.SetUser(sentry.User{ID: fullname})
			scope.SetTags(map[string]string{
				"format": Line.Events[0].Message.Type,
				"UserId":      fullname,
			})
		})

		event := sentry.NewEvent()
		event.Message = "Wrong Incoming Format"
		sentry.CaptureEvent(event)
	})



	//sentry.CaptureMessage("Wrong text format")
	//sentry.ConfigureScope(func(scope *sentry.Scope) {
	//	//scope.SetExtra("oristhis", "justfantasy")
	//	scope.SetTag("Wrong format:", Line.Events[0].Message.Type)
	//	scope.SetLevel(sentry.LevelWarning)
	//	scope.SetUser(sentry.User{
	//		ID: fullname,
	//	})
	//})
	return c.NoContent(http.StatusOK)
}

func (h *handler)GetNameByID(c echo.Context) error {
	id:=c.FormValue("id")
	idInt, err := strconv.Atoi(id)
	name,err:=h.sv.GetNameByIDService(idInt)

	if err!= nil{
		return c.JSON(http.StatusOK, "Get DB fail")
	}
	return c.JSON(http.StatusOK, name)
}











