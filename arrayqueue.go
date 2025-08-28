package main

import (
	"fmt"
	"errors"
	"strings"
)

// ============================================================================
// ARRAYQUEUE - IMPLEMENTAÇÃO BASEADA EM ARRAY CIRCULAR
// ============================================================================

// ArrayQueue implementa uma fila usando um array circular
// Características:
// - Enqueue/Dequeue são operações O(1)
// - Uso eficiente do espaço (reutiliza posições)
// - Evita necessidade de mover elementos
// - Capacidade fixa (pode ser redimensionada)
type ArrayQueue struct {
	data     []int // Array interno que armazena os elementos
	front    int   // Índice do primeiro elemento
	rear     int   // Índice da próxima posição livre
	size     int   // Número atual de elementos
	capacity int   // Capacidade máxima do array
}

// NewArrayQueue cria uma nova instância de ArrayQueue com capacidade inicial
func NewArrayQueue(initialCapacity int) *ArrayQueue {
	if initialCapacity <= 0 {
		initialCapacity = 10 // Capacidade padrão
	}
	return &ArrayQueue{
		data:     make([]int, initialCapacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: initialCapacity,
	}
}

// ============================================================================
// IMPLEMENTAÇÃO DA INTERFACE QUEUE
// ============================================================================

// Enqueue adiciona um elemento no final da fila
// Complexidade: O(1) - pode ser O(n) se redimensionar
func (q *ArrayQueue) Enqueue(element int) {
	// Verifica se precisa redimensionar
	if q.IsFull() {
		q.resize(q.capacity * 2)
	}
	
	q.data[q.rear] = element
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

// Dequeue remove e retorna o elemento do início da fila
// Complexidade: O(1)
func (q *ArrayQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não é possível fazer dequeue")
	}
	
	element := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	
	// Redimensiona para baixo se necessário (economiza memória)
	if q.size > 0 && q.size == q.capacity/4 {
		q.resize(q.capacity / 2)
	}
	
	return element, nil
}

// Front retorna o elemento do início sem removê-lo
// Complexidade: O(1)
func (q *ArrayQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento na frente")
	}
	return q.data[q.front], nil
}

// Rear retorna o elemento do final sem removê-lo
// Complexidade: O(1)
func (q *ArrayQueue) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento no final")
	}
	// rear aponta para próxima posição livre, então elemento atual está em rear-1
	rearIndex := (q.rear - 1 + q.capacity) % q.capacity
	return q.data[rearIndex], nil
}

// Size retorna o número de elementos na fila
// Complexidade: O(1)
func (q *ArrayQueue) Size() int {
	return q.size
}

// IsEmpty verifica se a fila está vazia
// Complexidade: O(1)
func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull verifica se a fila está cheia
// Complexidade: O(1)
func (q *ArrayQueue) IsFull() bool {
	return q.size == q.capacity
}

// Clear remove todos os elementos da fila
// Complexidade: O(1)
func (q *ArrayQueue) Clear() {
	q.front = 0
	q.rear = 0
	q.size = 0
	// Opcionalmente, pode redimensionar para capacidade inicial
	if q.capacity > 10 {
		q.resize(10)
	}
}

// ToSlice converte a fila para um slice (do início para o final)
// Complexidade: O(n)
func (q *ArrayQueue) ToSlice() []int {
	if q.IsEmpty() {
		return []int{}
	}
	
	result := make([]int, q.size)
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		result[i] = q.data[index]
	}
	return result
}

// String retorna uma representação em string da fila
// Complexidade: O(n)
func (q *ArrayQueue) String() string {
	if q.IsEmpty() {
		return "[vazia]"
	}
	
	var builder strings.Builder
	builder.WriteString("frente → [")
	
	for i := 0; i < q.size; i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		index := (q.front + i) % q.capacity
		builder.WriteString(fmt.Sprintf("%d", q.data[index]))
	}
	
	builder.WriteString("] ← final")
	return builder.String()
}

// ============================================================================
// MÉTODOS AUXILIARES ESPECÍFICOS DO ARRAYQUEUE
// ============================================================================

// Capacity retorna a capacidade atual do array interno
func (q *ArrayQueue) Capacity() int {
	return q.capacity
}

// resize redimensiona o array interno para uma nova capacidade
func (q *ArrayQueue) resize(newCapacity int) {
	if newCapacity < q.size {
		newCapacity = q.size // Não pode ser menor que o tamanho atual
	}
	
	newData := make([]int, newCapacity)
	
	// Copia elementos na ordem correta
	for i := 0; i < q.size; i++ {
		oldIndex := (q.front + i) % q.capacity
		newData[i] = q.data[oldIndex]
	}
	
	q.data = newData
	q.front = 0
	q.rear = q.size
	q.capacity = newCapacity
}

// TrimToSize reduz a capacidade para o tamanho atual (economiza memória)
func (q *ArrayQueue) TrimToSize() {
	if q.capacity > q.size {
		newCapacity := q.size
		if newCapacity == 0 {
			newCapacity = 1 // Mantém pelo menos 1 de capacidade
		}
		q.resize(newCapacity)
	}
}

// EnsureCapacity garante que a fila tenha pelo menos a capacidade especificada
func (q *ArrayQueue) EnsureCapacity(minCapacity int) {
	if q.capacity < minCapacity {
		q.resize(minCapacity)
	}
}

// EnqueueAll adiciona múltiplos elementos de uma vez
func (q *ArrayQueue) EnqueueAll(elements []int) {
	// Garante capacidade suficiente
	q.EnsureCapacity(q.size + len(elements))
	
	for _, element := range elements {
		q.Enqueue(element)
	}
}

// DequeueMultiple remove e retorna múltiplos elementos do início
func (q *ArrayQueue) DequeueMultiple(count int) ([]int, error) {
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
func (q *ArrayQueue) Contains(element int) bool {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		if q.data[index] == element {
			return true
		}
	}
	return false
}

// IndexOf procura um elemento e retorna sua posição a partir do início
// Retorna -1 se não encontrado
// Complexidade: O(n)
func (q *ArrayQueue) IndexOf(element int) int {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		if q.data[index] == element {
			return i
		}
	}
	return -1
}

// Clone cria uma cópia independente da fila
func (q *ArrayQueue) Clone() *ArrayQueue {
	newQueue := NewArrayQueue(q.capacity)
	newQueue.size = q.size
	newQueue.front = 0
	newQueue.rear = q.size
	
	// Copia elementos na ordem correta
	for i := 0; i < q.size; i++ {
		oldIndex := (q.front + i) % q.capacity
		newQueue.data[i] = q.data[oldIndex]
	}
	
	return newQueue
}

// Equals verifica se duas filas são iguais (mesmo conteúdo e ordem)
func (q *ArrayQueue) Equals(other *ArrayQueue) bool {
	if q.size != other.size {
		return false
	}
	
	for i := 0; i < q.size; i++ {
		index1 := (q.front + i) % q.capacity
		index2 := (other.front + i) % other.capacity
		if q.data[index1] != other.data[index2] {
			return false
		}
	}
	
	return true
}

// ============================================================================
// MÉTODOS DE ANÁLISE E ESTATÍSTICAS
// ============================================================================

// GetMemoryUsage retorna informações sobre uso de memória
func (q *ArrayQueue) GetMemoryUsage() (int, int, float64) {
	usedSlots := q.size
	totalSlots := q.capacity
	utilization := float64(usedSlots) / float64(totalSlots) * 100
	return usedSlots, totalSlots, utilization
}

// GetStatistics retorna estatísticas da fila
func (q *ArrayQueue) GetStatistics() map[string]interface{} {
	if q.IsEmpty() {
		return map[string]interface{}{
			"size":         0,
			"capacity":     q.capacity,
			"utilization": 0.0,
			"isEmpty":      true,
		}
	}
	
	// Calcula estatísticas
	sum := 0
	firstIndex := q.front
	min := q.data[firstIndex]
	max := q.data[firstIndex]
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		value := q.data[index]
		sum += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
	average := float64(sum) / float64(q.size)
	utilization := float64(q.size) / float64(q.capacity) * 100
	
	return map[string]interface{}{
		"size":         q.size,
		"capacity":     q.capacity,
		"utilization":  utilization,
		"isEmpty":      false,
		"sum":          sum,
		"average":      average,
		"min":          min,
		"max":          max,
		"frontIndex":   q.front,
		"rearIndex":    q.rear,
	}
}

// ============================================================================
// MÉTODOS AVANÇADOS
// ============================================================================

// Rotate rotaciona a fila k posições para a esquerda
func (q *ArrayQueue) Rotate(k int) {
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

// Reverse inverte a ordem dos elementos na fila
func (q *ArrayQueue) Reverse() {
	if q.size <= 1 {
		return
	}
	
	// Cria array temporário com elementos na ordem correta
	temp := q.ToSlice()
	
	// Inverte o array
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	
	// Reconstrói a fila
	q.Clear()
	q.EnqueueAll(temp)
}

// Filter cria uma nova fila com elementos que satisfazem uma condição
func (q *ArrayQueue) Filter(predicate func(int) bool) *ArrayQueue {
	result := NewArrayQueue(q.size)
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		value := q.data[index]
		if predicate(value) {
			result.Enqueue(value)
		}
	}
	
	return result
}

// Map aplica uma função a todos os elementos e retorna nova fila
func (q *ArrayQueue) Map(mapper func(int) int) *ArrayQueue {
	result := NewArrayQueue(q.size)
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		value := q.data[index]
		result.Enqueue(mapper(value))
	}
	
	return result
}

// Reduce aplica uma função de redução a todos os elementos
func (q *ArrayQueue) Reduce(reducer func(int, int) int, initialValue int) int {
	result := initialValue
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		value := q.data[index]
		result = reducer(result, value)
	}
	
	return result
}

// ForEach executa uma função para cada elemento (do início para o final)
func (q *ArrayQueue) ForEach(action func(int, int)) {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		value := q.data[index]
		action(value, i)
	}
}

// GetInternalState retorna informações sobre o estado interno (para debug)
func (q *ArrayQueue) GetInternalState() map[string]interface{} {
	return map[string]interface{}{
		"front":    q.front,
		"rear":     q.rear,
		"size":     q.size,
		"capacity": q.capacity,
		"data":     q.data,
	}
}