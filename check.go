package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination, port string) string {
	// address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", destination, port), timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is not reachable. \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable. \n From: %v \n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
