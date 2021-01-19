package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Download ...
func Download(w http.ResponseWriter, r *http.Request) {
	var (
		filename = mux.Vars(r)["filename"]
	)

	http.ServeFile(w, r, "/tmp/"+filename)
}
