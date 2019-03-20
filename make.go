package main 

import "fmt"

func main(){
	sa := make([]int, 5)

	show("a",sa)

	sb := make([]int, 0, 5)
	show("b", sb)

	sc := sb[0:2]
	show("c", sc)

	sd := sb[1:5]
	show("d", sd)

}

func show (s string, x []int){
	fmt.Println(s, "length=", len(x), "capacity=", cap(x), "s=", x)
}