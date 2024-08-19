package main

import (
	"fmt"

	"github.com/dns_service/pkg/dns"
)

func main() {
	fmt.Println("Running DNS Server")
	go dns.StartServer("127.0.0.1", "53", "udp")
	select {}
}
