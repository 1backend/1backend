package router

import (
	"fmt"
)

var port = "58231"

const address = "http://127.0.0.1"

func GetPort() string {
	return port
}

func SetPort(i int) {
	port = fmt.Sprintf("%v", i)
}

func SelfAddress() string {
	return fmt.Sprintf("%v:%v", address, GetPort())
}
