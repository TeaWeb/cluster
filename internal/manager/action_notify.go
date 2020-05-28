package manager

// cluster -> slave node
type NotifyAction struct {
	Action
}

func (this *NotifyAction) Name() string {
	return "notify"
}

func (this *NotifyAction) TypeId() int8 {
	return 4
}
