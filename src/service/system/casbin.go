package system

import (
	"Miyo_Assignment/global"
	model "Miyo_Assignment/model/User"
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strconv"
	"sync"
)

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type CasbinService struct{}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Sync.Once的作用是实现单例模式，速度比Init()好一些

func (CasbinService *CasbinService) AddRole(user *model.User, role string) {
	e := CasbinService.Casbin()
	e.AddRoleForUser(strconv.FormatInt(user.UID, 10), role)
	// 10 是10进制 这个函数还是需要记得的
	e.SavePolicy()
}
func (CasbinService *CasbinService) DeleteRole(user *model.User, role string) {
	e := CasbinService.Casbin()
	e.DeleteRoleForUser(strconv.FormatInt(user.UID, 10), role)
	e.SavePolicy()
}

func (CasbinService *CasbinService) DeletePermissionsForRole(role string, permissions []CasbinInfo) error {
	rules := [][]string{}
	for _, v := range permissions {
		rules = append(rules, []string{role, v.Path, v.Method})
	}
	e := CasbinService.Casbin()

	if ok, _ := e.RemovePolicies(rules); !ok {
		global.GV_LOG.Errorf("Rules don't exitest")
		return errors.New("please check possessing permissions! ")
	}
	e.SavePolicy()
	return nil
}
func (CasbinService *CasbinService) AddPermissionsForRole(role string, newPermissions []CasbinInfo) error {
	rules := [][]string{}
	for _, v := range newPermissions {
		rules = append(rules, []string{role, v.Path, v.Method})
	}
	e := CasbinService.Casbin()

	if ok, _ := e.AddPolicies(rules); !ok {
		global.GV_LOG.Errorf("Rules exitested")
		return errors.New("please check possessing permissions! ")
	}
	e.SavePolicy()
	return nil
}

func (casbinService *CasbinService) AddDefaultRole(user *model.User) {
	casbinService.AddRole(user, "level1")
}

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GV_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.GV_Config.Casbin, a)
		//为了方便修改，我把配置加到了global配置里，目前采用的是 "examples/rbac_model.conf"这个默认的配置，可以去Viper.yaml里修改
	})
	//Adapter采用的是单例，但是每次调用这个函数都会从数据库里加载
	syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
