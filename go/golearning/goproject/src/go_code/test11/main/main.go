package main

import "fmt"

func TypeJudge(items ...interface{}) []string {
	var typeSlice []string
	for _, item := range items {
		switch item.(type) {
		case bool:
			typeSlice = append(typeSlice, "bool")
		case string:
			typeSlice = append(typeSlice, "string")
		case int, int8, int16, int32, int64:
			typeSlice = append(typeSlice, "int")
		case float32, float64:
			typeSlice = append(typeSlice, "float")
		case nil:
			typeSlice = append(typeSlice, "nil")
		default:
			typeSlice = append(typeSlice, "unknown")
		}
	}
	return typeSlice
}

func main() {
	var a interface{}
	var i int

	typeSlice := TypeJudge(a, i)
	fmt.Println("typeSlice = ", typeSlice) // typeSlice =  [nil int]

}
