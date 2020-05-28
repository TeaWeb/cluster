package manager

import (
	"time"
)

type Action struct {
	Id          uint64
	RequestTime time.Time
	RequestId   uint64
}

func (this *Action) Execute(nodeConn *NodeConnection) error {
	return nil
}

func (this *Action) OnSuccess() error {
	return nil
}

func (this *Action) OnFail() error {
	return nil
}

func (this *Action) BaseAction() *Action {
	return this
}
