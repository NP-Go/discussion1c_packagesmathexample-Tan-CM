package main

import (
	"fmt"
	"math"
)

func main() {
	q := 34.0
	r := 8.0

	testing()
	s := math.Mod(q, r)
	fmt.Println(s)
}
