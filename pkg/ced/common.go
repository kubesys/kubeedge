/* Copyright (2021, ) Institute of Software, Chinese Academy of Sciences */
package ced

import (
	"github.com/kubesys/kubernetes-client-go/pkg/kubesys"
	"strings"
)

/*
 * author: wuheng@iscas.ac.cn
 * date:   2021-5-20
 */

type CEDHub struct {
	CloudHubClient *kubesys.KubernetesClient
	EdgeHubClient  *kubesys.KubernetesClient
	RegisterName   string
	RealName       string
}

func NewCEDHub(edgeName string, masterName string, cloudUrl string, cloudToken string, edgeUrl string, edgeToken string) *CEDHub {
	cloudHub := kubesys.NewKubernetesClient(cloudUrl, cloudToken)
	cloudHub.Init()
	edgeHub  := kubesys.NewKubernetesClient(edgeUrl, edgeToken)
	edgeHub.Init()
	hub := new(CEDHub)
	hub.CloudHubClient = cloudHub
	hub.EdgeHubClient = edgeHub
	hub.RegisterName = edgeName
	hub.RealName = masterName
	return hub
}

func GetNodeJSON (name string) string {
	str := "{\"apiVersion\": \"v1\",\"kind\": \"Node\",\"metadata\": {\"name\": \"#NAME\",\"namespace\": \"default\"}}"
	return strings.Replace(str, "#NAME", name, -1)
}