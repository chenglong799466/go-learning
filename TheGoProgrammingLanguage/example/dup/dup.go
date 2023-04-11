package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"jacklcheng/go-learning/TheGoProgrammingLanguage/example"
	"os"
	"strings"
)

// 查找重复的行

/*
dup的第一个版本打印标准输入中多次出现的行，以重复次数开头。该程序将引入if语句，map数据类型以及bufio包。
*/

// 必须在main函数中执行，
func Dup1() {

	content := make(map[string]int)

	// bufio包
	input := bufio.NewScanner(os.Stdin)

	// 一直轮询，需要跳出
	for input.Scan() {
		content[input.Text()]++
		if input.Text() == "ex" {
			break
		}
	}

	for key, value := range content {
		fmt.Println(fmt.Sprintf("%d:%s", value, key))
	}
}

/*
dup程序的下个版本读取标准输入或是使用os.Open打开各个具名文件，并操作它们。
*/

// 必须在main函数中执行，
func Dup2() {

	content := make(map[string]int)

	args := os.Args[1:]

	if len(args) == 0 {
		example.CountLine(os.Stdin, content)
	} else {
		for _, arg := range args {
			fmt.Println(fmt.Sprintf("args:%v", arg))
			// open默认只读的
			//open, err := os.Open(arg)
			// os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取。
			file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				fmt.Println(fmt.Sprintf("err:%v", err))
			}
			file.Name()
			example.CountLine(file, content)
			//writeLen, err := file.Write([]byte("jacklcheng"))
			//if err != nil {
			//	fmt.Println(fmt.Sprintf("err1:%v", err))
			//}
			//fmt.Println(fmt.Sprintf("write len:%v", writeLen))
			file.Close()
		}
	}
}

/*
一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。
*/

// 必须在main函数中执行，
func Dup3() {
	content := make(map[string]int)
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Println(fmt.Sprintf("args:%v", arg))
		// open默认只读的
		//open, err := os.Open(arg)
		// os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取。
		bytes, err := ioutil.ReadFile("test.txt")
		if err != nil {
			fmt.Println(fmt.Sprintf("err:%v", err))
		}
		split := strings.Split(string(bytes), "\n")
		for _, s := range split {
			content[s]++
		}
		fmt.Println(fmt.Sprintf("split:%v", split))
	}
}

func main() {
	Dup1()
	Dup2()
	Dup3()

}
