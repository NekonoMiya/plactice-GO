package main 

import "fmt"

func main(){
	for i:=1; i<=30; i++ {
		if i%3==0&&i%5==0 {
			fmt.Print("foobar\n") 
		}else if i%3==0 {
			fmt.Print("foo\n")
		}else if i%5==0 {
			fmt.Print("bar\n")
		}else{
			fmt.Println(i)
		}
		
	}
} 
