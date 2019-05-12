package manager

type SuccessAction struct {
	Action

	Message string
}

func (this *SuccessAction) Name() string {
	return "success"
}

func (this *SuccessAction) TypeId() int8 {
	return 1
}
