package main 

import (
	"fmt"
)

var a, b, c int = 1,2,3

const pi = 3.1415

func main(){
	i := 0.2222221
	j := int(i)

	pi = 888    //エラー

	fmt.Println(j)
	fmt.Println(a+b+c)
}
