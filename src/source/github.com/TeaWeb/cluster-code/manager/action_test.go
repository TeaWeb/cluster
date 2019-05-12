package manager

import (
	"github.com/iwind/TeaGo/logs"
	"reflect"
	"testing"
)

func TestNewAction(t *testing.T) {
	action := &RegisterAction{}
	action.ClusterId = "123456"
	action.NodeId = "654321"
	action.ClusterSecret = "abcdef"
	logs.PrintAsJSON(action, t)

	msg := NewActionMessage(action)
	data, err := msg.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))

	action1, err := Unmarshal("register", data[32:])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reflect.TypeOf(action1).Name())
	logs.PrintAsJSON(action1, t)
}
