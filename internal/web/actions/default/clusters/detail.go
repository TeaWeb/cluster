package clusters

import (
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/TeaWeb/cluster/internal/manager"
	"github.com/TeaWeb/cluster/internal/web/actions/default/clusters/clusterutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type DetailAction actions.Action

// cluster detail
func (this *DetailAction) RunGet(params struct {
	ClusterId string
}) {
	cluster := clusterutils.ClusterMenu(this)

	this.Data["secret"] = cluster.Secret

	lists.Sort(cluster.Nodes, func(i int, j int) bool {
		node1 := cluster.Nodes[i]
		if node1.IsMaster() {
			return true
		}
		return false
	})
	this.Data["nodes"] = lists.Map(cluster.Nodes, func(k int, v interface{}) interface{} {
		node := v.(*configs.NodeConfig)

		state, _ := manager.SharedManager.FindNodeState(cluster.Id, node.Id)

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
