package main

import (
	"errors"
	"fmt"
)

func divide(f1 int, f2 int) (float64, error) {
	if f2 == 0 {
		return 0, errors.New("分母不能为0")
	} else {
		result := float64(f1) / float64(f2)
		return result, nil
	}
}

func main() {
	var result float64
	var err error

	result, err = divide(100, 3)

	if err == nil {
		fmt.Printf("result: %f\n", result)
	} else {
		fmt.Printf("%s\n", err.Error())
	}

	result, err = divide(100, 0)
	if err == nil {
		fmt.Printf("result: %f\n", result)
	} else {
		fmt.Printf("%s\n", err.Error())
	}
}
