package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
	"time"
)

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
数组的初始化
*/
func Test1() {
	var arr1 = [3]int{1, 3, 4}            //指定长度
	var arr2 = [...]int{1, 3, 4}          //根据元素数量推算 [1 3 4]
	var arr3 = [...]int{2: 1, 1: 3, 3: 4} // 指定下标初始化  [0 3 1 4]，未指定的下标填充0值
	//var arr4 = [3]int{2: 1, 1: 3, 3: 4} // 指定了长度的话，编译会报错

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

/*
数组的比较
只有当两个数组的所有元素都是相等的时候数组才是相等的
*/
func Test2() {
	var arr1 = [3]int{1, 3, 4}   //指定长度
	var arr2 = [...]int{1, 3, 4} //根据元素数量推算 [1 3 4]
	var arr4 = [3]int{1, 4, 3}   //

	fmt.Println(arr1 == arr2) // true
	//fmt.Println(arr1 == arr3) 长度不同，编译会报错
	//fmt.Println(arr2 == arr3) 长度不同，编译会报错
	fmt.Println(arr1 == arr4) //false

}

/*
实际的数组比较的例子

对称性加密算法（AES、DES、3DES）：
对称式加密就是加密和解密使用同一个密钥。信息接收双方都需事先知道密匙和加解密算法且其密匙是相同的，之后便是对数据进行加解密了。对称加密算法用来对敏感数据等信息进行加密。

非对称算法（RSA、DSA、ECC）：
非对称式加密就是加密和解密所使用的不是同一个密钥，通常有两个密钥，称为"公钥"和"私钥"，它们两个必需配对使用，否则不能打开加密文件。发送双方A,B事先均生成一堆密匙，然后A将自己的公有密匙发送给B，B将自己的公有密匙发送给A，如果A要给B发送消 息，则先需要用B的公有密匙进行消息加密，然后发送给B端，此时B端再用自己的私有密匙进行消息解密，B向A发送消息时为同样的道理。

散列算法（MD5、SHA1、HMAC）：
散列算法，又称哈希函数，是一种单向加密算法。在信息安全技术中，经常需要验证消息的完整性，散列(Hash)函数提供了这一服务，它对不同长度的输入消息，产生固定长度的输出。这个固定长度的输出称为原输入消息的"散列"或"消息摘要"(Message digest)。散列算法不算加密算法，因为其结果是不可逆的，既然是不可逆的，那么当然不是用来加密的，而是签名。

*/
func Test3() {
	// 通过比较返回数组，可以知道加密的原值是否发生改变。类似的还有MD5、SHA1、HMAC
	c1 := sha256.Sum256([]byte("x")) // 会返回一个数组
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

/*
数组作为参数传递

*/
func Test4() {

	var a [1000000]int
	for i := 0; i < 1000000; i++ {
		a[i] = i
	}

	now := time.Now()
	incre(a)
	fmt.Println(fmt.Sprintf("0:%vms", time.Now().Sub(now).Milliseconds())) // 20ms 值传递需要复制是很低效的

	now1 := time.Now()
	increByP(&a)
	fmt.Println(fmt.Sprintf("1:%vms", time.Now().Sub(now1).Milliseconds())) //9ms
	//fmt.Println(a)    // [2 3 5] increByP没修改原值
	//fmt.Println(ints) // [2 3 5]
	//fmt.Println(intp) // [2 3 5]
}

func incre(arr [1000000]int) []int {

	arr1 := []int{}
	for i := range arr {
		arr1 = append(arr1, arr[i]+1)
	}

	return arr1
}

func increByP(arr *[1000000]int) []int {
	arr1 := []int{}
	for i := range *arr {
		arr1 = append(arr1, (*arr)[i]+1)
	}
	return arr1
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) // 存储的0-255的十进制数的对应的1的个数。
	}
}

/*
PopCount函数，用于返回一个数字中含二进制1bit的个数
*/
func Test5() {
	// byte is an alias for uint8 and is equivalent to uint8 in all ways.
	u := uint64(256)     // 超过unit8的范围, 00000001 00000000
	fmt.Println(byte(u)) // 转成byte，取后8位00000000 。 0
	u2 := u >> (1 * 8)   // 右移8位， 00000000 00000001
	fmt.Println(u2)      // 1
	fmt.Println(u)       // 256

	fmt.Println(PopCount(1))
	fmt.Println(PopCount(2))
	fmt.Println(PopCount(3))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
*/
func Test6() {

	c1 := sha256.Sum256([]byte("x")) // 会返回一个数组
	c2 := sha256.Sum256([]byte("X")) // 会返回一个数组

	/*	b1 := byte(255)
		b2 := byte(0)
		count := int(pc[b1^b2])

		num := 0
		//1个字节8个bit,移位运算，获取每个bit
		for m := 1; m <= 8; m++ {
			//比较每个bit是否相同
			if (b1 >> uint(m)) != (b2 >> uint(m)) {
				num++
			}
		}
		fmt.Println(fmt.Sprintf("notEqual count:%v", count))
		fmt.Println(fmt.Sprintf("notEqual num:%v", num))*/

	//bytes1 := c1[:]
	//bytes2 := c2[:]

	notEqualCount := 0
	for i := range c1 {
		count := int(pc[c1[i]^c2[i]])
		notEqualCount += count
	}
	fmt.Println(fmt.Sprintf("notEqual:%v", notEqualCount))
	fmt.Println(fmt.Sprintf("notEqual1:%v", compareSha256(c1, c2))) // 感觉这个算法有问题。
}

func compareSha256(str1 [32]byte, str2 [32]byte) int {
	num := 0
	//循环字节数组
	for i := 0; i < len(str1); i++ {
		//1个字节8个bit,移位运算，获取每个bit
		for m := 1; m <= 8; m++ {
			//比较每个bit是否相同
			if (str1[i] >> uint(m)) != (str2[i] >> uint(m)) {
				num++
			}
		}
	}
	return num
}

var encryName = flag.String("sha", "sha256", "input encry len")

/*
练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
*/
func Test7() {

	flag.Parse()

	switch strings.ToLower(*encryName) {
	case "sha256":
		sha256.Sum256([]byte("x"))
	case "sha384":
		sha512.Sum384([]byte("x"))
	case "sha512":
		sha512.Sum512([]byte("x"))
	default:
	}
}
