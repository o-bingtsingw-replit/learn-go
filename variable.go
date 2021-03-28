package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// bb := 5 函数外面不能用必须要用var关键字
// 包作用域 多行可以用一个var简写
var (
	aa = 3
	bb = true
	cc = "aaa"
)

// 每个类型都有固定的zero value
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

// 一般能用简写就用简写
func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	a = 5
	fmt.Println(a, b, c, s)
}

// 特殊的复数类型
func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	ret := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(ret)
	fmt.Printf("%.3f\n", ret)
}

// 有强制类型转换, 没有隐性类型转换
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const a, b = 4, 3
	var c int
	// 常量不定义类型的话，相当于占位，可以在用到的时候推测类型
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

// go 没有enum类型，可以用const代替
func enums() {
	// 普通常量
	const (
		java   = 0
		c      = 1
		golang = 2
		js     = 3
	)
	fmt.Println(java, c, golang, js)

	// 自增常量
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
	consts()
	enums()
}
