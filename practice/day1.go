package practice

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GetSum() {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n", sum)
}

func sum(a, b int) int {
	fmt.Println("sum函数中\na的值为%d", a)
	fmt.Println("sum函数中b的值为%d", b)
	num := a + b
	return num
}

func GetData() (int, int) {
	GetSum()
	return 100, 200
}

// 使用指针修改值,*操作符作为右值时，意义是取指针的值，作为左值时，也就是放在赋值操作符的左边时，
// 表示 a 指针指向的变量。其实归纳起来，*操作符的根本意义就是操作指针指向的变量。
// 当操作在右值时，就是取指向变量的值，当操作在左值时，就是将值设置给指向的变量。
func swap(a, b *int) {
	t := *a
	// ab交换值
	*a = *b
	*b = t
}

func day1() {
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a, b)
	c, _ := GetData()
	_, d := GetData()
	fmt.Println(c, d)
	fmt.Println(sum(a, b))
	fmt.Println(a == 200)
	fmt.Println(a == 300)

	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p \n", &cat, &str)

	var model = "chongqing zhonggaoyi nessy"
	// 取地址
	ptr := &model
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("ptr address: %p \n", ptr)
	// 指针取值
	value := *ptr
	fmt.Printf("value type is:%T\n", value)
	fmt.Printf("value is: %s\n", value)
	// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	x, y := 1, 2
	swap(&x, &y)
	fmt.Printf("x,y is: %d,%d \n", x, y)
	fmt.Println(math.Sin(45))
}

func printNessy() {
	println("hi nessy")
}
