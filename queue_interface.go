package main

import "fmt"

// ============================================================================
// INTERFACE QUEUE - TIPO ABSTRATO DE DADOS
// ============================================================================

// Queue define o contrato que todas as implementações de fila devem seguir
// Uma fila é uma estrutura de dados FIFO (First In, First Out)
// O primeiro elemento inserido é o primeiro a ser removido
type Queue interface {
	// Operações básicas de fila
	Enqueue(element int)        // Adiciona elemento no final da fila
	Dequeue() (int, error)      // Remove e retorna elemento do início
	Front() (int, error)        // Retorna elemento do início sem remover
	Rear() (int, error)         // Retorna elemento do final sem remover
	
	// Operações de consulta
	Size() int                  // Retorna número de elementos na fila
	IsEmpty() bool             // Verifica se a fila está vazia
	IsFull() bool              // Verifica se a fila está cheia (para implementações com limite)
	
	// Operações auxiliares
	Clear()                    // Remove todos os elementos
	ToSlice() []int            // Converte para slice (do início para o final)
	String() string            // Representação em string
}

// ============================================================================
// FUNÇÕES UTILITÁRIAS QUE TRABALHAM COM A INTERFACE
// ============================================================================

// PrintQueue imprime uma fila usando a interface
func PrintQueue(queue Queue, name string) {
	fmt.Printf("%s: %s (tamanho: %d)\n", name, queue.String(), queue.Size())
}

// CopyQueue copia elementos de uma fila para outra
// Mantém a ordem original (FIFO)
func CopyQueue(source Queue, destination Queue) {
	destination.Clear()
	
	// Fila auxiliar para manter a ordem
	aux := NewArrayQueue(source.Size())
	
	// Move elementos da origem para auxiliar
	for !source.IsEmpty() {
		value, _ := source.Dequeue()
		aux.Enqueue(value)
	}
	
	// Move elementos da auxiliar para destino e restaura origem
	for !aux.IsEmpty() {
		value, _ := aux.Dequeue()
		source.Enqueue(value)
		destination.Enqueue(value)
	}
}

// ReverseQueue inverte a ordem dos elementos na fila
func ReverseQueue(queue Queue) {
	if queue.IsEmpty() {
		return
	}
	
	// Remove elemento da frente
	value, _ := queue.Dequeue()
	
	// Recursivamente inverte o resto
	ReverseQueue(queue)
	
	// Adiciona elemento no final
	queue.Enqueue(value)
}

// FindInQueue procura um elemento na fila sem alterar sua estrutura
func FindInQueue(queue Queue, target int) bool {
	if queue.IsEmpty() {
		return false
	}
	
	// Fila auxiliar para preservar ordem
	aux := NewArrayQueue(queue.Size())
	found := false
	
	// Procura elemento movendo para auxiliar
	for !queue.IsEmpty() {
		value, _ := queue.Dequeue()
		if value == target {
			found = true
		}
		aux.Enqueue(value)
	}
	
	// Restaura fila original
	for !aux.IsEmpty() {
		value, _ := aux.Dequeue()
		queue.Enqueue(value)
	}
	
	return found
}

// QueueSum calcula a soma de todos os elementos na fila
func QueueSum(queue Queue) int {
	if queue.IsEmpty() {
		return 0
	}
	
	// Fila auxiliar para preservar ordem
	aux := NewArrayQueue(queue.Size())
	sum := 0
	
	// Soma elementos movendo para auxiliar
	for !queue.IsEmpty() {
		value, _ := queue.Dequeue()
		sum += value
		aux.Enqueue(value)
	}
	
	// Restaura fila original
	for !aux.IsEmpty() {
		value, _ := aux.Dequeue()
		queue.Enqueue(value)
	}
	
	return sum
}

// QueueMax encontra o maior elemento na fila
func QueueMax(queue Queue) (int, error) {
	if queue.IsEmpty() {
		return 0, fmt.Errorf("fila vazia")
	}
	
	// Fila auxiliar para preservar ordem
	aux := NewArrayQueue(queue.Size())
	first, _ := queue.Dequeue()
	max := first
	aux.Enqueue(first)
	
	// Encontra máximo movendo para auxiliar
	for !queue.IsEmpty() {
		value, _ := queue.Dequeue()
		if value > max {
			max = value
		}
		aux.Enqueue(value)
	}
	
	// Restaura fila original
	for !aux.IsEmpty() {
		value, _ := aux.Dequeue()
		queue.Enqueue(value)
	}
	
	return max, nil
}

// QueueMin encontra o menor elemento na fila
func QueueMin(queue Queue) (int, error) {
	if queue.IsEmpty() {
		return 0, fmt.Errorf("fila vazia")
	}
	
	// Fila auxiliar para preservar ordem
	aux := NewArrayQueue(queue.Size())
	first, _ := queue.Dequeue()
	min := first
	aux.Enqueue(first)
	
	// Encontra mínimo movendo para auxiliar
	for !queue.IsEmpty() {
		value, _ := queue.Dequeue()
		if value < min {
			min = value
		}
		aux.Enqueue(value)
	}
	
	// Restaura fila original
	for !aux.IsEmpty() {
		value, _ := aux.Dequeue()
		queue.Enqueue(value)
	}
	
	return min, nil
}

// ============================================================================
// ALGORITMOS CLÁSSICOS USANDO FILAS
// ============================================================================

// GenerateNumbers gera números de 1 a n usando fila
// Demonstra uso básico de fila para sequenciamento
func GenerateNumbers(n int) []int {
	queue := NewArrayQueue(n)
	result := make([]int, 0, n)
	
	for i := 1; i <= n; i++ {
		queue.Enqueue(i)
	}
	
	for !queue.IsEmpty() {
		value, _ := queue.Dequeue()
		result = append(result, value)
	}
	
	return result
}

// GenerateBinaryNumbers gera representações binárias de 1 a n usando fila
// Algoritmo clássico que demonstra o poder das filas
func GenerateBinaryNumbers(n int) []string {
	if n <= 0 {
		return []string{}
	}
	
	queue := NewLinkedQueue()
	result := make([]string, 0, n)
	
	// Inicia com "1"
	queue.Enqueue(1) // Representa "1" como int para simplicidade
	
	for i := 1; i <= n; i++ {
		// Remove da frente
		current, _ := queue.Dequeue()
		
		// Converte para string binária
		binary := fmt.Sprintf("%b", current)
		result = append(result, binary)
		
		// Gera próximos números: current*2 e current*2+1
		if current*2 <= n*2 { // Evita overflow
			queue.Enqueue(current * 2)
			queue.Enqueue(current*2 + 1)
		}
	}
	
	return result
}

// LevelOrderTraversal simula travessia em largura usando fila
// Exemplo conceitual para árvores (usando array para simplicidade)
func LevelOrderTraversal(tree []int) []int {
	if len(tree) == 0 {
		return []int{}
	}
	
	queue := NewArrayQueue(len(tree))
	result := make([]int, 0, len(tree))
	
	// Inicia com índice da raiz
	queue.Enqueue(0)
	
	for !queue.IsEmpty() {
		index, _ := queue.Dequeue()
		
		// Processa nó atual
		if index < len(tree) {
			result = append(result, tree[index])
			
			// Adiciona filhos (para árvore binária completa)
			leftChild := 2*index + 1
			rightChild := 2*index + 2
			
			if leftChild < len(tree) {
				queue.Enqueue(leftChild)
			}
			if rightChild < len(tree) {
				queue.Enqueue(rightChild)
			}
		}
	}
	
	return result
}

// FirstNonRepeatingCharacter encontra primeiro caractere não repetido em stream
// Usa fila para manter ordem de chegada
func FirstNonRepeatingCharacter(stream string) []rune {
	queue := NewLinkedQueue()
	frequency := make(map[rune]int)
	result := make([]rune, 0, len(stream))
	
	for _, char := range stream {
		// Atualiza frequência
		frequency[char]++
		
		// Adiciona à fila se primeira ocorrência
		if frequency[char] == 1 {
			queue.Enqueue(int(char))
		}
		
		// Remove caracteres repetidos da frente
		for !queue.IsEmpty() {
			front, _ := queue.Front()
			if frequency[rune(front)] > 1 {
				queue.Dequeue()
			} else {
				break
			}
		}
		
		// Resultado atual
		if queue.IsEmpty() {
			result = append(result, 0) // Nenhum caractere não repetido
		} else {
			front, _ := queue.Front()
			result = append(result, rune(front))
		}
	}
	
	return result
}

// RotateQueue rotaciona fila k posições para a esquerda
func RotateQueue(queue Queue, k int) {
	if queue.IsEmpty() || k <= 0 {
		return
	}
	
	size := queue.Size()
	k = k % size // Normaliza k
	
	// Move k elementos da frente para o final
	for i := 0; i < k; i++ {
		value, _ := queue.Dequeue()
		queue.Enqueue(value)
	}
}

// InterleaveQueue intercala primeira e segunda metade da fila
func InterleaveQueue(queue Queue) {
	if queue.Size() < 2 {
		return
	}
	
	size := queue.Size()
	half := size / 2
	
	// Fila auxiliar para primeira metade
	aux := NewArrayQueue(half)
	
	// Move primeira metade para auxiliar
	for i := 0; i < half; i++ {
		value, _ := queue.Dequeue()
		aux.Enqueue(value)
	}
	
	// Intercala elementos
	for !aux.IsEmpty() {
		// Adiciona da primeira metade
		value1, _ := aux.Dequeue()
		queue.Enqueue(value1)
		
		// Adiciona da segunda metade
		if !queue.IsEmpty() {
			value2, _ := queue.Dequeue()
			queue.Enqueue(value2)
		}
	}
}

// SortQueue ordena fila usando fila auxiliar
func SortQueue(queue Queue) {
	if queue.Size() <= 1 {
		return
	}
	
	aux := NewArrayQueue(queue.Size())
	
	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()
		
		// Move elementos maiores de volta para queue
		count := 0
		for !aux.IsEmpty() {
			top, _ := aux.Rear()
			if top > current {
				value, _ := aux.Dequeue()
				queue.Enqueue(value)
				count++
			} else {
				break
			}
		}
		
		// Insere elemento atual
		aux.Enqueue(current)
		
		// Move elementos de volta para aux
		for i := 0; i < count; i++ {
			value, _ := queue.Dequeue()
			aux.Enqueue(value)
		}
	}
	
	// Move todos de volta para queue original
	for !aux.IsEmpty() {
		value, _ := aux.Dequeue()
		queue.Enqueue(value)
	}
}