package test

import (
	"fmt"
	"gorm.io/gorm"
	"minh-chat-bot/core"
)

// TestStore TestStore
type TestStore struct {
	db *gorm.DB
}

//ITestStore controlled interface
type ITestStore interface {
	getNameByID() (*TestTable, error)
}

//NewTestStore return test st instance
func NewTestStore(ctx *core.BaseContext) *TestStore {
	return &TestStore{
		db: ctx.Mysql,
	}
}

//InsertNewItem InsertNewItem store
func (st *TestStore) getNameByID() (*TestTable, error) {
	MinhDB:=&TestTable{}
	//result := st.db.Debug().Table("MinhDatabase.test_tables").Where("id=1").Find(MinhDB)
	st.db.Debug().Where("id = ?", 1).Find(MinhDB)
	fmt.Println(MinhDB)
	//st.db.Table("test_tables").Where(&TestTable{id: 1}).Find(MinhDB)
	return MinhDB, nil
}
