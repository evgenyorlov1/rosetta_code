package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func pi(n int) float64 {
	var count int
	for i := 1; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if math.Hypot(x, y) <= 1 {
			count++
		}
	}
	return 4 * float64(count) / float64(n)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int

	fmt.Println("Enter number of throws: ")
	fmt.Scan(&n)
	fmt.Println("Pi estimation: ", pi(n))
}
