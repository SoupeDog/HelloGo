package 基本语法

import (
	"fmt"
	"strconv"
)

//P_的事情很无奈,Go 里面 "public" 方法得字母开头大写 后不赘述
func P_声明变量() {
	//1. 写法1
	var 强迫症写法 string = "这是强迫症格式"

	//2. 写法2
	index1 := 0

	//3. 写法3
	index2, index3 := 1, 2

	//4. 写法4  (常量)   iota 这种东西弄谷歌这是得多喜欢语法糖！！！
	const (
		constVar1 = iota // 从 0 自增
		constVar2
		constVar3
	)

	fmt.Println("写法1：" + 强迫症写法)
	fmt.Println("写法2：" + strconv.Itoa(index1))
	fmt.Println("写法3：" + strconv.Itoa(index2) + "  " + strconv.Itoa(index3))
	fmt.Println("写法3：" + strconv.Itoa(constVar1) + "  " + strconv.Itoa(constVar2) + "  " + strconv.Itoa(constVar3))

}
