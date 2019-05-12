package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/configs"
	"source/github.com/TeaWeb/cluster-code/web/actions/default/clusters/clusterutils"
)

type UpdateAction actions.Action

// update cluster
func (this *UpdateAction) RunGet(params struct {
	ClusterId string
}) {
	clusterutils.ClusterMenu(this)

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	ClusterId string
	Name      string
	Must      *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")

	cluster := configs.FindCluster(params.ClusterId)
	if cluster == nil {
		this.Fail("找不到集群")
	}

	cluster.Name = params.Name
	err := cluster.Save()
	if err != nil {
		this.Fail("保存失败：" + err.Error())
	}

	this.Success()
}
