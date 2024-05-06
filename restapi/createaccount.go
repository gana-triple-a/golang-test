package restapi

import (
	"net/http"

	"github.com/go-chi/render"

	"example.com/testgolang/domain"
)

func CreateAccountHandler(mgr *domain.Mgr) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		reqData := &domain.CreateAccountPayload{}
		if err := render.Bind(req, reqData); err != nil {
			return
		}

		mgr.CreateAccount(reqData)

		render.Status(req, http.StatusCreated)
	}
}
