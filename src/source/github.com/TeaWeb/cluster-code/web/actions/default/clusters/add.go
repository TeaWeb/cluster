package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/configs"
)

type AddAction actions.Action

func (this *AddAction) RunGet(params struct{}) {
	this.Show()
}

func (this *AddAction) RunPost(params struct {
	Name string
	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入名称")

	cluster := configs.NewCluster()
	cluster.Name = params.Name
	err := cluster.Save()
	if err != nil {
		this.Fail("保存失败：" + err.Error())
	}

	clusterList := configs.SharedClusterList()
	clusterList.AddCluster(cluster.Id)
	err = clusterList.Save()
	if err != nil {
		cluster.Delete()

		this.Fail("保存失败：" + err.Error())
	}

	this.Success()
}
