package main

import (
	
	"fmt"
	"math"
)

var x = 0.0001

func Looping() (r string) {
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	
	return "Code.education Rocks!"
}

func main() {
	
	fmt.Println(Looping())
	
}