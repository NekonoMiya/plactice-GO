package main 
 
import "fmt"

func main(){
	array := [6]int{1,2,3,4,5,6}
	length := len(array)

	fmt.Println(len(array))

	s := array[2:4]
	length = len(s)
	capacity := cap(s)
	fmt.Println(s, length,capacity)


}