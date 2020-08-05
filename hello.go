package main

import (
	"fmt"
)

func main()  {
	//i := 10
	//switch {
	//case i > 1:
	//	fmt.Println(1)
	//	fallthrough
	//case i > 2:
	//	fmt.Println(2)
	//case i > 3:
	//	fmt.Println(3)
	//default:
	//	fmt.Println(111)
	//}
	//var data = [10]int{1,2,3,4,5,6,7,8,9,0}
	//fmt.Println(data[1])
	var name = "zhaosy"
	var age = 1
	fmt.Println(name,age)
	_,age = test()
	fmt.Println(age)
}

func test() (string,int) {
	return "zhaosy",10
}