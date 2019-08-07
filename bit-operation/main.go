package main

import (
	"fmt"
)

func main() {
	//原码 0000 0011
	//反码 0000 0011
	//补码 0000 0011
	var a int8 = 3
	//原码 0000 0010
	//反码 1111 1101
	//补码 1111 1110
	var b int8 = -2

	// 左移和右移 操作的是原码, 结果为乘以2的n次方
	fmt.Println(a << 2) // 0000 1100
	fmt.Println(b << 2) // 0000 1000
	// 右移， 正数的最小值是0， 负数的最小值是-1
	fmt.Println(a >> 2) // 0000 0000
	fmt.Println(b >> 2) // 0000 0000

	//位运算符使用的是补码
	fmt.Println(a | b) // 1111 1111 -> 1111 1110 -> 0000 0001
	fmt.Println(a & b) // 0000 0010 -> 0000 0010 -> 0000 0010
}
