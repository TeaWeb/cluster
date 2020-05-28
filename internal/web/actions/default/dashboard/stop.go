package dashboard

import (
	"github.com/TeaWeb/cluster/internal/manager"
	"github.com/iwind/TeaGo/actions"
)

type StopAction actions.Action

// stop the manager
func (this *StopAction) RunPost(params struct{}) {
	if !manager.SharedManager.IsListening() {
		this.Success()
	}

	err := manager.SharedManager.Stop()
	if err != nil {
		this.Fail("ERROR:" + err.Error())
	}

	this.Success()
}
