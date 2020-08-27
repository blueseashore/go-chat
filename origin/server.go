package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func handler(ws *websocket.Conn) {
	// 获取请求头里的授权信息
	var auth = ws.Request().Header.Get("Authorization")
	if auth == "" {
		fmt.Println("授权信息为空")
	} else {
		fmt.Println("授权信息:" + auth)
	}
	fmt.Printf("新连接")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
		fmt.Println(string(buf))
		_, _ = ws.Write([]byte("i know"))
	}
	fmt.Printf("=>关闭连接")
	_ = ws.Close()
}

func main() {
	http.Handle("/websocket", websocket.Handler(handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
