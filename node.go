package bidirectional

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func CreateNode(prev *Node, next *Node, value interface{}) *Node {

	return &Node{
		prev:  prev,
		next:  next,
		value: value,
	}
}

func (n *Node) SetValue(newValue interface{}) {

	n.value = newValue
}

func (n Node) GetValue() interface{} {

	return n.value
}

func (n *Node) SetPrev(newPrevNode *Node) {

	n.prev = newPrevNode
}

func (n Node) GetPrev() *Node {

	return n.prev
}

func (n *Node) SetNext(newNextNode *Node) {

	n.next = newNextNode
}

func (n Node) GetNext() *Node {

	return n.next
}
