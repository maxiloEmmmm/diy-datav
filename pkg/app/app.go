package app

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"log"
)

var (
	Db *model.Client
)

func init() {
	db, err := model.Open("sqlite3", "file:./ent.db?cache=shared&mode=rwc&_fk=1", model.Debug())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	err = db.Schema.Create(
		context.Background(),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	Db = db
}
