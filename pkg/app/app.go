package app

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/maxiloEmmmm/diy-datav/pkg/model"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/menu"
	"github.com/maxiloEmmmm/diy-datav/pkg/model/user"
	"github.com/maxiloEmmmm/diy-datav/pkg/permission"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"github.com/maxiloEmmmm/go-web/contact"
	"github.com/pkg/errors"
	"log"
)

var (
	Db       *model.Client
	dbDriver *sql.Driver
	Engine   *gin.Engine
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

type ConfigOption struct {
	DB         DBConfig         `yaml:"db"`
	Permission PermissionConfig `yaml:"permission"`
}

type DBConfig struct {
	Type   string `yaml:"type"`
	Source string `yaml:"source"`
}

type PermissionConfig struct {
	Admin    string `yaml:"admin"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}

var Config *ConfigOption

func Init(ctx context.Context) {
	contact.Init()
	Config = &ConfigOption{}
	contact.ConfigFile("./app.yaml", Config)

	initDB()
	policy()
	menuInit(ctx)
}

func InitPermission(ctx context.Context, engine *gin.Engine) {
	if Config.Permission.Admin == "" || Config.Permission.Password == "" || Config.Permission.Role == "" {
		log.Fatal("permission admin empty")
	}
	adminUser := Db.User.Query().Where(user.UsernameEQ(Config.Permission.Admin)).FirstX(ctx)
	if adminUser == nil {
		adminUser = Db.User.Create().SetUsername(Config.Permission.Admin).SetPassword(go_tool.Md5(Config.Permission.Password)).SetEnable(&contact.BoolField{Bool: true}).SaveX(ctx)
	}

	contact.Permission.AddGroupingPolicy(contact.DefaultRoleUser, Config.Permission.Role)
	contact.Permission.AddRoleForUser(fmt.Sprintf("%d", adminUser.ID), Config.Permission.Role)

	for _, menu := range Db.Menu.Query().AllX(ctx) {
		AddAdminPolicy(&permission.PRMenu{Menu: menu}, permission.AccessMenuAction)
	}
	for _, route := range engine.Routes() {
		// TODO: make permission pair
		AddAdminPolicy(&permission.PRRouter{route.Path}, permission.NewAction(route.Method, route.Method))
	}
	for _, view := range Db.View.Query().AllX(ctx) {
		AddAdminPolicy(&permission.PRView{view}, permission.GetInfoAction)
	}
}

func Close() {
	Db.Close()
}

func initDB() {
	drv, err := sql.Open(Config.DB.Type, Config.DB.Source)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	dbDriver = drv

	db := model.NewClient(model.Driver(drv), model.Debug())

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

func policy() {
	entAdapter, err := contact.NewEntPolicyAdapter(&contact.EntPolicyAdapterOption{
		Db:     dbDriver.DB(),
		Driver: Config.DB.Type,
		Ctx:    context.Background(),
	})
	if err != nil {
		log.Fatalf("ent.policy.adapter", err.Error())
	}
	contact.InitPermission("", entAdapter)
}

func User(ctx context.Context) string {
	auth := contact.AuthByContext(ctx)
	if auth != nil {
		return auth.Id
	}

	return ""
}

func AddAdminPolicy(resource permission.Resource, action *permission.Action) {
	contact.Permission.AddPolicy(Config.Permission.Role, resource.Key(), action.Key, "allow")
}

func RemoveAdminPolicy(resource permission.Resource, action *permission.Action) {
	contact.Permission.RemovePolicy(Config.Permission.Role, resource.Key(), action.Key, "allow")
}

type menuItem struct {
	Title string
	Href  string
}

func menuInit(ctx context.Context) {
	menus := []*menuItem{
		{Title: "可视化", Href: "/sys/view"},
		{Title: "素材", Href: "/sys/assets"},
		{Title: "数据源", Href: "/sys/tc"},
		{Title: "用户", Href: "/sys/user"},
		{Title: "菜单", Href: "/sys/menu"},
		{Title: "角色", Href: "/sys/role"},
	}

	for _, m := range menus {
		if !Db.Menu.Query().Where(menu.And(
			menu.Title(m.Title),
			menu.URL(m.Href),
		)).ExistX(ctx) {
			mObj := Db.Menu.Create().SetTitle(m.Title).SetURL(m.Href).SaveX(ctx)
			AddAdminPolicy(&permission.PRMenu{Menu: mObj}, permission.AccessMenuAction)
		}
	}
}
