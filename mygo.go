package main

import (
	"fmt"
)

func main() {
	//var i = 1
	//println(&i)
	//j := i
	//println(&j)
	//var a, b int
	//var c string
	//a, b, c = 1, 2, "3"
	//qq := 0x67 & 0x1F
	//println(qq)
	//println(a, b, c)
	//const (
	//	a = iota
	//	b
	//	c
	//	d = "d"
	//	e = "e"
	//	f = iota
	//	g
	//)
	//println(a, b, c, d, e, f, g)
	//const (
	//	i = 1 << iota // 1 << 0
	//	j = 3 << iota // 3 << 1
	//	k             // 3 << 2
	//	l             // 3 << 3
	//)
	//println(i, j, k, l)
	//println(1<<0, 3<<1, 3<<2, 3<<3)
	//var c1, c2, c3 chan int
	//var i1, i2 int
	//select {
	//case i1 = <-c1:
	//	println("received", i1, "from c1")
	//case c2 <- i2:
	//	println("sent", i2, "to c2")
	//case i3, ok := (<-c3):
	//	if ok {
	//		println("received ", i3, " from c3")
	//	} else {
	//		println("c3 is closed")
	//	}
	//default:
	//	println("no communication")
	//}
	//sum := 1
	//for sum <= 10 {
	//	sum += sum
	//	print(sum, " ")
	//}
	//println(sum)

	//strs := []string{"1", "2", "3"}
	//for i, s := range strs {
	//	//println("%d->%s", i, s)
	//	fmt.Printf(" %d -> %s\n", i, s)
	//}
	//a, b := 1, 2
	//a, b = swap(a, b)
	//println(a, b)

	//nums := []int{1, 23, 41, 5, 9, 10} //6
	//for i := 0; i < len(nums)-1; i++ { // 0 -> 5
	//	for j := 0; j < len(nums)-1-i; j++ {
	//		if nums[j] < nums[j+1] {
	//			nums[j], nums[j+1] = nums[j+1], nums[j]
	//		}
	//	}
	//
	//}
	//for i := range nums {
	//	print(nums[i], " ")
	//}

	//a := 100
	//b := 200
	//println(a, b)
	////swap(&a, &b)
	//s := swap
	//s(&a, &b)
	//println(a, b)
	//fmt.Println("")

	//var arr = [5]int{1, 2, 3, 4, 5}
	//println(arr[0])
	//
	//arr2 := [3]int{1, 2}
	//println(arr2[1])
	//
	//arr3 := [...]int{1, 2, 3, 4, 5}
	//println(len(arr3))
	//
	//arr4 := [...]int{1: 1, 2: 2}
	//println(arr4[2])
	//
	//arr5 := [20]int{}
	//for i := 0; i < len(arr5); i++ {
	//	arr5[i] = i * i
	//}
	//for i := 0; i < len(arr5); i++ {
	//	print(arr5[i], " ")
	//}
	//println()
	//
	//for i := 0; i < len(arr5)-1; i++ {
	//	arr5[i], arr5[len(arr5)-i-1] = arr5[len(arr5)-i-1], arr5[i]
	//}
	//
	//for i := 0; i < len(arr5); i++ {
	//	print(arr5[i], " ")
	//}

	//var a int = 10
	//fmt.Printf("%x", &a)

	//var a int = 20 /* 声明实际变量 */
	//var ip *int    /* 声明指针变量 */
	//ip = &a        /* 为指针变量存储地址 */
	//fmt.Printf("a变量值：%d\n", a)
	//fmt.Printf("a变量地址：%x\n", &a)
	//fmt.Printf("ip指针地址：%x\n", ip) /* 指针变量的存储地址 */
	//fmt.Printf("ip指针值：%d\n", *ip) /* 使用指针访问值，指针变量前面加*访问其具体值 */

	//a := []int{10, 100, 1000}
	//var ptr [MAX]*int
	//for i := 0; i < MAX; i++ {
	//	ptr[i] = &a[i]
	//}
	//
	//for i := 0; i < MAX; i++ {
	//	fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	//}

	//a := 1000
	//
	//var ptr *int
	//var pptr **int
	//
	//ptr = &a
	//pptr = &ptr
	//
	//fmt.Printf("a = %d\n", a)
	//fmt.Printf("ptr = %d\n", *ptr)
	//fmt.Printf("pptr = %d\n", **pptr)

	s := Books{"Java教程","山哥","编程",123}
	s.title = "hhh"
	fmt.Println(s.title)

}

type Books struct {
	title string
	author string
	subject string
	bookId int
}

func (c *Books) get1(){
	c.getTest()
}

func (c *Books) getTest(){

}


func getFun() func() func() int {
	i := 0
	return func() func() int {
		i++
		return func() int {
			i++
			return i
		}
	}
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func test(string2 string) (string, string, int) {
	var code = 123
	var date = "2020-12-31"
	var url = "Code=%d&date=%s"
	var str = fmt.Sprintf(url, code, date)
	fmt.Println(string2)
	return str, "", 1
}
