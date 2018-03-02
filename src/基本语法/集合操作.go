package 基本语法

import (
	"fmt"
	"strconv"
	"container/list"
)

func P_集合操作() {
	// 1. map 操作
	myMap := make(map[int]string)
	myMap[1] = "张三"
	myMap[2] = "李四"

	for mapKey := 0; mapKey < 3; mapKey++ {
		itemVal, flag := myMap[mapKey]
		if flag { // () 括号写上提示多余贼气！
			fmt.Println(itemVal)
		} else {
			fmt.Println("不存在")
		}
	}

	// 2. 定长数组操作           plan B     array := [10]string{0:"a", 2:"b", 3:"c"}   下标0 2 3 特定值其余默认值
	array := [3]string{}
	for i := 0; i < 3; i++ {
		array[i] = strconv.Itoa(i)
	}
	fmt.Println(array)


	// 3. 切片
	slice := array[0:1] // 切取 array index 0~1 包括0 不包括1

	for i, item := range slice {
		fmt.Printf("slice[%d]=%s\n", i, item) // Fomat 没有 printfln ？
	}

	// 4 list 链表操作
	list := list.New()
	list.PushBack("我是第一个")
	list.PushBack("我是第二个")
	list.PushBack("我是第三个")

	//      循环初始化变量            判断条件             循环完成后执行
	for itemPoint := list.Front(); itemPoint != nil; itemPoint = itemPoint.Next() {
		fmt.Println(itemPoint.Value)
	}

}
