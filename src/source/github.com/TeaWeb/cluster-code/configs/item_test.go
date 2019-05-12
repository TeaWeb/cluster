package configs

import (
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func TestItem(t *testing.T) {
	item := NewItem()
	item.Id = "123456"
	item.Sum = "abcde"
	item.Flags = []int{1, 2, 3, 4, 5, 6}
	item.Data = []byte("ABCDEF")
	data := item.Marshal()
	t.Log(data)
	t.Log(string(data))

	t.Log("===")
	item1 := UnmarshalItem(data)
	logs.PrintAsJSON(item1, t)
	t.Log(string(item1.Data))
}
