package clusters

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/utils/time"
	"source/github.com/TeaWeb/cluster-code/configs"
	"source/github.com/TeaWeb/cluster-code/manager"
	"source/github.com/TeaWeb/cluster-code/web/actions/default/clusters/clusterutils"
)

type NodeDetailAction actions.Action

// node detail
func (this *NodeDetailAction) RunGet(params struct {
	ClusterId string
	NodeId    string
}) {
	cluster := clusterutils.ClusterMenu(this)

	node := cluster.FindNode(params.NodeId)
	if node == nil {
		this.Fail("找不到节点")
	}

	state, _ := manager.SharedManager.FindNodeState(cluster.Id, node.Id)
	this.Data["node"] = maps.Map{
		"id":       node.Id,
		"name":     node.Name,
		"addr":     node.Addr,
		"isOn":     node.IsOn,
		"isActive": state.IsActive,
		"role":     node.Role,
		"isMaster": node.IsMaster(),
	}

	// items
	items, _ := manager.SharedItemManager.ReadMasterItems(params.ClusterId)
	this.Data["items"] = lists.Map(items, func(k int, v interface{}) interface{} {
		item := v.(*configs.Item)
		return maps.Map{
			"id":   item.Id,
			"size": len(item.Data),
		}
	})

	// push | pull logs
	this.Data["syncTime"] = ""
	this.Data["logs"] = []maps.Map{}
	result, err := manager.SharedLogManager.ReadAll(cluster.Id, node.Id)
	if err != nil {
		logs.Error(err)
	} else {
		if len(result) > 0 {
			lists.Reverse(result)
			this.Data["syncTime"] = timeutil.Format("Y-m-d H:i:s", result[0].Time())

			if len(result) > 5 {
				result = result[:5]
			}

			this.Data["logs"] = lists.Map(result, func(k int, v interface{}) interface{} {
				nodeLog := v.(*configs.NodeLog)
				return maps.Map{
					"datetime":    timeutil.Format("Y-m-d H:i:s", nodeLog.Time()),
					"action":      nodeLog.Action,
					"object":      nodeLog.Object,
					"description": nodeLog.Description,
				}
			})
		}
	}

	this.Show()
}
