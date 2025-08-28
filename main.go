package main

import (
	"fmt"
	"time"
)

// ============================================================================
// PROGRAMA PRINCIPAL - DEMONSTRAÇÃO E TESTES
// ============================================================================

func main() {
	fmt.Println("=== DEMONSTRAÇÃO DE ESTRUTURAS DE DADOS ==="\n")
	
	// Demonstrações básicas
	demonstrateArrayList()
	demonstrateLinkedList()
	
	// Demonstrações de pilhas
	demonstrateArrayStack()
	demonstrateLinkedStack()
	
	// Demonstrações de filas
	demonstrateArrayQueue()
	demonstrateLinkedQueue()
	
	// Comparação de performance
	comparePerformance()
	compareStackPerformance()
	compareQueuePerformance()
	
	// Demonstração da interface
	demonstrateInterface()
	demonstrateStackInterface()
	demonstrateQueueInterface()
	
	// Algoritmos usando a interface
	demonstrateAlgorithms()
	demonstrateStackAlgorithms()
	demonstrateQueueAlgorithms()
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
// DEMONSTRAÇÃO ARRAYQUEUE
// ============================================================================

func demonstrateArrayQueue() {
	fmt.Println("=== DEMONSTRAÇÃO ARRAYQUEUE ===")
	
	// Criação e inicialização
	aq := NewArrayQueue(5)
	fmt.Printf("ArrayQueue criado com capacidade: %d\n", aq.Capacity())
	
	// Adicionando elementos (Enqueue)
	fmt.Println("\nAdicionando elementos de 1 a 8...")
	for i := 1; i <= 8; i++ {
		aq.Enqueue(i)
		fmt.Printf("Enqueue %d: %s\n", i, aq.String())
	}
	fmt.Printf("Tamanho: %d, Capacidade: %d\n", aq.Size(), aq.Capacity())
	
	// Testando Front e Rear
	fmt.Println("\nTestando Front e Rear:")
	front, _ := aq.Front()
	rear, _ := aq.Rear()
	fmt.Printf("Elemento na frente: %d\n", front)
	fmt.Printf("Elemento no final: %d\n", rear)
	
	// Removendo elementos (Dequeue)
	fmt.Println("\nRemovendo 3 elementos:")
	for i := 0; i < 3; i++ {
		value, err := aq.Dequeue()
		if err == nil {
			fmt.Printf("Dequeue: %d, Fila: %s\n", value, aq.String())
		}
	}
	
	// Testando operações auxiliares
	fmt.Println("\nOperações auxiliares:")
	fmt.Printf("Contém 5? %t\n", aq.Contains(5))
	fmt.Printf("Índice do 5: %d\n", aq.IndexOf(5))
	fmt.Printf("ToSlice: %v\n", aq.ToSlice())
	
	// Testando rotação
	fmt.Println("\nTestando rotação (2 posições):")
	fmt.Printf("Antes: %s\n", aq.String())
	aq.Rotate(2)
	fmt.Printf("Depois: %s\n", aq.String())
	
	// Estatísticas
	fmt.Println("\nEstatísticas:")
	stats := aq.GetStatistics()
	for key, value := range stats {
		fmt.Printf("%s: %v\n", key, value)
	}
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO LINKEDQUEUE
// ============================================================================

func demonstrateLinkedQueue() {
	fmt.Println("=== DEMONSTRAÇÃO LINKEDQUEUE ===")
	
	// Criação
	lq := NewLinkedQueue()
	fmt.Println("LinkedQueue criado")
	
	// Adicionando elementos
	fmt.Println("\nAdicionando elementos de 10 a 16...")
	for i := 10; i <= 16; i++ {
		lq.Enqueue(i)
		fmt.Printf("Enqueue %d: %s\n", i, lq.String())
	}
	
	// Testando operações
	fmt.Println("\nTestando operações:")
	front, _ := lq.Front()
	rear, _ := lq.Rear()
	fmt.Printf("Frente: %d, Final: %d\n", front, rear)
	fmt.Printf("Tamanho: %d\n", lq.Size())
	fmt.Printf("Vazia? %t\n", lq.IsEmpty())
	
	// Clonando fila
	fmt.Println("\nClonando fila:")
	clone := lq.Clone()
	fmt.Printf("Original: %s\n", lq.String())
	fmt.Printf("Clone: %s\n", clone.String())
	fmt.Printf("São iguais? %t\n", lq.Equals(clone))
	
	// Removendo alguns elementos
	fmt.Println("\nRemovendo 4 elementos:")
	for i := 0; i < 4; i++ {
		value, _ := lq.Dequeue()
		fmt.Printf("Dequeue: %d, Fila: %s\n", value, lq.String())
	}
	
	// Testando métodos avançados
	fmt.Println("\nMétodos avançados:")
	filtered := lq.Filter(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Elementos pares: %s\n", filtered.String())
	
	mapped := lq.Map(func(x int) int { return x * 2 })
	fmt.Printf("Elementos dobrados: %s\n", mapped.String())
	
	sum := lq.Reduce(func(acc, x int) int { return acc + x }, 0)
	fmt.Printf("Soma dos elementos: %d\n", sum)
	
	// Testando split
	fmt.Println("\nTestando split no índice 1:")
	first, second := lq.Split(1)
	fmt.Printf("Primeira parte: %s\n", first.String())
	fmt.Printf("Segunda parte: %s\n", second.String())
	
	fmt.Println()
}

// ============================================================================
// COMPARAÇÃO DE PERFORMANCE - QUEUES
// ============================================================================

func compareQueuePerformance() {
	fmt.Println("=== COMPARAÇÃO DE PERFORMANCE - QUEUES ===")
	
	const numOperations = 100000
	
	fmt.Printf("Testando %d operações Enqueue/Dequeue...\n\n", numOperations)
	
	// ArrayQueue
	fmt.Println("ArrayQueue:")
	benchmarkFunction("  Enqueue", func() {
		aq := NewArrayQueue(10)
		for i := 0; i < numOperations; i++ {
			aq.Enqueue(i)
		}
	})
	
	benchmarkFunction("  Enqueue+Dequeue", func() {
		aq := NewArrayQueue(10)
		for i := 0; i < numOperations; i++ {
			aq.Enqueue(i)
		}
		for i := 0; i < numOperations; i++ {
			aq.Dequeue()
		}
	})
	
	// LinkedQueue
	fmt.Println("\nLinkedQueue:")
	benchmarkFunction("  Enqueue", func() {
		lq := NewLinkedQueue()
		for i := 0; i < numOperations; i++ {
			lq.Enqueue(i)
		}
	})
	
	benchmarkFunction("  Enqueue+Dequeue", func() {
		lq := NewLinkedQueue()
		for i := 0; i < numOperations; i++ {
			lq.Enqueue(i)
		}
		for i := 0; i < numOperations; i++ {
			lq.Dequeue()
		}
	})
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO DA INTERFACE QUEUE
// ============================================================================

func demonstrateQueueInterface() {
	fmt.Println("=== DEMONSTRAÇÃO DA INTERFACE QUEUE ===")
	
	// Criando diferentes implementações
	var queues []Queue
	queues = append(queues, NewArrayQueue(5))
	queues = append(queues, NewLinkedQueue())
	
	names := []string{"ArrayQueue", "LinkedQueue"}
	
	// Testando polimorfismo
	for i, queue := range queues {
		fmt.Printf("\nTestando %s:\n", names[i])
		
		// Adicionando elementos
		for j := 1; j <= 5; j++ {
			queue.Enqueue(j * 10)
		}
		
		PrintQueue(queue, names[i])
		
		// Testando operações
		front, _ := queue.Front()
		rear, _ := queue.Rear()
		fmt.Printf("Frente: %d, Final: %d\n", front, rear)
		
		// Removendo elementos
		for j := 0; j < 2; j++ {
			value, _ := queue.Dequeue()
			fmt.Printf("Removido: %d\n", value)
		}
		
		PrintQueue(queue, names[i])
	}
	
	fmt.Println()
}

// ============================================================================
// ALGORITMOS USANDO QUEUES
// ============================================================================

func demonstrateQueueAlgorithms() {
	fmt.Println("=== ALGORITMOS USANDO QUEUES ===")
	
	// 1. Geração de números binários
	fmt.Println("\n1. Geração de Números Binários (1 a 10):")
	binaryNumbers := GenerateBinaryNumbers(10)
	for i, binary := range binaryNumbers {
		fmt.Printf("%d -> %s\n", i+1, binary)
	}
	
	// 2. Busca em largura (simulação)
	fmt.Println("\n2. Busca em Largura (árvore como array):")
	tree := []int{1, 2, 3, 4, 5, 6, 7} // Árvore binária completa
	traversal := LevelOrderTraversal(tree)
	fmt.Printf("Árvore: %v\n", tree)
	fmt.Printf("Travessia em largura: %v\n", traversal)
	
	// 3. Primeiro caractere não repetido
	fmt.Println("\n3. Primeiro Caractere Não Repetido:")
	streams := []string{"abccba", "abcabc", "aabc"}
	for _, stream := range streams {
		result := FirstNonRepeatingCharacter(stream)
		fmt.Printf("Stream: %s\n", stream)
		fmt.Print("Resultado: ")
		for _, char := range result {
			if char == 0 {
				fmt.Print("- ")
			} else {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}
	
	// 4. Rotação de fila
	fmt.Println("\n4. Rotação de Fila:")
	queue := NewArrayQueue(10)
	for i := 1; i <= 6; i++ {
		queue.Enqueue(i)
	}
	fmt.Printf("Original: %s\n", queue.String())
	
	RotateQueue(queue, 2)
	fmt.Printf("Após rotação de 2: %s\n", queue.String())
	
	// 5. Intercalar fila
	fmt.Println("\n5. Intercalar Fila:")
	queue2 := NewLinkedQueue()
	for i := 1; i <= 6; i++ {
		queue2.Enqueue(i)
	}
	fmt.Printf("Antes de intercalar: %s\n", queue2.String())
	
	InterleaveQueue(queue2)
	fmt.Printf("Após intercalar: %s\n", queue2.String())
	
	// 6. Estatísticas da fila
	fmt.Println("\n6. Estatísticas da Fila:")
	queue3 := NewArrayQueue(10)
	for i := 10; i <= 50; i += 10 {
		queue3.Enqueue(i)
	}
	fmt.Printf("Fila: %s\n", queue3.String())
	
	max, _ := QueueMax(queue3)
	min, _ := QueueMin(queue3)
	sum := QueueSum(queue3)
	fmt.Printf("Máximo: %d\n", max)
	fmt.Printf("Mínimo: %d\n", min)
	fmt.Printf("Soma: %d\n", sum)
	
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

// ============================================================================
// DEMONSTRAÇÃO ARRAYSTACK
// ============================================================================

func demonstrateArrayStack() {
	fmt.Println("=== DEMONSTRAÇÃO ARRAYSTACK ===")
	
	// Criação e inicialização
	as := NewArrayStack(5)
	fmt.Printf("ArrayStack criado com capacidade: %d\n", as.Capacity())
	
	// Adicionando elementos (Push)
	fmt.Println("\nAdicionando elementos de 1 a 8...")
	for i := 1; i <= 8; i++ {
		as.Push(i)
		fmt.Printf("Push %d: %s\n", i, as.String())
	}
	fmt.Printf("Tamanho: %d, Capacidade: %d\n", as.Size(), as.Capacity())
	
	// Testando Peek
	fmt.Println("\nTestando Peek:")
	top, err := as.Peek()
	if err == nil {
		fmt.Printf("Elemento no topo: %d\n", top)
	}
	
	// Removendo elementos (Pop)
	fmt.Println("\nRemovendo 3 elementos:")
	for i := 0; i < 3; i++ {
		value, err := as.Pop()
		if err == nil {
			fmt.Printf("Pop: %d, Pilha: %s\n", value, as.String())
		}
	}
	
	// Testando operações auxiliares
	fmt.Println("\nOperações auxiliares:")
	fmt.Printf("Contém 3? %t\n", as.Contains(3))
	fmt.Printf("Posição do 3: %d\n", as.Search(3))
	fmt.Printf("ToSlice: %v\n", as.ToSlice())
	
	// Estatísticas
	fmt.Println("\nEstatísticas:")
	stats := as.GetStatistics()
	for key, value := range stats {
		fmt.Printf("%s: %v\n", key, value)
	}
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO LINKEDSTACK
// ============================================================================

func demonstrateLinkedStack() {
	fmt.Println("=== DEMONSTRAÇÃO LINKEDSTACK ===")
	
	// Criação
	ls := NewLinkedStack()
	fmt.Println("LinkedStack criado")
	
	// Adicionando elementos
	fmt.Println("\nAdicionando elementos de 10 a 16...")
	for i := 10; i <= 16; i++ {
		ls.Push(i)
		fmt.Printf("Push %d: %s\n", i, ls.String())
	}
	
	// Testando operações
	fmt.Println("\nTestando operações:")
	top, _ := ls.Peek()
	fmt.Printf("Topo: %d\n", top)
	fmt.Printf("Tamanho: %d\n", ls.Size())
	fmt.Printf("Vazia? %t\n", ls.IsEmpty())
	
	// Clonando pilha
	fmt.Println("\nClonando pilha:")
	clone := ls.Clone()
	fmt.Printf("Original: %s\n", ls.String())
	fmt.Printf("Clone: %s\n", clone.String())
	fmt.Printf("São iguais? %t\n", ls.Equals(clone))
	
	// Removendo alguns elementos
	fmt.Println("\nRemovendo 4 elementos:")
	for i := 0; i < 4; i++ {
		value, _ := ls.Pop()
		fmt.Printf("Pop: %d, Pilha: %s\n", value, ls.String())
	}
	
	// Testando métodos avançados
	fmt.Println("\nMétodos avançados:")
	filtered := ls.Filter(func(x int) bool { return x%2 == 0 })
	fmt.Printf("Elementos pares: %s\n", filtered.String())
	
	mapped := ls.Map(func(x int) int { return x * 2 })
	fmt.Printf("Elementos dobrados: %s\n", mapped.String())
	
	sum := ls.Reduce(func(acc, x int) int { return acc + x }, 0)
	fmt.Printf("Soma dos elementos: %d\n", sum)
	
	fmt.Println()
}

// ============================================================================
// COMPARAÇÃO DE PERFORMANCE - STACKS
// ============================================================================

func compareStackPerformance() {
	fmt.Println("=== COMPARAÇÃO DE PERFORMANCE - STACKS ===")
	
	const numOperations = 100000
	
	fmt.Printf("Testando %d operações Push/Pop...\n\n", numOperations)
	
	// ArrayStack
	fmt.Println("ArrayStack:")
	benchmarkFunction("  Push", func() {
		as := NewArrayStack(10)
		for i := 0; i < numOperations; i++ {
			as.Push(i)
		}
	})
	
	benchmarkFunction("  Push+Pop", func() {
		as := NewArrayStack(10)
		for i := 0; i < numOperations; i++ {
			as.Push(i)
		}
		for i := 0; i < numOperations; i++ {
			as.Pop()
		}
	})
	
	// LinkedStack
	fmt.Println("\nLinkedStack:")
	benchmarkFunction("  Push", func() {
		ls := NewLinkedStack()
		for i := 0; i < numOperations; i++ {
			ls.Push(i)
		}
	})
	
	benchmarkFunction("  Push+Pop", func() {
		ls := NewLinkedStack()
		for i := 0; i < numOperations; i++ {
			ls.Push(i)
		}
		for i := 0; i < numOperations; i++ {
			ls.Pop()
		}
	})
	
	fmt.Println()
}

// ============================================================================
// DEMONSTRAÇÃO DA INTERFACE STACK
// ============================================================================

func demonstrateStackInterface() {
	fmt.Println("=== DEMONSTRAÇÃO DA INTERFACE STACK ===")
	
	// Criando diferentes implementações
	var stacks []Stack
	stacks = append(stacks, NewArrayStack(5))
	stacks = append(stacks, NewLinkedStack())
	
	names := []string{"ArrayStack", "LinkedStack"}
	
	// Testando polimorfismo
	for i, stack := range stacks {
		fmt.Printf("\nTestando %s:\n", names[i])
		
		// Adicionando elementos
		for j := 1; j <= 5; j++ {
			stack.Push(j * 10)
		}
		
		PrintStack(stack, names[i])
		
		// Testando operações
		top, _ := stack.Peek()
		fmt.Printf("Topo: %d\n", top)
		
		// Removendo elementos
		for j := 0; j < 2; j++ {
			value, _ := stack.Pop()
			fmt.Printf("Removido: %d\n", value)
		}
		
		PrintStack(stack, names[i])
	}
	
	fmt.Println()
}

// ============================================================================
// ALGORITMOS USANDO STACKS
// ============================================================================

func demonstrateStackAlgorithms() {
	fmt.Println("=== ALGORITMOS USANDO STACKS ===")
	
	// 1. Verificação de parênteses balanceados
	fmt.Println("\n1. Verificação de Parênteses Balanceados:")
	testCases := []string{
		"()",
		"()[]{}",
		"([{}])",
		"([)]",
		"(((",
		")))",
	}
	
	for _, test := range testCases {
		result := IsValidParentheses(test)
		fmt.Printf("'%s' -> %t\n", test, result)
	}
	
	// 2. Avaliação de expressão pós-fixa
	fmt.Println("\n2. Avaliação de Expressão Pós-fixa:")
	postfixExpressions := [][]string{
		{"3", "4", "+"},                    // 3 + 4 = 7
		{"3", "4", "+", "2", "*"},          // (3 + 4) * 2 = 14
		{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}, // Complexa
	}
	
	for _, expr := range postfixExpressions {
		result, err := EvaluatePostfix(expr)
		if err == nil {
			fmt.Printf("%v -> %d\n", expr, result)
		} else {
			fmt.Printf("%v -> Erro: %v\n", expr, err)
		}
	}
	
	// 3. Inversão usando pilha
	fmt.Println("\n3. Inversão usando Pilha:")
	original := NewArrayStack(10)
	for i := 1; i <= 5; i++ {
		original.Push(i)
	}
	fmt.Printf("Original: %s\n", original.String())
	
	ReverseStack(original)
	fmt.Printf("Invertida: %s\n", original.String())
	
	// 4. Busca em pilha
	fmt.Println("\n4. Busca em Pilha:")
	stack := NewLinkedStack()
	for i := 10; i <= 50; i += 10 {
		stack.Push(i)
	}
	fmt.Printf("Pilha: %s\n", stack.String())
	
	targets := []int{30, 60, 10}
	for _, target := range targets {
		found := FindInStack(stack, target)
		fmt.Printf("Buscar %d: %t\n", target, found)
	}
	
	// 5. Estatísticas da pilha
	fmt.Println("\n5. Estatísticas da Pilha:")
	max, _ := StackMax(stack)
	sum := StackSum(stack)
	fmt.Printf("Máximo: %d\n", max)
	fmt.Printf("Soma: %d\n", sum)
	
	fmt.Println()
}