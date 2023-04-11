package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	TestStr1()
	TestStr2()
	TestStr3()
	TestStr4()
	TestStr5()
	TestStr6()
	TestStr7()
	TestStr8()

}

func TestStr1() {
	s := "Hello, 世界"
	/*
		1. 返回的字节的长度，utf-8对于一个字符采用1到4个字节表示，汉字在utf-8编码里面使用三个字节表示
		2. 表示的规则：
		高位0表示需要一个字节表示，1表示需要两个字节表示，11表示需要三个字节表示
		// 1110xxxx 10xxxxxx 10xxxxxx

	*/
	fmt.Println(len(s)) // "13"
	/*
		返回的rune的个数，即unicode的码点数
	*/
	fmt.Println(utf8.RuneCountInString(s)) // "9"
}

func TestStr2() {
	s := "Hello, 世界"
	// 如何取出字符串的中文字符呢？
	// 1. 取出对应字节数
	fmt.Println()
	fmt.Printf("%q\n", s[10:])

	// 2.字符串遍历的时候会自动，使用utf-8编码。因此i的步长不是固定的，c则是一个unicode的字符
	for i, c := range s {
		//fmt.Println(fmt.Sprintf("i:%v,c:%v", i, c))
		fmt.Printf("%d\t%q\t%d\n", i, c, c)

	}
}

/*
通过8进制或者16进制，输入字面量

一个十六进制的转义形式是`\xhh`，其中两个h表示十六进制数字（大写或小写都可以）。
一个八进制转义形式是`\ooo`，包含三个八进制的o数字（0到7），但是不能超过\377（译注：对应一个字节的范围，十进制为255）。
*/
func TestStr3() {
	s := "Hello, 世界"

	// 2.字符串遍历的时候会自动，使用utf-8编码。因此i的步长不是固定的，c则是一个unicode的字符
	for i, c := range s {
		//fmt.Println(fmt.Sprintf("i:%v,c:%v", i, c))
		fmt.Printf("%d\t%q\t%x\n", i, c, c) // 打印字符的unicode编码的16进制形式

	}

	/*
		Go语言字符串面值中的Unicode转义字符让我们可以通过Unicode码点输入特殊的字符。
		有两种形式：\uhhhh对应16bit的码点值，\Uhhhhhhhh对应32bit的码点值，其中h是一个十六进制数字；一般很少需要使用32bit的形式。
		每一个对应码点的UTF8编码。例如：下面的字母串面值都表示相同的值：
	*/

	/*
		"世界"
		"\xe4\xb8\x96\xe7\x95\x8c"
		"\u4e16\u754c"
		"\U00004e16\U0000754c"
	*/
	s1 := "世界"
	s2 := "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(s2, s1 == s2)
}

/*
每一个UTF8字符解码，不管是显式地调用utf8.DecodeRuneInString解码或是在range循环中隐式地解码，
如果遇到一个错误的UTF8编码输入，将生成一个特别的Unicode字符\uFFFD，
在印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号"?"。
当程序遇到这样的一个字符，通常是一个危险信号，说明输入并不是一个完美没有错误的UTF8字符串。
*/
func TestStr4() {
	s := "Hello, 世界"

	s2 := s[12:]
	fmt.Println(s2) // 不合法的utf-8编码，会打印黑色六角或钻石形状，里面包含一个白色的问号"?"

	/*
		utf-8包包含对unicode字符的操作
	*/
	utf8.RuneCountInString(s) // 9
	fmt.Println(utf8.RuneCountInString(s))
}

/*
[]rune 序列
*/
func TestStr5() {
	s := "Hello, 世界"
	fmt.Printf("% x\n", s) // 打印每个字节的16进制表示

	// s转成rune序列存储的unicode码点
	runes := []rune(s)
	fmt.Printf("%x", runes) // 打印每个字节的16进制表示 [48 65 6c 6c 6f 2c 20 4e16 754c]

	// rune序列转成string，对rune序列进行utf-8编码
	fmt.Println(string(runes)) //  "Hello, 世界"

	// 将整型转成string，将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

}

/*
实现basename
fmt.Println(basename("a/b/c.go")) // "c"
fmt.Println(basename("c.d.go"))   // "c.d"
fmt.Println(basename("abc"))      // "abc"
*/

func TestStr6() {

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))

}

func basename(full string) string {
	if len(full) == 0 {
		return ""
	}

	runes := []rune(full)
	lastIndex := 0
	lastIndex1 := 0
	for i, r := range runes {
		if r == '.' {
			lastIndex = i
		}
		if r == '/' {
			lastIndex1 = i
		}
	}
	if lastIndex == 0 && lastIndex1 == 0 {
		return full
	}
	if lastIndex == 0 && lastIndex1 != 0 {
		return string(runes[lastIndex1+1:])
	}
	if lastIndex != 0 && lastIndex1 == 0 {
		return string(runes[:lastIndex])
	}
	if lastIndex != 0 && lastIndex1 != 0 {
		if lastIndex1+1 <= lastIndex {
			return string(runes[lastIndex1+1 : lastIndex])
		} else {
			fmt.Println("err")
		}
	}

	return ""

}

/*
strings.LastIndex
*/
func basename1(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

/*

 */
func TestStr7() {
	fmt.Println(comma("123456778"))                   //递归实现
	fmt.Println(commaByByteBuff("123456778"))         // bytes.buffer实现，异常
	fmt.Println(commaByByteBuff1("123456778"))        // bytes.buffer实现
	fmt.Println(commaByByteBuff2("-123456778.44444")) // bytes.buffer实现，支持浮点数和可选正负号
	fmt.Println(compare("123456778", "123456787"))    // 比较两个字符串是打乱顺序的 true
	fmt.Println(compare("12345程龙", "12345龙程"))        // 比较两个字符串是打乱顺序的 true
	fmt.Println(compare("12345程龙", "12345程龙"))        // 比较两个字符串是打乱顺序的 false
	fmt.Println(compare1("12345程龙", "12345程馟"))       // compare1方法是有问题的，没有转成rune序列，因此对中文可能有问题 。期望是fasle，实际为true
}

/*
“12345”处理后成为“12,345”
*/
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

/*
“12345”处理后成为“12,345”
*/
// comma inserts commas in a non-negative decimal integer string.
func commaByByteBuff(s string) string {
	buffer := bytes.Buffer{}

	if len(s) < 3 {
		return s
	}
	var arr []string
	for i := range s {
		if i+3 >= len(s) {
			arr = append(arr, s[i:])
			break
		}
		arr = append(arr, s[i:i+3])
		i = i + 3 // 傻逼了
	}
	for i, s2 := range arr {
		if i+1 == len(arr) {
			continue
		}
		buffer.WriteString(s2 + ",")
	}
	return string(buffer.Bytes())
}

/*
练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
*/
// comma inserts commas in a non-negative decimal integer string.
func commaByByteBuff1(s string) string {
	buffer := bytes.Buffer{}

	if len(s) < 3 {
		return s
	}
	for i := 0; i < len(s); i = i + 3 {
		if i+3 >= len(s) {
			buffer.WriteString(s[i:])
			break
		}
		buffer.WriteString(s[i:i+3] + ",")
	}
	return string(buffer.Bytes())
}

/*
练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

*/
// comma inserts commas in a non-negative decimal integer string.
func commaByByteBuff2(s string) string {
	buffer := bytes.Buffer{}

	index0 := 0
	if s[:1] == "-" {
		index0 = 1
	}

	index := 0
	for i := range s {
		if s[i] == '.' {
			index = i
			break
		}
	}

	sub := s[index0:index]
	// 负数
	if len(sub) < 3 {
		return s[:index0] + sub + s[index:]
	}
	for i := 0; i < len(sub); i = i + 3 {
		if i+3 >= len(sub) {
			buffer.WriteString(sub[i:])
			break
		}
		buffer.WriteString(sub[i:i+3] + ",")
	}
	return s[:index0] + string(buffer.Bytes()) + s[index:]
}

/*
练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
*/

func compare(s string, s1 string) bool {

	if s == s1 {
		return false
	}

	runes := []rune(s)
	runes1 := []rune(s1)
	// hha ahh
	m := make(map[int32]int)
	m1 := make(map[int32]int)
	for _, i := range runes {
		m[i]++
	}
	for _, i := range runes1 {
		m1[i]++
	}

	if len(m) == len(m1) {
		for key, value := range m {
			value1, ok := m1[key]
			if ok {
				if value1 == value {
					return true
				}
			}

		}

	}

	return false
}

func compare1(s string, s1 string) bool {

	if s == s1 {
		return false
	}

	// 【龙】的unicode码 9f99
	fmt.Printf("%x\n", []rune(s))    // 打印每个字节的16进制表示
	fmt.Printf("%s", string(0x9f99)) //
	fmt.Printf("%s", string(0x999f)) //

	// hha ahh
	m := make(map[int32]int)
	m1 := make(map[int32]int)
	for _, i := range s {
		m[i]++
	}
	for _, i := range s1 {
		m1[i]++
	}

	if len(m) == len(m1) {
		for key, value := range m {
			value1, ok := m1[key]
			if ok {
				if value1 == value {
					return true
				}
			}

		}

	}

	return false
}

/*

strconv 包 使用

*/
func TestStr8() {

	// 把数值类型转成字符串
	// 1. strconv
	fmt.Println(strconv.FormatInt(7, 16))                              // base代表进制 // 7
	fmt.Println(strconv.FormatInt(7, 10))                              // base代表进制 // 7
	fmt.Println(strconv.FormatInt(7, 8))                               // base代表进制 // 7
	fmt.Println(strconv.FormatInt(7, 2))                               // base代表进制 // 111
	fmt.Println(strconv.FormatBool(true))                              // base代表进制 // "true"
	fmt.Println(strconv.FormatUint(7, 2))                              // base代表进制  和FormatInt一样，参数不同
	fmt.Println(strconv.FormatFloat(70.2444444333333333, 'f', -1, 64)) //fmt格式化的范式，prec精度,bitSize 内存精度大小64或者32。64位 70.24444443333333
	fmt.Println(strconv.FormatFloat(70.2444444333333333, 'f', -1, 32)) //32位  70.244446

	// reassign value
	num := -17.96
	// returns a string type
	fmt.Println(strconv.FormatFloat(num, 'G', 2, 64))
	fmt.Println()

	// reassign value
	num = 0.235
	// returns a string type
	fmt.Println(strconv.FormatFloat(num, 'E', -1, 64))

	// 2. fmt %b、%d、%o和%x
	fmt.Println(fmt.Sprintf("b=%b", 4)) // base 2
	fmt.Println(fmt.Sprintf("d=%d", 4)) // base 10
	fmt.Println(fmt.Sprintf("o=%o", 4)) // base 8
	fmt.Println(fmt.Sprintf("x=%x", 4)) // base 16

	// 字符串转成数值类型
	fmt.Println(strconv.ParseInt("711111111111111111111", 10, 32)) // value out of range
	fmt.Println(strconv.ParseInt("7111111111111111", 10, 32))      // value out of range
	fmt.Println(strconv.ParseInt("7111111111111111", 10, 64))      // 7111111111111111
	fmt.Println(strconv.ParseFloat("7111111111.33", 64))           // 7.11111111133e+09
	fmt.Println(strconv.Atoi("711111111111111111111"))             //   Atoi is equivalent to ParseInt(s, 10, 0), converted to type int. bitsize=0，在uint的时候升级为int64
}
