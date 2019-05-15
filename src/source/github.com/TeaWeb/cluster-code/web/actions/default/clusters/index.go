package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"source/github.com/TeaWeb/cluster-code/configs"
	"source/github.com/TeaWeb/cluster-code/manager"
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
