package main

import "fmt"

func main(){
	i:=3
	j:=4

	fmt.Println(kisuu(i))
	fmt.Println(kisuu(j))
}

func kisuu(x int) string {
	if x%2==0 {
		return "偶数"
	}
	return "奇数"
}