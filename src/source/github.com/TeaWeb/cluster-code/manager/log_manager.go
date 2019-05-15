package manager

import (
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"source/github.com/TeaWeb/cluster-code/configs"
	"time"
)

var SharedLogManager = NewLogManager()

type LogManager struct {
	db *leveldb.DB
}

func NewLogManager() *LogManager {
	return &LogManager{}
}

func (this *LogManager) SetDB(db *leveldb.DB) {
	this.db = db
}

func (this *LogManager) Write(clusterId string, nodeId string, nodeLog *configs.NodeLog) error {
	if this.db == nil {
		return errors.New("LogManager db should not be nil")
	}
	return this.db.Put([]byte("/cluster/"+clusterId+"/node/"+nodeId+"/"+fmt.Sprintf("%d", time.Now().Unix())), nodeLog.Marshal(), nil)
}

func (this *LogManager) ReadAll(clusterId string, nodeId string) (result []*configs.NodeLog, err error) {
	if this.db == nil {
		return nil, errors.New("'logsDB' should not be nil")
	}
	it := this.db.NewIterator(util.BytesPrefix([]byte("/cluster/"+clusterId+"/node/"+nodeId+"/")), nil)
	for it.Next() {
		result = append(result, configs.UnmarshalNodeLog(it.Value()))
	}
	it.Release()
	return
}
