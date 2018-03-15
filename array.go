package main

import "fmt"

func main(){

	arr :=[...] int {1,2,3,4,5}

	for index, value := range arr {
		fmt.Printf("arr[%d]=%d \n", index, value)
	}

	for index := 0; index < len(arr); index++ {
		fmt.Printf("arr[%d]=%d \n", index, arr[index])
	}

}