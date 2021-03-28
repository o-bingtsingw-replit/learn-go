package main

import (
	"fmt"
	"os"
)

// 默认是break, 如果不需要break则使用关键字fallthrough
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		fmt.Printf("Wrong score: %d", score)
		panic(0)
	}

	return g
}

func main() {
	// if 条件里可以赋值, 赋值的作用域就在这个if语句里
	if dir, err := os.Getwd(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}

	fmt.Println(
		grade(0),
		grade(88),
		grade(66),
		grade(100),
		//grade(1000),
	)
}
