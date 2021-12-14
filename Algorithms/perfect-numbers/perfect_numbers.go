package perfect

import "errors"

// Define the Classification type here.
type Classification string

const (
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
	ClassificationDeficient Classification = "deficient"
)

var ErrOnlyPositive = errors.New("not a natural number")

func Classify(n int64) (Classification, error) {

	if n > 0 {
		sum := int64(0)
		for i := int64(1); i < n; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		if sum == n {
			return ClassificationPerfect, nil
		} else if sum > n {
			return ClassificationAbundant, nil
		} else {
			return ClassificationDeficient, nil
		}
	} else {

		return "", ErrOnlyPositive
	}

	panic("Please implement the Classify function")
}
