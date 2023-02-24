package practice

import (
	"fmt"
	"math"
)

//函数特点
//• 无需声明原型。
//    • 支持不定 变参。
//    • 支持多返回值。
//    • 支持命名返回参数。
//    • 支持匿名函数和闭包。
//    • 函数也是一种类型，一个函数可以赋值给变量。
//
//    • 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
//    • 不支持 重载 (overload)
//    • 不支持 默认参数 (default parameter)。

// 交换函数
func swap1(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func SwapCase() {
	x, y := 1, 2
	fmt.Println(x, y)
	swap1(&x, &y)
	fmt.Println(x, y)
}

/**
在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。

  func myfunc(args ...int) {    //0个或多个参数
  }

  func add(a int, args…int) int {    //1个或多个参数
  }

  func add(a int, b int, args…int) int {    //2个或多个参数
  }
*/

/**
函数返回
"_"标识符，用来忽略函数的某个返回值

Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。

返回值的名称应当具有一定的意义，可以作为文档使用。

没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。

直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/

func add(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2

	return
}

/**
匿名函数是指不需要定义函数名的一种函数实现方式。1958年LISP首先采用匿名函数。

在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。

匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
*/

func NiMing() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
}
