package node

type Node struct {
	item string
	next *Node
}

func New(item string) *Node {
	return &Node{item, nil}
}

func (n *Node) Item() string {
	return n.item
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetItem(item string) string {
	n.item = item
	return item
}

func (n *Node) SetNext(nextNode *Node) *Node {
	n.next = nextNode
	return n.next
}
