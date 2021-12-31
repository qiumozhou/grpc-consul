package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main(){
	consulConfig := api.DefaultConfig()
	consulClient,err := api.NewClient(consulConfig)
	if err != nil{
		fmt.Println("deregister, api.newclient err:",err)
		return
	}
	consulClient.Agent().ServiceDeregister("wbw1_id")
}
