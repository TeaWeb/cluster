package manager

import (
	"errors"
	"fmt"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"source/github.com/TeaWeb/cluster-code/configs"
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

	if node.Role != configs.NodeRoleMaster {
		return errors.New("node '" + nodeConn.ClusterId + "/" + nodeConn.NodeId + "' not master")
	}

	if itemsDB == nil {
		return errors.New("'itemsDB' should not be nil")
	}

	// current version
	versionBytes, err := itemsDB.Get([]byte("/cluster/"+cluster.Id+"/info/version"), nil)
	if err != nil {
		if err != leveldb.ErrNotFound {
			logs.Error(err)
		}
		versionBytes = []byte("0")
	}
	version := types.Int(string(versionBytes))

	// TODO archive current version, so we can rollback later

	// remove DELETED items
	itemsMap := map[string]*configs.Item{}
	for _, item := range this.Items {
		itemsMap[item.Id] = item
	}

	it := itemsDB.NewIterator(util.BytesPrefix([]byte("/cluster/"+nodeConn.ClusterId+"/master/")), nil)
	for it.Next() {
		item := configs.UnmarshalItem(it.Value())
		if item != nil {
			_, ok := itemsMap[item.Id]
			if !ok {
				err := itemsDB.Delete(it.Key(), nil)
				if err != nil {
					logs.Error(err)
				}
			}
		}
	}
	it.Release()

	// write items
	for _, item := range this.Items {
		err := itemsDB.Put([]byte("/cluster/"+nodeConn.ClusterId+"/master/"+item.Id), item.Marshal(), nil)
		if err != nil {
			return err
		}
	}

	// /cluster/${clusterId}/info/*
	err = itemsDB.Put([]byte("/cluster/"+cluster.Id+"/info/time"), []byte(fmt.Sprintf("%d", time.Now().Unix())), nil)
	if err != nil {
		logs.Error(err)
	}

	err = itemsDB.Put([]byte("/cluster/"+cluster.Id+"/info/version"), []byte(fmt.Sprintf("%d", version+1)), nil)
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

		state, conn := SharedManager.FindNodeState(node.Id)
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
	return 5
}
