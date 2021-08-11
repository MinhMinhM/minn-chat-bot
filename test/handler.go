package test

import (


	"github.com/labstack/echo/v4"
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

func (h *handler) webhook(c echo.Context) error {
	a:=h.sv.testSv()
	return a
}


