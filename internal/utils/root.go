package utils

import (
	"github.com/iwind/TeaGo/Tea"
	"os"
	"path/filepath"
)

func InitRoot() {
	webIsSet := false
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
	} else {
		pwd, ok := os.LookupEnv("PWD")
		if ok {
			pwd = lookupModRoot(pwd)
			webIsSet = true
			Tea.SetPublicDir(pwd + Tea.DS + "web" + Tea.DS + "public")
			Tea.SetViewsDir(pwd + Tea.DS + "web" + Tea.DS + "views")
			Tea.SetTmpDir(pwd + Tea.DS + "web" + Tea.DS + "tmp")

			Tea.Root = pwd + Tea.DS + "build"
		}
	}

	if !webIsSet {
		Tea.SetPublicDir(Tea.Root + Tea.DS + "web" + Tea.DS + "public")
		Tea.SetViewsDir(Tea.Root + Tea.DS + "web" + Tea.DS + "views")
		Tea.SetTmpDir(Tea.Root + Tea.DS + "web" + Tea.DS + "tmp")
	}
	Tea.SetConfigDir(Tea.Root + Tea.DS + "configs")

	// Rlimit
	_ = SetRLimit(100 * 4096)
}

func lookupModRoot(pwd string) string {
	f := pwd + Tea.DS + "go.mod"
	_, err := os.Stat(f)
	if err != nil {
		parent := filepath.Dir(pwd)
		if len(parent) == 0 || parent == pwd {
			return pwd
		}
		return lookupModRoot(parent)
	}

	return pwd
}
