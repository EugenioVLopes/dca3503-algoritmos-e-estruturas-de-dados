package main

import (
	"fmt"
	"errors"
	"strings"
)

// ============================================================================
// LINKEDQUEUE - IMPLEMENTAÇÃO BASEADA EM LISTA LIGADA
// ============================================================================

// QueueNode representa um nó na fila ligada
type QueueNode struct {
	data int        // Valor armazenado no nó
	next *QueueNode // Ponteiro para o próximo nó
}

// LinkedQueue implementa uma fila usando lista ligada
// Características:
// - Enqueue/Dequeue são sempre O(1)
// - Não há limite de capacidade (limitado apenas pela memória)
// - Uso dinâmico de memória (aloca conforme necessário)
// - Mantém ponteiros para frente e final para eficiência
type LinkedQueue struct {
	front *QueueNode // Ponteiro para o primeiro nó
	rear  *QueueNode // Ponteiro para o último nó
	size  int        // Contador de elementos
}

// NewLinkedQueue cria uma nova instância de LinkedQueue
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// ============================================================================
// IMPLEMENTAÇÃO DA INTERFACE QUEUE
// ============================================================================

// Enqueue adiciona um elemento no final da fila
// Complexidade: O(1)
func (q *LinkedQueue) Enqueue(element int) {
	newNode := &QueueNode{
		data: element,
		next: nil,
	}
	
	if q.IsEmpty() {
		// Primeiro elemento
		q.front = newNode
		q.rear = newNode
	} else {
		// Adiciona no final
		q.rear.next = newNode
		q.rear = newNode
	}
	
	q.size++
}

// Dequeue remove e retorna o elemento do início da fila
// Complexidade: O(1)
func (q *LinkedQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não é possível fazer dequeue")
	}
	
	value := q.front.data
	q.front = q.front.next
	q.size--
	
	// Se ficou vazia, atualiza rear também
	if q.front == nil {
		q.rear = nil
	}
	
	return value, nil
}

// Front retorna o elemento do início sem removê-lo
// Complexidade: O(1)
func (q *LinkedQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento na frente")
	}
	return q.front.data, nil
}

// Rear retorna o elemento do final sem removê-lo
// Complexidade: O(1)
func (q *LinkedQueue) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento no final")
	}
	return q.rear.data, nil
}

// Size retorna o número de elementos na fila
// Complexidade: O(1)
func (q *LinkedQueue) Size() int {
	return q.size
}

// IsEmpty verifica se a fila está vazia
// Complexidade: O(1)
func (q *LinkedQueue) IsEmpty() bool {
	return q.front == nil
}

// IsFull verifica se a fila está cheia
// Para LinkedQueue, nunca está "cheia" pois usa alocação dinâmica
// Complexidade: O(1)
func (q *LinkedQueue) IsFull() bool {
	return false // Lista ligada nunca está "cheia"
}

// Clear remove todos os elementos da fila
// Complexidade: O(1) - apenas redefine ponteiros, GC limpa os nós
func (q *LinkedQueue) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
	// O Garbage Collector do Go automaticamente limpa os nós órfãos
}

// ToSlice converte a fila para um slice (do início para o final)
// Complexidade: O(n)
func (q *LinkedQueue) ToSlice() []int {
	if q.IsEmpty() {
		return []int{}
	}
	
	result := make([]int, q.size)
	current := q.front
	index := 0
	
	for current != nil {
		result[index] = current.data
		current = current.next
		index++
	}
	
	return result
}

// String retorna uma representação em string da fila
// Complexidade: O(n)
func (q *LinkedQueue) String() string {
	if q.IsEmpty() {
		return "[vazia]"
	}
	
	var builder strings.Builder
	builder.WriteString("frente → [")
	
	current := q.front
	first := true
	
	for current != nil {
		if !first {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("%d", current.data))
		current = current.next
		first = false
	}
	
	builder.WriteString("] ← final")
	return builder.String()
}

// ============================================================================
// MÉTODOS AUXILIARES ESPECÍFICOS DO LINKEDQUEUE
// ============================================================================

// EnqueueAll adiciona múltiplos elementos de uma vez
// Complexidade: O(n) onde n é o número de elementos
func (q *LinkedQueue) EnqueueAll(elements []int) {
	for _, element := range elements {
		q.Enqueue(element)
	}
}

// DequeueMultiple remove e retorna múltiplos elementos do início
// Complexidade: O(n) onde n é o número de elementos a remover
func (q *LinkedQueue) DequeueMultiple(count int) ([]int, error) {
	if count <= 0 {
		return []int{}, nil
	}
	
	if count > q.size {
		return nil, fmt.Errorf("não há elementos suficientes: solicitado %d, disponível %d", count, q.size)
	}
	
	result := make([]int, count)
	for i := 0; i < count; i++ {
		value, _ := q.Dequeue()
		result[i] = value
	}
	
	return result, nil
}

// Contains verifica se a fila contém um elemento específico
// Complexidade: O(n)
func (q *LinkedQueue) Contains(element int) bool {
	current := q.front
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

// IndexOf procura um elemento e retorna sua posição a partir do início
// Retorna -1 se não encontrado
// Complexidade: O(n)
func (q *LinkedQueue) IndexOf(element int) int {
	current := q.front
	index := 0
	
	for current != nil {
		if current.data == element {
			return index
		}
		current = current.next
		index++
	}
	
	return -1
}

// Clone cria uma cópia independente da fila
// Complexidade: O(n)
func (q *LinkedQueue) Clone() *LinkedQueue {
	newQueue := NewLinkedQueue()
	
	if q.IsEmpty() {
		return newQueue
	}
	
	current := q.front
	for current != nil {
		newQueue.Enqueue(current.data)
		current = current.next
	}
	
	return newQueue
}

// Equals verifica se duas filas são iguais (mesmo conteúdo e ordem)
// Complexidade: O(n)
func (q *LinkedQueue) Equals(other *LinkedQueue) bool {
	if q.size != other.size {
		return false
	}
	
	current1 := q.front
	current2 := other.front
	
	for current1 != nil && current2 != nil {
		if current1.data != current2.data {
			return false
		}
		current1 = current1.next
		current2 = current2.next
	}
	
	return true
}

// Reverse inverte a ordem dos elementos na fila
// Complexidade: O(n)
func (q *LinkedQueue) Reverse() {
	if q.size <= 1 {
		return
	}
	
	// Coleta todos os elementos
	elements := q.ToSlice()
	
	// Limpa a fila
	q.Clear()
	
	// Adiciona elementos na ordem inversa
	for i := len(elements) - 1; i >= 0; i-- {
		q.Enqueue(elements[i])
	}
}

// ============================================================================
// MÉTODOS DE ANÁLISE E ESTATÍSTICAS
// ============================================================================

// GetStatistics retorna estatísticas da fila
// Complexidade: O(n)
func (q *LinkedQueue) GetStatistics() map[string]interface{} {
	if q.IsEmpty() {
		return map[string]interface{}{
			"size":    0,
			"isEmpty": true,
		}
	}
	
	// Calcula estatísticas
	sum := 0
	min := q.front.data
	max := q.front.data
	current := q.front
	
	for current != nil {
		value := current.data
		sum += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		current = current.next
	}
	
	average := float64(sum) / float64(q.size)
	
	return map[string]interface{}{
		"size":    q.size,
		"isEmpty": false,
		"sum":     sum,
		"average": average,
		"min":     min,
		"max":     max,
	}
}

// GetMemoryInfo retorna informações sobre uso de memória
// Estima o uso de memória baseado no número de nós
func (q *LinkedQueue) GetMemoryInfo() map[string]interface{} {
	// Cada nó tem: int (8 bytes) + ponteiro (8 bytes) = 16 bytes
	nodeSize := 16
	totalMemory := q.size * nodeSize
	
	return map[string]interface{}{
		"nodes":           q.size,
		"estimatedBytes":  totalMemory,
		"bytesPerNode":    nodeSize,
		"memoryEfficient": q.size > 0, // Sem overhead de capacidade não utilizada
	}
}

// ============================================================================
// MÉTODOS AVANÇADOS
// ============================================================================

// Filter cria uma nova fila com elementos que satisfazem uma condição
// Mantém a ordem relativa dos elementos
func (q *LinkedQueue) Filter(predicate func(int) bool) *LinkedQueue {
	result := NewLinkedQueue()
	
	current := q.front
	for current != nil {
		if predicate(current.data) {
			result.Enqueue(current.data)
		}
		current = current.next
	}
	
	return result
}

// Map aplica uma função a todos os elementos e retorna nova fila
func (q *LinkedQueue) Map(mapper func(int) int) *LinkedQueue {
	result := NewLinkedQueue()
	
	current := q.front
	for current != nil {
		result.Enqueue(mapper(current.data))
		current = current.next
	}
	
	return result
}

// Reduce aplica uma função de redução a todos os elementos
func (q *LinkedQueue) Reduce(reducer func(int, int) int, initialValue int) int {
	result := initialValue
	current := q.front
	
	for current != nil {
		result = reducer(result, current.data)
		current = current.next
	}
	
	return result
}

// ForEach executa uma função para cada elemento (do início para o final)
func (q *LinkedQueue) ForEach(action func(int, int)) {
	current := q.front
	index := 0
	
	for current != nil {
		action(current.data, index)
		current = current.next
		index++
	}
}

// Partition divide a fila em duas baseado em um predicado
// Retorna (elementos que satisfazem, elementos que não satisfazem)
func (q *LinkedQueue) Partition(predicate func(int) bool) (*LinkedQueue, *LinkedQueue) {
	true_queue := NewLinkedQueue()
	false_queue := NewLinkedQueue()
	
	current := q.front
	for current != nil {
		if predicate(current.data) {
			true_queue.Enqueue(current.data)
		} else {
			false_queue.Enqueue(current.data)
		}
		current = current.next
	}
	
	return true_queue, false_queue
}

// TakeWhile retorna nova fila com elementos do início que satisfazem condição
func (q *LinkedQueue) TakeWhile(predicate func(int) bool) *LinkedQueue {
	result := NewLinkedQueue()
	
	current := q.front
	for current != nil {
		if !predicate(current.data) {
			break
		}
		result.Enqueue(current.data)
		current = current.next
	}
	
	return result
}

// DropWhile retorna nova fila removendo elementos do início que satisfazem condição
func (q *LinkedQueue) DropWhile(predicate func(int) bool) *LinkedQueue {
	result := NewLinkedQueue()
	
	current := q.front
	// Pula elementos que satisfazem o predicado
	for current != nil && predicate(current.data) {
		current = current.next
	}
	
	// Adiciona o resto
	for current != nil {
		result.Enqueue(current.data)
		current = current.next
	}
	
	return result
}

// Rotate rotaciona a fila k posições para a esquerda
func (q *LinkedQueue) Rotate(k int) {
	if q.IsEmpty() || k <= 0 {
		return
	}
	
	k = k % q.size // Normaliza k
	
	// Move k elementos da frente para o final
	for i := 0; i < k; i++ {
		value, _ := q.Dequeue()
		q.Enqueue(value)
	}
}

// Split divide a fila em duas partes no índice especificado
func (q *LinkedQueue) Split(index int) (*LinkedQueue, *LinkedQueue) {
	first := NewLinkedQueue()
	second := NewLinkedQueue()
	
	if index <= 0 {
		return first, q.Clone()
	}
	if index >= q.size {
		return q.Clone(), second
	}
	
	current := q.front
	currentIndex := 0
	
	for current != nil {
		if currentIndex < index {
			first.Enqueue(current.data)
		} else {
			second.Enqueue(current.data)
		}
		current = current.next
		currentIndex++
	}
	
	return first, second
}

// Merge combina esta fila com outra, alternando elementos
func (q *LinkedQueue) Merge(other *LinkedQueue) *LinkedQueue {
	result := NewLinkedQueue()
	
	current1 := q.front
	current2 := other.front
	
	// Alterna entre as duas filas
	for current1 != nil || current2 != nil {
		if current1 != nil {
			result.Enqueue(current1.data)
			current1 = current1.next
		}
		if current2 != nil {
			result.Enqueue(current2.data)
			current2 = current2.next
		}
	}
	
	return result
}

// GetNth retorna o n-ésimo elemento (0-indexado) sem removê-lo
func (q *LinkedQueue) GetNth(n int) (int, error) {
	if n < 0 || n >= q.size {
		return 0, fmt.Errorf("índice fora dos limites: %d", n)
	}
	
	current := q.front
	for i := 0; i < n; i++ {
		current = current.next
	}
	
	return current.data, nil
}