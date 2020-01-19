/*
 * @Date: 2020-01-19 15:04:31
 * @LastEditors  : lifangdi
 * @LastEditTime : 2020-01-19 15:36:04
 */
package main

import (
	"fmt"
)

// 数组是值类型
// [10]int 和 [20]int是不同类型
// 调用func f(arr [10]int)会拷贝数组（值传递）
// go语言中一般不直接使用数组
func printArray(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	// 一维数组定义
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 5, 6, 3}

	// 二维数组 行 列
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 遍历数组
	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Println(arr3[i])
	// }

	// 遍历数组range关键字, 可以获得下标和值
	// for i, v := range arr3 {
	// 	fmt.Println(i, v)
	// }

	fmt.Println("printArray arr1")
	printArray(&arr1)
	fmt.Println("printArray arr3")
	printArray(&arr3)
	fmt.Println("arr1, arr3")
	fmt.Println(arr1, arr3)
}
