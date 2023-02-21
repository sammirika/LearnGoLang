package practice

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
	//使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	//    delete(map, key)
	delete(scoreMap, "小明")
	fmt.Println(scoreMap)

	// 按顺序遍历
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var orderMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		orderMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range orderMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, orderMap[key])
	}
}
