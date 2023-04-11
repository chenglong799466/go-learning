package example

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// test文件的方法不能被其他包引用。。
func CountLine(at io.Reader, content map[string]int) {

	file, ok := at.(*os.File)
	if !ok {
		fmt.Errorf("type err")
	}

	// bufio包,NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines.
	input := bufio.NewScanner(at)

	// 一直轮询，需要跳出
	for input.Scan() {
		content[input.Text()]++
		if input.Text() == "ex" {
			break
		}
	}

	for key, value := range content {
		fmt.Println(fmt.Sprintf("%d:%s:%s", value, key, file.Name()))
	}

	// 出现重复打印文件名
	for key, value := range content {
		fmt.Println(fmt.Sprintf("%d:%s:%s", value, key, file.Name()))
	}
}
