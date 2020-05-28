package manager

// node -> cluster
type PingAction struct {
	Action

	Version int64
}

func (this *PingAction) Name() string {
	return "ping"
}

func (this *PingAction) Execute(nodeConn *NodeConnection) error {
	if !nodeConn.IsAuthenticated() {
		nodeConn.Reply(this, &FailAction{
			Message: "client was not authenticated",
		})
		return nil
	}
	return nil
}

func (this *PingAction) TypeId() int8 {
	return 7
}
