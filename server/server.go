package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro ao criar o servidor:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Servidor aguardando conexões...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Erro na conexão:", err)
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro na leitura:", err)
		return
	}

	message := string(buffer)
	fmt.Printf("Mensagem recebida: %s", message)
	conn.Write([]byte("Pong"))

	// Adicione um loop para lidar com mensagens bidirecionais
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erro na leitura:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Printf("Mensagem recebida: %s", message)
	}
}
