package main

import (
	"computing_cluster/computing_cluster"
	"computing_cluster/node"
	"fmt"
)

func main(){

	var nodes_per_rack int
	fmt.Println("enter the number of nodes per rack")
	fmt.Scanln(&nodes_per_rack)
	comp_cluster:=computing_cluster.New(nodes_per_rack)
	for i:=0;i<650;i++{
		comp_cluster.AddNode(node.New(64,1))
	}
	for j:=0;j<16;j++{
		comp_cluster.AddNode(node.New(1024,2))
	}

	fmt.Printf("nodes with atleast 32 GB memory: %d \n",comp_cluster.NodeCount(32))
	fmt.Printf("nodes with atleast 64 GB memory: %d \n",comp_cluster.NodeCount(64))
	fmt.Printf("nodes with atleast 128 GB memory: %d \n",comp_cluster.NodeCount(128))

	fmt.Printf("total number of processors: %d \n",comp_cluster.TotalProc())

	fmt.Println("total number of racks: ",comp_cluster.Rack_count())

}