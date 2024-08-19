package dns

import (
	"fmt"
	"log"
	"net"
)

func StartServer(loop_back_ip string, port string, serverType string) {
	address := fmt.Sprintf("%s:%s", loop_back_ip, port)
	server, err := net.Listen(serverType, address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()
	log.Printf("DNS server listening on %s", address)

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		log.Println("Client connected")
		go handleClient(connection)
	}
}

func handleClient(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 512) // DNS messages can be up to 512 bytes for UDP
	mLen, err := connection.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}

	log.Printf("Received: %s", string(buffer[:mLen]))

	response := "Transacting: " + string(buffer[:mLen])
	_, err = connection.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
	}
}
