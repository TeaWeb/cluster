package main

import (
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/sessions"
	"source/github.com/TeaWeb/cluster-code/manager"
	"source/github.com/TeaWeb/cluster-code/shell"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/clusters"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/dashboard"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/index"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/login"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/logout"
)

func main() {
	// set web root
	sh := new(shell.Shell)
	sh.Start()
	if sh.ShouldStop {
		return
	}

	// start manager
	manager.SharedManager.StartInBackground()

	// start server
	server := TeaGo.NewServer(false).
		AccessLog(false).
		Session(sessions.NewFileSessionManager(
			86400,
			"c9f5ee602110028c8b7c9aa10af0b000",
		))
	server.Start()
}
