package configs

import (
	"bytes"
	"fmt"
	"github.com/iwind/TeaGo/types"
	"time"
)

type NodeLog struct {
	Timestamp   int64
	Action      string
	Object      string
	Description string
}

func UnmarshalNodeLog(data []byte) *NodeLog {
	if len(data) == 0 {
		return &NodeLog{}
	}
	pieces := bytes.SplitN(data, []byte{'|'}, 4)
	c := len(pieces)
	switch c {
	case 1:
		return &NodeLog{
			Timestamp: types.Int64(string(pieces[0])),
		}
	case 2:
		return &NodeLog{
			Timestamp: types.Int64(string(pieces[0])),
			Action:    string(pieces[1]),
		}
	case 3:
		return &NodeLog{
			Timestamp: types.Int64(string(pieces[0])),
			Action:    string(pieces[1]),
			Object:    string(pieces[2]),
		}
	case 4:
		return &NodeLog{
			Timestamp:   types.Int64(string(pieces[0])),
			Action:      string(pieces[1]),
			Object:      string(pieces[2]),
			Description: string(pieces[3]),
		}
	}
	return &NodeLog{}
}

func (this *NodeLog) Marshal() []byte {
	return []byte(fmt.Sprintf("%d", this.Timestamp) + "|" + this.Action + "|" + this.Object + "|" + this.Description)
}

func (this *NodeLog) Time() time.Time {
	return time.Unix(this.Timestamp, 0)
}
