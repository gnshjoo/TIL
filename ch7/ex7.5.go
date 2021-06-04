package main

import "fmt"

// 반환할 변수명 명시 시 반환할 값을 지정하지 않고 return 가능

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
}