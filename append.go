package main

import "fmt"

func main(){
	slice := make([]int, 0, 3)
	show("元のslice", slice)

	apslice := append(slice, 1)
	show("１を追加したslice", apslice)

	apslice = append(apslice, 2, 3)
	show("２と3を追加したslice", apslice)

}


func show (s string, x []int){
	fmt.Println(s, "length=", len(x), "capacity=", cap(x), "スライス=", x)
}