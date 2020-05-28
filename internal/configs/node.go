package configs

type NodeConfig struct {
	Id         string      `yaml:"id"`
	IsOn       bool        `yaml:"isOn"`
	Name       string      `yaml:"name"`
	Namespaces []Namespace `yaml:"namespaces"`
	CreatedAt  int64       `yaml:"createdAt"`
	Role       NodeRole    `yaml:"role" json:"role"`
	Addr       string

	state *NodeState
}

func NewNode() *NodeConfig {
	return &NodeConfig{}
}

func (this *NodeConfig) SetState(state *NodeState) {
	this.state = state
}

func (this *NodeConfig) IsMaster() bool {
	return this.Role == NodeRoleMaster
}
