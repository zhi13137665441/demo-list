package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// websocket 升级允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var port string

func handleGet(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("我是server_%s", port)
	w.Write([]byte(message))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 可以通过参数传输token来验证用户
	fmt.Println("token为: ", r.FormValue("token"))
	token := r.FormValue("token")
	// checkToken方法，来验证连接是否安全 token可以转为是用户。
	if token == "" {
		//data := map[string]string{"err_reason": "unauthorized"}
		//errData, err := json.Marshal(data)
		//if err != nil {
		//	w.WriteHeader(401)
		//	w.Write(errData)
		//}
		return
	}
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 服务端逻辑代码
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
			message := fmt.Sprintf("server_%s:message :%d", port, count)
			log.Println("发送了: " + message)
			err = conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println(err)
				return
			}
			count += 1
		}
	}()
}

func main() {
	flag.StringVar(&port, "port", "8765", "端口号")
	flag.Parse()
	http.HandleFunc("/websocket", handleWebSocket)
	http.HandleFunc("/get", handleGet)
	log.Println(fmt.Sprintf("%s", port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
