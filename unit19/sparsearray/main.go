package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main()  {
	// 1、创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑
	chessMap[2][3] = 2 // 蓝

    // 2、输出原始数组
    fmt.Println("原始数组：")
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3、转成稀疏数组
	// -> 遍历 chessMap ，如果发现有一个元素的值不为 0，创建一个 node 结构体
	// -> 将其放入到对应的切片即可
    var sparseArr []ValNode
    // 记录行和列
    valNode := ValNode {
        row : 11,
        col : 11,
        val : 0,
    }
    sparseArr = append(sparseArr, valNode)
    for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
                // 创建一个 ValNode 值结点
                valNode := ValNode {
                    row : i,
                    col : j,
                    val : v2,
                }
                sparseArr = append(sparseArr, valNode)
            }
		}
    }
    
    // 输出稀疏数组
    fmt.Println("当前稀疏数组：")
    for i, valNode := range sparseArr {
        fmt.Printf("%d -> %d\t%d\t%d\n", i, valNode.row, valNode.col, valNode.val)
    }
    // 存盘
    // 打开文件恢复

    // 创建一个原始数组
    var chessMap2 [11][11]int
    // 遍历稀疏数组
    for i, valNode := range sparseArr {
        // 跳过第一行
        if i != 0 {
            chessMap2[valNode.row][valNode.col] = valNode.val
        }
    }
    fmt.Println("恢复后数组：")
    for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}