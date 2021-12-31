package main

import (
	pb "com.qmz.dev/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
	"context"
)

func main(){
	consulConfig := api.DefaultConfig()
	conculClient,err := api.NewClient(consulConfig)
	if err!=nil{
		fmt.Println("client, api.newclient err:",err)
		return
	}
	services,_,err := conculClient.Health().Service("wbw001_grpc_consul","wbw1",true,nil )
	addr := services[0].Service.Address+":"+strconv.Itoa(services[0].Service.Port)
	grpcConn,err := grpc.Dial(addr,grpc.WithInsecure())
	if err != nil{
		fmt.Println("dial err:",err)
	}
	defer grpcConn.Close()
	grpcClient := pb.NewHelloClient(grpcConn)

	var person pb.Person
	person.Name = ",qmz!!!"
	person.Age = 18
	p,err := grpcClient.SayHello(context.TODO(),&person)
	fmt.Println(p,err)
}
