package manager

import (
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"source/github.com/TeaWeb/cluster-code/configs"
	"time"
)

var SharedLogManager = NewLogManager()

type LogManager struct {
}

func NewLogManager() *LogManager {
	return &LogManager{}
}

func (this *LogManager) Write(clusterId string, nodeId string, nodeLog *configs.NodeLog) error {
	if logsDB == nil {
		return errors.New("'logsDB' should not be nil")
	}
	return logsDB.Put([]byte("/cluster/"+clusterId+"/node/"+nodeId+"/"+fmt.Sprintf("%d", time.Now().Unix())), nodeLog.Marshal(), nil)
}

func (this *LogManager) ReadAll(clusterId string, nodeId string) (result []*configs.NodeLog, err error) {
	if logsDB == nil {
		return nil, errors.New("'logsDB' should not be nil")
	}
	it := logsDB.NewIterator(util.BytesPrefix([]byte("/cluster/"+clusterId+"/node/"+nodeId+"/")), nil)
	for it.Next() {
		result = append(result, configs.UnmarshalNodeLog(it.Value()))
	}
	it.Release()
	return
}
