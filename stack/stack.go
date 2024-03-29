package stack

import (
	"errors"
)

//Item : Hack it into what data you want to store
type Item interface{}

//ErrNoItem #
var ErrNoItem = errors.New("No Item to pop")

//blockBasedListStack is simple LIFO data structure
// It uses block based list to store data
type blockBasedListStack struct {
	cap       int
	len       int
	size      int
	baseBlock *block
}

//Stack is simple LIFO data structure
type Stack interface {
	Pop() (Item, error)
	Peek() (Item, error)
	Len() int
	Cap() int
	IsEmpty() bool
	Push(Item)
}

type block struct {
	len      int
	base     []Item
	previous *block
}

func (stb *block) add(item Item) {
	stb.base[stb.len] = item
	stb.len++
}

func (stb *block) isEmpty() bool {
	return stb.isBlockEmpty() && stb.previous == nil
}

func (stb *block) isBlockEmpty() bool {
	return stb.len == 0
}
func (st *blockBasedListStack) addBlock() {
	newblock := newBlock(st.size)
	if st.baseBlock != nil {
		newblock.previous = st.baseBlock
	}
	st.baseBlock = newblock
	st.cap += 5
}

func newBlock(size int) *block {
	stb := new(block)
	stb.base = make([]Item, size)
	return stb
}

//NewBlockBasedStack creates new Stack with size
func NewBlockBasedStack(size int) Stack {
	st := new(blockBasedListStack)
	st.size = size
	return st
}

//New creates new Stack with size
func New(size int) Stack {
	return NewBlockBasedStack(size)
}

//IsEmpty checks whether stack is empty
func (st *blockBasedListStack) IsEmpty() bool {
	return st.baseBlock.isEmpty()
}

//Push pushes new item on the top of the stack
func (st *blockBasedListStack) Push(i Item) {
	if st.baseBlock == nil || st.baseBlock.len >= st.size {
		st.addBlock()
	}
	st.baseBlock.add(i)
	st.len++
}

//Pop removes Last in item and return it
func (st *blockBasedListStack) Pop() (Item, error) {
	baseBlock := st.baseBlock
	if baseBlock.isBlockEmpty() {
		old := baseBlock.previous
		if old != nil {
			st.baseBlock = nil
			st.cap -= st.size
			st.baseBlock = baseBlock.previous
		} else {
			return 0, ErrNoItem
		}
	}
	st.len--
	st.baseBlock.len--
	return st.baseBlock.base[st.baseBlock.len], nil
}

//Peek lets peek top of the stack
func (st *blockBasedListStack) Peek() (Item, error) {
	baseBlock := st.baseBlock
	if baseBlock.isBlockEmpty() {
		old := baseBlock.previous
		if old != nil {
			return old.base[old.len-1], nil
		}
		return 0, ErrNoItem
	}
	return baseBlock.base[baseBlock.len-1], nil
}

//Cap returns current capacity of the stack
func (st *blockBasedListStack) Cap() int {
	return st.cap
}

//Len returns length of stack
func (st *blockBasedListStack) Len() int {
	return st.len
}
