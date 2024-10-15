package common

import (
	"LiveData/config"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 全局CasbinEnforcer
var CasbinEnforcer *casbin.Enforcer

// 初始化casbin策略管理器
func InitCasbinEnforcer() {
	e, err := mysqlCasbin()
	if err != nil {
		Log.Panicf("初始化Casbin失败：%v", err)
		panic(fmt.Sprintf("初始化Casbin失败：%v", err))
	}

	CasbinEnforcer = e
	Log.Info("初始化Casbin完成!")
}

func mysqlCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(DB)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	e, err := casbin.NewEnforcer(config.Conf.Casbin.ModelPath, a)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	Log.Debug(e)
	err = e.LoadPolicy()
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	return e, nil
}
