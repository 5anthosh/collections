package arraystack

import (
	"errors"
	"math"

	"github.com/5anthosh/collections/stack"
)

//Item : Change it into what data you want to store
type Item int

const maxSize = math.MaxUint32

//ErrMaxCap #
var ErrMaxCap = errors.New("cap out of range")

//ArrayStack is simple LIFO data structure
// it uses a growable buffer
type ArrayStack struct {
	cap  int
	top  int
	base []Item
}

//New #
func New(size int) ArrayStack {
	var as ArrayStack
	as.base = make([]Item, size)
	as.cap = size
	return as
}

//IsEmpty #
func (as *ArrayStack) IsEmpty() bool {
	return as.top == 0
}

func (as *ArrayStack) grow(minCap int) error {
	newCap := as.cap << 1
	if minCap > newCap {
		newCap = minCap
	}
	if newCap > maxSize {
		return ErrMaxCap
	}
	newArray := make([]Item, newCap)
	copy(newArray, as.base)
	as.base = newArray
	as.cap = newCap
	return nil
}

//Push #
func (as *ArrayStack) Push(item Item) error {
	var err error

	if as.top == as.cap {
		err = as.grow(as.cap + 1)
		if err != nil {
			return err
		}
	}
	as.base[as.top] = item
	as.top++
	return nil
}

//Pop #
func (as *ArrayStack) Pop() (Item, error) {
	var item Item
	if as.IsEmpty() {
		return item, stack.ErrNoItem
	}
	as.top--
	return as.base[as.top], nil
}

//Peek #
func (as *ArrayStack) Peek() (Item, error) {
	var item Item
	if as.IsEmpty() {
		return item, stack.ErrNoItem
	}
	return as.base[as.top-1], nil
}

//Cap returns current capacity of the stack
func (as *ArrayStack) Cap() int {
	return as.cap
}

//Len returns length of stack
func (as *ArrayStack) Len() int {
	return as.top
}
