package main

import "fmt"

// Variavel a nivel de packcage, variaveis que sao declaradas fora de uma funcao estao disponiveis em todo o package
var nome string = "Ana"

// Forma de declarar mais de uma variavel global
var (
	nome1 string = "Ana"
	nome2 string = "Maria"
	idade int    = 20
)

func main() {
	// Variavel em go deve ser declarada e usada (utilizada)
	// Declaracao de variavel local
	nome1 = "Joao"

	// Declaracao de variavel local com inferencia de tipo e atribuicao de valor
	idade2 := 30
	nome3 := "Pedro"

	var total int = 10
	total++

	//Declaracao de variavel local sem atribuicao de valor
	var total2 int
	fmt.Println("total:", total, "total2:", total2, "ciclo 1")
	// Declaraçao com multiplas atribuicoes
	total2, total = total, total2
	var a, b, c = true, 2, "string"

	fmt.Println("total:", total, "total2:", total2)
	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("Hello", nome, nome1, nome2, idade, idade2, nome3)
}
