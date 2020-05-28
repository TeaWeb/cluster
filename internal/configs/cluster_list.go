package configs

import (
	"github.com/go-yaml/yaml"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/files"
	"github.com/iwind/TeaGo/logs"
)

var clusterListFile = "clusters.yml"

type ClusterList struct {
	ClusterIds []string `yaml:"clusterIds"`
}

// always return a not-nil object
func SharedClusterList() *ClusterList {
	lock()
	defer unlock()

	configFile := files.NewFile(Tea.ConfigFile(clusterListFile))
	if !configFile.Exists() {
		return &ClusterList{}
	}

	data, err := configFile.ReadAll()
	if err != nil {
		logs.Error(err)
		return &ClusterList{}
	}

	config := &ClusterList{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		logs.Error(err)
		return &ClusterList{}
	}

	return config
}

func (this *ClusterList) AddCluster(clusterId string) {
	this.ClusterIds = append(this.ClusterIds, clusterId)
}

func (this *ClusterList) RemoveCluster(clusterId string) {
	result := []string{}
	for _, s := range this.ClusterIds {
		if s == clusterId {
			continue
		}
		result = append(result, s)
	}
	this.ClusterIds = result
}

func (this *ClusterList) FindAllClusters() []*ClusterConfig {
	result := []*ClusterConfig{}
	for _, clusterId := range this.ClusterIds {
		cluster := FindCluster(clusterId)
		if cluster == nil {
			continue
		}
		result = append(result, cluster)
	}
	return result
}

func (this *ClusterList) Save() error {
	lock()
	defer unlock()

	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}
	configFile := files.NewFile(Tea.ConfigFile(clusterListFile))
	return configFile.Write(data)
}
