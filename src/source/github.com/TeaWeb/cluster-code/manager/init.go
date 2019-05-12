package manager

import (
	"errors"
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/files"
	"github.com/iwind/TeaGo/logs"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

func init() {
	// register actions
	RegisterActionType(
		new(SuccessAction),
		new(FailAction),
		new(RegisterAction),
		new(NotifyAction),
		new(PushAction),
		new(PullAction),
		new(PingAction),
		new(SyncAction),
	)

	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		dataDir := files.NewFile(Tea.Root + "/data")
		if !dataDir.Exists() {
			err := dataDir.Mkdir()
			if err != nil {
				logs.Error(err)
				os.Exit(0)
			}
		}

		{
			db1, err := leveldb.OpenFile(Tea.Root+"/data/items.db", nil)
			if err != nil {
				logs.Error(errors.New("[items.db]" + err.Error()))
				os.Exit(0)
			}
			itemsDB = db1

			db2, err := leveldb.OpenFile(Tea.Root+"/data/logs.db", nil)
			if err != nil {
				logs.Error(errors.New("[logs.db]" + err.Error()))
				os.Exit(0)
			}
			logsDB = db2
		}
	})
}
