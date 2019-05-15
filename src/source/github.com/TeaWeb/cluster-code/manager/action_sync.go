package manager

import (
	"source/github.com/TeaWeb/cluster-code/configs"
)

// cluster -> node
type SyncAction struct {
	Action

	ItemActions []*configs.ItemAction
}

func (this *SyncAction) Name() string {
	return "sync"
}

func (this *SyncAction) TypeId() int8 {
	return 8
}
