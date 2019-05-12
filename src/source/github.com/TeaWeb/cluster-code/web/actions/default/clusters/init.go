package clusters

import (
	"github.com/iwind/TeaGo"
	"source/github.com/TeaWeb/cluster-code/web/actions/default/helpers"
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
			EndAll()
	})
}
