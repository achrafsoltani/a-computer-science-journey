package main

type Stack interface {
	Push(string)
	Pop() string
	isFull() bool
	isEmpty() bool
}

type ArrayStack struct {
	data []string
	top  int
	cap  int
}

func CreateArrayStack(cap int) *ArrayStack {
	arrayStack := new(ArrayStack)
	arrayStack.data = make([]string, cap)
	arrayStack.cap = cap
	arrayStack.top = -1
	return arrayStack
}

func (arrayStack *ArrayStack) isEmpty() bool {
	return arrayStack.top == -1
}

func (arrayStack *ArrayStack) isFull() bool {
	return arrayStack.top == arrayStack.cap
}

func (arrayStack *ArrayStack) Push(value string) bool {
	if arrayStack.isFull() {
		println("Stack Overflow")
		return false
	}
	arrayStack.top++
	arrayStack.data[arrayStack.top] = value
	return true
}

func (arrayStack *ArrayStack) Pop() string {
	if arrayStack.isEmpty() {
		println("Stack is Empty")
		return ""
	}
	value := arrayStack.data[arrayStack.top]
	arrayStack.top--
	return value
}

func main() {
	stack := CreateArrayStack(5)
	stack.Push("Apple")
	stack.Push("Oranges")
	stack.Push("Bananas")
	println(stack.Pop())
	for !stack.isEmpty() {
		println(stack.Pop())
	}
}
