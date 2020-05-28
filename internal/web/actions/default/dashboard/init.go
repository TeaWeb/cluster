package dashboard

import (
	"github.com/TeaWeb/cluster/internal/web/actions/default/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(new(helpers.AuthHelper)).
			Helper(new(Helper)).
			Prefix("/dashboard").
			Get("", new(IndexAction)).
			Post("/start", new(StartAction)).
			Post("/stop", new(StopAction)).
			Post("/restart", new(RestartAction)).
			EndAll()
	})
}
