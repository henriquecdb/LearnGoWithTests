package main

import "fmt"

const prefixHelloPortuguese = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixHelloPortuguese + nome
}

func main() {
	fmt.Println(Ola("Chris"))
}

