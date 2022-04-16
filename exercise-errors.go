package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	z := float64(1)

	if x < 0 {
		return z, ErrNegativeSqrt(x)
	}

	var t float64

	for {
		//Check how close z² is to x
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-6 {
			break
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
https://stackoverflow.com/questions/27474907/why-would-a-call-to-fmt-sprinte-inside-the-error-method-result-in-an-infinit
*/
