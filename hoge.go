package main

import "fmt"

func pl(x int, y int) int {
	return x+y
}

func main(){
	fmt.Println(pl(3,4))
	a, b := i("hello","world")
	fmt.Println(a,b)
}

func i(x,y string) (string, string) {
	return y, x
}