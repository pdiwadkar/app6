package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	PowerOfTwo()
}

func PowerOfTwo() {

	var result int64
	var power int
	fmt.Println("What power of two?")
	fmt.Scanf("%d", &power)
	if power < 0 {
		fmt.Println("Power must be >= 0")
		os.Exit(1)
	}
	result = int64(math.Pow(2, float64(power)))
	fmt.Printf("%s %d %s %d\n", "Two to the power of ", power, " is ", result)

}
