/* Copyright (2021, ) Institute of Software, Chinese Academy of Sciences */
package ced

import (
	"encoding/json"
	"fmt"
	"github.com/kubesys/kubernetes-client-go/pkg/kubesys"
	"time"
)

/*
 * author: wuheng@iscas.ac.cn
 * date:   2021-5-20
 */

type EdgeWatch struct {
	CloudHubClient *kubesys.KubernetesClient
	RegisterName   string
}

func NewEdgeWatch(cloudHubClient *kubesys.KubernetesClient, registerName string) *EdgeWatch {
	edgeWatch := new(EdgeWatch)
	edgeWatch.CloudHubClient = cloudHubClient
	edgeWatch.RegisterName = registerName
	return edgeWatch
}

func (p EdgeWatch) DoAdded(obj map[string]interface{}) {
	fmt.Println("Edge connected...")
}
func (p EdgeWatch) DoModified(obj map[string]interface{}) {

	jsonObj, _ := p.CloudHubClient.GetResource("Node", "", p.RegisterName)
	dataObj := jsonObj.Object
	dataObj["status"] = obj["status"]
	data, _ := json.Marshal(dataObj)
	p.CloudHubClient.UpdateResourceStatus(string(data))
	fmt.Println("Updated edge status...")
}
func (p EdgeWatch) DoDeleted(obj map[string]interface{}) {
	fmt.Println("Edge disconnected...")
}

func (hub *CEDHub) Report() {

	_, err := hub.CloudHubClient.GetResource("Node", "", hub.RegisterName)

	if err != nil {
		edgeNode, err := hub.EdgeHubClient.GetResource("Node", "", hub.RealName)
		if err != nil {
			fmt.Println("Wrong edge realname, please login in edge and execute 'kubectl get no'")
			return
		}

		cloudNode := make(map[string]interface{})
		cloudNode["spec"] = edgeNode.Object["spec"]

		specBytes, _ := json.Marshal(cloudNode)
		hub.CloudHubClient.CreateResource(string(specBytes))
	}

	go hub.EdgeHubClient.WatchResource("Node", "", hub.RealName,
		kubesys.NewKubernetesWatcher(hub.EdgeHubClient,
			NewEdgeWatch(hub.CloudHubClient, hub.RegisterName)))

	for {
		targetObj,  _ := hub.EdgeHubClient.GetResource("Node", "", hub.RealName)
		mappingObj, _ := hub.CloudHubClient.GetResource("Node", "", hub.RegisterName)
		mappingObj.Object["status"] = targetObj.Object["status"]
		data, _ := json.Marshal(mappingObj.Object)
		_, err = hub.CloudHubClient.UpdateResourceStatus(string(data))
		fmt.Println(err)
		fmt.Println(string(data))
		fmt.Println("Updated...")
		time.Sleep(20 * time.Second)
	}

}

