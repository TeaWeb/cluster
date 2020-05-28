package manager

import (
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/iwind/TeaGo/logs"
	"strings"
	"time"
)

type RegisterAction struct {
	Action

	ClusterId     string
	ClusterSecret string
	NodeId        string
	NodeName      string
	NodeRole      string
}

func (this *RegisterAction) Name() string {
	return "register"
}

func (this *RegisterAction) Execute(nodeConn *NodeConnection) error {
	// check cluster id
	cluster := configs.FindCluster(this.ClusterId)
	if cluster == nil {
		return nodeConn.Reply(this, &FailAction{
			Message: "fail to register node, cluster is not found",
		})
	}

	// check cluster secret
	if cluster.Secret != this.ClusterSecret {
		return nodeConn.Reply(this, &FailAction{
			Message: "fail to register node, cluster secret is incorrect",
		})
	}

	// check node id
	if len(this.NodeId) == 0 {
		return nodeConn.Reply(this, &FailAction{
			Message: "fail to register node, node id is empty",
		})
	}

	// check role
	if !configs.ExistNodeRole(this.NodeRole) {
		return nodeConn.Reply(this, &FailAction{
			Message: "fail to register node, node role is empty",
		})
	}

	// should be one master only
	if this.NodeRole == configs.NodeRoleMaster {
		masterNode := cluster.FindMasterNode()
		if masterNode != nil && masterNode.Id != this.NodeId {
			return nodeConn.Reply(this, &FailAction{
				Message: "fail to register node, there can be only one master node",
			})
		}
	}

	// if not registered, save to config file
	node := cluster.FindNode(this.NodeId)

	remoteAddr := nodeConn.Conn.RemoteAddr().String()
	portIndex := strings.LastIndex(remoteAddr, ":")
	if portIndex > -1 {
		remoteAddr = remoteAddr[:portIndex]
	}

	if node == nil {
		node = configs.NewNode()
		node.CreatedAt = time.Now().Unix()
		node.IsOn = true
		node.Name = this.NodeName
		node.Id = this.NodeId
		node.Addr = remoteAddr
		node.Role = this.NodeRole

		cluster.AddNode(node)
	} else {
		node.Name = this.NodeName
		node.Addr = remoteAddr
		node.Role = this.NodeRole
	}
	err := cluster.Save()
	if err != nil {
		logs.Error(err)
		return nodeConn.Reply(this, &FailAction{
			Message: "cluster occurs error:" + err.Error(),
		})
	}

	nodeConn.ClusterId = this.ClusterId
	nodeConn.NodeId = this.NodeId

	return nodeConn.Reply(this, &SuccessAction{
		Message: "ok",
	})
}

func (this *RegisterAction) TypeId() int8 {
	return ActionCodeRegister
}
