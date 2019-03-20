package main 

import "fmt"

func main(){
	i := 1
	j := 2
	fmt.Println(i, j)

	pi := &i
	fmt.Println(*pi)

	pj := &j
	fmt.Println(*pj)

	fmt.Println(pi)

	*pi = 3
	fmt.Println(*pi)

	*pj = 5
	fmt.Println(*pj)

}