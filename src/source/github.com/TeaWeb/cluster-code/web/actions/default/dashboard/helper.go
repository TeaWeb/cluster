package dashboard

import "github.com/iwind/TeaGo/actions"

type Helper struct {
}

func (this *Helper) BeforeAction(action *actions.ActionObject) {
	action.Data["teaMenu"] = "dashboard"
}
