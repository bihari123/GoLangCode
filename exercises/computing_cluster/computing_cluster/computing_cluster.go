package computing_cluster

import (
	"computing_cluster/node"
	"computing_cluster/rack"
)
type computingCluster struct{
	rack_list []rack.Rack
	nodes_per_rack int
	listEnd int
}
func New(nodes_per_rack int,) computingCluster {
	var rack_list []rack.Rack
	n:= computingCluster{rack_list,nodes_per_rack,0}

	return n
}

func(c *computingCluster) AddNode (n node.Node){
	last_elem:=&c.listEnd
	list:=&c.rack_list
	nodes_per_rack:=&c.nodes_per_rack

	if len(*list) == 0{
		*list=append(*list,rack.New(*nodes_per_rack))
		(*list)[*last_elem].AddNode(n)
	}else if (*list)[*last_elem].IsFull(){
		*last_elem=*last_elem+1
		*list=append(*list,rack.New(*nodes_per_rack))
		(*list)[*last_elem].AddNode(n)

	}else{
		(*list)[*last_elem].AddNode(n)
	}


}

func (c computingCluster) NodeCount(memory_size int) int{
	var num_of_nodes int
	num_of_nodes=0
	num_of_nodes=0
	for _,my_rack:= range c.rack_list{
		for _,my_node:= range my_rack.Nodes_list{
            if my_node.GetMemSize() > memory_size{
            	num_of_nodes++
			}
		}
	}
	return num_of_nodes
}

func (c computingCluster) TotalProc() int {
	sum:=0
	for _,my_list:= range c.rack_list{
		for _,my_node:= range my_list.Nodes_list{
			sum = sum + my_node.GetProNo()
		}
	}
	return sum
}

func (c computingCluster) Rack_count() int{
	return c.listEnd + 1
}