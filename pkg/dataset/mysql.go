package dataset

import (
	"context"
	sql2 "database/sql"
	"database/sql/driver"
	"encoding/json"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"reflect"
	"sync"
)

const Mysql = "dataset.mysql"

func init() {
	Reg[Mysql] = &mysql{}
}

type mysql struct {}

type mysqlConfig struct {
	Sql string
	Engine int
}

func (h *mysql) Load(ctx context.Context, config string) (interface{}, error) {
	mc := &mysqlConfig{}
	err := json.Unmarshal([]byte(config), mc)
	if err != nil {
		return EmptyData{}, err
	}

	engine, err := NewMysqlEngine(ctx, mc.Engine)
	if err != nil {
		return EmptyData{}, err
	}

	rows := &sql.Rows{}
	if err = (*engine).Query(ctx, mc.Sql, []interface{}{}, rows); err != nil {
		return EmptyData{}, err
	}

	result := make([]*DataDesc, 0)
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

type MysqlEngineType *sql.Driver

//todo: close engine by channel
var mysqlEngine = make(map[int]MysqlEngineType, 0)
var mysqlEngineLock = sync.Mutex{}

type mysqlOpenConfig struct {
	Host string
	Port int
	Pass string
	User string
	DB string
}

func NewMysqlEngine(ctx context.Context, id int) (MysqlEngineType, error) {
	tc, err := app.Db.TypeConfig.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if tc.Type != Mysql {
		return nil, errors.New("engine type not eq")
	}

	mysqlEngineLock.Lock()
	defer mysqlEngineLock.Unlock()
	if engine, ok := mysqlEngine[id]; ok {
		return engine, nil
	}

	moc := &mysqlOpenConfig{}
	err = json.Unmarshal([]byte(tc.Config), moc)
	if err != nil {
		return nil, err
	}

	drv, err := sql.Open(dialect.MySQL, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", moc.User, moc.Pass, moc.Host, moc.Port, moc.DB))
	if err != nil {
		return nil, err
	}

	mysqlEngine[id] = drv

	return drv, nil
}