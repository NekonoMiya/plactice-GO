package main 

import (
	"fmt"
	"strconv"
)


func main(){

	slice := make([]string, 0, 30)

	apslice := append(slice)

	for i:=1; i<=30; i++ {
		if i%3==0&&i%5==0 {
			apslice = append(apslice, "foobar")
		}else if i%3==0 {
			apslice = append(apslice, "foo")
		}else if i%5==0 {
			apslice = append(apslice, "bar")
		}else{
			apslice = append(apslice, strconv.Itoa(i))
		}
		
	}

	fmt.Print(apslice)

} 
