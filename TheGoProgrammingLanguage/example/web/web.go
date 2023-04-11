package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/count", ServeHTTP1)
	http.ListenAndServe("localhost:8080", nil)
}

var mu sync.Mutex
var count int

func ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	mu.Lock()
	count++
	mu.Unlock()

	//// 直接写入write中
	//sprintf := fmt.Sprintf("path:%s", request.URL.Path)
	//_, err := writer.Write([]byte(sprintf))
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("err:%v", err))
	//}
	// 这个函数会将/hello这个路径从请求的URL中解析出来，然后把其发送到响应中，这里我们用的是标准输出流的fmt.Fprintf
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

/*
对请求的次数进行计算
*/

func ServeHTTP1(writer http.ResponseWriter, request *http.Request) {
	// 请求和返回一致
	sprintf := fmt.Sprintf("path:%s,count:%d", request.URL.Path, count)
	// 直接写入write中
	_, err := writer.Write([]byte(sprintf))
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%v", err))
	}
	// 这个函数会将/hello这个路径从请求的URL中解析出来，然后把其发送到响应中，这里我们用的是标准输出流的fmt.Fprintf
	//fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}
