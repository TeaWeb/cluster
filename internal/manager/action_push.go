package manager

import (
	"errors"
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/iwind/TeaGo/logs"
	"time"
)

// master -> cluster
type PushAction struct {
	Action

	Items []*configs.Item
}

func (this *PushAction) Name() string {
	return "push"
}

func (this *PushAction) Execute(nodeConn *NodeConnection) error {
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

	// current version
	version, err := SharedItemManager.FindClusterVersion(cluster.Id)
	if err != nil {
		return err
	}

	// TODO archive current version, so we can rollback later

	err = SharedItemManager.WriteClusterItems(cluster.Id, this.Items)
	if err != nil {
		return err
	}

	// /cluster/${clusterId}/info/*
	err = SharedItemManager.UpdateClusterTime(cluster.Id)
	if err != nil {
		logs.Error(err)
	}

	err = SharedItemManager.UpdateClusterVersion(cluster.Id, version+1)
	if err != nil {
		logs.Error(err)
	}

	// log
	err = SharedLogManager.Write(nodeConn.ClusterId, nodeConn.NodeId, &configs.NodeLog{
		Timestamp: time.Now().Unix(),
		Action:    "push",
	})
	if err != nil {
		logs.Error(err)
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

		notifyAction := &NotifyAction{}
		err := conn.Write(notifyAction)
		if err != nil {
			logs.Error(err)
		}
	}

	nodeConn.Reply(this, &SuccessAction{
		Message: "push ok",
	})

	return nil
}

func (this *PushAction) AddItem(item *configs.Item) {
	this.Items = append(this.Items, item)
}

func (this *PushAction) TypeId() int8 {
	return ActionCodePush
}
