package practice

import "fmt"

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
}
