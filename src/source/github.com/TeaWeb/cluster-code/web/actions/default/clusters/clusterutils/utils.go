package clusterutils

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"source/github.com/TeaWeb/cluster-code/configs"
)

func ClusterMenu(actionWrapper actions.ActionWrapper) *configs.ClusterConfig {
	action := actionWrapper.Object()
	clusterId := action.ParamString("clusterId")
	if len(clusterId) == 0 {
		action.Fail("找不到集群")
	}

	cluster := configs.FindCluster(clusterId)
	if cluster == nil {
		action.Fail("找不到集群 '" + clusterId + "'")
	}

	action.Data["cluster"] = maps.Map{
		"id":   cluster.Id,
		"name": cluster.Name,
	}

	return cluster
}
