package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/logs"
	"source/github.com/TeaWeb/cluster-code/configs"
	"source/github.com/TeaWeb/cluster-code/manager"
)

type NodeDeleteAction actions.Action

// delete node
func (this *NodeDeleteAction) RunPost(params struct {
	ClusterId string
	NodeId    string
}) {
	cluster := configs.FindCluster(params.ClusterId)
	if cluster == nil {
		this.Fail("找不到集群")
	}

	node := cluster.FindNode(params.NodeId)
	if node == nil {
		this.Fail("找不到节点")
	}

	cluster.RemoveNode(node.Id)
	err := cluster.Save()
	if err != nil {
		this.Fail("删除失败：" + err.Error())
	}

	// 关闭连接
	err = manager.SharedManager.CloseNode(node.Id)
	if err != nil {
		logs.Error(err)
	}

	this.Success()
}
