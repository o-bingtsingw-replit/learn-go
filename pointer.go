package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func swapByPointer(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)

	swapByPointer(&a, &b)
	fmt.Println(a, b)
}
