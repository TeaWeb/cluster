package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"source/github.com/TeaWeb/cluster-code/configs"
	"source/github.com/TeaWeb/cluster-code/manager"
	"source/github.com/TeaWeb/cluster-code/web/actions/default/clusters/clusterutils"
)

type DetailAction actions.Action

// cluster detail
func (this *DetailAction) RunGet(params struct {
	ClusterId string
}) {
	cluster := clusterutils.ClusterMenu(this)

	this.Data["secret"] = cluster.Secret
	this.Data["nodes"] = lists.Map(cluster.Nodes, func(k int, v interface{}) interface{} {
		node := v.(*configs.NodeConfig)

		state, _ := manager.SharedManager.FindNodeState(node.Id)

		return maps.Map{
			"id":       node.Id,
			"name":     node.Name,
			"addr":     node.Addr,
			"isActive": state.IsActive,
			"role":     node.Role,
		}
	})

	this.Show()
}
