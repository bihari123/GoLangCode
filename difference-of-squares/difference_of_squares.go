package diffsquares

import "math"

func SquareOfSum(n int) int {
    return int(math.Pow(float64(n*(n+1)/2), 2.0))
	panic("Please implement the SquareOfSum function")
}

func SumOfSquares(n int) int {
	return (n*(n+1)*(2*n + 1))/6
	panic("Please implement the SumOfSquares function")
}

func Difference(n int) int {
    return int(math.Abs(float64((SumOfSquares(n) - SquareOfSum(n)))))
	panic("Please implement the Difference function")
}
