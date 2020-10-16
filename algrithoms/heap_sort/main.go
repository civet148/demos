package main

import "fmt"

func main() {

	var slice = []int{3, 7, 8, 6, 1, 30, 4, 50, 3, 2, 9, 11}

	slice = HeapSort(slice, false)
	CheckHeap(slice)
	slice = HeapSort(slice, true)
	CheckHeap(slice)
}

//堆排序(大根堆)
//算法：从后往前进行堆调整
func HeapSort(src []int, big bool) (dest []int) {
	var size = len(src)

	//构造堆
	for i := size/2 - 1; i >= 0; i-- {
		HeapAdjust(src, i, size-1, big)
	}

	//排序
	for i := size - 1; i > 0; i-- {
		Swap(src, 0, i)
		HeapAdjust(src, 0, i-1, big)
	}
	return src
}

//调整
func HeapAdjust(src []int, i, size int, big bool) {
	var m = src[i] //临时变量
	for j := i*2 + 1; j <= size; {

		if big { //大根堆
			if j < size && src[j] > src[j+1] { //找左右孩子中最大值
				j++
			}
			if m > src[j] {
				Swap(src, i, j)
				i = j
				j = 2*j + 1
			} else {
				break
			}
		} else { //小根堆
			if j < size && src[j] < src[j+1] { //找左右孩子中最小值
				j++
			}
			if m < src[j] {
				Swap(src, i, j)
				i = j
				j = 2*j + 1
			} else {
				break
			}
		}
	}
}

//交换
func Swap(src []int, i, j int) {
	m := src[i]
	src[i] = src[j]
	src[j] = m
}

//取出头部元素，将末尾元素调位至顶点进行重新排序，返回取出的值和排序切片
func Pop(src []int, big bool) (n int, dest []int) {

	return
}

//尾部插入元素，返回排序切片
func Push(src []int, n int, big bool) (dest []int) {

	return
}

func CheckHeap(s []int) {
	if CheckRootHeap(s, false) {
		fmt.Println("小根堆[YES]")
	} else {
		if CheckRootHeap(s, true) {
			fmt.Println("大根堆[YES]")
		} else {
			fmt.Println("大/小根堆[NO]")
		}
	}
	fmt.Printf("slice %+v \n", s)
}

//检查是否大根堆
//规则：从最后一个切片元素往前跟自己的父节点(i-1)/2的值比对，如果大于父节点则返回false
func CheckRootHeap(s []int, big bool) (ok bool) {
	size := len(s)
	if size == 0 {
		return
	}
	for i := size - 1; i >= 0; i-- {
		pi := (i - 1) / 2
		var compare bool
		if big {
			compare = (s[i] > s[pi])
		} else {
			compare = (s[i] < s[pi])
		}
		if compare {
			return false
		}
	}
	return true
}
