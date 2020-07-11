// You can edit this code!
// Click here and start typing.
package main

import "fmt"

/**
* https://my.oschina.net/u/2612999/blog/908114
* 这篇文章对range语法糖的深入分析
**/

func IndexArray() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := range a {
		a[3] = 100
		if i == 3 {
			fmt.Println("IndexArray", i, a[i])
		}
	}
}

func IndexValueArray() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range a {
		a[3] = 100
		if i == 3 {
			fmt.Println("IndexValueArray", i, v)
		}
	}
}

func IndexValueArrayPtr() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, v := range &a {
		a[3] = 100
		if i == 3 {
			fmt.Println("IndexValueArrayPtr", i, v)
		}
	}
}

type Foo struct {
	bar string
}

func main() {
	v := []int{1, 2, 3}
	for _, val := range v {
		v = append(v, val)
	}
	fmt.Println("aaa")

	IndexArray()
	IndexValueArray()
	IndexValueArrayPtr()

	fmt.Println("====================================")

	// 复制整个数组
	var a [10]int
	acopy := a
	a[0] = 10
	fmt.Println("a", a)
	fmt.Println("acopy", acopy)
	// 只复制了 slice 的结构体，并没有复制成员指针指向的数组
	s := make([]int, 10)
	s[0] = 10
	scopy := s
	fmt.Println("s", s)
	fmt.Println("scopy", scopy)
	// 只复制了 map 的指针
	m := make(map[string]int)
	mcopy := m
	m["0"] = 10
	fmt.Println("m", m)
	fmt.Println("mcopy", mcopy)

	fmt.Println("====================================")

	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	// 在Go的for…range循环中，Go始终使用值拷贝的方式代替被遍历的元素本身，
	// 简单来说，就是for…range中那个value，是一个值拷贝，而不是元素本身。
	// 这样一来，当我们期望用&获取元素的地址时，实际上只是取到了value这个临时变量的地址，
	// 而非list中真正被遍历到的某个元素的地址。而在整个for…range循环中，value这个临时变量会被重复使用，
	// 所以，在上面的例子中，list2被填充了三个相同的地址，其实都是value的地址。
	// 而在最后一次循环中，value被赋值为{c}。因此，list2输出的时候显示出了三个&{c}。

	list2 := make([]*Foo, len(list))
	for i, value := range list {
		list2[i] = &value
	}
	list3 := make([]*Foo, len(list))
	for i := range list {
		list3[i] = &list[i]
	}

	fmt.Println(list[0], list[1], list[2])
	fmt.Println(list2[0], list2[1], list2[2])
	fmt.Println(list3[0], list3[1], list3[2])
}
