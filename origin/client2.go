package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

type Rwc interface {
	io.ReadWriteCloser
}



func main() {
	var location = &url2.URL{
		Scheme:     "",
		Opaque:     "",
		User:       nil,
		Host:       "ws://localhost:8888",
		Path:       "/websocket",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}
	var origin = &url2.URL{
		Scheme:     "",
		Opaque:     "",
		User:       nil,
		Host:       "http://localhost/",
		Path:       "",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}
	var header = http.Header{}
	header.Add("Authorization", "你好啊")

	var config = &websocket.Config{
		Location:  location,
		Origin:    origin,
		Protocol:  nil,
		Version:   13,
		TlsConfig: nil,
		Header:    header,
		Dialer:    nil,
	}
	var rwc = io.ReadWriteCloser{}
	ws, err := websocket.NewClient(config, rwc)
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
