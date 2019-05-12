package dashboard

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/manager"
)

type RestartAction actions.Action

// restart manager
func (this *RestartAction) RunPost(params struct{}) {
	err := manager.SharedManager.Stop()
	if err != nil {
		this.Fail("ERROR:" + err.Error())
	}

	manager.SharedManager.StartInBackground()

	this.Success()
}
