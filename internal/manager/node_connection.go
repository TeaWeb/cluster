package manager

import (
	"errors"
	"fmt"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/timers"
	"github.com/vmihailenco/msgpack"
	"io"
	"net"
	"sync"
	"time"
)

type NodeConnection struct {
	ClusterId string
	NodeId    string

	ConnId int64
	Conn   net.Conn

	encoder     *msgpack.Encoder
	decoder     *msgpack.Decoder
	queueLocker sync.Mutex
}

func NewNodeConnection(connId int64) *NodeConnection {
	return &NodeConnection{
		ConnId: connId,
	}
}

func (this *NodeConnection) Listen() error {
	if this.Conn == nil {
		return errors.New("[node]no connection ready")
	}

	this.encoder = msgpack.NewEncoder(this.Conn)
	this.decoder = msgpack.NewDecoder(this.Conn)

	// if not authenticated in 30 seconds, we will close the connection
	timers.Delay(30*time.Second, func(timer *time.Timer) {
		if len(this.NodeId) == 0 {
			this.Close()
		}
	})

	// read actions
	this.Read(func(action ActionInterface) {
		if len(this.NodeId) > 0 {
			if action.Name() != "ping" {
				logs.Println("[manager]["+this.Conn.RemoteAddr().String()+"]["+this.NodeId+"]receive action:", action.Name())
			}
		} else {
			logs.Println("[manager]["+this.Conn.RemoteAddr().String()+"]receive action:", action.Name())
		}
		err := action.Execute(this)
		if err != nil {
			logs.Error(err)
		}
	})

	return nil
}

// read action from cluster
func (this *NodeConnection) Read(f func(action ActionInterface)) {
	for {
		typeId, _, err := this.decoder.DecodeExtHeader()
		if err != nil {
			if err == io.EOF {
				break
			}
			//logs.Error(err) // always is a closed error
			break
		}
		instance := FindActionInstance(typeId)
		if instance == nil {
			logs.Error(errors.New("invalid action type '" + fmt.Sprintf("%d", typeId) + "'"))
			continue
		}
		err = this.decoder.Decode(instance)
		if err != nil {
			if err == io.EOF {
				break
			}
			logs.Error(err)
			break
		}
		f(instance)
	}
}

func (this *NodeConnection) Write(action ActionInterface) error {
	if this.Conn == nil {
		return errors.New("no connection to node")
	}
	this.queueLocker.Lock()
	action.BaseAction().Id = GenerateActionId()
	err := this.encoder.Encode(action)
	this.queueLocker.Unlock()
	return err
}

func (this *NodeConnection) Reply(fromAction ActionInterface, newAction ActionInterface) error {
	newAction.BaseAction().RequestId = fromAction.BaseAction().Id
	return this.Write(newAction)
}

func (this *NodeConnection) Close() error {
	if this.Conn != nil {
		err := this.Conn.Close()
		this.Conn = nil
		return err
	}
	return nil
}

func (this *NodeConnection) IsAuthenticated() bool {
	return len(this.ClusterId) > 0 && len(this.NodeId) > 0
}
