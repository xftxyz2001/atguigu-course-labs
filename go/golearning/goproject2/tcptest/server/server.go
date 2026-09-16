package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("监听失败:", err.Error())
	}
	fmt.Println("等待客户端连接...")
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受连接失败:", err.Error())
			continue // 等待下次连接
		}
		fmt.Println("接受到客户端连接.", conn.RemoteAddr())
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		// 读取数据
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			// fmt.Println("读取数据失败:", err.Error())
			if err.Error() == "EOF" {
				fmt.Println("客户端退出.")
			} else {
				fmt.Println("读取数据失败:", err.Error())
			}

			break
		}
		fmt.Print(conn.RemoteAddr(), ":", string(data[:n]))
	}
}
