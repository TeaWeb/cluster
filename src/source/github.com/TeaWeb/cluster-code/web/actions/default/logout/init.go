package logout

import "github.com/iwind/TeaGo"

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Prefix("/logout").
			GetPost("", new(IndexAction)).
			EndAll()
	})
}
