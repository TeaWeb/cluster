package manager

import (
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/utils/string"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"source/github.com/TeaWeb/cluster-code/configs"
	"testing"
)

func TestSyncAction_Execute(t *testing.T) {
	cluster := configs.NewCluster()
	cluster.Id = "123456"
	cluster.AddNode(&configs.NodeConfig{
		Id:   "abc",
		IsOn: true,
		Role: configs.NodeRoleMaster,
	})
	err := cluster.Save()
	if err != nil {
		t.Fatal(err)
	}
	defer cluster.Delete()

	itemsDB, _ = leveldb.OpenFile(Tea.Root+"/data/items.db", nil)

	action := &PushAction{}
	action.AddItem(&configs.Item{
		Id:    "123",
		Sum:   stringutil.Md5("123456"),
		Flags: []int{0777},
		Data:  []byte("Hello, World"),
	})
	action.AddItem(&configs.Item{
		Id:    "100",
		Sum:   stringutil.Md5("123456123"),
		Flags: []int{0777},
		Data:  []byte("Hello, World123"),
	})
	action.AddItem(&configs.Item{
		Id:    "101",
		Sum:   stringutil.Md5("123456123"),
		Flags: []int{0777},
		Data:  []byte("Hello, World123"),
	})
	action.AddItem(&configs.Item{
		Id:    "300",
		Sum:   stringutil.Md5("123456123"),
		Flags: []int{0777},
		Data:  []byte("Hello, World123"),
	})
	err = action.Execute(&NodeConnection{
		ClusterId: "123456",
		NodeId:    "abc",
	})
	if err != nil {
		t.Fatal(err)
	}

	it := itemsDB.NewIterator(util.BytesPrefix([]byte("/cluster")), nil)
	for it.Next() {
		t.Log(string(it.Key()), string(it.Value()))
	}
	it.Release()
}
