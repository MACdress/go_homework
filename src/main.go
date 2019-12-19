package src

import "fmt"

/**
冒泡排序
*/
func BubbleSort(slice []int) {
	if len(slice) > 0 {
		for i := 0; i < len(slice)-1; i++ {
			//标识位，减少无用的循环
			tag := true
			for j := 0; j < len(slice)-i-1; j++ {
				if slice[j] > slice[j+1] {
					temp := slice[j]
					slice[j] = slice[j+1]
					slice[j+1] = temp
					tag = false
				}
			}
			if tag {
				break
			}
		}
	}
}

/**
选择排序
优化：一次性既找出最大的又找出最小的
*/
func SelectSort(slice []int) {
	if len(slice) > 0 {
		left := 0
		right := len(slice) - 1
		for ; left < right; {
			min := left
			max := right
			for i := min; i <= right; i++ {
				if slice[min] > slice[i] {
					min = i
				}
				if slice[max] < slice[i] {
					max = i
				}
			}
			//查找的区间的数大小一样
			if min == max {
				break
			}
			if min != left {
				temp := slice[left]
				slice[left] = slice [min]
				slice[min] = temp
			}

			//如果最大的数的下标是最左边的数，则现在最大的数的下标是min
			if max == left {
				max = min
			}
			if max != right {
				temp := slice[max]
				slice[max] = slice [right]
				slice[right] = temp
			}
			right--
			left++
		}
	}
}

/**
插入排序
这里选择了折半插入
*/
//func InsertSort(slice []int) {
//	if len(slice) > 0 {
//		for i := 0; i < len(slice)-1; i++ {
//			//标识位，减少无用的循环
//			j := 0
//			index := (j + i) / 2
//			while(true)
//			{
//				if
//			}
//			if min != i {
//				temp := slice[i]
//				slice[i] = slice [min]
//				slice[min] = temp
//			}
//		}
//	}
//}

/**
希尔排序
这里选择了折半插入
*/
//func ShellSort(slice []int) {
//	if len(slice) > 0 {
//		for i := 0; i < len(slice)-1; i++ {
//			//标识位，减少无用的循环
//			j := 0
//			index := (j + i) / 2
//			while(true)
//			{
//				if
//			}
//			if min != i {
//				temp := slice[i]
//				slice[i] = slice [min]
//				slice[min] = temp
//			}
//		}
//	}
//}
//
///**
//堆排序
//这里选择了折半插入
//*/
//func HeapSort(slice []int) {
//	if len(slice) > 0 {
//		for i := 0; i < len(slice)-1; i++ {
//			//标识位，减少无用的循环
//			j := 0
//			index := (j + i) / 2
//			while(true)
//			{
//				if
//			}
//			if min != i {
//				temp := slice[i]
//				slice[i] = slice [min]
//				slice[min] = temp
//			}
//		}
//	}
//}
//
///**
//归并排序
//这里选择了折半插入
//*/
//func MergeSort(slice []int) {
//	if len(slice) > 0 {
//		for i := 0; i < len(slice)-1; i++ {
//			//标识位，减少无用的循环
//			j := 0
//			index := (j + i) / 2
//			while(true)
//			{
//				if
//			}
//			if min != i {
//				temp := slice[i]
//				slice[i] = slice [min]
//				slice[min] = temp
//			}
//		}
//	}
//}
//
///**
//快速排序
//这里选择了折半插入
//*/
//func QuickSort(slice []int) {
//	if len(slice) > 0 {
//		for i := 0; i < len(slice)-1; i++ {
//			//标识位，减少无用的循环
//			j := 0
//			index := (j + i) / 2
//			while(true)
//			{
//				if
//			}
//			if min != i {
//				temp := slice[i]
//				slice[i] = slice [min]
//				slice[min] = temp
//			}
//		}
//	}
//}
func main() {
	slice := []int{3, 2, 5, 8, 4, 7, 6, 9}
	SelectSort(slice)
	for _, data := range slice {
		fmt.Print(data, "  ")
	}
	fmt.Println()
}
