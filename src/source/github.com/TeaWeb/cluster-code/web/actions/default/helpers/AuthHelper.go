package helpers

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"source/github.com/TeaWeb/cluster-code/consts"
)

type AuthHelper struct {
}

func (this *AuthHelper) BeforeAction(action *actions.ActionObject) {
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
}
