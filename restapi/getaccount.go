package restapi

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"

	"example.com/testgolang/domain"
)

type GetAccountResponse domain.Account

func (gpResp *GetAccountResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func GetAccountHandler(mgr *domain.Mgr) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		// get account_id from url
		accountId, err := strconv.Atoi(chi.URLParam(req, "accountId"))
		if err != nil {
			return
		}

		account := mgr.GetAccount(req.Context(), accountId)

		gpr := GetAccountResponse(*account)

		render.Status(req, http.StatusOK)

		if errR := render.Render(resp, req, &gpr); errR != nil {
			oplog := httplog.LogEntry(req.Context())
			oplog.Error("render error", "errRender", errR)
		}
	}
}
