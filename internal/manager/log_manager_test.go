package manager

import (
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"source/github.com/TeaWeb/cluster-code/configs"
	"testing"
	"time"
)

func TestLogManager_Write(t *testing.T) {
	logsDB1, err := leveldb.OpenFile(Tea.Root+"/data/test.db", nil)
	if err != nil {
		t.Fatal(err)
	}
	logsDB = logsDB1
	err = SharedLogManager.Write("123456", "123456", &configs.NodeLog{
		Timestamp:   time.Now().Unix(),
		Action:      "test",
		Object:      "hello",
		Description: "this is test log",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLogManager_ReadAll(t *testing.T) {
	logsDB1, err := leveldb.OpenFile(Tea.Root+"/data/test.db", nil)
	if err != nil {
		t.Fatal(err)
	}
	logsDB = logsDB1
	result, err := SharedLogManager.ReadAll("123456", "123456")
	if err != nil {
		t.Fatal(err)
	}
	logs.PrintAsJSON(result, t)
}
