package stack

type listBasedStack struct {
	top      int
	baseNode *node
}

type node struct {
	item Item
	old  *node
}

func ListBasedStack() Stack {
	return new(listBasedStack)
}

func (lbs *listBasedStack) IsEmpty() bool {
	return lbs.baseNode == nil
}

func (lbs *listBasedStack) Push(i Item) {
	newNode := new(node)
	newNode.item = i
	newNode.old = lbs.baseNode
	lbs.baseNode = newNode
	lbs.top++
}

func (lbs *listBasedStack) Pop() (Item, error) {
	var item Item
	if lbs.IsEmpty() {
		return item, ErrNoItem
	}
	item = lbs.baseNode.item
	lbs.baseNode = lbs.baseNode.old
	lbs.top--
	return item, nil
}

func (lbs *listBasedStack) Peek() (Item, error) {
	var item Item
	if lbs.IsEmpty() {
		return item, ErrNoItem
	}
	return lbs.baseNode.item, nil
}
func (lbs *listBasedStack) Len() int {
	return lbs.top
}

func (lbs *listBasedStack) Cap() int {
	return lbs.top
}
