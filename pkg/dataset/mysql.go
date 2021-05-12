package dataset

import (
	"encoding/json"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/maxiloEmmmm/diy-datav/pkg/app"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/schema"
	"sync"
)

const Mysql = "dataset.mysql"

func init() {
	Reg[Mysql] = &mysql{}
}

type mysql struct {}

type mysqlConfig struct {
	Sql string
	Engine schema.TypeKey
}

func (h *mysql) Load(config string) (interface{}, error) {
	mc := &mysqlConfig{}
	err := json.Unmarshal([]byte(config), mc)
	if err != nil {
		return EmptyData{}, err
	}

	engine, err := NewMysqlEngine(mc.Engine)
	if err != nil {
		return EmptyData{}, err
	}

	s := &DataDesc{}
	err = (*engine).Query(nil, mc.Sql, nil, s)
	return s, err
}

type MysqlEngineType *sql.Driver

var mysqlEngine = make(map[schema.TypeKey]MysqlEngineType, 0)
var mysqlEngineLock = sync.Mutex{}

type mysqlOpenConfig struct {
	Host string
	Port int
	Pass string
	User string
	DB string
}

func NewMysqlEngine(id schema.TypeKey) (MysqlEngineType, error) {
	mysqlEngineLock.Lock()
	defer mysqlEngineLock.Unlock()
	if engine, ok := mysqlEngine[id]; ok {
		return engine, nil
	}

	tc, err := app.Db.TypeConfig.Get(nil, id)
	if err != nil {
		return nil, err
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