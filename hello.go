/*
 * @Author: lifangdi_i
 * @Date: 2019-12-12 10:31:30
 * @LastEditors: lifangdi_i
 * @LastEditTime: 2019-12-12 11:23:02
 */
package main

import (
	"fmt"
)

// 函数外定义变量不能用冒号
var aa = 3
var (
	k = 1
	b = "test"
)

func variableZeroValue() {
	// 变量名 在前，变量类型在后
	// 在复合变量定义时有优势
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "aaa"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 1, 2, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 第一次定义变量用冒号
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello go", aa, b, k)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
