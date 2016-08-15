// All material is licensed under the Apache License Version 2.0, January 2016
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	temp := Test(rand.Intn(10))
	fmt.Printf("%d", temp)

}

func Test(x int) int {
	for i := 0; i < 10000; i++ {
		y := i + x
		x = 2 * y
	}
	return x
}
