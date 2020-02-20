package main;

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	//proverb, err := os.Create("./proverb.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//proverb.Close()

	//写入文件
	proverb, err := os.Create("./proverb.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := "don't communicate by sharing memory share memory by communicating"
	_, _ = proverb.Write([]byte(data))
	proverb.Close()

	//读取文件内容
	proverb, err = os.Open("./proverb.txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 1024)
	for {
		lean, _ := proverb.Read(buf)
		if lean == 0 {
			break;
		}
		fmt.Println(string(buf))
	}
	proverb.Close()
}