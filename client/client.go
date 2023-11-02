package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro na conexão:", err)
		return
	}
	defer conn.Close()

	message := "Ping"
	conn.Write([]byte(message))
	fmt.Println("Mensagem enviada para o servidor:", message)

	// Leitura da resposta do servidor
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println("Resposta do servidor:", string(buffer))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Digite sua mensagem: ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))

	}

}
