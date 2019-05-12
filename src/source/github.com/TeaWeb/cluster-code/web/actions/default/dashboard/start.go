package dashboard

import (
	"github.com/iwind/TeaGo/actions"
	"source/github.com/TeaWeb/cluster-code/manager"
)

type StartAction actions.Action

// start the manager
func (this *StartAction) RunPost(params struct{}) {
	if manager.SharedManager.IsListening() {
		this.Success()
	}

	manager.SharedManager.StartInBackground()

	this.Success()
}
