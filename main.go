package main

import (
	"LearnGoLang/practice"
	"fmt"
)

func main() {
	fmt.Println()
	practice.GetSum()
	b := [5]int{1, 3, 5, 8, 7}
	practice.GetIndex(b, 8)
	practice.ShowSlice()
	practice.ZhiZhen()
	practice.MapCase()
	practice.PrintStruct()
	practice.DelStruct()
	practice.SortMap()
	practice.Inherit()
	practice.SwitchCase()
	practice.RangeCase()
	practice.SwapCase()
}
