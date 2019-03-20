package main 

import "fmt"

func main(){
	array:=[6]int {1,2,3,4,5,6}

	fmt.Println(array)

	var s []int = array[1:5]
	var s2 []int = array[2:4]

	fmt.Println(s)
	fmt.Println(s2)

	s2[1] = 30
	fmt.Println(array, s, s2)

}