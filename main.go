package main

import (
	"fmt"
	"time"
)

// ============================================================================
// PROGRAMA PRINCIPAL - DEMONSTRAÇÃO E TESTES
// ============================================================================

func main() {
	fmt.Println("=== DEMONSTRAÇÃO DE ESTRUTURAS DE DADOS ===\n")
	
	// Demonstrações básicas
	demonstrateArrayList()
	demonstrateLinkedList()
	
	// Comparação de performance
	comparePerformance()
	
	// Demonstração da interface
	demonstrateInterface()
	
	// Algoritmos usando a interface
	demonstrateAlgorithms()
}

// ============================================================================
// DEMONSTRAÇÃO ARRAYLIST
// ============================================================================

func demonstrateArrayList() {
	fmt.Println("=== DEMONSTRAÇÃO ARRAYLIST ===")
	
	// Criação e inicialização
	al := NewArrayList(5)
	fmt.Printf("ArrayList criado com capacidade: %d\n", al.Capacity())
	
	// Adicionando elementos
	fmt.Println("\nAdicionando elementos de 1 a 10...")
	for i := 1; i <= 10; i++ {
		al.Add(i)
	}
	fmt.Printf("Após inserções: %s\n", al.String())
	fmt.Printf("Tamanho: %d, Capacidade: %d\n", al.Size(), al.Capacity())
	
	// Testando acesso
	fmt.Println("\nTestando acesso por índice:")
	for i := 0; i < 3; i++ {
		val, _ := al.Get(i)
		fmt.Printf("Elemento no índice %d: %d\n", i, val)
	}
	
	// Inserção em posição específica
	fmt.Println("\nInserindo -1 no início...")
	al.AddOnIndex(-1, 0)
	fmt.Printf("Após inserção: %s\n", al.String())
	
	// Inserção no meio
	fmt.Println("\nInserindo 99 no meio (índice 5)...")
	al.AddOnIndex(99, 5)
	fmt.Printf("Após inserção: %s\n", al.String())
	
	// Remoção
	fmt.Println("\nRemovendo primeiro elemento...")
	al.Remove(0)
	fmt.Printf("Após remoção: %s\n", al.String())
	
	// Busca
	fmt.Println("\nTestando busca:")
	fmt.Printf("Contém 99? %t\n", al.Contains(99))
	fmt.Printf("Índice do 99: %d\n", al.IndexOf(99))
	fmt.Printf("Contém 999? %t\n", al.Contains(999))
	
	// Operações em lote
	fmt.Println("\nAdicionando elementos em lote...")
	al.AddAll([]int{100, 200, 300})
	fmt.Printf("Após AddAll: %s\n", al.String())
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO LINKEDLIST
// ============================================================================

func demonstrateLinkedList() {
	fmt.Println("=== DEMONSTRAÇÃO LINKEDLIST ===")
	
	// Criação
	ll := NewLinkedList()
	fmt.Println("LinkedList criada")
	
	// Adicionando elementos no final
	fmt.Println("\nAdicionando elementos de 1 a 5 no final...")
	for i := 1; i <= 5; i++ {
		ll.Add(i)
	}
	fmt.Printf("Após inserções: %s\n", ll.String())
	
	// Adicionando elementos no início
	fmt.Println("\nAdicionando elementos no início...")
	ll.AddFirst(0)
	ll.AddFirst(-1)
	fmt.Printf("Após inserções no início: %s\n", ll.String())
	
	// Testando acesso
	fmt.Println("\nTestando acesso por índice:")
	for i := 0; i < 3; i++ {
		val, _ := ll.Get(i)
		fmt.Printf("Elemento no índice %d: %d\n", i, val)
	}
	
	// Inserção em posição específica
	fmt.Println("\nInserindo 99 no meio (índice 3)...")
	ll.AddOnIndex(99, 3)
	fmt.Printf("Após inserção: %s\n", ll.String())
	
	// Remoção
	fmt.Println("\nRemovendo primeiro elemento...")
	removedVal, _ := ll.RemoveFirst()
	fmt.Printf("Elemento removido: %d\n", removedVal)
	fmt.Printf("Após remoção: %s\n", ll.String())
	
	// Busca
	fmt.Println("\nTestando busca:")
	fmt.Printf("Contém 99? %t\n", ll.Contains(99))
	fmt.Printf("Índice do 99: %d\n", ll.IndexOf(99))
	
	// Algoritmos especiais
	fmt.Println("\nAlgoritmos especiais:")
	middle, _ := ll.GetMiddle()
	fmt.Printf("Elemento do meio: %d\n", middle)
	
	// Invertendo a lista
	fmt.Printf("Antes de inverter: %s\n", ll.String())
	ll.Reverse()
	fmt.Printf("Após inverter: %s\n", ll.String())
	
	// Testando duplicatas
	ll.Add(1)
	ll.Add(2)
	ll.Add(1)
	fmt.Printf("Com duplicatas: %s\n", ll.String())
	ll.RemoveDuplicates()
	fmt.Printf("Sem duplicatas: %s\n", ll.String())
	
	fmt.Println()
}

// ============================================================================
// COMPARAÇÃO DE PERFORMANCE
// ============================================================================

func comparePerformance() {
	fmt.Println("=== COMPARAÇÃO DE PERFORMANCE ===")
	
	const numElements = 10000
	
	// Teste 1: Inserção no final
	fmt.Printf("\nTeste 1: Inserção de %d elementos no final\n", numElements)
	
	// ArrayList
	al := NewArrayList(100)
	start := time.Now()
	for i := 0; i < numElements; i++ {
		al.Add(i)
	}
	arrayListTime := time.Since(start)
	fmt.Printf("ArrayList: %v\n", arrayListTime)
	
	// LinkedList
	ll := NewLinkedList()
	start = time.Now()
	for i := 0; i < numElements; i++ {
		ll.Add(i)
	}
	linkedListTime := time.Since(start)
	fmt.Printf("LinkedList: %v\n", linkedListTime)
	
	if arrayListTime < linkedListTime {
		fmt.Printf("ArrayList é %.2fx mais rápido\n", float64(linkedListTime)/float64(arrayListTime))
	} else {
		fmt.Printf("LinkedList é %.2fx mais rápido\n", float64(arrayListTime)/float64(linkedListTime))
	}
	
	// Teste 2: Inserção no início
	fmt.Printf("\nTeste 2: Inserção de 1000 elementos no início\n")
	const numInsertions = 1000
	
	// ArrayList
	al2 := NewArrayList(10)
	start = time.Now()
	for i := 0; i < numInsertions; i++ {
		al2.AddOnIndex(i, 0)
	}
	arrayListTime = time.Since(start)
	fmt.Printf("ArrayList: %v\n", arrayListTime)
	
	// LinkedList
	ll2 := NewLinkedList()
	start = time.Now()
	for i := 0; i < numInsertions; i++ {
		ll2.AddFirst(i)
	}
	linkedListTime = time.Since(start)
	fmt.Printf("LinkedList: %v\n", linkedListTime)
	
	if arrayListTime < linkedListTime {
		fmt.Printf("ArrayList é %.2fx mais rápido\n", float64(linkedListTime)/float64(arrayListTime))
	} else {
		fmt.Printf("LinkedList é %.2fx mais rápido\n", float64(arrayListTime)/float64(linkedListTime))
	}
	
	// Teste 3: Acesso aleatório
	fmt.Printf("\nTeste 3: 1000 acessos aleatórios\n")
	const numAccesses = 1000
	
	// ArrayList
	start = time.Now()
	for i := 0; i < numAccesses; i++ {
		index := i % al.Size()
		al.Get(index)
	}
	arrayListTime = time.Since(start)
	fmt.Printf("ArrayList: %v\n", arrayListTime)
	
	// LinkedList
	start = time.Now()
	for i := 0; i < numAccesses; i++ {
		index := i % ll.Size()
		ll.Get(index)
	}
	linkedListTime = time.Since(start)
	fmt.Printf("LinkedList: %v\n", linkedListTime)
	
	if arrayListTime < linkedListTime {
		fmt.Printf("ArrayList é %.2fx mais rápido\n", float64(linkedListTime)/float64(arrayListTime))
	} else {
		fmt.Printf("LinkedList é %.2fx mais rápido\n", float64(arrayListTime)/float64(linkedListTime))
	}
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO DA INTERFACE
// ============================================================================

func demonstrateInterface() {
	fmt.Println("=== DEMONSTRAÇÃO DA INTERFACE LIST ===")
	
	// Criando diferentes implementações
	var list1 List = NewArrayList(5)
	var list2 List = NewLinkedList()
	
	// Adicionando elementos usando a interface
	for i := 1; i <= 5; i++ {
		list1.Add(i * 10)
		list2.Add(i * 20)
	}
	
	PrintList(list1, "ArrayList")
	PrintList(list2, "LinkedList")
	
	// Copiando de uma lista para outra
	fmt.Println("\nCopiando ArrayList para LinkedList...")
	CopyList(list1, list2)
	PrintList(list2, "LinkedList após cópia")
	
	// Testando funções utilitárias
	fmt.Println("\nTestando funções utilitárias:")
	max, _ := FindMax(list1)
	min, _ := FindMin(list1)
	fmt.Printf("Máximo: %d, Mínimo: %d\n", max, min)
	fmt.Printf("Soma: %d, Média: %.2f\n", Sum(list1), Average(list1))
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO DE ALGORITMOS
// ============================================================================

func demonstrateAlgorithms() {
	fmt.Println("=== DEMONSTRAÇÃO DE ALGORITMOS ===")
	
	// Criando lista para ordenação
	var list List = NewArrayList(10)
	elements := []int{64, 34, 25, 12, 22, 11, 90, 5, 77, 30}
	
	for _, elem := range elements {
		list.Add(elem)
	}
	
	fmt.Printf("Lista original: %s\n", list.String())
	fmt.Printf("Está ordenada? %t\n", IsSorted(list))
	
	// Ordenando com bubble sort
	fmt.Println("\nOrdenando com Bubble Sort...")
	BubbleSort(list)
	fmt.Printf("Lista ordenada: %s\n", list.String())
	fmt.Printf("Está ordenada? %t\n", IsSorted(list))
	
	// Busca binária
	fmt.Println("\nTestando busca binária:")
	targets := []int{25, 77, 100}
	for _, target := range targets {
		index := BinarySearch(list, target)
		if index != -1 {
			fmt.Printf("Elemento %d encontrado no índice %d\n", target, index)
		} else {
			fmt.Printf("Elemento %d não encontrado\n", target)
		}
	}
	
	// Mesclando listas ordenadas
	fmt.Println("\nMesclando duas listas ordenadas:")
	var list1 List = NewArrayList(5)
	var list2 List = NewLinkedList()
	var result List = NewArrayList(10)
	
	// Lista 1: números pares
	for i := 2; i <= 10; i += 2 {
		list1.Add(i)
	}
	
	// Lista 2: números ímpares
	for i := 1; i <= 9; i += 2 {
		list2.Add(i)
	}
	
	PrintList(list1, "Lista 1 (pares)")
	PrintList(list2, "Lista 2 (ímpares)")
	
	MergeSorted(list1, list2, result)
	PrintList(result, "Lista mesclada")
	
	// Removendo elementos
	fmt.Println("\nTestando remoção de elementos:")
	list.Add(25) // Adicionar duplicata
	list.Add(25) // Adicionar outra duplicata
	fmt.Printf("Com duplicatas: %s\n", list.String())
	
	removed := RemoveAll(list, 25)
	fmt.Printf("Removidas %d ocorrências de 25\n", removed)
	fmt.Printf("Após remoção: %s\n", list.String())
	
	fmt.Println()
}

// ============================================================================
// FUNÇÕES AUXILIARES PARA DEMONSTRAÇÃO
// ============================================================================

// demonstrateMemoryUsage mostra o uso de memória das estruturas
func demonstrateMemoryUsage() {
	fmt.Println("=== ANÁLISE DE USO DE MEMÓRIA ===")
	
	const numElements = 1000
	
	// ArrayList
	al := NewArrayList(numElements)
	for i := 0; i < numElements; i++ {
		al.Add(i)
	}
	
	// LinkedList
	ll := NewLinkedList()
	for i := 0; i < numElements; i++ {
		ll.Add(i)
	}
	
	fmt.Printf("ArrayList - Elementos: %d, Capacidade: %d\n", al.Size(), al.Capacity())
	fmt.Printf("Uso teórico de memória:\n")
	fmt.Printf("  ArrayList: ~%d bytes\n", al.Capacity()*4+16) // 4 bytes por int + overhead
	fmt.Printf("  LinkedList: ~%d bytes\n", ll.Size()*12+16)   // 4 bytes + 8 bytes ponteiro + overhead
	
	fmt.Println()
}

// benchmarkFunction executa uma função e mede o tempo
func benchmarkFunction(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s: %v\n", name, duration)
}