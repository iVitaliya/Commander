package utils

import (
	"fmt"
)

func Stringify[T any](value T) string {
	return fmt.Sprint(value)
}

func JoinArray[T any](array []T) string {
	var result string
	for i := 0; i < len(array); i++ {
		length := len(array)-1
		if i == length {
			result = result + fmt.Sprint(array[i])
		}

		result = result + fmt.Sprint(array[i]) + " "
	}

	return result
}

func AppendArray[T]() {

}