package main

import (
	"fmt"
)

/*
* 放在类型前面可以标明是个指针，如果加在变量前面代表指针取值（对指针取值即内存地址）
& 取地址

变量都可以取地址，

指针的使用：
主要经过三个步骤：声明、赋值和访问指针指向的变量的值

p是指针变量*int

&p，意思是取指针变量本身的地址，不是取指针变量指向的值。
*p，意思是取指针变量的值（另一个变量的地址）的值
&*p,上诉两个步骤的叠加。返回的是指针变量的值&int
*&p，上诉两个步骤的叠加。返回的是指针变量p的值&int

返回值的类型判断，可以根据**抵消
&*和*&是个互补的操作，也可以抵消

&p

*/
func main() {
	// 指针的声明，赋值，取值
	point1()
	// 通过指针修改指针引用的值
	point2()

}

func point1() {
	var intVar int                // 声明变量，int类型
	var pointerVar *int           // 声明变量，指针类型，指向一个int值
	var pointerToPointerVar **int // 声明变量, 指针类型，指向一个指针值

	intVar = 100                      // 变量intVar赋值，值为100
	pointerVar = &intVar              // 变量pointerVar赋值，值为一个地址
	pointerToPointerVar = &pointerVar // 变量pointerVar赋值，值为一个地址

	/*
		变量和指针变量赋值
	*/
	fmt.Println("1")
	fmt.Println("intVar:\t\t\t", intVar) // intVar变量的值（intVar存储的值为100）100
	if intVar != 100 {
		fmt.Errorf("err")
	}
	fmt.Println("pointerVar:\t\t", pointerVar) // pointerVar变量的值（pointerVar的存储的值为（intVar的内存地址） ）                   0xc000016260
	if pointerVar != &intVar {
		fmt.Errorf("err")
	}
	fmt.Println("pointerToPointerVar:\t", pointerToPointerVar) // pointerToPointerVar变量的值（pointerToPointerVar存储的值为（pointerVar变量的内存地址））  0xc00000e038
	if pointerToPointerVar != &pointerVar {
		fmt.Errorf("err")
	}

	/*
		变量和指针变量取地址
	*/
	fmt.Println("2")
	fmt.Println("&intVar:\t\t", &intVar) // intVar变量自身的内存地址（pointerVar变量的值） 0xc000016260
	if &intVar != pointerVar {
		fmt.Errorf("err")
	}

	fmt.Println("&pointerVar:\t\t", &pointerVar) // pointerVar变量自身的内存地址（pointerToPointerVar变量的值）   0xc00000e038
	if &pointerVar != pointerToPointerVar {
		fmt.Errorf("err")
	}
	fmt.Println("&pointerToPointerVar:\t", &pointerToPointerVar) // pointerToPointerVar变量自身的内存地址  新

	/*
		变量和指针变量取值
	*/
	fmt.Println("3")
	fmt.Println("*pointerVar:\t\t", *pointerVar) // *取值是取pointerVar变量的值（intVar变量的内存地址）的值 100
	if *pointerVar != intVar {
		fmt.Errorf("err")
	}
	fmt.Println("*pointerToPointerVar:\t", *pointerToPointerVar) // *取值是取pointerToPointerVar变量的值（pointerVar变量的内存地址&pointerVar）的值 。可以用变量类型判断**抵消，最终**pointerToPointerVar的类型应该为*int，&intVar
	if *pointerToPointerVar != &intVar {
		fmt.Errorf("err")
	}
	fmt.Println("**pointerToPointerVar:\t", **pointerToPointerVar) // **多次取值（可以理解为套娃）,1.取pointerToPointerVar变量的值的值 &intVar 2.取&intVar的值，intVar  。 也可以用变量类型判断**抵消，最终**pointerToPointerVar的类型应该为int
	if **pointerToPointerVar != intVar {
		fmt.Errorf("err")
	}

	/*
		变量和指针变量取值+取地址
	*/
	fmt.Println("4")
	fmt.Println("*pointerVar:\t\t", &*pointerVar) // *取值是取pointerVar变量的值（intVar变量的内存地址）的值 100
	if &*pointerVar != &intVar {
		fmt.Errorf("err")
	}
	fmt.Println("*pointerToPointerVar:\t", *&pointerVar) // *取值是取pointerToPointerVar变量的值（pointerVar变量的内存地址&pointerVar）的值 。可以用变量类型判断**抵消，最终**pointerToPointerVar的类型应该为*int，&intVar
	if *&pointerVar != &intVar {
		fmt.Errorf("err")
	}
}

func point2() {

	/*
		每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了新的别名。
		例如，*p就是变量v的别名。指针特别有价值的地方在于我们可以不用名字而访问一个变量，但是这是一把双刃剑：
		要找到一个变量的所有访问者并不容易，我们必须知道变量全部的别名（译注：这是Go语言的垃圾回收器所做的工作）。
		不仅仅是指针会创建别名，很多其他引用类型也会创建别名，例如slice、map和chan，甚至结构体、数组和接口都会创建所引用变量的别名。
	*/
	i := 0
	var p *int // 声明
	p = &i     // 赋值

	incr(p)

	fmt.Println(*p) // 1
}

func incr(p *int) {
	*p++
}

/*
通过new()函数创建变量
*/
func point3() {

	/*
		每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了新的别名。
		例如，*p就是变量v的别名。指针特别有价值的地方在于我们可以不用名字而访问一个变量，但是这是一把双刃剑：
		要找到一个变量的所有访问者并不容易，我们必须知道变量全部的别名（译注：这是Go语言的垃圾回收器所做的工作）。
		不仅仅是指针会创建别名，很多其他引用类型也会创建别名，例如slice、map和chan，甚至结构体、数组和接口都会创建所引用变量的别名。
	*/
	var p = new(int) // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p)  // "0"
	*p = 2           // 设置 int 匿名变量的值为 2
	fmt.Println(*p)  // "2"
}
