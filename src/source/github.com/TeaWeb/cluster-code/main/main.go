package main

import (
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/sessions"
	"os"
	"path/filepath"
	"source/github.com/TeaWeb/cluster-code/manager"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/clusters"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/dashboard"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/index"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/login"
	_ "source/github.com/TeaWeb/cluster-code/web/actions/default/logout"
)

func main() {
	// set web root
	setRoot()

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

func setRoot() {
	if !Tea.IsTesting() {
		exePath, err := os.Executable()
		if err != nil {
			exePath = os.Args[0]
		}
		link, err := filepath.EvalSymlinks(exePath)
		if err == nil {
			exePath = link
		}
		fullPath, err := filepath.Abs(exePath)
		if err == nil {
			Tea.UpdateRoot(filepath.Dir(filepath.Dir(fullPath)))
		}
	}
	Tea.SetPublicDir(Tea.Root + Tea.DS + "web" + Tea.DS + "public")
	Tea.SetViewsDir(Tea.Root + Tea.DS + "web" + Tea.DS + "views")
	Tea.SetTmpDir(Tea.Root + Tea.DS + "web" + Tea.DS + "tmp")
}
