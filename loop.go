package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("forever")
	}
}

// 没有while, for就能代替while的功能
func main() {
	fmt.Println(
		convertToBin(3),
		convertToBin(11),
		convertToBin(3342345643),
		convertToBin(0),
	)

	if _, filename, _, ok := runtime.Caller(0); ok {
		printFile(filename)
	}

	//forever()
}
