package practice

import "fmt"

// make和new的区别
/**
*   1.二者都是用来做内存分配的。
    2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
    3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
func createNew() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)
}

func createMake() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}

//程序定义一个int变量num的地址并打印
//将num的地址赋给指针ptr，并通过ptr去修改num的值
func ZhiZhen() {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}

// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
//    make(map[KeyType]ValueType, [cap])
func MapCase() {
	scoreMap := make(map[string]int, 8)
	scoreMap["nessy"] = 100
	scoreMap["dandan"] = 90
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["nessy"])
	fmt.Printf("scoreMap type is %T", scoreMap)
	//map也支持在声明的时候填充元素，例如：
	userInfo := map[string]string{
		"lihaoyu":   "zhenjiang",
		"sunyalong": "chongqing",
	}
	fmt.Println(userInfo)

	//Go语言中有个判断map中键是否存在的特殊写法，格式如下:
	//    value, ok := map[key]
	v, ok := scoreMap["nessy"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没这个人")
	}

	//Go语言中使用for range遍历map。
}
