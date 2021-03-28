package main

import (
	"fmt"
	"math"
)

func op(operator func(int, int) int, a int, b int) int {
	return operator(a, b)
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	// 函数作为参数
	fmt.Println(op(pow, 2, 10))

	// 匿名函数
	fmt.Println(op(
		func(a int, b int) int {
			return pow(a, b)
		}, 2, 11))

	// 可变参数列表
	fmt.Println(sum(1, 3, 5, 7, 9))
}
