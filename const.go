package main

import "fmt"

//声明常量 const url
const url = "http://www.netkiller.cn"

const nickname string = "netkiller"

const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

const (
	New, Pending, Done = 1, 2, 3
	Canceled, Failed, Scuess = 4, 5, 6
	)

// 生成枚举值时候可以使用关键字：itoa, 值将依次递增
const (
	Sunday = iota 	// 0
	Monday        	// 1
	Tuesday       	// 2
	Wednesday		// 3
	​Thursday		 // 4
	Friday			// 5
	Saturday		// 6
	)

func main(){
	
	fmt.Println(url);

}