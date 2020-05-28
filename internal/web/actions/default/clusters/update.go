package clusters

import (
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/TeaWeb/cluster/internal/web/actions/default/clusters/clusterutils"
	"github.com/iwind/TeaGo/actions"
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
