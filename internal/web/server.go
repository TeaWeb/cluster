package web

import (
	"github.com/TeaWeb/cluster/internal/manager"
	"github.com/TeaWeb/cluster/internal/shell"
	_ "github.com/TeaWeb/cluster/internal/web/actions/default/clusters"
	_ "github.com/TeaWeb/cluster/internal/web/actions/default/dashboard"
	_ "github.com/TeaWeb/cluster/internal/web/actions/default/index"
	_ "github.com/TeaWeb/cluster/internal/web/actions/default/login"
	_ "github.com/TeaWeb/cluster/internal/web/actions/default/logout"
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/sessions"
)

func Start() {
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
