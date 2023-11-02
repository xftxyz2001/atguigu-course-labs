package main

import (
	"fmt"
	"mymodule/src/go_code/factorytest/model"
)

func main() {
	s := model.NewStudent("tom", 18)
	fmt.Println(s)
}
