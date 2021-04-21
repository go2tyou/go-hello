package main

import (
	"fmt"
)

// 走出迷宫
// i,j -> 表示对哪个点测试
func FindWay(mmap *[8][7]int, i, j int) bool {
	// 什么情况下找到了出路
	if mmap[6][5] == 2 {
		return true
	}
	// 继续找
	// 确保可以走
	if mmap[i][j] == 0 {
		// 假设这个点可以通过
		mmap[i][j] = 2
		// 策略 -> 下右上左
		if FindWay(mmap, i + 1, j) {
			return true
		} else if FindWay(mmap, i, j + 1) {
			return true
		} else if FindWay(mmap, i - 1, j) {
			return true
		} else if FindWay(mmap, i, j - 1) {
			return true
		} else {
			// 天坑
			mmap[i][j] = 3
			return false
		}
	} else {
		// 墙
		return false
	}
}

func main()  {
	// 先创建一个二维数组模拟迷宫
	// 1 -> 墙 
	// 0 -> 未走过
	// 2 -> ok
	// 3 -> 走过，不通
	var mmap [8][7]int
	// 建墙
	mmap[3][1] = 1
	mmap[3][2] = 1
	// mmap[1][2] = 1
	// mmap[2][2] = 1
	for i := 0; i < 7; i++ {
		mmap[0][i] = 1
		mmap[7][i] = 1
		mmap[i][0] = 1
		mmap[i][6] = 1
	}
	
	// 输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mmap[i][j], " ")
		}
		fmt.Println()
	}

	FindWay(&mmap, 1, 1)
	fmt.Println()
	fmt.Println("-------start-------")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mmap[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("-------end-------")
}