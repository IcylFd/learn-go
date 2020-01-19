/*
 * @Date: 2020-01-19 14:17:00
 * @LastEditors  : lifangdi
 * @LastEditTime : 2020-01-19 14:32:41
 */
package main

import (
	"fmt"
)

// go指针不能运算
// go参数传递：只有值传递
// （值传递 引用传递：值传递将参数做了一份拷贝，main函数里的值不变）

// 值传递 和 指针 配合实现相当于引用传递

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	// var a int = 2
	// var pa *int = &a
	// *pa = 3
	// fmt.Println(a) // 3
	a, b := 3, 4
	// swap(&a, &b)
	// fmt.Println(a, b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
