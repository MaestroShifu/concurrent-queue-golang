package node

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
		prev: nil,
	}

}

func (n *Node) SetPrev(node *Node) {
	n.prev = node
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

func (n *Node) GetNext() *Node {
	return n.prev.next
}

func (n *Node) GetData() interface{} {
	return n.data
}
