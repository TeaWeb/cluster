package manager

type ActionInterface interface {
	Name() string
	Execute(nodeConn *NodeConnection) error
	OnSuccess() error
	OnFail() error
	TypeId() int8
	BaseAction() *Action
}
