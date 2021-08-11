package test

// TestStore TestStore
type TestStore struct {

}

//ITestStore controlled interface
type ITestStore interface {
	InsertNewItem() (bool, error)
}

//NewTestStore return test st instance
func NewTestStore() *TestStore {
	return &TestStore{

	}
}

//InsertNewItem InsertNewItem store
func (st *TestStore) InsertNewItem() (bool, error) {
	return true, nil
}
