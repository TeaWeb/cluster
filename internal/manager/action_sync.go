package manager

import (
	"github.com/TeaWeb/cluster/internal/configs"
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
