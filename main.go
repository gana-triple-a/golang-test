package main

import (
	"fmt"
	"net/http"

	"example.com/testgolang/domain"
	"example.com/testgolang/restapi"
)

//nolint:gochecknoglobals // only allowed global vars - filled at build time - do not change
var (
	BuildTime  = "dev"
	CommitHash = "dev"
)

func main() {
	mgr, err := domain.NewManager()
	if err != nil {
		fmt.Errorf("Error initializing manager")
	}

	http.HandleFunc("/accounts", restapi.CreateAccountHandler(mgr))
	http.HandleFunc("/accounts/{accountId}", restapi.GetAccountHandler(mgr))
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
