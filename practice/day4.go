package practice

import (
	"fmt"
	"sort"
)

type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}

// 打印嵌套结构体

func PrintStruct() {
	user1 := User{
		Name:   "nessy",
		Gender: "gril",
		Address: Address{
			Province: "Chongqing",
			City:     "Yuzhong",
		},
	}
	fmt.Printf("User1=%#v\n", user1)

	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
	// 删除操作
}

//Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func Inherit() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

//删除map类型的结构体
func DelStruct() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Printf("删除前")
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Printf("删除后")
	fmt.Println(ce)
}

// 实现map有序输出
func SortMap() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	//利用数组排序
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	// 对数组排序
	sort.Ints(sli)
	fmt.Println(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}

}
