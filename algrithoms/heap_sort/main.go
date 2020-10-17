package main

import "fmt"

func main() {

	var max = true
	var s = []int{9, 8, 5, 4, 2, 3, 7, 6, 1}
	s = HeapSort(s, true) //大根堆排序
	n := len(s)
	for i := 0; i < n; i++ {
		_, s = Pop(s, max)
		if len(s) > 0 {
			CheckHeap(s)
		} else {
			break
		}
	}
}

//构造堆
func HeapBuild(src []int, max bool) (dest []int) {
	var n = len(src)
	for i := n/2 - 1; i >= 0; i-- {
		dest = HeapAdjust(src, i, n-1, max)
	}
	return
}

//堆排序(大根堆)
//算法：从后往前进行堆调整
//src=排序切片 max=大根堆/小根堆
func HeapSort(src []int, max bool) (dest []int) {
	var n = len(src)
	if n == 1 {
		return src
	}
	//构造堆
	src = HeapBuild(src, max)
	//调整排序
	for i := n - 1; i > 0; i-- {
		Swap(src, 0, i)                     //末端元素跟对顶元素交换
		dest = HeapAdjust(src, 0, i-1, max) //重新调整
	}
	return src
}

//调整(src=排序切片 i=当前节点位置 n=元素个数 max=大根堆/小根堆)
func HeapAdjust(src []int, i, n int, max bool) (dest []int) {
	var m = src[i] //临时变量
	for j := i*2 + 1; j <= n; {

		if max { //大根堆
			if j < n && src[j] > src[j+1] { //找左右孩子中最大值
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
			if j < n && src[j] < src[j+1] { //找左右孩子中最小值
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
	return src
}

//交换
func Swap(src []int, i, j int) {
	m := src[i]
	src[i] = src[j]
	src[j] = m
}

//取出头部元素，将末尾元素调位至顶点进行重新排序，返回取出的值和排序切片
func Pop(src []int, max bool) (n int, dest []int) {

	var size = len(src)
	if size == 0 {
		return
	}
	if size <= 1 {
		n = src[0]
		src = nil
		return
	}
	n = src[0]           //取出堆顶元素
	src[0] = src[size-1] //将末端元素移动至堆顶
	dest = src[:size-1]  //重新切片(切掉末端元素位置)
	HeapSort(src, max)   //重新排序
	return n, dest
}

//尾部插入元素，返回排序切片
func Push(src []int, n int, max bool) {
	src = append(src, n) //将新增元素append到切片末尾
	HeapSort(src, max)   //重新排序
}

func CheckHeap(s []int) {

	if CheckRootHeap(s, false) {
		fmt.Printf("小根堆[YES] %+v \n", s)
	} else {
		if CheckRootHeap(s, true) {
			fmt.Printf("大根堆[YES] %+v \n", s)
		} else {
			fmt.Printf("大/小根堆[NO] %+v \n", s)
		}
	}
}

//检查是否大根堆
//规则：从最后一个切片元素往前跟自己的父节点(i-1)/2的值比对，如果大于父节点则返回false
func CheckRootHeap(s []int, max bool) (ok bool) {
	n := len(s)
	if n == 0 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		pi := (i - 1) / 2 //父节点位置
		var compare bool
		if max {
			compare = (s[i] > s[pi]) //判断当前节点是否比父节点大
		} else {
			compare = (s[i] < s[pi]) //判断当前节点是否比父节点小
		}
		if compare {
			return false
		}
	}
	return true
}
