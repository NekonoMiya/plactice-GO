package main

import "fmt"

func pl(x int, y int) int {
	return x+y
}

func main(){
	fmt.Println(pl(3,4))

	a, b := i("hello","world")
	fmt.Println(a,b)

	c, d := kw(4,2)
	fmt.Println("剰余演算",c,d)
}

func i(x,y string) (string, string) {
	return y, x
}

func kw(x, y int) (a, b int){
	a = x*y
	b = x/y
	return 
}