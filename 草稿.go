package main

import (
	"fmt"
	"math"
)

func main() {
	size := 0.8
	a := 0.19
	stepX := 0.025
	stepY := 4 * stepX // 一行的高度大概能竖着排四个`*`

	// 坐标系 从上到下y是递减的，从左到右x是递增的
	for y := size; y >= -size; y -= stepY {
		for x := -size; x <= size; x += stepX {
			// 拟合心形函数
			// math.Pow(x,n) 求x的n次方
			res := math.Pow(x*x+y*y-a, 3) - x*x*y*y*y
			if res <= 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}