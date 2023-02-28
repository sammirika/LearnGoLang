package practice

import "fmt"

/**
闭包的应该都听过，但到底什么是闭包呢？
闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)
“官方”的解释是：所谓“闭包”，指的是一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也是该表达式的一部分。
*/

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func BiBao() {
	c := a()
	c()
	c()
	c()

	a() //不会输出i
}

// 递归hanshu

func JieCheng(i int) int {
	if i == 1 {
		return i
	}
	return i * JieCheng(i-1)
}

func ShowNum() {
	i := 7
	fmt.Printf("i的阶乘为%d", JieCheng(i))
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
