package controller_test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/frozentech/logs"
	"github.com/frozentech/merchant/controller"
	"github.com/frozentech/merchant/mock"
	"github.com/frozentech/merchant/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var App *controller.App

// GetTextContext ...
func GetTextContext() context.Context {
	return context.WithValue(context.Background(), "x-amzn-trace-id", "Root=fakeid; Parent=reqid; Sampled=1")
}

func TestMain(m *testing.M) {
	var (
		err error
		tdb *sqlx.DB
	)

	controller.Log = logs.New()

	tdb, err = sqlx.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	model.DB = tdb
	mock.SetSQLiteDB(GetTextContext(), model.DB)

	App = controller.NewApp()
	App.SetUpRouter()

	os.Exit(m.Run())
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	App.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
