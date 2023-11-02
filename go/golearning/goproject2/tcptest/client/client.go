package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("链接服务器失败:", err.Error())
	}
	fmt.Println("链接服务器成功.")
	defer conn.Close()

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取用户输入失败:", err.Error())
			break
		}
		// 去除空白字符
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("退出客户端.")
			break
		}
		// 发送数据
		line = line + "\n"
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("发送数据失败:", err.Error())
			break
		}
	}

}
