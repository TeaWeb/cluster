package configs

import (
	"github.com/iwind/TeaGo/logs"
	"testing"
	"time"
)

func TestNodeLog_Marshal(t *testing.T) {
	nodeLog := &NodeLog{
		Timestamp:   time.Now().Unix(),
		Action:      "add",
		Object:      "abc.txt",
		Description: "Add abc.txt",
	}
	data := nodeLog.Marshal()
	t.Log(string(data))

	logs.PrintAsJSON(UnmarshalNodeLog(data), t)
}
