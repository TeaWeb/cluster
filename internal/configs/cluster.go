package configs

import (
	"errors"
	"github.com/go-yaml/yaml"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/files"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/utils/string"
)

var clusterFile = "cluster.${id}.yml"

type ClusterConfig struct {
	Id         string        `yaml:"id" json:"id"`
	IsOn       bool          `yaml:"isOn" json:"isOn"`
	Name       string        `yaml:"name" json:"name"`
	Namespaces []Namespace   `yaml:"namespaces" json:"namespaces"`
	Nodes      []*NodeConfig `yaml:"nodes" json:"nodes"`
	Secret     string        `yaml:"secret" json:"secret"`
}

func FindCluster(clusterId string) *ClusterConfig {
	if len(clusterId) == 0 {
		return nil
	}

	if !idRegexp.MatchString(clusterId) {
		return nil
	}

	filename := mapVars(clusterFile, map[string]string{
		"id": clusterId,
	})
	file := files.NewFile(Tea.ConfigFile(filename))

	lock()
	defer unlock()

	if !file.Exists() {
		return nil
	}

	data, err := file.ReadAll()
	if err != nil {
		logs.Error(err)
		return nil
	}

	config := &ClusterConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil
	}
	return config
}

func NewCluster() *ClusterConfig {
	return &ClusterConfig{
		Id:     stringutil.Rand(16),
		Secret: stringutil.Rand(32),
		IsOn:   true,
	}
}

func (this *ClusterConfig) FindNode(nodeId string) *NodeConfig {
	for _, node := range this.Nodes {
		if node.Id == nodeId {
			return node
		}
	}
	return nil
}

func (this *ClusterConfig) FindMasterNode() *NodeConfig {
	for _, node := range this.Nodes {
		if node.Role == NodeRoleMaster {
			return node
		}
	}
	return nil
}

func (this *ClusterConfig) AddNode(node *NodeConfig) {
	lock()
	defer unlock()

	if node == nil {
		return
	}

	for _, n := range this.Nodes {
		if n.Id == node.Id {
			return
		}
	}
	this.Nodes = append(this.Nodes, node)
}

func (this *ClusterConfig) RemoveNode(nodeId string) {
	lock()
	defer unlock()

	result := []*NodeConfig{}
	for _, n := range this.Nodes {
		if n.Id == nodeId {
			continue
		}
		result = append(result, n)
	}
	this.Nodes = result
}

func (this *ClusterConfig) Save() error {
	if len(this.Id) == 0 {
		return errors.New("no 'id' found")
	}

	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}

	filename := mapVars(clusterFile, map[string]string{
		"id": this.Id,
	})

	lock()
	defer unlock()

	return files.NewFile(Tea.ConfigFile(filename)).Write(data)
}

func (this *ClusterConfig) Delete() error {
	if len(this.Id) == 0 {
		return errors.New("no 'id' found")
	}
	filename := mapVars(clusterFile, map[string]string{
		"id": this.Id,
	})
	lock()
	defer unlock()
	file := files.NewFile(Tea.ConfigFile(filename))
	return file.Delete()
}
