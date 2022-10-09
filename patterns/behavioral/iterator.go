package behavioral

// Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.).

// https://refactoring.guru/design-patterns/iterator

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should not be here!")
}

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewLeafNode(value int) *Node {
	return &Node{Value: value}
}

func NewNode(value int, left, right *Node) *Node {
	n := &Node{
		Value: value,
		left:  left,
		right: right,
	}
	if left != nil {
		left.parent = n
	}
	if right != nil {
		right.parent = n
	}
	return n
}

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{
		Current:       root,
		root:          root,
		returnedStart: false,
	}
	return i.traverseToLeftMost()
}

func (i *InOrderIterator) traverseToLeftMost() *InOrderIterator {
	for i.Current.left != nil { // traverse to the left-most node
		i.Current = i.Current.left
	}
	return i
}

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}
}

func (i *InOrderIterator) Value() int {
	return i.Current.Value
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}
