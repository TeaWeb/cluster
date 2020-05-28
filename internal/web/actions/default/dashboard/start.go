package dashboard

import (
	"github.com/TeaWeb/cluster/internal/manager"
	"github.com/iwind/TeaGo/actions"
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
