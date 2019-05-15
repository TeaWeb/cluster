package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/manager"
)

type NodeSyncAction actions.Action

// sync node
func (this *NodeSyncAction) RunPost(params struct {
	ClusterId string
	NodeId    string
}) {
	state, conn := manager.SharedManager.FindNodeState(params.ClusterId, params.NodeId)
	if state.IsActive {
		conn.Write(&manager.NotifyAction{})
	} else {
		this.Fail("当前节点不在线，无法同步")
	}
	this.Success()
}
