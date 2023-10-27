package main

import "math/rand"

var top = -1

// 以数组模拟扑克牌洗牌以及发牌的过程
// 以随机数取得扑克牌后放入堆栈，放满 52 张牌后开始发牌，同样使用堆栈来发给 4 个人
func main() {
	card := make([]int, 52)
	stack := make([]int, 52)
	k := 0
	var test int

	for i := 0; i < 52; i++ {
		card[i] = i

		println("[正在洗牌...]")
		for k < 30 {
			for i := 0; i < 51; i++ {
				for j := i + 1; j < 52; j++ {
					if rand.Intn(5) == 2 {
						test = card[i]
						card[i] = card[j]
						card[j] = test
					}
				}
			}

			k++
		}
		i = 0

		for i != 52 {
			push(&stack, 52, card[i])
			i++
		}

		println("[逆时针发牌]")
	}
}

func push(stack *[]int, MAX, val int) {
	if top > MAX-1 {
		panic("[堆栈已经满了]")
	} else {
		top++
		(*stack)[top] = val
	}
}

func pop(stack *[]int) int {
	if top < 0 {
		panic("[堆栈已经空了]")
	} else {
		top--

		return (*stack)[top]
	}
}
