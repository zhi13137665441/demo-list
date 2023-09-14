package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func sendEvent(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定数据流的 MIME 类型和字符集
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 设置允许跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 模拟数据推送
	for i := 1; i <= 10; i++ {
		// 按照指定格式发送数据给客户端
		fmt.Fprintf(w, "data: Event %d\n\n", i)
		w.(http.Flusher).Flush() // 立即将数据发送到客户端

		time.Sleep(1 * time.Second) // 模拟推送间隔
	}
}
func main() {
	fmt.Println(os.Getwd())
	fs := http.FileServer(http.Dir("static/"))
	mux1 := mux.NewRouter()
	mux1.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	mux1.HandleFunc("/events/", sendEvent)
	log.Println("http://127.0.0.1:8081/static/")
	server := http.Server{
		Addr:    ":8081",
		Handler: mux1,
	}
	log.Fatal(server.ListenAndServe())
	//log.Fatal(http.ListenAndServe(":8081", nil))
}
