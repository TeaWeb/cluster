package clusters

import (
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/TeaWeb/cluster/internal/manager"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction actions.Action

func (this *IndexAction) RunGet(params struct{}) {
	// clusters
	clusters := configs.SharedClusterList().FindAllClusters()
	this.Data["clusters"] = lists.Map(clusters, func(k int, v interface{}) interface{} {
		cluster := v.(*configs.ClusterConfig)
		return maps.Map{
			"id":               cluster.Id,
			"name":             cluster.Name,
			"countNodes":       len(cluster.Nodes),
			"countActiveNodes": manager.SharedManager.CountClusterNodes(cluster.Id),
		}
	})

	this.Show()
}
