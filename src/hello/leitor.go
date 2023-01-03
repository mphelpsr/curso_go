package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)

	linha, _ := leitor.ReadString('\n')
	fmt.Println(linha)

}
