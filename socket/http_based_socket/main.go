package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	var request = make([]byte, 1000)
	_, err := conn.Read(request)
	if err != nil {
		log.Println("failed to read request contents")
		return err
	}
	requestData := string(request)
	headers := strings.Split(requestData, "\r\n")
	fmt.Println(headers)
	fmt.Println(headers[0])
	route := strings.Split(headers[0], " ")[1]
	fmt.Println(route)
	// 实现文本的静态文件服务器
	if route == "/" {
		route = "index.html"
	} else {
		route = route[1:]
	}
	route = "static/" + route
	data, err := ioutil.ReadFile(route)
	if err != nil {
		data := "FILE NOT FOUND"
		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 404 NOT FOUND\r\nContent-Type:text/html\r\nContent-Length:%d\r\n\r\n%s", len(data), data)))
		if err != nil {
			log.Println("failed to write response contents")
			return err
		}
	}
	// 这里修改Content-Type即可实现静态资源服务器
	_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type:text/html\r\nContent-Length:%d\r\n\r\n%s", len(data), string(data))))
	if err != nil {
		log.Println("failed to write response contents")
		return err
	}
	conn.Close()
	return nil
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8081")
	log.Println("http://127.0.0.1:8081")
	if err != nil {
		panic("error listening on port 8080")
	}
	for {
		conn, err := ln.Accept()
		log.Println("received connection")
		if err != nil {
			panic("failed to accept connection")
		}
		handleConnection(conn)
	}
}
