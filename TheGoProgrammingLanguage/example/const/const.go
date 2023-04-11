package main

import (
	"fmt"
	"time"
)

func main() {
	Test()
	Test2()

}

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

type MyStr struct {
	str string
}

func (m MyStr) GoString() string {
	return m.str + "jacklcheng"
}

func Test() {
	// 常量的类型,[1]代表取值顺序
	fmt.Printf("%[2]d，%[1]d\n", 1, 2) // 2,1

	// %#v会打印GoString()方法的实现
	fmt.Printf("%#v\n", MyStr{}) // jacklcheng

	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s" //Duration类型实现了String()方法，所以会输出5m0s
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
}

const (
	a = 1
	b
	c = 2
	d
)

func Test2() {
	// 没有赋值的话=上一个常量值
	fmt.Printf("a:%v", a) //1
	fmt.Printf("b:%v", b) //1
	fmt.Printf("c:%v", c) //2
	fmt.Printf("d:%v", d) //2
}

/*
iota 常量生成器
*/

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)
