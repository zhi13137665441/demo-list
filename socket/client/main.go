package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()
	for {
		fmt.Println("请输入信息，回车结束输入")
		reader := bufio.NewReader(os.Stdin)
		//终端读取用户回车，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		line = strings.TrimSpace(line)
		//将line发送给服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		go func() {
			for {
				//创建一个切片
				buf := make([]byte, 1024)
				//1.等待服务端通过conn发送信息学
				//2.如果服务端没有wirte[发送]，那么协程就阻塞在这里
				//fmt.Printf("客户端在等待服务端%s发送信息\n",conn.RemoteAddr().String())
				n, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						fmt.Println("the connetction is closed")
						conn.Close()
					} else {
						fmt.Printf("Read Error: %s\n", err)
					}
					return
				}
				//3.显示服务端发送的内容到客户端的终端
				fmt.Printf("服务端%s 发送信息%s\n", conn.RemoteAddr().String(), string(buf[:n]))
			}
		}()
		fmt.Printf("发送的内容了%d文字\n", n)
	}
}
