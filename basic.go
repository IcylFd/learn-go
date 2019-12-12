/*
 * @Author: lifangdi_i
 * @Date: 2019-12-12 10:31:30
 * @LastEditors: lifangdi_i
 * @LastEditTime: 2019-12-12 20:11:39
 */
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 欧拉公式
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义 const (一般不全部大写)
// const数值可以作为各种类型使用
func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota // 自增枚举值
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	a := 3 + 4i
	fmt.Println(a)
	euler()
	triangle()
	consts()
	enums()
}
