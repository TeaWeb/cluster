package manager

import (
	"errors"
	"github.com/iwind/TeaGo/logs"
	"io"
	"net"
	"source/github.com/TeaWeb/cluster-code/configs"
	"sync"
)

var SharedManager = NewManager()

type Manager struct {
	addr        string
	listener    net.Listener
	isListening bool
	error       string

	connList map[int64]*NodeConnection // connId => nodeConnection

	connId int64
	locker sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		connList: map[int64]*NodeConnection{},
	}
}

func (this *Manager) StartInBackground() {
	go func() {
		err := this.Start()
		if err != nil {
			logs.Error(err)
		}
	}()
}

func (this *Manager) Start() error {
	config := configs.SharedConfig()
	if len(config.Bind) == 0 {
		return errors.New("[manager]'bind' in config should not be empty")
	}

	this.addr = config.Bind

	server, err := net.Listen("tcp", this.addr)
	if err != nil {
		this.error = err.Error()
		return errors.New("[manager]" + err.Error())
	}

	logs.Println("[manager]start listener", this.addr)
	this.listener = server

	this.isListening = true

	for {
		listener := this.listener // forbidden change listener concurrently

		if listener == nil {
			break
		}

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		this.locker.Lock()
		this.connId ++

		logs.Println("[manager]receive a node connection", this.connId)

		nodeConn := NewNodeConnection(this.connId)
		nodeConn.Conn = conn

		this.connList[this.connId] = nodeConn
		this.locker.Unlock()

		go func(nodeConn *NodeConnection) {
			err = nodeConn.Listen()
			if err != nil && err != io.EOF {
				logs.Error(err)
			}

			this.locker.Lock()
			delete(this.connList, nodeConn.ConnId)
			this.locker.Unlock()

			logs.Println("[manager]close node connection", this.connId)
		}(nodeConn)
	}

	this.isListening = false

	return nil
}

func (this *Manager) IsListening() bool {
	return this.isListening
}

func (this *Manager) Addr() string {
	return this.addr
}

func (this *Manager) Error() string {
	return this.error
}

func (this *Manager) Stop() error {
	logs.Println("[manager]stop listener", this.addr)

	for _, conn := range this.connList {
		conn.Close()
	}

	this.connList = map[int64]*NodeConnection{}

	if this.listener != nil {
		err := this.listener.Close()
		this.listener = nil
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO cache nodeId -> conn mapping to improve performance
func (this *Manager) FindNodeState(clusterId string, nodeId string) (state *configs.NodeState, conn *NodeConnection) {
	this.locker.Lock()
	defer this.locker.Unlock()

	state = configs.NewNodeState()

	for _, nodeConn := range this.connList {
		if nodeConn.ClusterId == clusterId && nodeConn.NodeId == nodeId {
			state.IsActive = true
			conn = nodeConn
		}
	}

	return
}

func (this *Manager) CloseNode(nodeId string) error {
	this.locker.Lock()
	defer this.locker.Unlock()

	for _, conn := range this.connList {
		if conn.NodeId == nodeId {
			return conn.Close()
		}
	}

	return nil
}

func (this *Manager) CountNodes() int {
	this.locker.Lock()
	defer this.locker.Unlock()

	count := 0
	for _, nodeConn := range this.connList {
		if len(nodeConn.NodeId) > 0 {
			count ++
		}
	}
	return count
}

func (this *Manager) CountClusterNodes(clusterId string) int {
	this.locker.Lock()
	defer this.locker.Unlock()

	count := 0
	for _, nodeConn := range this.connList {
		if len(nodeConn.NodeId) > 0 && nodeConn.ClusterId == clusterId {
			count ++
		}
	}
	return count
}
