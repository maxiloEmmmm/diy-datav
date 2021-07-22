package dataset

import (
	"context"
	sql2 "database/sql"
	"database/sql/driver"
	"encoding/json"
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"reflect"
	"sync"
)

const Sql = "dataset.sql"

func init() {
	Reg[Sql] = &sql{}
}

type sql struct{}

type mysqlConfig struct {
	Sql    string
	Engine int
}

func (h *sql) Load(ctx context.Context, config string) (interface{}, error) {
	mc := &mysqlConfig{}
	err := json.Unmarshal([]byte(config), mc)
	if err != nil {
		return EmptyData{}, err
	}

	engine, err := NewEngine(ctx, mc.Engine)
	if err != nil {
		return EmptyData{}, err
	}

	rows := &entSql.Rows{}
	if err = (*engine).Query(ctx, mc.Sql, []interface{}{}, rows); err != nil {
		return EmptyData{}, err
	}

	result := EmptyData{}
	cts, err := rows.ColumnTypes()
	if err != nil {
		return EmptyData{}, err
	}

	columns, _ := rows.Columns()
	for rows.Next() {
		values := make([]interface{}, len(columns))

		for idx, columnType := range cts {
			if columnType.ScanType() != nil {
				values[idx] = reflect.New(reflect.PtrTo(columnType.ScanType())).Interface()
			} else {
				values[idx] = new(interface{})
			}
		}

		err = rows.Scan(values...)
		if err != nil {
			return EmptyData{}, err
		}

		item := make(map[string]interface{}, 0)
		for idx, column := range columns {
			if reflectValue := reflect.Indirect(reflect.Indirect(reflect.ValueOf(values[idx]))); reflectValue.IsValid() {
				item[column] = reflectValue.Interface()
				if valuer, ok := item[column].(driver.Valuer); ok {
					item[column], _ = valuer.Value()
				} else if b, ok := item[column].(sql2.RawBytes); ok {
					item[column] = string(b)
				}
			} else {
				item[column] = nil
			}
		}

		dd := DataDesc(item)
		result = append(result, &dd)
	}

	return result, rows.Err()
}

type sqlEngineType *entSql.Driver

//todo: close engine by channel
var sqlEngine = make(map[int]sqlEngineType, 0)
var sqlEngineLock = sync.Mutex{}

type sqlOpenConfig struct {
	Host string
	Port int
	Pass string
	User string
	DB   string
	Type string
}

func NewEngine(ctx context.Context, id int) (sqlEngineType, error) {
	tc, err := app.Db.TypeConfig.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if tc.Type != Sql {
		return nil, errors.New("engine type not eq")
	}

	sqlEngineLock.Lock()
	defer sqlEngineLock.Unlock()
	if engine, ok := sqlEngine[id]; ok {
		return engine, nil
	}

	moc := &sqlOpenConfig{}
	err = json.Unmarshal([]byte(tc.Config), moc)
	if err != nil {
		return nil, err
	}

	switch moc.Type {
	case dialect.MySQL, dialect.SQLite, dialect.Postgres, dialect.Gremlin:
	default:
		return nil, fmt.Errorf("unknown sql db type(%s)", moc.Type)
	}

	drv, err := entSql.Open(moc.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", moc.User, moc.Pass, moc.Host, moc.Port, moc.DB))
	if err != nil {
		return nil, err
	}

	sqlEngine[id] = drv

	return drv, nil
}
