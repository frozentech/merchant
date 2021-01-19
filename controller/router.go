package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router *mux.Router
}

// NewApp ...
func NewApp() *App {
	return &App{
		Router: mux.NewRouter().StrictSlash(true),
	}
}

// SetUpRouter ...
func (a *App) SetUpRouter() {

	a.Router.HandleFunc("/assets/{filename}", Download)
	a.Router.HandleFunc("/merchant", Merchants)
	a.Router.HandleFunc("/merchant/{merchantId}", Merchant)
	a.Router.HandleFunc("/merchant/{merchantId}/upload", Upload)
	a.Router.HandleFunc("/merchant/{merchantId}/member", Members)
	a.Router.HandleFunc("/merchant/{merchantId}/member/{memberId}", Member)
}

// Run ...
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
