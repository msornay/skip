package main

import (
	"fmt"

	"github.com/pkg/taptun"
)

func main() {
	tun, err := taptun.OpenTun()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	defer tun.Close()
	packet := make([]byte, 2000)
	for {
		n, err := tun.Read(packet)
		if err != nil {
			fmt.Printf("%+v", err)
			return
		}
		fmt.Printf("Read: %d\n", n)
		fmt.Println(packet[:n])
	}
}
