package fileIO

import (
	"bufio"
	"fmt"
	"os"
)

func write() {
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer file.Close()
	str := "Byebye, Gardian\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush() // 真正将缓存数据写入文件
}
