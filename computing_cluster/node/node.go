package node

type Node struct {
	memory_size int
	processor_number int
}

func New(memory_size int, processor_number int) Node {
	n:= Node{memory_size,processor_number}
	return n
}

func (n Node) GetMemSize() int{
	return n.memory_size
}

func (n Node) GetProNo() int{
	return n.processor_number
}
