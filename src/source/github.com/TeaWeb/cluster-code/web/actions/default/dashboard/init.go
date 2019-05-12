package dashboard

import (
	"github.com/iwind/TeaGo"
	"source/github.com/TeaWeb/cluster-code/web/actions/default/helpers"
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
