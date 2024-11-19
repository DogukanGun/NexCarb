package utils

import "fmt"

func RunWithHandlingError(err error) {
	if err != nil {
		fmt.Printf("error cached: %s", err.Error())
	}
}
