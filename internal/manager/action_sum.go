package manager

import "github.com/iwind/TeaGo/maps"

// master|node -> cluster
type SumAction struct {
	Action
}

func (this *SumAction) Name() string {
	return "sum"
}

func (this *SumAction) Execute(nodeConn *NodeConnection) error {
	if !nodeConn.IsAuthenticated() {
		return nodeConn.Reply(this, &FailAction{
			Message: "client was not authenticated",
		})
	}

	sum, err := SharedItemManager.FindClusterSum(nodeConn.ClusterId)
	if err != nil {
		return err
	}

	return nodeConn.Reply(this, &SuccessAction{
		Message: "ok",
		Data: maps.Map{
			"sum": sum,
		},
	})
}

func (this *SumAction) TypeId() int8 {
	return 9
}
