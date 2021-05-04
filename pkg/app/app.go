package app

import (
	"database/sql"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
)

var (
	Db *model.Client
	DbDrive *sql.DB
)
