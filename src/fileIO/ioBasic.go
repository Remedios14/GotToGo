package fileIO

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	const (
		defaultBufsize = 4096
	)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到换行符为止
		if err == io.EOF { // 
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束...")

	// 读取完整文件
	content, err := ioutil.ReadFile("d:/test.txt")
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content)) // 本身是个 []byte
}
