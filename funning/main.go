package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomNumber := rand.IntN(10)

	fmt.Println(randomNumber)
}
