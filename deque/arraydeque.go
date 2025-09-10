package deque

import (
	"errors"
	"fmt"
	"strings"
)

// ============================================================================
// ARRAYDEQUE - IMPLEMENTAÇÃO BASEADA EM ARRAY CIRCULAR
// ============================================================================

// ArrayDeque implementa um deque (double-ended queue) usando um array circular
// Características:
// - Operações em ambas extremidades são O(1)
// - Uso eficiente do espaço (reutiliza posições)
// - Evita necessidade de mover elementos
// - Capacidade fixa (pode ser redimensionada)
type ArrayDeque struct {
	data     []int // Array interno que armazena os elementos
	front    int   // Índice do primeiro elemento
	rear     int   // Índice da próxima posição livre
	size     int   // Número atual de elementos
	capacity int   // Capacidade máxima do array
}

// NewArrayDeque cria uma nova instância de ArrayDeque com capacidade inicial
func NewArrayDeque(initialCapacity int) *ArrayDeque {
	if initialCapacity <= 0 {
		initialCapacity = 10 // Capacidade padrão
	}
	return &ArrayDeque{
		data:     make([]int, initialCapacity),
		front:    0,
		rear:     0,
		size:     0,
		capacity: initialCapacity,
	}
}

// ============================================================================
// MÉTODOS PRINCIPAIS
// ============================================================================

// EnqueueFront adiciona elemento no início da fila (comportamento de deque)
// Complexidade: O(1)
func (q *ArrayDeque) EnqueueFront(element int) {
	if q.IsFull() {
		q.resize(q.capacity * 2)
	}
	
	q.front = (q.front - 1 + q.capacity) % q.capacity
	q.data[q.front] = element
	q.size++
}

// EnqueueRear adiciona elemento no final da fila
// Complexidade: O(1) - pode ser O(n) se redimensionar
func (q *ArrayDeque) EnqueueRear(element int) {
	// Verifica se precisa redimensionar
	if q.IsFull() {
		q.resize(q.capacity * 2)
	}
	
	q.data[q.rear] = element
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

// DequeueFront remove e retorna elemento do início da fila
// Complexidade: O(1)
func (q *ArrayDeque) DequeueFront() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não é possível fazer dequeue")
	}
	
	value := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	
	return value, nil
}

// DequeueRear remove e retorna elemento do final da fila (comportamento de deque)
// Complexidade: O(1)
func (q *ArrayDeque) DequeueRear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não é possível fazer dequeue")
	}
	
	q.rear = (q.rear - 1 + q.capacity) % q.capacity
	value := q.data[q.rear]
	q.size--
	
	return value, nil
}

// Front retorna o elemento do início sem removê-lo
// Complexidade: O(1)
func (q *ArrayDeque) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento na frente")
	}
	return q.data[q.front], nil
}

// Rear retorna o elemento do final sem removê-lo
// Complexidade: O(1)
func (q *ArrayDeque) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento no final")
	}
	rearIndex := (q.rear - 1 + q.capacity) % q.capacity
	return q.data[rearIndex], nil
}

// Size retorna o número de elementos na fila
// Complexidade: O(1)
func (q *ArrayDeque) Size() int {
	return q.size
}

// IsEmpty verifica se a fila está vazia
// Complexidade: O(1)
func (q *ArrayDeque) IsEmpty() bool {
	return q.size == 0
}

// IsFull verifica se a fila está cheia
// Complexidade: O(1)
func (q *ArrayDeque) IsFull() bool {
	return q.size == q.capacity
}

// Clear remove todos os elementos da fila
// Complexidade: O(1)
func (q *ArrayDeque) Clear() {
	q.front = 0
	q.rear = 0
	q.size = 0
	// Não precisa limpar o array, apenas resetar os índices
}

// ============================================================================
// MÉTODOS AUXILIARES
// ============================================================================

// ToSlice converte a fila para um slice (do início para o final)
// Complexidade: O(n)
func (q *ArrayDeque) ToSlice() []int {
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
func (q *ArrayDeque) String() string {
	if q.IsEmpty() {
		return "[vazia]"
	}
	
	var builder strings.Builder
	builder.WriteString("início → [")
	
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

// Capacity retorna a capacidade atual da fila
func (q *ArrayDeque) Capacity() int {
	return q.capacity
}

// resize redimensiona o array interno
func (q *ArrayDeque) resize(newCapacity int) {
	if newCapacity < q.size {
		return // Não pode reduzir abaixo do tamanho atual
	}
	
	newData := make([]int, newCapacity)
	
	// Copia elementos na ordem correta
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		newData[i] = q.data[index]
	}
	
	q.data = newData
	q.front = 0
	q.rear = q.size
	q.capacity = newCapacity
}

// TrimToSize reduz a capacidade para o tamanho atual
func (q *ArrayDeque) TrimToSize() {
	if q.size < q.capacity {
		q.resize(q.size)
	}
}

// EnsureCapacity garante que a fila tenha pelo menos a capacidade mínima
func (q *ArrayDeque) EnsureCapacity(minCapacity int) {
	if minCapacity > q.capacity {
		q.resize(minCapacity)
	}
}

// ============================================================================
// MÉTODOS AVANÇADOS
// ============================================================================

// EnqueueAll adiciona múltiplos elementos de uma vez
func (q *ArrayDeque) EnqueueAll(elements []int) {
	for _, element := range elements {
		q.EnqueueRear(element)
	}
}

// DequeueMultiple remove e retorna múltiplos elementos do início
func (q *ArrayDeque) DequeueMultiple(count int) ([]int, error) {
	if count <= 0 {
		return []int{}, nil
	}
	
	if count > q.size {
		return nil, fmt.Errorf("não há elementos suficientes: solicitado %d, disponível %d", count, q.size)
	}
	
	result := make([]int, count)
	for i := 0; i < count; i++ {
		value, _ := q.DequeueFront()
		result[i] = value
	}
	
	return result, nil
}

// Contains verifica se a fila contém um elemento específico
func (q *ArrayDeque) Contains(element int) bool {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		if q.data[index] == element {
			return true
		}
	}
	return false
}

// IndexOf procura um elemento e retorna sua posição a partir do início
func (q *ArrayDeque) IndexOf(element int) int {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		if q.data[index] == element {
			return i
		}
	}
	return -1
}

// Clone cria uma cópia independente da fila
func (q *ArrayDeque) Clone() *ArrayDeque {
	return q.CloneArrayDeque()
}

// CloneArrayDeque cria uma cópia independente retornando *ArrayDeque
func (q *ArrayDeque) CloneArrayDeque() *ArrayDeque {
	newQueue := NewArrayDeque(q.capacity)
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		newQueue.EnqueueRear(q.data[index])
	}
	
	return newQueue
}

// Equals verifica se duas filas são iguais
func (q *ArrayDeque) Equals(other *ArrayDeque) bool {
	return q.EqualsArrayDeque(other)
}

// EqualsArrayDeque verifica se duas filas ArrayDeque são iguais
func (q *ArrayDeque) EqualsArrayDeque(other *ArrayDeque) bool {
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

// GetMemoryUsage retorna informações sobre uso de memória
func (q *ArrayDeque) GetMemoryUsage() (int, int, float64) {
	usedSlots := q.size
	totalSlots := q.capacity
	utilization := float64(usedSlots) / float64(totalSlots) * 100
	return usedSlots, totalSlots, utilization
}

// GetStatistics retorna estatísticas da fila
func (q *ArrayDeque) GetStatistics() map[string]interface{} {
	if q.IsEmpty() {
		return map[string]interface{}{
			"size":     0,
			"capacity": q.capacity,
			"isEmpty":  true,
		}
	}
	
	// Calcula estatísticas
	sum := 0
	min := q.data[q.front]
	max := q.data[q.front]
	
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
	usedSlots, totalSlots, utilization := q.GetMemoryUsage()
	
	return map[string]interface{}{
		"size":        q.size,
		"capacity":    q.capacity,
		"isEmpty":     false,
		"sum":         sum,
		"average":     average,
		"min":         min,
		"max":         max,
		"usedSlots":   usedSlots,
		"totalSlots":  totalSlots,
		"utilization": fmt.Sprintf("%.1f%%", utilization),
	}
}

// Rotate rotaciona a fila k posições para a esquerda
func (q *ArrayDeque) Rotate(k int) {
	if q.IsEmpty() || k <= 0 {
		return
	}
	
	k = k % q.size // Normaliza k
	
	// Move k elementos da frente para o final
	for i := 0; i < k; i++ {
		value, _ := q.DequeueFront()
		q.EnqueueRear(value)
	}
}

// Reverse inverte a ordem dos elementos na fila
func (q *ArrayDeque) Reverse() {
	if q.size <= 1 {
		return
	}
	
	// Coleta todos os elementos
	elements := q.ToSlice()
	
	// Limpa a fila
	q.Clear()
	
	// Adiciona elementos na ordem inversa
	for i := len(elements) - 1; i >= 0; i-- {
		q.EnqueueRear(elements[i])
	}
}

// Filter cria uma nova fila com elementos que satisfazem uma condição
func (q *ArrayDeque) Filter(predicate func(int) bool) *ArrayDeque {
	result := NewArrayDeque(q.capacity)
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		value := q.data[index]
		if predicate(value) {
			result.EnqueueRear(value)
		}
	}
	
	return result
}

// Map aplica uma função a todos os elementos e retorna nova fila
func (q *ArrayDeque) Map(mapper func(int) int) *ArrayDeque {
	result := NewArrayDeque(q.capacity)
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		result.EnqueueRear(mapper(q.data[index]))
	}
	
	return result
}

// Reduce aplica uma função de redução a todos os elementos
func (q *ArrayDeque) Reduce(reducer func(int, int) int, initialValue int) int {
	result := initialValue
	
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		result = reducer(result, q.data[index])
	}
	
	return result
}

// ForEach executa uma função para cada elemento
func (q *ArrayDeque) ForEach(action func(int, int)) {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		action(q.data[index], i)
	}
}

// GetInternalState retorna o estado interno da fila (para debug)
func (q *ArrayDeque) GetInternalState() map[string]interface{} {
	return map[string]interface{}{
		"front":    q.front,
		"rear":     q.rear,
		"size":     q.size,
		"capacity": q.capacity,
		"data":     q.data,
	}
}

// ============================================================================
// MÉTODOS COMPATÍVEIS COM INTERFACE IDEQUE
// ============================================================================

// CloneIDeque implementa IDeque.Clone() retornando IDeque
func (q *ArrayDeque) CloneIDeque() IDeque {
	return q.CloneArrayDeque()
}

// EqualsIDeque implementa IDeque.Equals() aceitando IDeque
func (q *ArrayDeque) EqualsIDeque(other IDeque) bool {
	if q.Size() != other.Size() {
		return false
	}
	
	slice1 := q.ToSlice()
	slice2 := other.ToSlice()
	
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	
	return true
}

// FilterIDeque implementa IDeque.Filter() retornando IDeque
func (q *ArrayDeque) FilterIDeque(predicate func(int) bool) IDeque {
	return q.Filter(predicate)
}

// MapIDeque implementa IDeque.Map() retornando IDeque
func (q *ArrayDeque) MapIDeque(mapper func(int) int) IDeque {
	return q.Map(mapper)
}