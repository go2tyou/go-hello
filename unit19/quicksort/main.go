package main

import (
	"fmt"
)

// 快速排序
// left  -> 数组左边的下标
// right -> 数组右边的下标
// arr   -> 要排序的数组
func QuickSort(left int, right int, arr *[6]int) {
	l := left
	r := right
	pivot := arr[(left + right) / 2]
	// for => 比 pivot 小放左边，比 pivot 大方右边
	for ; l < r; {
		// 从 pivot 左边找到 > pivot 的值
		for ; arr[l] < pivot; {
			l++
		}
		// 从 pivot 右边找到 < pivot 的值
		for ; arr[r] > pivot; {
			r--
		}
		// 表明本次分解任务完成
		if l >= r {
			break
		}
		// 交换
		arr[l], arr[r] = arr[r], arr[l]
		// 优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	fmt.Println("交换后 =>", *arr)

	// 再移动一下
	if l == r {
		l++
		r--
	}

	// 向左递归
	if left < r {
		QuickSort(left, r, arr)
	}

	// 向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main()  {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	fmt.Println("main before =>", arr)
	QuickSort(0, 5, &arr)
	fmt.Println("main after =>", arr)

}