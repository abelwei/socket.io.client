package main

import (
	"bufio"
	"fmt"
	ioclient "go-socket.io-client"
	"log"
	"os"
)


func main() {

	opts := &ioclient.Options{
		//Transport:"polling",
		Transport:"websocket",
		Query:     make(map[string]string),
	}
	//opts.Query["user"] = "user"
	opts.Query["token"] = "b978ea49-7a18-4755-9ae0-e181eddf006d"
	uri := "http://127.0.0.1:8890"

	client, err := ioclient.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		str := `{}`
		str = fmt.Sprintf(str, command, command)
		client.Emit("sendMsg", str)
		log.Printf("send message:%v\n", command)
	}
}
