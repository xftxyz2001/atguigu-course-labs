package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.ParseInt("66hello", 10, 64)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("i: %v\n", i)
	}
	// fmt.Printf("i: %v\n", i)
	// fmt.Printf("err: %v\n", err)

	/* func ParseBool(str string) (bool, error) {
		switch str {
		case "1", "t", "T", "true", "TRUE", "True":
			return true, nil
		case "0", "f", "F", "false", "FALSE", "False":
			return false, nil
		}
		return false, syntaxError("ParseBool", str)
	} */
	b, err2 := strconv.ParseBool("t")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("b: %v\n", b)
	}

}
