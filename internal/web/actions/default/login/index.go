package login

import (
	"github.com/TeaWeb/cluster/internal/configs"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction actions.Action

func (this *IndexAction) RunGet(params struct{}) {
	this.Show()
}

// 提交登录
func (this *IndexAction) RunPost(params struct {
	Username string
	Password string
	Must     *actions.Must
}) {
	params.Must.
		Field("username", params.Username).
		Require("请输入用户名").
		Field("password", params.Password).
		Require("请输入密码")

	adminConfig := configs.SharedAdminConfig()
	user := adminConfig.FindActiveUser(params.Username)
	if user != nil {
		// 错误次数
		if user.CountLoginTries() >= 3 {
			this.Fail("登录失败已超过3次，系统被锁定，需要重启服务后才能继续")
		}

		// 密码错误
		if user.Password != params.Password {
			user.IncreaseLoginTries()
			this.Fail("登录失败，请检查用户名密码")
		}

		user.ResetLoginTries()

		// Session
		this.Session().Write("username", user.Username)

		this.Next("/", nil, "").Success()
		return
	}

	this.Fail("登录失败，请检查用户名密码")
}
