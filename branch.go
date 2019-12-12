/*
 * @Author: lifangdi_i
 * @Date: 2019-12-12 20:13:08
 * @LastEditors: lifangdi_i
 * @LastEditTime: 2019-12-12 20:29:44
 */
package main

import (
	"fmt"
	"io/ioutil"
)

// switch会自动break, 除非使用fallthrough
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator: " + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	// switch后面可以不加条件，在逻辑里增加条件
	switch {
	case score < 0 || score > 100:
		// panic会中断语句执行，报错
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	// if的条件里可以赋值
	// if的条件里赋值的变量作用域就在这个if语句里
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(eval(1, 3, "+"))
	// eval(1, 3, "_")

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(100),
		grade(92),
		grade(101),
	)
}
