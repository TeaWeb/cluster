package helpers

import (
	"github.com/TeaWeb/cluster/internal/consts"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type AuthHelper struct {
}

func (this *AuthHelper) BeforeAction(action *actions.ActionObject) (goNext bool) {
	var session = action.Session()
	var username = session.GetString("username")
	if len(username) == 0 {
		this.login(action)
		return false
	}

	action.Data["teaVersion"] = consts.Version
	action.Data["teaName"] = "TeaWeb Cluster"
	action.Data["teaUserAvatar"] = ""
	action.Data["teaUsername"] = ""
	action.Data["teaIsSuper"] = true

	// modules
	modules := []maps.Map{}
	modules = append(modules, maps.Map{
		"code":     "dashboard",
		"menuName": "服务状态",
		"icon":     "dashboard",
		"url":      "/dashboard",
	})
	modules = append(modules, maps.Map{
		"code":     "clusters",
		"menuName": "集群管理",
		"icon":     "clone",
		"url":      "/clusters",
	})

	action.Data["teaModules"] = modules
	action.Data["teaMenu"] = "dashboard"

	return true
}

func (this *AuthHelper) login(action *actions.ActionObject) {
	action.RedirectURL("/login")
}
