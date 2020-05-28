package manager

import (
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/iwind/TeaGo/logs"
	"time"
)

// node <- cluster
type PullAction struct {
	Action

	LocalItems []*configs.Item // items without data
}

func (this *PullAction) Name() string {
	return "pull"
}

func (this *PullAction) Execute(nodeConn *NodeConnection) error {
	if !nodeConn.IsAuthenticated() {
		nodeConn.Reply(this, &FailAction{
			Message: "client was not authenticated",
		})
		return nil
	}

	items, err := SharedItemManager.ReadMasterItems(nodeConn.ClusterId)
	if err != nil {
		return err
	}

	clusterItemsMap := map[string]*configs.Item{}
	for _, item := range items {
		clusterItemsMap[item.Id] = item
	}

	// local items map
	localItemsMap := map[string]*configs.Item{}
	for _, item := range this.LocalItems {
		localItemsMap[item.Id] = item
	}

	itemActions := []*configs.ItemAction{}

	// add | change
	for _, clusterItem := range clusterItemsMap {
		localItem, ok := localItemsMap[clusterItem.Id]
		if !ok {
			// add
			itemActions = append(itemActions, configs.NewItemAction(clusterItem.Id, configs.ItemActionAdd, clusterItem))
		} else if localItem.Sum != clusterItem.Sum {
			// change
			itemActions = append(itemActions, configs.NewItemAction(clusterItem.Id, configs.ItemActionChange, clusterItem))
		}
	}

	// remove
	for _, localItem := range localItemsMap {
		_, ok := clusterItemsMap[localItem.Id]
		if !ok {
			itemActions = append(itemActions, configs.NewItemAction(localItem.Id, configs.ItemActionRemove, nil))
		}
	}

	// log
	err = SharedLogManager.Write(nodeConn.ClusterId, nodeConn.NodeId, &configs.NodeLog{
		Timestamp: time.Now().Unix(),
		Action:    "pull",
	})
	if err != nil {
		logs.Error(err)
	}

	return nodeConn.Reply(this, &SyncAction{
		ItemActions: itemActions,
	})
}

func (this *PullAction) TypeId() int8 {
	return 6
}
