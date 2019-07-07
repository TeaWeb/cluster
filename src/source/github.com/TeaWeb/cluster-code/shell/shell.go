package shell

import (
	"fmt"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/files"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/types"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"source/github.com/TeaWeb/cluster-code/consts"
	"strings"
	"syscall"
)

// command line implementation
type Shell struct {
	ShouldStop bool
}

// start shell
func (this *Shell) Start() {
	// reset root
	this.resetRoot()

	// execute arguments
	if this.execArgs() {
		this.ShouldStop = true
		return
	}

	// write current pid
	files.NewFile(Tea.Root + Tea.DS + "bin" + Tea.DS + "pid").
		WriteString(fmt.Sprintf("%d", os.Getpid()))

	// log
	if !Tea.IsTesting() {
		fp, err := os.OpenFile(Tea.Root+"/logs/run.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			log.SetOutput(fp)
		} else {
			logs.Println("[error]" + err.Error())
		}
	}
}

// reset root
func (this *Shell) resetRoot() {
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

// check command line arguments
func (this *Shell) execArgs() bool {
	if len(os.Args) == 1 {
		// check process pid
		proc := this.checkPid()
		if proc != nil {
			fmt.Println("TeaWeb Cluster is already running, pid:", proc.Pid)
			return true
		}
		return false
	}
	args := os.Args[1:]
	if lists.ContainsAny(args, "?", "help", "-help", "h", "-h") {
		return this.execHelp()
	} else if lists.ContainsAny(args, "-v", "version", "-version") {
		return this.execVersion()
	} else if lists.ContainsString(args, "start") {
		return this.execStart()
	} else if lists.ContainsString(args, "stop") {
		return this.execStop()
	} else if lists.ContainsString(args, "restart") {
		return this.execRestart()
	} else if lists.ContainsString(args, "status") {
		return this.execStatus()
	}

	if len(args) > 0 {
		fmt.Println("Unknown command option '" + strings.Join(args, " ") + "', run './bin/teaweb-cluster -h' to lookup the usage.")
		return true
	}
	return false
}

// command line helps
func (this *Shell) execHelp() bool {
	fmt.Println("TeaWeb Cluster v" + consts.Version)
	fmt.Println("Usage:", "\n   ./bin/teaweb-cluster [option]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h", "\n     print this help")
	fmt.Println("  -v", "\n     print version")
	fmt.Println("  start", "\n     start the server in background")
	fmt.Println("  stop", "\n     stop the server")
	fmt.Println("  restart", "\n     restart the server")
	fmt.Println("  status", "\n     print server status")
	fmt.Println("")
	fmt.Println("To run the server in foreground:", "\n   ./bin/teaweb-cluster")

	return true
}

// version
func (this *Shell) execVersion() bool {
	fmt.Println("TeaWeb Cluster v"+consts.Version, "(build: "+runtime.Version(), runtime.GOOS, runtime.GOARCH+")")
	return true
}

// start the server
func (this *Shell) execStart() bool {
	proc := this.checkPid()
	if proc != nil {
		fmt.Println("TeaWeb Cluster already started, pid:", proc.Pid)
		return true
	}

	cmd := exec.Command(os.Args[0])
	err := cmd.Start()
	if err != nil {
		fmt.Println("TeaWeb Cluster start failed:", err.Error())
		return true
	}
	fmt.Println("TeaWeb Cluster started ok, pid:", cmd.Process.Pid)

	return true
}

// stop the server
func (this *Shell) execStop() bool {
	proc := this.checkPid()
	if proc == nil {
		fmt.Println("TeaWeb Cluster not started")
		return true
	}

	err := proc.Kill()
	if err != nil {
		fmt.Println("TeaWeb Cluster stop error:", err.Error())
		return true
	}

	files.NewFile(Tea.Root + "/bin/pid").Delete()
	fmt.Println("TeaWeb Cluster stopped ok, pid:", proc.Pid)

	return true
}

// restart the server
func (this *Shell) execRestart() bool {
	proc := this.checkPid()
	if proc != nil {
		err := proc.Kill()
		if err != nil {
			fmt.Println("TeaWeb Cluster stop error:", err.Error())
			return true
		}
	}

	cmd := exec.Command(os.Args[0])
	err := cmd.Start()
	if err != nil {
		fmt.Println("TeaWeb Cluster restart failed:", err.Error())
		return true
	}
	fmt.Println("TeaWeb Cluster restarted ok, pid:", cmd.Process.Pid)

	return true
}

// server status
func (this *Shell) execStatus() bool {
	proc := this.checkPid()
	if proc == nil {
		fmt.Println("TeaWeb Cluster not started yet")
	} else {
		fmt.Println("TeaWeb Cluster is running, pid:" + fmt.Sprintf("%d", proc.Pid))
	}
	return true
}

// check process pid
func (this *Shell) checkPid() *os.Process {
	// check pid file
	pidFile := files.NewFile(Tea.Root + "/bin/pid")
	if !pidFile.Exists() {
		return nil
	}
	pidString, err := pidFile.ReadAllString()
	if err != nil {
		return nil
	}
	pid := types.Int(pidString)

	if pid <= 0 {
		return nil
	}

	// if pid equals current pid
	if pid == os.Getpid() {
		return nil
	}

	proc, err := os.FindProcess(pid)
	if err != nil || proc == nil {
		return nil
	}

	if runtime.GOOS == "windows" {
		return proc
	}

	err = proc.Signal(syscall.Signal(0))
	if err != nil {
		return nil
	}

	// ps?
	ps, err := exec.LookPath("ps")
	if err != nil {
		return proc
	}

	cmd := exec.Command(ps, "-p", pidString, "-o", "command=")
	output, err := cmd.Output()
	if err != nil {
		return proc
	}

	if len(output) == 0 {
		return nil
	}

	outputString := string(output)
	index := strings.LastIndex(outputString, "/")
	if index > -1 {
		outputString = outputString[index+1:]
	}
	index2 := strings.LastIndex(outputString, "\\")
	if index2 > 0 {
		outputString = outputString[index2+1:]
	}
	if strings.Contains(outputString, "teaweb-cluster") {
		return proc
	}

	return nil
}
