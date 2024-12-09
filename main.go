package main

import (
	"fmt"
)

func gerarCombinacoes(tamanho int, contador int) []string {
	var combinacoes []string

	base := make([]rune, tamanho)

	for i := 0; i < tamanho; i++ {
		base[i] = '+'
	}

	gerarCombinacoesRecursiva(&combinacoes, base, contador, 0)

	return combinacoes
}

func gerarCombinacoesRecursiva(combinacoes *[]string, base []rune, contador, start int) {
	if contador == 0 {
		*combinacoes = append(*combinacoes, string(base))
		return
	}

	for i := start; i < len(base); i++ {
		base[i] = '*'
		gerarCombinacoesRecursiva(combinacoes, base, contador-1, i+1)
		base[i] = '+'
	}
}

func main() {
	// Exemplo de uso da função
	tamanho := 5
	contador := 2

	// Gerar as combinações
	combinacoes := gerarCombinacoes(tamanho, contador)

	// Imprimir as combinações
	for _, combinacao := range combinacoes {
		fmt.Println(combinacao)
	}
}
