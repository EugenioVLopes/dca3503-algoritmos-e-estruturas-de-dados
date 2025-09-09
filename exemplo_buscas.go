package main

import (
	"fmt"
	"dca3503/buscas"
)

func main() {
	// Criando uma lista ordenada de exemplo
	lista := []int{10, 20, 30, 40, 50}

	// Elemento a ser buscado
	elemento := 30

	fmt.Println("=== DEMONSTRAÇÃO DE ALGORITMOS DE BUSCA ===")
	
	// Utilizando a função de busca sequencial do pacote buscas
	fmt.Println("\nBusca Sequencial:")
	indiceSeq := buscas.BuscaSequencial(lista, elemento)

	// Exibindo o resultado da busca sequencial
	if indiceSeq != -1 {
		fmt.Printf("Elemento %d encontrado na posição %d\n", elemento, indiceSeq)
	} else {
		fmt.Printf("Elemento %d não encontrado na lista\n", elemento)
	}

	// Utilizando a função de busca binária do pacote buscas
	fmt.Println("\nBusca Binária:")
	indiceBin := buscas.BuscaBinaria(lista, elemento)

	// Exibindo o resultado da busca binária
	if indiceBin != -1 {
		fmt.Printf("Elemento %d encontrado na posição %d\n", elemento, indiceBin)
	} else {
		fmt.Printf("Elemento %d não encontrado na lista\n", elemento)
	}

	// Demonstrando busca de elemento inexistente
	elementoInexistente := 35
	fmt.Printf("\nBuscando elemento inexistente %d:\n", elementoInexistente)
	
	indiceSeq = buscas.BuscaSequencial(lista, elementoInexistente)
	indiceBin = buscas.BuscaBinaria(lista, elementoInexistente)
	
	fmt.Printf("Busca Sequencial: %d\n", indiceSeq)
	fmt.Printf("Busca Binária: %d\n", indiceBin)
}