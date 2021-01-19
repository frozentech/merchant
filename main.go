package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/frozentech/logs"
	"github.com/frozentech/merchant/controller"
	"github.com/frozentech/merchant/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func setupHandler() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/upload", controller.Upload)
	router.HandleFunc("/merchant", controller.Merchants)
	router.HandleFunc("/merchant/{merchantId}", controller.Merchant)
	router.HandleFunc("/merchant/{merchantId}/member", controller.Members)
	router.HandleFunc("/merchant/{merchantId}/member/{memberId}", controller.Member)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("ENV_PORT"), router))
}

// setupDB ...
func setupDB() {
	var (
		err error
		tdb *sqlx.DB
	)

	tdb, err = sqlx.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	model.DB = tdb

	b, err := ioutil.ReadFile("create-schema.sqlite.sql")

	if err != nil {
		fmt.Println("error: cannot load create-schema.sqlite.sql")
		panic(err)
	}

	model.DB.MustExecContext(context.Background(), string(b))
}

func main() {
	var (
		err error
	)

	controller.Log = logs.New()

	controller.Log.Print("enviroment: loading")
	err = godotenv.Load(".env")
	if err != nil {
		controller.Log.Print("enviroment: ", err.Error())
		log.Fatal("Error loading .env file")
	}
	controller.Log.Print("enviroment: loaded")

	defer func() {
		controller.Log.Dump(true)
	}()

	if os.Getenv("ENV_DB_ENGINE") == "sqlite3" {
		setupDB()
	} else {
		defer func() {
			if model.DB != nil {
				model.DB.Close()
				model.IsClose = true
			}
		}()

		if model.DB == nil || model.IsClose == true {
			controller.Log.Print("database: connecting")
			model.DB, err = model.ConnectDB()
			if err != nil {
				controller.Log.Print("database: ", err.Error())
				panic(err)
			}
			controller.Log.Print("database: connected")
		}
	}

	app := controller.NewApp()
	app.SetUpRouter()
	app.Run(os.Getenv("ENV_APP_PORT"))

}
