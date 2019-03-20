package main

import "fmt"

func main(){
	double := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	double [0][0] = "X"
	double [1][1] = "X"
	double [2][2] = "X"
	double [0][2] = "o"
	double [2][0] = "o"

	for i:=0; i<len(double); i++ {
		fmt.Println(double[i])
	}
}