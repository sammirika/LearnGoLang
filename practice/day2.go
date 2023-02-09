package practice

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组
// 求元素和
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

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

//找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func GetIndex(a [5]int, target int) {
	fmt.Printf("下标分别为\n")
	for i := 0; i < len(a); i++ {
		temp := target - a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] == temp {
				fmt.Printf("%d,%d \n", i, j)
			}
		}
	}
}

// 切片
func ShowSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[2:8]
	var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
	var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
	var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
	var slice4 = arr[:len(arr)-1]
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)

}
