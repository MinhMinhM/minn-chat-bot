package test

import (
	"gorm.io/gorm"
	"minh-chat-bot/core"
)

// TestStore TestStore
type TestStore struct {
	db *gorm.DB
}

//ITestStore controlled interface
type ITestStore interface {
	ProceedAddForm(form1 *Form1Request) (*Form1Request, error)
	ProceedGetCustomerInfoByUid(uid int,agentEmail string) (*Form1Response, *gorm.DB)
}

//NewTestStore return test st instance
func NewTestStore(ctx *core.BaseContext) *TestStore {
	return &TestStore{
		db: ctx.Mysql,
	}
}



func (st *TestStore) ProceedAddForm(form1 *Form1Request) (*Form1Request, error) {
	err:=st.db.Create(form1).Error
	return form1, err
}

func (st *TestStore) ProceedGetCustomerInfoByUid(uid int,agentEmail string) (*Form1Response, *gorm.DB) {
	data:=&Form1Response{UniqueId: uid,AgentEmail: agentEmail}
	err:=st.db.Where(data).Find(data)
	return data, err
}