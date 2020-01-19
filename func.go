/*
 * @LastEditors  : lifangdi
 */
package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsopported option:" + op)
	}
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d, %d\n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a, b)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 可变参数列表(像数组一样使用)
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "+"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// 匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	// fmt.Println(eval(3, 4, "x"))

	// q, r := div(13, 3)
	// fmt.Println(q, r)
	fmt.Println(sum(3, 4, 5))
}
