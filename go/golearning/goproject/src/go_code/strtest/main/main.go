package main

import "strings"

func main() {
	i := strings.Count("aaaa", "aa")
	println(i)
	s := "a" + "b" + "c"
	println("abc" == s)
}
