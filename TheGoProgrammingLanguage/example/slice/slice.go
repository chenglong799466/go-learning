package main

import "fmt"

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
	Test7()

}

/*
切片是对底层数组的间接引用,直接引用是生成的底层数组的匿名变量
*/
func Test1() {
	var a, b int

	var p *int /*定义指针变量p*/
	p = &b     /*将变量b的地址放在变量p中*/

	a = 3 /*直接引用变量a*/

	*p = 5 /* 间接引用变量b*/
	fmt.Println(a, b)
}

/*
 slice的cap和底层数组的长度有关.
 1.对底层数组切片,则cap=切片的起始坐标到底层数组的最后一个元素的长度
 2.切片操作超出cap(s)的上限将导致一个panic异常
*/
func Test2() {

	var arr [20]string
	sli := arr[2:3]
	fmt.Println(cap(sli)) // 18
	fmt.Println(len(sli)) // 1

}

/*
 数组翻转

 1.slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的元素。
*/
func Test3() {

	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	// 左旋转
	b := [...]int{0, 1, 2, 3, 4, 5}
	reverse(b[:2])
	reverse(b[2:])
	reverse(b[:])
	fmt.Println(b) // "[2, 3, 4, 5, 0, 1]"

	// 右旋转
	c := [...]int{0, 1, 2, 3, 4, 5}
	reverse(c[:])
	reverse(c[:2])
	reverse(c[2:])
	fmt.Println(c) // "[4, 5, 0, 1, 2, 3]"

}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
 切片和底层数据

*/
func Test4() {

	a := [...]int{0, 1, 2, 3, 4, 5}

	ints0 := a[:3]
	ints1 := a[1:4]
	fmt.Printf("len %v,cap %v,ints0:%v \n", len(ints0), cap(ints0), ints0)  // 3,6
	fmt.Printf("len %v,cap %v,ints1:%v  \n", len(ints1), cap(ints1), ints1) // 3,5

	// append函数可以直接修改底层数组的值
	/*
		ints0和ints1共用同一个底层数据，ints0 append函数直接修改了底层数组的值。
		导致ints1的末位也发生改变
	*/
	ints0 = append(ints0, 1111)
	fmt.Printf("len %v,cap %v,ints0:%v  \n", len(ints0), cap(ints0), ints0) // 4,6, [0 1 2 1111]
	fmt.Printf("len %v,cap %v,ints1:%v  \n", len(ints1), cap(ints1), ints1) // 3,5, [1 2 1111]

	// append函数可以直接修改底层数组的值
	/*
		1.如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。因此，输入的x和输出的z共享相同的底层数组。
		2.如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。
	*/
	a1 := [...]int{0, 1, 2, 3, 4, 5}

	ints2 := a1[2:]
	ints3 := a1[1:]
	ints3 = append(ints3, 1111)
	fmt.Printf("len %v,cap %v,ints2:%v  \n", len(ints2), cap(ints2), ints2) // 4,6, [2 3 4 5]
	fmt.Printf("len %v,cap %v,ints3:%v  \n", len(ints3), cap(ints3), ints3) // 3,5, [1 2 3 4 5 1111]

	ints4 := a1[2:]
	fmt.Printf("len %v,cap %v,ints4:%v  \n", len(ints4), cap(ints3), ints4) // 3,5, [1 2 3 4 5 1111]
}

/*
练习 4.3： 重写reverse函数，使用数组指针代替slice。
*/

func Test5() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	reversePrt(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"

}

func reversePrt(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。

*/

func Test6() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 2)

}

func rotate(s []int, r int) []int {

	intNew := make([]int, 0)
	intNew = append(intNew, s[r:]...)
	intNew = append(intNew, s[:r]...)
	fmt.Println(intNew)

	lens := len(s)
	//创建一个空的指定长度的slice
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	fmt.Println(res)
	return res
}

/*
 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
*/

func Test7() {
	a := [...]string{"1", "0", "0", "1", "2", "2", "4", "5", "5"}
	remove(a[:])
	removeMap(a[:])
	fmt.Println(a)

}

func remove(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			// append函数直接修改了底层数据的值
			fmt.Println(fmt.Sprintf("s[:i]:cap %d,len %d", cap(s[:i]), len(s[:i])))
			s = append(s[:i], s[i+1:]...)
			fmt.Println(s)
		}
	}
	return s
}

func removeMap(s []string) []string {
	m := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		_, ok := m[s[i]]
		if ok {
			i++
		} else {
			m[s[i]] = i
		}
	}
	for key := range m {
		i := 0
		s[i] = key
		i++
	}
	return s
}
