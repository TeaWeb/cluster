package dashboard

import (
	"github.com/TeaWeb/cluster/internal/manager"
	"github.com/iwind/TeaGo/actions"
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
