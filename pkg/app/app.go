package app

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/pkg/errors"
	"log"
)

var (
	Db *model.Client
)

func WithTx(ctx context.Context, fn func(tx *model.Tx) error) error {
	tx, err := Db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

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
