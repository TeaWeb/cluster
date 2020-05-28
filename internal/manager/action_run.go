package manager

import (
	"errors"
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
)

type RunAction struct {
	Action

	Cmd  string
	Data maps.Map
}

func (this *RunAction) Name() string {
	return "run"
}

func (this *RunAction) Execute(nodeConn *NodeConnection) error {
	if !nodeConn.IsAuthenticated() {
		nodeConn.Reply(this, &FailAction{
			Message: "client was not authenticated",
		})
		return nil
	}

	if nodeConn == nil {
		return errors.New("'nodeConn' should not be nil")
	}
	cluster := configs.FindCluster(nodeConn.ClusterId)
	if cluster == nil {
		return errors.New("cluster '" + nodeConn.ClusterId + "' not found")
	}

	node := cluster.FindNode(nodeConn.NodeId)
	if node == nil {
		return errors.New("node '" + nodeConn.ClusterId + "/" + nodeConn.NodeId + "' not found")
	}

	if !node.IsMaster() {
		return errors.New("node '" + nodeConn.ClusterId + "/" + nodeConn.NodeId + "' not master")
	}

	// notify all nodes
	for _, node := range cluster.Nodes {
		if node.IsMaster() {
			continue
		}

		state, conn := SharedManager.FindNodeState(cluster.Id, node.Id)
		if !state.IsActive {
			continue
		}

		err := conn.Write(this)
		if err != nil {
			logs.Error(err)
		}
	}

	nodeConn.Reply(this, &SuccessAction{
		Message: "run ok",
	})

	return nil
}

func (this *RunAction) TypeId() int8 {
	return ActionCodeRun
}
