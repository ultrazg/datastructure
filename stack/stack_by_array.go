package main

// 数组模拟堆栈
type stackByArray struct {
	stack []int // 数组
	top   int   // 堆栈顶端的索引
}

func (stack *stackByArray) push(data int) {
	stack.stack = append(stack.stack, data)
	stack.top = len(stack.stack) - 1
}

func (stack *stackByArray) empty() bool {
	if len(stack.stack) == 0 {
		return true
	} else {
		return false
	}
}

func (stack *stackByArray) pop() int {
	if stack.empty() {
		panic("堆栈为空")
	}

	item := stack.stack[stack.top]
	stack.stack = stack.stack[:stack.top]
	stack.top--

	return item
}

func main() {
	stack := stackByArray{}

	for i := 0; i < 10; i++ {
		stack.push(i)
	}

	for !stack.empty() {
		println("stack pop", stack.pop())
	}
}
