/*
 * @Author: lifangdi_i
 * @Date: 2019-12-13 11:40:40
 * @LastEditors: lifangdi_i
 * @LastEditTime: 2019-12-13 18:16:06
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 可以省略起始条件、递增条件、结束条件

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 省略递增条件 相当于while（go里没有while）
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

// 死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(342424),
		convertToBin(0),
		printFile("abc.txt"),
	)
}
