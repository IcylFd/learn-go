/*
 * @Date: 2020-01-19 15:36:18
 * @LastEditors  : lifangdi
 * @LastEditTime : 2020-02-14 23:33:09
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

	// extending slice
	arr[0], arr[2] = 0, 2
	fmt.Println("arr=", s1)
	s1 = arr[2:6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 = s1[3:5]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	// ptr 指向开头
	// len
	// cap 容量 从ptr开始到整个arr，不超过cap都可以向后扩展，不可以向前扩展
	// s[i]不可以超过len(s)，但可以向后扩展到cap(s)
}
