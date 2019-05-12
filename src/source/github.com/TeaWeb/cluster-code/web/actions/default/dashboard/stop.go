package dashboard

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/manager"
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
