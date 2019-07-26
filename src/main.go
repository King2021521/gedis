package main

import (
	"tcp"
	"fmt"
	"template"
)

func main() {
	conn := tcp.Connect("10.10.5.239:6379")
	authResult := template.Auth("123456", conn)
	fmt.Println("auth result:" + authResult)
	sendResult := template.Set("name", "james", conn)
	fmt.Println("send result:" + sendResult)
	result := template.Get("name", conn)
	fmt.Println("get result:" + result)
}
