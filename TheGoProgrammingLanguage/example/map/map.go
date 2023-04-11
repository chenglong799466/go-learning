package main

import "fmt"

/*
一个map就是一个哈希表的引用
*/

func main() {

	Test1()
}

/*
map删除元素
*/
func Test1() {

	m2 := map[string]int{
		"chenglong":  1,
		"jacklcheng": 2,
	}
	delete(m2, "chenglong")
	fmt.Println(m2)
}

/*
和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。
*/
func Test2() {

	m2 := map[string]int{
		"chenglong":  1,
		"jacklcheng": 2,
	}

	m1 := map[string]int{
		"chenglong":  1,
		"jacklcheng": 2,
	}

	// 报错
	/*if m1 == m2 {

	}*/
	// 可以和nil比较
	if m1 == nil {

	}
	// 必须通过循环比较map
	if equal(m1, m2) {
		fmt.Println("==")
	}

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

/*
go程序中set的实现
Go程序员将这种忽略value的map当作一个字符串集合[set]，并非所有map[string]bool类型value都是无关紧要的；有一些则可能会同时包含true和false的值。
*/

/*
map的key必须是可比较类型的。
如果是非可以比较类型的key，定义一个辅助函数转换。
第一步，定义一个辅助函数k，将slice转为map对应的string类型的key，确保只有x和y相等时k(x) == k(y)才成立。然后创建一个key为string类型的map，在每次对map操作时先用k辅助函数将slice转化为string类型。
*/

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
