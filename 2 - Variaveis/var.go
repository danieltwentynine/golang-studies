package main

import "fmt"

func main() {
	var variavel1 string = "isso é uma string com o tipo explicito" // declaração de variavel com tipo explicito
	variavel2 := "isso é uma variavel com inferencia de tipo" // declaração de variavel com inferencia de tipo
	fmt.Println(variavel1 + "\n----------\n" + variavel2) // print de todas as variaveis

	var ( // declaração de multiplas variaveis
		var1 string = "var1"
		var2 string = "var2"
	)

	fmt.Println(var1, var2)

	var3, var4, var5 := "var3", "var4", "var5" // declaração de multiplas variaveis com inferencia de tipo

	const constante string = "| uma constante" // declaração de constante

	fmt.Println(var3, var4, var5, constante) // print de todas as variaveis e constante

	var3, var4 = var4, var3 // troca de valores entre variaveis

	fmt.Println(var3, var4) // print das variaveis trocadas
}
