package main

import "fmt"

func main() {
	//	数组是固定长度
	//	声明固定长度的数组
	var ages [10]int
	for idx := 0; idx < 10; idx++ {
		ages[idx] = idx
	}
	//	初始化非固定长度的数组
	noLengthNames := [...]string{"I am a new gopher", "You are a new gopher"}
	fmt.Println(len(noLengthNames))
	//	切片是不固定长度
	//	使用make创建一个带有初始容量的切片
	names := make([]string, 10)
	names[0] = "I am a new gopher"
	//	初始一个切片
	newNames := []string{"I am a new gopher"}
	fmt.Println(len(newNames))
	//	从数组中得到一个切片
	var arr = []int{1, 2, 3, 4, 5, 6}
	sl := arr[2:5]
	fmt.Printf("%v\n", sl)

	//	区别
	//	实现
	// 数组是一个定长，有特定数据类型的数据结构
	// 切片是一个指向底层数组的动态指向
	// 传参
	// 数组是值传递(副本)
	// 切片是指针传递(自身)

	modifySlice(names)
	modifyArray(noLengthNames)

	fmt.Printf("%v\n", names)
	fmt.Printf("%v\n", noLengthNames)
}

func modifySlice(name []string) {
	name[0] = "change"
}

func modifyArray(name [2]string) {
	name[0] = "change"
}
