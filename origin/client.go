package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

// 设置默认的监听地址
var url = flag.String("u", "ws://localhost:8888/websocket", "监听地址")

// 设置默认的协议
var protocol = flag.String("p", "", "协议")

// 设置默认的来源
var origin = flag.String("o", "http://localhost/", "来源")

func main() {
	// flag解析
	flag.Parse()
	// 建立连接
	ws, err := websocket.Dial(*url, *protocol, *origin)
	if err != nil {
		log.Fatal(err)
	}
	// 处理消息
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
