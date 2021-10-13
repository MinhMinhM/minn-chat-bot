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
}

//NewTestStore return test st instance
func NewTestStore(ctx *core.BaseContext) *TestStore {
	return &TestStore{
		db: ctx.Mysql,
	}
}

//InsertNewItem InsertNewItem store
func (st *TestStore) ProceedgetNameByID(id int) (*TestTable, error) {
	MinhDB:=&TestTable{Id: id}
	err:=st.db.Where(MinhDB).Find(MinhDB).Error
	return MinhDB, err
}
