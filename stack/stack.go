package stack

import (
	"errors"
)

//ErrNoItem #
var ErrNoItem = errors.New("No Item to pop")

//Stack is simple LIFO data structure
type Stack struct {
	cap       int
	len       int
	size      int
	baseBlock *block
}

type block struct {
	len      int
	base     []interface{}
	previous *block
}

func (stb *block) add(item interface{}) {
	stb.base[stb.len] = item
	stb.len++
}

func (stb *block) isEmpty() bool {
	return stb.isBlockEmpty() && stb.previous == nil
}

func (stb *block) isBlockEmpty() bool {
	return stb.len == 0
}
func (st *Stack) addBlock() {
	newblock := newBlock(st.size)
	if st.baseBlock != nil {
		newblock.previous = st.baseBlock
	}
	st.baseBlock = newblock
	st.cap += 5
}

func newBlock(size int) *block {
	stb := new(block)
	stb.base = make([]interface{}, size)
	return stb
}

//New creates new Stack with size
func New(size int) *Stack {
	st := new(Stack)
	st.size = size
	return st
}

//IsEmpty checks whether stack is empty
func (st *Stack) IsEmpty() bool {
	return st.baseBlock.isEmpty()
}

//Push pushes new item on the top of the stack
func (st *Stack) Push(i interface{}) {
	if st.baseBlock == nil || st.baseBlock.len >= st.size {
		st.addBlock()
	}
	st.baseBlock.add(i)
	st.len++
}

//Pop removes Last in item and return it
func (st *Stack) Pop() (interface{}, error) {
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
func (st *Stack) Peek() (interface{}, error) {
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
func (st *Stack) Cap() int {
	return st.cap
}

//Len returns length of stack
func (st *Stack) Len() int {
	return st.len
}
