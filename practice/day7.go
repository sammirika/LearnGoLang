package practice

/**
给你一个非负整数数组 nums 。在一步操作中，你必须：
选出一个正整数 x ，x 需要小于或等于 nums 中 最小 的 非零 元素。
nums 中的每个正整数都减去 x。
返回使 nums 中所有元素都等于 0 需要的 最少 操作数。
示例 1：

输入：nums = [1,5,0,3,5]
输出：3
解释：
第一步操作：选出 x = 1 ，之后 nums = [0,4,0,2,4] 。
第二步操作：选出 x = 2 ，之后 nums = [0,2,0,0,2] 。
第三步操作：选出 x = 2 ，之后 nums = [0,0,0,0,0] 。
示例 2：

输入：nums = [0]
输出：0
解释：nums 中的每个元素都已经是 0 ，所以不需要执行任何操作。
*/

func minimumOperations(nums []int) (ans int) {
	s := [101]bool{true}
	for _, x := range nums {
		if !s[x] {
			s[x] = true
			ans++
		}
	}
	return
}

//Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
//
//for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
//for key, value := range oldMap {
//    newMap[key] = value
//}
//for 和 for range有什么区别?
//主要是使用场景不同
//for可以
//遍历array和slice
//遍历key为整型递增的map
//遍历string
//for range可以完成所有for可以做的事情，却能做到for不能做的，包括
//遍历key为string类型的map并同时获取key和value
//遍历channel

func RangeCase() {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		println(s[i])
	}
	// 忽略 index。
	for _, c := range s {
		println(c)
	}
	// 忽略全部返回值，仅迭代。
	for range s {

	}

	m := map[string]int{"a": 1, "b": 2}
	// 返回 (key, value)。
	for k, v := range m {
		println(k, v)
	}
}
