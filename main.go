package main

import (
	"fmt"

	"gomponents/maths"
)

func main() {
	n, err := maths.ClosestPowerOf2(1536); 
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}