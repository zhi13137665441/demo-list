package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		fmt.Println("开始监听消息")
		for {
			// 读取客户端发送的消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				log.Println("退出监听")
				return
			}
			fmt.Println("接收到了消息，内容是: ", string(message))
		}
	}()
	go func() {
		fmt.Println("定时发送消息")
		// 在新连接接收和发送数据
		count := 1
		for {
			time.Sleep(time.Second)
			// 回复消息给客户端
			err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("message :%d", count)))
			if err != nil {
				log.Println(err)
				return
			}
			count += 1
		}
	}()
}

func main() {
	http.HandleFunc("/websocket", handleWebSocket)
	dir := "./static"
	prefix := "/static/"
	fileServer := http.FileServer(http.Dir(dir))
	handler := http.StripPrefix(prefix, fileServer)
	http.Handle(prefix, handler)
	log.Println("http://localhost:8765/static/")
	log.Fatal(http.ListenAndServe("localhost:8765", nil))
}
