package test

type testService struct {
	st ITestStore
}

//NewTestService service prototype
func NewTestService(st ITestStore) *testService {
	return &testService{
		st: st,
	}
}

func (sv *testService) testSv() error {
	return nil
}
