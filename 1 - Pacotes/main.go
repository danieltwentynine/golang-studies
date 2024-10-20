package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail" 
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	validacao := checkmail.ValidateFormat("exemplo@gmail.com")
	fmt.Println(validacao)
}