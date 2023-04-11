package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// 命令行参数

/*
练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
*/
func Test1() {
	var s = os.Args[0]
	fmt.Println("s:" + s)
}

/*
练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。

*/
func Test2() {
	for i, arg := range os.Args {
		fmt.Println(fmt.Sprintf("Arg[%d]:%s", i, arg))
	}
}

/*
练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
*/
func Test3() {

	average1()
	average2()

}

func average1() {
	var spans []int64
	for i := 0; i < 1000; i++ {
		now := time.Now()
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		spans = append(spans, time.Now().Sub(now).Microseconds())
	}

	var average float64
	var total float64
	for _, span := range spans {
		total = total + float64(span)
	}
	average = total / 1000
	fmt.Println(fmt.Sprintf("average1:%fms", average))
}
func average2() {
	var spans []int64
	for i := 0; i < 1000; i++ {
		now := time.Now()
		strings.Join(os.Args[1:], " ")
		spans = append(spans, time.Now().Sub(now).Microseconds())
	}

	var average float64
	var total float64
	for _, span := range spans {
		total = total + float64(span)
	}
	average = total / 1000
	fmt.Println(fmt.Sprintf("average2:%fms", average))
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
}

/*
go.flag包
*/

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func Test4() {
	// 必须调用Parse()方法，否则n,sep指针指向的为默认值
	flag.Parse()

	fmt.Println("s:" + *sep)
	fmt.Println(fmt.Sprintf("n:%v", *n))
}

func init() {

}
func init() {

}
func init() {

}
func init() {

}
