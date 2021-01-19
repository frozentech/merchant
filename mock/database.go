package mock

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

// SetSQLiteDB set sql lite
func SetSQLiteDB(ctx context.Context, db *sqlx.DB) {
	b, err := ioutil.ReadFile("../create-schema.sqlite.sql")

	if err != nil {
		fmt.Println("error: cannot load ../create-schema.sqlite.sql")
		panic(err)
	}

	db.MustExecContext(ctx, string(b))
}
