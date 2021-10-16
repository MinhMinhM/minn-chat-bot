package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"minh-chat-bot/core"
	"net/http"
	"errors"
)

type testService struct {
	st ITestStore
	ctx *core.BaseContext
}

//NewTestService service prototype
func NewTestService(ctx *core.BaseContext,st ITestStore) *testService {
	return &testService{
		st: st,
		ctx: ctx,
	}
}

func (sv *testService)replyMessageLine(Message ReplyMessage,ChannelToken string) error {
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

	//log.Println("response Status:", resp.Status)
	//log.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//log.Println("response Body:", string(body))

	return err
}

func (sv *testService)getProfile(userId string,ChannelToken string) string {

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
	//log.Println(profile.DisplayName)
	return profile.DisplayName

}

func (sv *testService)AddFormToDB(form1 *Form1Request) (*Form1Request,error) {
	result,err:=sv.st.ProceedAddForm(form1)
	return result,err
}

func (sv *testService)GetCustomerInfoByUid(uid int,agentEmail string) (*Form1Response,error) {
	data,err:=sv.st.ProceedGetCustomerInfoByUid(uid,agentEmail)
	if err.Error!=nil{
		return nil,errors.New("DB error")
	}else if err.RowsAffected==0 {
		return nil,errors.New("data not found")
	}
	return data,nil
}