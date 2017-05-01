package models

import (
	"errors"
	"fmt"

	"github.com/FlyCynomys/gear/conf"
	"github.com/FlyCynomys/tools/log"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/FlyCynomys/gear/models/auth"
	"github.com/FlyCynomys/gear/models/group"
	"github.com/FlyCynomys/gear/models/support"
	"github.com/FlyCynomys/gear/models/todo"
	"github.com/FlyCynomys/gear/models/user"
)

func Init(conf *conf.Config) (bool, error) {
	if conf == nil {
		return false, errors.New("config is nil")
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	log.Info("start init database ", conf)
	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8", conf.DbUser, conf.DbPassword, conf.DbAddress)
	log.Info("database dsn : ", dsn)
	err := orm.RegisterDataBase("default", "mysql", dsn, 10, 10)
	if err != nil {
		return false, err
	}

	orm.RegisterModel(new(auth.Auth))
	orm.RegisterModel(new(group.Group))
	orm.RegisterModel(new(support.Location))
	orm.RegisterModel(new(support.Company))
	orm.RegisterModel(new(support.Job))
	orm.RegisterModel(new(todo.TodoPlan))
	orm.RegisterModel(new(user.User))

	orm.RunSyncdb("default", true, true)
	orm.Debug = true
	log.Info("end init database ")
	return true, nil
}
