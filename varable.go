package main

import "fmt"

//声明变量number为int数据，此时默认复制为0
var number int 	

// 定义多个变量
var(
	name string
	age int
	sex bool
)

func main(){

	//直接声明变量number赋值为2
	number := 2
	fmt.Println(number);

	name = "Neo"
	fmt.Println(name);
	
	url := "http://www.netkiller.cn"
	fmt.Println(url);

}