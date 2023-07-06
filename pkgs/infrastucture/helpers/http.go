package helpers

import (
	"fmt"

)

func AbortError(err error) {

	if err != nil {
		fmt.Println(err)
		// c.AbortWithStatus(http.StatusInternalServerError, err)

		// panic(err)
	}
}

func SecondMustBeNull[T any](data T, err error) T {
	AbortError(err)
	return data
}

func FirstMustBeNull[T any](data T, err error) T {
	AbortError(err)
	return data
}
