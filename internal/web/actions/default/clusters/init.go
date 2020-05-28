package clusters

import (
	"github.com/TeaWeb/cluster/internal/web/actions/default/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(new(helpers.AuthHelper)).
			Helper(new(Helper)).
			Prefix("/clusters").
			Get("", new(IndexAction)).
			GetPost("/add", new(AddAction)).
			Post("/delete", new(DeleteAction)).
			GetPost("/update", new(UpdateAction)).
			Get("/detail", new(DetailAction)).
			Post("/node/delete", new(NodeDeleteAction)).
			Get("/node/detail", new(NodeDetailAction)).
			Post("/node/sync", new(NodeSyncAction)).
			EndAll()
	})
}
