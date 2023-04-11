package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	//http1()
	http2()

}

func http1() {
	args := os.Args[1:]

	for _, arg := range args {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Println(fmt.Sprintf("err:%v", err))
		}

		readCloser := resp.Body
		all, err := ioutil.ReadAll(readCloser)
		if err != nil {
			fmt.Println(fmt.Sprintf("err:%v", err))
			// exit
			os.Exit(1)
		}
		readCloser.Close()
		fmt.Println(fmt.Sprintf("resp:%s", all))
	}
}

/*
函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中
*/

func http2() {
	args := os.Args[1:]

	for _, arg := range args {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Println(fmt.Sprintf("err:%v", err))
		}

		readCloser := resp.Body
		writeLen, err := io.Copy(os.Stdout, readCloser)
		if err != nil {
			fmt.Println(fmt.Sprintf("err:%v", err))
		}
		readCloser.Close()
		fmt.Println(fmt.Sprintf("resp:%s", writeLen))
	}
}

/*
并发获取多个URL
*/

func http3() {
	args := os.Args[1:]
	respChan := make(chan string)
	for _, arg := range args {
		go call(arg, respChan)
	}

	for s := range respChan {
		fmt.Println(fmt.Sprintf("resp:%s", s))
	}
}

func call(arg string, respChan chan<- string) {
	now := time.Now()
	resp, err := http.Get(arg)
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%v", err))
	}
	readCloser := resp.Body
	_, err = io.ReadAll(readCloser)
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%v", err))
	}
	readCloser.Close()

	respChan <- fmt.Sprintf("escape:%v,url:%v", time.Now().Sub(now), arg)
}
