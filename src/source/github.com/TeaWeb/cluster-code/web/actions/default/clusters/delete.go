package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/configs"
)

type DeleteAction actions.Action

// delete cluster
func (this *DeleteAction) RunPost(params struct {
	ClusterId string
}) {
	if len(params.ClusterId) == 0 {
		this.Fail("请输入集群ID")
	}

	cluster := configs.FindCluster(params.ClusterId)
	if cluster == nil {
		this.Fail("找不到集群")
	}

	clusterList := configs.SharedClusterList()
	clusterList.RemoveCluster(cluster.Id)
	err := clusterList.Save()
	if err != nil {
		this.Fail("删除失败：" + err.Error())
	}

	err = cluster.Delete()
	if err != nil {
		clusterList.AddCluster(cluster.Id)
		clusterList.Save()

		this.Fail("删除失败：" + err.Error())
	}

	this.Success()
}
