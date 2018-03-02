package 基本语法

import (
	"fmt"
	"strconv"
)

var 全局变量 = 10

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

func P_多返回值() (int, string) {
	return 1, "第二个返回值"
}

func P_多返回值接收(第一个参数 int, 第二个参数 string) {
	fmt.Print("第一个:")
	fmt.Println(第一个参数)
	fmt.Print("第二个:")
	fmt.Println(第二个参数)
}

func P_打印全局变量() {
	fmt.Println(全局变量)
}

func P_局部变量覆盖全局变量() {
	全局变量 := 5
	fmt.Println(全局变量)
}
