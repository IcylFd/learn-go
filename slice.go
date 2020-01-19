/*
 * @Date: 2020-01-19 15:36:18
 * @LastEditors  : lifangdi
 * @LastEditTime : 2020-01-19 15:49:33
 */
package main

import (
	"fmt"
)

// 不加长度就是slice
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// slice是对array的一个view（引用）
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("after update slice s1")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("after update slice s2")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
}
