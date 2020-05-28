package manager

import (
	"errors"
	"fmt"
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/iwind/TeaGo/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"strings"
	"sync"
	"time"
)

var SharedItemManager = NewItemManager()

type ItemManager struct {
	db     *leveldb.DB
	locker sync.Mutex
}

func NewItemManager() *ItemManager {
	return &ItemManager{}
}

func (this *ItemManager) SetDB(db *leveldb.DB) {
	this.db = db
}

func (this *ItemManager) ReadMasterItems(clusterId string) (items []*configs.Item, err error) {
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.db == nil {
		err = errors.New("ItemManager db should not be nil")
		return
	}
	it := this.db.NewIterator(util.BytesPrefix([]byte("/cluster/"+clusterId+"/master/")), nil)
	for it.Next() {
		item := configs.UnmarshalItem(it.Value())
		if item != nil {
			items = append(items, item)
		}
	}
	it.Release()
	return items, nil
}

func (this *ItemManager) FindClusterVersion(clusterId string) (version int64, err error) {
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.db == nil {
		err = errors.New("ItemManager db should not be nil")
		return
	}

	versionBytes, err := this.db.Get([]byte("/cluster/"+clusterId+"/info/version"), nil)
	if err != nil {
		if err != leveldb.ErrNotFound {
			return 0, err
		}
		versionBytes = []byte("0")
		err = nil
	}
	version = types.Int64(string(versionBytes))
	return
}

func (this *ItemManager) UpdateClusterVersion(clusterId string, version int64) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.db == nil {
		return errors.New("ItemManager db should not be nil")
	}

	return this.db.Put([]byte("/cluster/"+clusterId+"/info/version"), []byte(fmt.Sprintf("%d", version)), nil)
}

func (this *ItemManager) UpdateClusterTime(clusterId string) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.db == nil {
		return errors.New("ItemManager db should not be nil")
	}

	return this.db.Put([]byte("/cluster/"+clusterId+"/info/time"), []byte(fmt.Sprintf("%d", time.Now().Unix())), nil)
}

func (this *ItemManager) WriteClusterItems(clusterId string, newItems []*configs.Item) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.db == nil {
		return errors.New("ItemManager db should not be nil")
	}

	// remove DELETED items
	itemsMap := map[string]*configs.Item{}
	for _, item := range newItems {
		itemsMap[item.Id] = item
	}

	oldItemsSum := map[string]string{} // id => sum
	{
		// id => sum
		it := this.db.NewIterator(util.BytesPrefix([]byte("/cluster/"+clusterId+"/sum/")), nil)
		for it.Next() {
			key := string(it.Key())
			id := strings.TrimPrefix(key, "/cluster/"+clusterId+"/sum/")
			_, ok := itemsMap[id]
			if ok {
				oldItemsSum[id] = string(it.Value())
				continue
			}

			err := this.db.Delete([]byte("/cluster/"+clusterId+"/master/"+id), nil)
			if err != nil {
				it.Release()
				return err
			}

			err = this.db.Delete(it.Key(), nil)
			if err != nil {
				it.Release()
				return err
			}
		}
		it.Release()
	}

	// write items
	for _, item := range newItems {
		oldItemSum, ok := oldItemsSum[item.Id]
		if ok && oldItemSum == item.Sum {
			continue
		}
		err := this.db.Put([]byte("/cluster/"+clusterId+"/master/"+item.Id), item.Marshal(), nil)
		if err != nil {
			return err
		}

		err = this.db.Put([]byte("/cluster/"+clusterId+"/sum/"+item.Id), []byte(item.Sum), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *ItemManager) FindClusterSum(clusterId string) (map[string]string, error) {
	this.locker.Lock()
	defer this.locker.Unlock()

	if this.db == nil {
		return map[string]string{}, errors.New("ItemManager db should not be nil")
	}

	// id => sum
	result := map[string]string{}
	it := this.db.NewIterator(util.BytesPrefix([]byte("/cluster/"+clusterId+"/sum/")), nil)
	for it.Next() {
		key := string(it.Key())
		id := strings.TrimPrefix(key, "/cluster/"+clusterId+"/sum/")
		result[id] = string(it.Value())
	}
	it.Release()

	return result, nil
}
