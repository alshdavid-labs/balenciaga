package node

type Node struct {
	Address string
	Connected int64
}

func NewNode(address string) *Node {
	return &Node{
		Address: address,
		Connected: 0,
	}
}