package main

import (
	"fmt"
	"math"
)

/*
算法描述：翻转数值（假设给出的数值大于0，例如 123 翻转后321）
解题思路：

*/

func main() {
	var num int32 = 2147483647 //前提条件：正整数
	fmt.Printf("reverse %d = %d \n", num, rotateNumber(num))
}

func rotateNumber(x int32) int32 {

	var y int64 //先用int64存储，否则无法跟MaxInt32/MinInt32进行比较
	for {
		if x == 0 {
			break
		}
		y = y*10 + int64(x%10)
		x = x / 10
	}

	fmt.Printf("y = int64(%d) \n", y)
	if y > math.MaxInt32 || y < math.MinInt32 { // -2147483648 ~ 2147483647
		return 0
	}
	return int32(y)
}
