package manager

import (
	"github.com/syndtr/goleveldb/leveldb/util"
	"source/github.com/TeaWeb/cluster-code/configs"
)

var SharedItemManager = NewItemManager()

type ItemManager struct {
}

func NewItemManager() *ItemManager {
	return &ItemManager{}
}

func (this *ItemManager) ReadMasterItems(clusterId string) (items []*configs.Item, err error) {
	it := itemsDB.NewIterator(util.BytesPrefix([]byte("/cluster/"+clusterId+"/master/")), nil)
	for it.Next() {
		item := configs.UnmarshalItem(it.Value())
		if item != nil {
			items = append(items, item)
		}
	}
	it.Release()
	return items, nil
}
