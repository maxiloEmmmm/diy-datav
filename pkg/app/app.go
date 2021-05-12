package app

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"log"
)

var (
	Db *model.Client
)

func init() {
	db, err := model.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	Db = db
}
