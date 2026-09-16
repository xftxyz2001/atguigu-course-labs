package main

import (
	"fmt"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int           { return len(hs) }
func (hs HeroSlice) Swap(i, j int)      { hs[i], hs[j] = hs[j], hs[i] }
func (hs HeroSlice) Less(i, j int) bool { return hs[i].Age < hs[j].Age }

func main() {
	// intSlice := []int{0, -1, 10, 7, 90}
	// sort.Ints(intSlice)
	// fmt.Println("intSlice=", intSlice)

	heroSlice := []Hero{
		{"宋江", 23},
		{"卢俊义", 30},
		{"林冲", 25},
		{"吴用", 28},
	}
	sort.Sort(HeroSlice(heroSlice))
	fmt.Println("heroSlice=", heroSlice)

}
