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
	ProceedGetNameByID(id int) (*TestTable, error)
	ProceedGetAddName(name string,sirname string) (*TestTable, error)
	ProceedAddForm(form1 *Form1) (*Form1, error)
	ProceedGetCustomerInfoByUid(uid int) (*Form1Request, error)
}

//NewTestStore return test st instance
func NewTestStore(ctx *core.BaseContext) *TestStore {
	return &TestStore{
		db: ctx.Mysql,
	}
}

//Get Name by ID
func (st *TestStore) ProceedGetNameByID(id int) (*TestTable, error) {
	MinhDB:=&TestTable{Id: id}
	err:=st.db.Where(MinhDB).Find(MinhDB).Error
	return MinhDB, err
}
//Insert Name
func (st *TestStore) ProceedGetAddName(name string,surname string) (*TestTable, error) {
	MinhDB:=&TestTable{Name: name,Surname: surname}
	err:=st.db.Create(MinhDB).Error
	return MinhDB, err
}

func (st *TestStore) ProceedAddForm(form1 *Form1) (*Form1, error) {

	err:=st.db.Create(form1).Error
	return form1, err
}

func (st *TestStore) ProceedGetCustomerInfoByUid(uid int) (*Form1Request, error) {
	data:=&Form1Request{UniqueId: uid}
	err:=st.db.Where(data).Find(data).Error
	return data, err
}