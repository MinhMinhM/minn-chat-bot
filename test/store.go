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
	ProceedgetNameByID(id int) (*TestTable, error)
	ProceedgetAddName(name string,sirname string) (*TestTable, error)
}

//NewTestStore return test st instance
func NewTestStore(ctx *core.BaseContext) *TestStore {
	return &TestStore{
		db: ctx.Mysql,
	}
}

//Get Name by ID
func (st *TestStore) ProceedgetNameByID(id int) (*TestTable, error) {
	MinhDB:=&TestTable{Id: id}
	err:=st.db.Where(MinhDB).Find(MinhDB).Error
	return MinhDB, err
}
//Insert Name
func (st *TestStore) ProceedgetAddName(name string,surname string) (*TestTable, error) {
	MinhDB:=&TestTable{Name: name,Surname: surname}
	err:=st.db.Create(MinhDB).Error
	return MinhDB, err
}