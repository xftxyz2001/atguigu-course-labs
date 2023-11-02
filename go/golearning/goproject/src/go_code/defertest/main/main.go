package main

import "fmt"

func test(n int) int {
	fmt.Println("defer", n)
	n++
	return n

}

func test2(n *int) int {
	fmt.Println("defer", *n)
	*n++
	return *n

}

type student struct {
	Name string
}

func test3(stu student) {
	stu.Name = "jack"
	defer fmt.Println(stu.Name)
	stu.Name = "tom"
	defer fmt.Println(stu.Name)
}

func main() {
	var i int = 10
	i = test(i)
	fmt.Println(i)

	j := 10
	j = test2(&j)
	fmt.Println(j)

	stu := student{
		Name: "mary",
	}
	test3(stu)
	fmt.Println(stu.Name)
	
}
