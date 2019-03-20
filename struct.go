package main

import "fmt"
type ju struct {
	kata string
	m float32
	x float32
	y float32
	z int
}

func main(){
	box := ju{"箱", 32.8, 44.3, 23.5, 14}
	xbox := ju{"円柱", 44.6, 23.5, 11.9, 16}

	fmt.Println(box)
	fmt.Println(xbox)

	fmt.Println(box.kata, xbox.kata)

	box.kata = "四角柱"

	fmt.Println(box.kata, xbox.kata)

}