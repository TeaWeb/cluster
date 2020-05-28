package logout

import "github.com/iwind/TeaGo/actions"

type IndexAction actions.Action

func (this *IndexAction) RunGet(params struct{}) {
	this.Session().Delete()
	this.RedirectURL("/")
}
