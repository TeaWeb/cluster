package dashboard

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"source/github.com/TeaWeb/cluster-code/manager"
)

type IndexAction actions.Action

func (this *IndexAction) RunGet(params struct{}) {
	m := manager.SharedManager
	this.Data["manager"] = maps.Map{
		"addr":        m.Addr(),
		"isListening": m.IsListening(),
		"error":       m.Error(),
		"countNodes":  m.CountNodes(),
	}

	this.Show()
}
