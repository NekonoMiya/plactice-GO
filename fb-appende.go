package main 

import (
	"fmt"
	"strconv"
)


func main(){

	slice := make([]string, 0, 30)

	apslice := append(slice)
	s := ""

	for i:=1; i<=30; i++ {
		if i%3==0&&i%5==0 {
			apslice = append(apslice, "foobar")
		}else if i%3==0 {
			apslice = append(apslice, "foo")
		}else if i%5==0 {
			apslice = append(apslice, "bar")
		}else{
			s = strconv.Itoa(i)
			apslice = append(apslice, s)
		}
		
	}

	fmt.Print(apslice)

} 
