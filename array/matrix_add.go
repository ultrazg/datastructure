package main

// 矩阵相加
func main() {
	A := [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}
	B := [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	C := make([][]int, 3)

	for i, _ := range C {
		C[i] = make([]int, 3)
	}

	println("[矩阵A的各个元素]")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(A[i][j])
			println()
		}
	}

	println("[矩阵B的各个元素]")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(B[i][j])
			println()
		}
	}

	matrixAdd(A, B, &C)

	println("[矩阵C的各个元素]")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(C[i][j])
			println()
		}
	}
}

// matrixAdd 函数实现了两个矩阵的相加
func matrixAdd(arr1, arr2 [][]int, r *[][]int) {
	if len(arr1) != len(arr2) || len(arr1[0]) != len(arr2[0]) {
		println("矩阵维度不匹配")
		return
	}

	for i, row := range arr1 {
		for j, val := range row {
			(*r)[i][j] = val + arr2[i][j]
		}
	}
}
