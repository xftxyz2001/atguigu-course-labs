package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	No    int
	Name  string
	Score float64
}

type Stus []Stu

func (a Stus) Len() int           { return len(a) }
func (a Stus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Stus) Less(i, j int) bool { return a[i].Score > a[j].Score }

func main() {
	a := Stus{
		Stu{
			No:    1,
			Name:  "stu01",
			Score: 90.0,
		},
		Stu{
			No:    2,
			Name:  "stu02",
			Score: 80.0,
		},
		Stu{
			No:    3,
			Name:  "stu03",
			Score: 70.0,
		},
	}
	sort.Sort(a)

	for _, v := range a {
		fmt.Println(v.No, v.Name, v.Score)
	}
}
