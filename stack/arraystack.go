package stack

import (
	"errors"
	"math"
)

const maxSize = math.MaxUint32

//ErrMaxCap #
var ErrMaxCap = errors.New("cap out of range")

//arrayStack is simple LIFO data structure
// it uses a growable buffer
type arrayStack struct {
	cap  int
	top  int
	base []Item
}

//ArrayBasedStack #
func ArrayBasedStack(size int) Stack {
	as := new(arrayStack)
	as.base = make([]Item, size)
	as.cap = size
	return as
}

//IsEmpty #
func (as *arrayStack) IsEmpty() bool {
	return as.top == 0
}

func (as *arrayStack) grow(minCap int) error {
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
func (as *arrayStack) Push(item Item) error {
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
func (as *arrayStack) Pop() (Item, error) {
	var item Item
	if as.IsEmpty() {
		return item, ErrNoItem
	}
	as.top--
	return as.base[as.top], nil
}

//Peek #
func (as *arrayStack) Peek() (Item, error) {
	var item Item
	if as.IsEmpty() {
		return item, ErrNoItem
	}
	return as.base[as.top-1], nil
}

//Cap returns current capacity of the stack
func (as *arrayStack) Cap() int {
	return as.cap
}

//Len returns length of stack
func (as *arrayStack) Len() int {
	return as.top
}
