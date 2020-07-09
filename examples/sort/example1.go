package main

import (
	"fmt"
	"sort"
)

// https://yourbasic.org/golang/how-to-sort-in-go/

// Person custom struct
type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	// 3种排序规则

	// 1:3种基本类型的常规排序
	s := []int{10, 2, 5, 6}
	sort.Ints(s)
	fmt.Println(s)

	s2 := []float64{3.4, 2.3, 4.5, 7.8}
	sort.Float64s(s2)
	fmt.Println(s2)

	s3 := []string{"abc", "def", "aaa", "aac"}
	sort.Strings(s3)
	fmt.Println(s3)

	// 2:采用sort.Slice和sort.SliceStable进行排序
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)

	// 3:自定义数据类型，实现排序接口
	family2 := []Person{
		{"bAlice", 23},
		{"bDavid", 2},
		{"bEve", 2},
		{"bBob", 25},
	}
	sort.Sort(ByAge(family2))
	fmt.Println(family2)

}
