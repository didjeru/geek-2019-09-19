package main

import (
	"fmt"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, errNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-2))
}
