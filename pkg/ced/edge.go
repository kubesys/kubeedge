/* Copyright (2021, ) Institute of Software, Chinese Academy of Sciences */
package ced

import (
	"encoding/json"
	"fmt"
	"github.com/kubesys/kubernetes-client-go/pkg/kubesys"
)

/*
 * author: wuheng@iscas.ac.cn
 * date:   2021-5-20
 */

type EdgeWatch struct {
	CloudHubClient *kubesys.KubernetesClient
}

func NewEdgeWatch(cloudHubClient *kubesys.KubernetesClient) *EdgeWatch {
	edgeWatch := new(EdgeWatch)
	edgeWatch.CloudHubClient = cloudHubClient
	return edgeWatch
}

func (p EdgeWatch) DoAdded(obj map[string]interface{}) {
	fmt.Println("add Node")
}
func (p EdgeWatch) DoModified(obj map[string]interface{}) {
	data, _ := json.Marshal(obj)
	p.CloudHubClient.UpdateResource(string(data))
	fmt.Println("update Node")
}
func (p EdgeWatch) DoDeleted(obj map[string]interface{}) {
	fmt.Println("delete Node")
}

func (hub *CEDHub) Report() {

	_, err := hub.CloudHubClient.GetResource("Node", "", hub.Name)

	if err != nil {
		node := GetNodeJSON(hub.Name)
		fmt.Println(node)
		hub.CloudHubClient.CreateResource(node)
	}

	hub.EdgeHubClient.WatchResource("Node", "", hub.Name,
		kubesys.NewKubernetesWatcher(hub.EdgeHubClient, NewEdgeWatch(hub.CloudHubClient)))
}

