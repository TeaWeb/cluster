package configs

import "testing"

func TestSharedClusterList(t *testing.T) {
	list := SharedClusterList()
	//list.AddCluster("123")
	//list.AddCluster("456")
	err := list.Save()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list.ClusterIds)
}
