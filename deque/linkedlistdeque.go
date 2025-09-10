package deque

import (
	"errors"
	"fmt"
	"strings"
)

// ============================================================================
// LINKEDLISTDEQUE - IMPLEMENTAÇÃO BASEADA EM LISTA LIGADA
// ============================================================================

// LinkedDequeNode representa um nó no deque ligado
type LinkedDequeNode struct {
	data int               // Valor armazenado no nó
	next *LinkedDequeNode  // Ponteiro para o próximo nó
}

// LinkedListDeque implementa um deque (double-ended queue) usando lista ligada
// Características:
// - Operações em ambas extremidades são O(1) (exceto DequeueRear que é O(n))
// - Não há limite de capacidade (limitado apenas pela memória)
// - Uso dinâmico de memória (aloca conforme necessário)
// - Mantém ponteiros para frente e final para eficiência
type LinkedListDeque struct {
	front *LinkedDequeNode // Ponteiro para o primeiro nó
	rear  *LinkedDequeNode // Ponteiro para o último nó
	size  int        // Contador de elementos
}

// NewLinkedListDeque cria uma nova instância de LinkedListDeque
func NewLinkedListDeque() *LinkedListDeque {
	return &LinkedListDeque{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// ============================================================================
// MÉTODOS PRINCIPAIS
// ============================================================================

// EnqueueFront adiciona elemento no início da fila (comportamento de deque)
// Complexidade: O(1)
func (q *LinkedListDeque) EnqueueFront(element int) {
	newNode := &LinkedDequeNode{
		data: element,
		next: q.front,
	}
	
	if q.IsEmpty() {
		// Primeiro elemento
		q.front = newNode
		q.rear = newNode
	} else {
		// Adiciona no início
		q.front = newNode
	}
	
	q.size++
}

// EnqueueRear adiciona elemento no final da fila
// Complexidade: O(1)
func (q *LinkedListDeque) EnqueueRear(element int) {
	newNode := &LinkedDequeNode{
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

// DequeueFront remove e retorna elemento do início da fila
// Complexidade: O(1)
func (q *LinkedListDeque) DequeueFront() (int, error) {
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

// DequeueRear remove e retorna elemento do final da fila (comportamento de deque)
// Complexidade: O(n) - precisa percorrer até o penúltimo nó
func (q *LinkedListDeque) DequeueRear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não é possível fazer dequeue")
	}
	
	if q.size == 1 {
		// Único elemento
		value := q.rear.data
		q.front = nil
		q.rear = nil
		q.size--
		return value, nil
	}
	
	// Encontra o penúltimo nó
	current := q.front
	for current.next != q.rear {
		current = current.next
	}
	
	value := q.rear.data
	current.next = nil
	q.rear = current
	q.size--
	
	return value, nil
}

// Front retorna elemento do início sem removê-lo
// Complexidade: O(1)
func (q *LinkedListDeque) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento na frente")
	}
	return q.front.data, nil
}

// Rear retorna elemento do final sem removê-lo
// Complexidade: O(1)
func (q *LinkedListDeque) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("fila vazia: não há elemento no final")
	}
	return q.rear.data, nil
}

// Size retorna o número de elementos na fila
// Complexidade: O(1)
func (q *LinkedListDeque) Size() int {
	return q.size
}

// IsEmpty verifica se a fila está vazia
// Complexidade: O(1)
func (q *LinkedListDeque) IsEmpty() bool {
	return q.front == nil
}

// IsFull verifica se a fila está cheia
// Para LinkedListDeque, nunca está "cheia" pois usa alocação dinâmica
// Complexidade: O(1)
func (q *LinkedListDeque) IsFull() bool {
	return false // Lista ligada nunca está "cheia"
}

// Clear remove todos os elementos da fila
// Complexidade: O(1) - apenas redefine ponteiros, GC limpa os nós
func (q *LinkedListDeque) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
	// O Garbage Collector do Go automaticamente limpa os nós órfãos
}

// ============================================================================
// MÉTODOS AUXILIARES
// ============================================================================

// ToSlice converte a fila para um slice (do início para o final)
// Complexidade: O(n)
func (q *LinkedListDeque) ToSlice() []int {
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
func (q *LinkedListDeque) String() string {
	if q.IsEmpty() {
		return "[vazia]"
	}
	
	var builder strings.Builder
	builder.WriteString("início → [")
	
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
// MÉTODOS AVANÇADOS
// ============================================================================

// EnqueueAll adiciona múltiplos elementos de uma vez
// Complexidade: O(m) onde m é o número de elementos
func (q *LinkedListDeque) EnqueueAll(elements []int) {
	for _, element := range elements {
		q.EnqueueRear(element)
	}
}

// DequeueMultiple remove e retorna múltiplos elementos do início
// Complexidade: O(n) onde n é o número de elementos a remover
func (q *LinkedListDeque) DequeueMultiple(count int) ([]int, error) {
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
// Complexidade: O(n)
func (q *LinkedListDeque) Contains(element int) bool {
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
func (q *LinkedListDeque) IndexOf(element int) int {
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
func (q *LinkedListDeque) Clone() *LinkedListDeque {
	return q.CloneLinkedListDeque()
}

// CloneLinkedListDeque cria uma cópia independente retornando *LinkedListDeque
func (q *LinkedListDeque) CloneLinkedListDeque() *LinkedListDeque {
	newQueue := NewLinkedListDeque()
	
	if q.IsEmpty() {
		return newQueue
	}
	
	current := q.front
	for current != nil {
		newQueue.EnqueueRear(current.data)
		current = current.next
	}
	
	return newQueue
}

// Equals verifica se duas filas são iguais (mesmo conteúdo e ordem)
// Complexidade: O(n)
func (q *LinkedListDeque) Equals(other *LinkedListDeque) bool {
	return q.EqualsLinkedListDeque(other)
}

// EqualsLinkedListDeque verifica se duas filas LinkedListDeque são iguais
func (q *LinkedListDeque) EqualsLinkedListDeque(other *LinkedListDeque) bool {
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
func (q *LinkedListDeque) Reverse() {
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

// GetStatistics retorna estatísticas da fila
// Complexidade: O(n)
func (q *LinkedListDeque) GetStatistics() map[string]interface{} {
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
func (q *LinkedListDeque) GetMemoryInfo() map[string]interface{} {
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

// Filter cria uma nova fila com elementos que satisfazem uma condição
// Mantém a ordem relativa dos elementos
func (q *LinkedListDeque) Filter(predicate func(int) bool) *LinkedListDeque {
	result := NewLinkedListDeque()
	
	current := q.front
	for current != nil {
		if predicate(current.data) {
			result.EnqueueRear(current.data)
		}
		current = current.next
	}
	
	return result
}

// Map aplica uma função a todos os elementos e retorna nova fila
func (q *LinkedListDeque) Map(mapper func(int) int) *LinkedListDeque {
	result := NewLinkedListDeque()
	
	current := q.front
	for current != nil {
		result.EnqueueRear(mapper(current.data))
		current = current.next
	}
	
	return result
}

// Reduce aplica uma função de redução a todos os elementos
func (q *LinkedListDeque) Reduce(reducer func(int, int) int, initialValue int) int {
	result := initialValue
	current := q.front
	
	for current != nil {
		result = reducer(result, current.data)
		current = current.next
	}
	
	return result
}

// ForEach executa uma função para cada elemento (do início para o final)
func (q *LinkedListDeque) ForEach(action func(int, int)) {
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
func (q *LinkedListDeque) Partition(predicate func(int) bool) (*LinkedListDeque, *LinkedListDeque) {
	true_queue := NewLinkedListDeque()
	false_queue := NewLinkedListDeque()
	
	current := q.front
	for current != nil {
		if predicate(current.data) {
			true_queue.EnqueueRear(current.data)
		} else {
			false_queue.EnqueueRear(current.data)
		}
		current = current.next
	}
	
	return true_queue, false_queue
}

// TakeWhile retorna nova fila com elementos do início que satisfazem condição
func (q *LinkedListDeque) TakeWhile(predicate func(int) bool) *LinkedListDeque {
	result := NewLinkedListDeque()
	
	current := q.front
	for current != nil {
		if !predicate(current.data) {
			break
		}
		result.EnqueueRear(current.data)
		current = current.next
	}
	
	return result
}

// DropWhile retorna nova fila removendo elementos do início que satisfazem condição
func (q *LinkedListDeque) DropWhile(predicate func(int) bool) *LinkedListDeque {
	result := NewLinkedListDeque()
	
	current := q.front
	// Pula elementos que satisfazem o predicado
	for current != nil && predicate(current.data) {
		current = current.next
	}
	
	// Adiciona o resto
	for current != nil {
		result.EnqueueRear(current.data)
		current = current.next
	}
	
	return result
}

// Rotate rotaciona a fila k posições para a esquerda
func (q *LinkedListDeque) Rotate(k int) {
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

// Split divide a fila em duas partes no índice especificado
func (q *LinkedListDeque) Split(index int) (*LinkedListDeque, *LinkedListDeque) {
	first := NewLinkedListDeque()
	second := NewLinkedListDeque()
	
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
			first.EnqueueRear(current.data)
		} else {
			second.EnqueueRear(current.data)
		}
		current = current.next
		currentIndex++
	}
	
	return first, second
}

// Merge combina esta fila com outra, alternando elementos
func (q *LinkedListDeque) Merge(other *LinkedListDeque) *LinkedListDeque {
	result := NewLinkedListDeque()
	
	current1 := q.front
	current2 := other.front
	
	// Alterna entre as duas filas
	for current1 != nil || current2 != nil {
		if current1 != nil {
			result.EnqueueRear(current1.data)
			current1 = current1.next
		}
		if current2 != nil {
			result.EnqueueRear(current2.data)
			current2 = current2.next
		}
	}
	
	return result
}

// GetNth retorna o n-ésimo elemento (0-indexado) sem removê-lo
func (q *LinkedListDeque) GetNth(n int) (int, error) {
	if n < 0 || n >= q.size {
		return 0, fmt.Errorf("índice fora dos limites: %d", n)
	}
	
	current := q.front
	for i := 0; i < n; i++ {
		current = current.next
	}
	
	return current.data, nil
}

// ============================================================================
// MÉTODOS COMPATÍVEIS COM INTERFACE IDEQUE
// ============================================================================

// CloneIDeque implementa IDeque.Clone() retornando IDeque
func (q *LinkedListDeque) CloneIDeque() IDeque {
	return q.CloneLinkedListDeque()
}

// EqualsIDeque implementa IDeque.Equals() aceitando IDeque
func (q *LinkedListDeque) EqualsIDeque(other IDeque) bool {
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
func (q *LinkedListDeque) FilterIDeque(predicate func(int) bool) IDeque {
	return q.Filter(predicate)
}

// MapIDeque implementa IDeque.Map() retornando IDeque
func (q *LinkedListDeque) MapIDeque(mapper func(int) int) IDeque {
	return q.Map(mapper)
}