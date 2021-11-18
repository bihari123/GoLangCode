package rack

import (
	"computing_cluster/node"
	"fmt"
)
type Rack struct{
	max_nodes int
	Nodes_list [] node.Node
}
 func New(max_nodes int) Rack{
	var list[] node.Node
 	r:= Rack{max_nodes,list}
 	return r
 }
 
 func (r *Rack)AddNode(node node.Node){
	 if len(r.Nodes_list) < r.max_nodes {
		 r.Nodes_list=append(r.Nodes_list,node)
		 //fmt.Println("node added")
	 }else {
	 	fmt.Println("not enough space in Rack")
	 }
 }
 func (r Rack) IsFull() bool{
 	if len(r.Nodes_list)<r.max_nodes{
 		return false
	}else{
		return true
	}
 }