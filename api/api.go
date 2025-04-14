package api

import "fmt"

func HelloWorld() string {
	msg := "Hello from Go!"
	fmt.Println(msg)
	return msg
}
