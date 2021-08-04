package main

import (
	"fmt"
	"strings"
)

var m = 100

func main() {
	//iota计数器
	const (
		n0 = iota
		n1
		n2
	)
	const n3 = 3
	fmt.Printf("n0:%v\nn1:%v\nn2:%v\nn3:%v\n", n0, n1, n2, n3)

	//字符串拼接
	str := fmt.Sprintf("code = %v", 12)
	fmt.Println(str)
	//字符串分割
	split := strings.Split(str, "=")
	fmt.Println(split)

	//判断是否包含
	contains := strings.Contains(str, "e")
	fmt.Println(contains)

	//前缀判断
	if strings.HasPrefix("str", "s") {
		fmt.Println("s开头")
	}
	//后缀判断
	if strings.HasSuffix("str", "r") {
		fmt.Println("r结尾")
	}

	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}

	//修改字符串
	s1 := "hello"
	//强转
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))
	for _, i2 := range byteS1 {
		fmt.Printf("%s", string(i2))
	}
	fmt.Println()
	d := [...]struct {
		name string
		age  int
	}{
		{"sos", 12},
	}

	fmt.Println(d)

	//改变数组中的值
	arr := [2]int{1, 2}
	fmt.Printf("arr: %p\n", &arr)
	arrTest(&arr)
	fmt.Println(arr)

	//切片
	var nums = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(nums), cap(nums))
	PrintNums(nums[2:3])
	PrintNums(nums[0:2])
	PrintNums(nums[1:6])
}

func PrintNums(nums []int) {
	fmt.Printf("len = %d,cap =  %d, nums=%v\n", len(nums), cap(nums), nums)
}

func arrTest(x *[2]int) {
	fmt.Printf("x: %p\n", x)
	x[0] = 1000
}
