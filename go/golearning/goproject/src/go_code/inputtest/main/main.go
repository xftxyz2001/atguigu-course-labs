package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var sal float32
	var isPass bool

	// fmt.Print("请输入姓名:")
	// fmt.Scanln(&name)
	// fmt.Print("请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Print("请输入薪水:")
	// fmt.Scanln(&sal)
	// fmt.Print("请输入是否通过考试:")
	// fmt.Scanln(&isPass)

	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("name=%v age=%v sal=%v isPass=%v", name, age, sal, isPass)
}
