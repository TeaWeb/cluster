package configs

type NodeState struct {
	IsActive bool
}

func NewNodeState() *NodeState {
	return &NodeState{}
}
