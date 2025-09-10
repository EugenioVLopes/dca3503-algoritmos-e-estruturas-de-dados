package deque

import (
	"errors"
	"fmt"
	"strings"
)

// ============================================================================
// DEQUE - IMPLEMENTAÇÃO BASEADA EM LISTA DUPLAMENTE LIGADA
// ============================================================================

// DequeNode representa um nó no deque
type DequeNode struct {
	data int        // Valor armazenado no nó
	next *DequeNode // Ponteiro para o próximo nó
	prev *DequeNode // Ponteiro para o nó anterior
}

// Deque implementa uma fila de duas extremidades (Double-ended queue)
// Características:
// - Inserção e remoção em ambas extremidades O(1)
// - Combina funcionalidades de pilha e fila
// - Uso dinâmico de memória
// - Navegação bidirecional
type Deque struct {
	front *DequeNode // Ponteiro para o primeiro nó
	rear  *DequeNode // Ponteiro para o último nó
	size  int        // Contador de elementos
}

// NewDeque cria uma nova instância de Deque
func NewDeque() *Deque {
	return &Deque{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

// ============================================================================
// MÉTODOS PRINCIPAIS DO DEQUE
// ============================================================================

// EnqueueFront adiciona elemento no início do deque
// Complexidade: O(1)
func (d *Deque) EnqueueFront(value int) {
	newNode := &DequeNode{
		data: value,
		next: d.front,
		prev: nil,
	}
	
	if d.IsEmpty() {
		// Primeiro elemento
		d.front = newNode
		d.rear = newNode
	} else {
		// Conecta novo nó ao início
		d.front.prev = newNode
		d.front = newNode
	}
	
	d.size++
}

// EnqueueRear adiciona elemento no final do deque
// Complexidade: O(1)
func (d *Deque) EnqueueRear(value int) {
	newNode := &DequeNode{
		data: value,
		next: nil,
		prev: d.rear,
	}
	
	if d.IsEmpty() {
		// Primeiro elemento
		d.front = newNode
		d.rear = newNode
	} else {
		// Conecta novo nó ao final
		d.rear.next = newNode
		d.rear = newNode
	}
	
	d.size++
}

// DequeueFront remove e retorna elemento do início do deque
// Complexidade: O(1)
func (d *Deque) DequeueFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque vazio: não é possível remover do início")
	}
	
	value := d.front.data
	
	if d.size == 1 {
		// Último elemento
		d.front = nil
		d.rear = nil
	} else {
		// Move front para o próximo
		d.front = d.front.next
		d.front.prev = nil
	}
	
	d.size--
	return value, nil
}

// DequeueRear remove e retorna elemento do final do deque
// Complexidade: O(1)
func (d *Deque) DequeueRear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque vazio: não é possível remover do final")
	}
	
	value := d.rear.data
	
	if d.size == 1 {
		// Último elemento
		d.front = nil
		d.rear = nil
	} else {
		// Move rear para o anterior
		d.rear = d.rear.prev
		d.rear.next = nil
	}
	
	d.size--
	return value, nil
}

// Front retorna elemento do início sem removê-lo
// Complexidade: O(1)
func (d *Deque) Front() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque vazio: não há elemento no início")
	}
	return d.front.data, nil
}

// Rear retorna elemento do final sem removê-lo
// Complexidade: O(1)
func (d *Deque) Rear() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque vazio: não há elemento no final")
	}
	return d.rear.data, nil
}

// IsEmpty verifica se o deque está vazio
// Complexidade: O(1)
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

// Size retorna o número de elementos no deque
// Complexidade: O(1)
func (d *Deque) Size() int {
	return d.size
}

// ============================================================================
// MÉTODOS AUXILIARES
// ============================================================================

// Clear remove todos os elementos do deque
// Complexidade: O(1)
func (d *Deque) Clear() {
	d.front = nil
	d.rear = nil
	d.size = 0
}

// ToSlice converte o deque para um slice (do início para o final)
// Complexidade: O(n)
func (d *Deque) ToSlice() []int {
	if d.IsEmpty() {
		return []int{}
	}
	
	result := make([]int, d.size)
	current := d.front
	index := 0
	
	for current != nil {
		result[index] = current.data
		current = current.next
		index++
	}
	
	return result
}

// ToSliceReverse converte o deque para um slice (do final para o início)
// Complexidade: O(n)
func (d *Deque) ToSliceReverse() []int {
	if d.IsEmpty() {
		return []int{}
	}
	
	result := make([]int, d.size)
	current := d.rear
	index := 0
	
	for current != nil {
		result[index] = current.data
		current = current.prev
		index++
	}
	
	return result
}

// String retorna uma representação em string do deque
// Complexidade: O(n)
func (d *Deque) String() string {
	if d.IsEmpty() {
		return "[vazio]"
	}
	
	var builder strings.Builder
	builder.WriteString("início → [")
	
	current := d.front
	first := true
	
	for current != nil {
		if !first {
			builder.WriteString(" ⇄ ")
		}
		builder.WriteString(fmt.Sprintf("%d", current.data))
		current = current.next
		first = false
	}
	
	builder.WriteString("] ← final")
	return builder.String()
}

// ============================================================================
// MÉTODOS ESPECÍFICOS DO DEQUE
// ============================================================================

// PushFront é um alias para EnqueueFront (comportamento de pilha no início)
func (d *Deque) PushFront(value int) {
	d.EnqueueFront(value)
}

// PushRear é um alias para EnqueueRear (comportamento de pilha no final)
func (d *Deque) PushRear(value int) {
	d.EnqueueRear(value)
}

// PopFront é um alias para DequeueFront (comportamento de pilha no início)
func (d *Deque) PopFront() (int, error) {
	return d.DequeueFront()
}

// PopRear é um alias para DequeueRear (comportamento de pilha no final)
func (d *Deque) PopRear() (int, error) {
	return d.DequeueRear()
}

// PeekFront é um alias para Front
func (d *Deque) PeekFront() (int, error) {
	return d.Front()
}

// PeekRear é um alias para Rear
func (d *Deque) PeekRear() (int, error) {
	return d.Rear()
}

// ============================================================================
// MÉTODOS AVANÇADOS
// ============================================================================

// Contains verifica se o deque contém um elemento específico
// Complexidade: O(n)
func (d *Deque) Contains(element int) bool {
	current := d.front
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
func (d *Deque) IndexOf(element int) int {
	current := d.front
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

// GetAt retorna o elemento na posição especificada (0-indexado)
// Complexidade: O(n) no pior caso, O(n/2) em média (busca bidirecional)
func (d *Deque) GetAt(index int) (int, error) {
	if index < 0 || index >= d.size {
		return 0, fmt.Errorf("índice fora dos limites: %d", index)
	}
	
	// Otimização: escolhe direção mais próxima
	if index < d.size/2 {
		// Busca do início
		current := d.front
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current.data, nil
	} else {
		// Busca do final
		current := d.rear
		for i := d.size - 1; i > index; i-- {
			current = current.prev
		}
		return current.data, nil
	}
}

// RemoveAt remove elemento na posição especificada
// Complexidade: O(n) no pior caso, O(n/2) em média
func (d *Deque) RemoveAt(index int) (int, error) {
	if index < 0 || index >= d.size {
		return 0, fmt.Errorf("índice fora dos limites: %d", index)
	}
	
	// Casos especiais para extremidades
	if index == 0 {
		return d.DequeueFront()
	}
	if index == d.size-1 {
		return d.DequeueRear()
	}
	
	// Encontra o nó
	var nodeToRemove *DequeNode
	if index < d.size/2 {
		// Busca do início
		nodeToRemove = d.front
		for i := 0; i < index; i++ {
			nodeToRemove = nodeToRemove.next
		}
	} else {
		// Busca do final
		nodeToRemove = d.rear
		for i := d.size - 1; i > index; i-- {
			nodeToRemove = nodeToRemove.prev
		}
	}
	
	// Remove o nó
	value := nodeToRemove.data
	nodeToRemove.prev.next = nodeToRemove.next
	nodeToRemove.next.prev = nodeToRemove.prev
	d.size--
	
	return value, nil
}

// InsertAt insere elemento na posição especificada
// Complexidade: O(n) no pior caso, O(n/2) em média
func (d *Deque) InsertAt(index int, value int) error {
	if index < 0 || index > d.size {
		return fmt.Errorf("índice fora dos limites: %d", index)
	}
	
	// Casos especiais para extremidades
	if index == 0 {
		d.EnqueueFront(value)
		return nil
	}
	if index == d.size {
		d.EnqueueRear(value)
		return nil
	}
	
	// Encontra posição de inserção
	var nextNode *DequeNode
	if index < d.size/2 {
		// Busca do início
		nextNode = d.front
		for i := 0; i < index; i++ {
			nextNode = nextNode.next
		}
	} else {
		// Busca do final
		nextNode = d.rear
		for i := d.size - 1; i > index; i-- {
			nextNode = nextNode.prev
		}
	}
	
	// Cria e insere novo nó
	newNode := &DequeNode{
		data: value,
		next: nextNode,
		prev: nextNode.prev,
	}
	
	nextNode.prev.next = newNode
	nextNode.prev = newNode
	d.size++
	
	return nil
}

// Reverse inverte a ordem dos elementos no deque
// Complexidade: O(n)
func (d *Deque) Reverse() {
	if d.size <= 1 {
		return
	}
	
	current := d.front
	
	// Troca ponteiros prev e next de cada nó
	for current != nil {
		next := current.next
		current.next = current.prev
		current.prev = next
		current = next
	}
	
	// Troca front e rear
	d.front, d.rear = d.rear, d.front
}

// Clone cria uma cópia independente do deque
// Complexidade: O(n)
func (d *Deque) Clone() *Deque {
	return d.CloneDeque()
}

// CloneDeque cria uma cópia independente retornando *Deque
func (d *Deque) CloneDeque() *Deque {
	newDeque := NewDeque()
	
	current := d.front
	for current != nil {
		newDeque.EnqueueRear(current.data)
		current = current.next
	}
	
	return newDeque
}

// Equals verifica se dois deques são iguais
// Complexidade: O(n)
func (d *Deque) Equals(other *Deque) bool {
	return d.EqualsDeque(other)
}

// EqualsDeque verifica se dois deques Deque são iguais
func (d *Deque) EqualsDeque(other *Deque) bool {
	if d.size != other.size {
		return false
	}
	
	current1 := d.front
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

// ============================================================================
// MÉTODOS DE ITERAÇÃO
// ============================================================================

// ForEachForward executa uma função para cada elemento (do início para o final)
func (d *Deque) ForEachForward(action func(int, int)) {
	current := d.front
	index := 0
	
	for current != nil {
		action(current.data, index)
		current = current.next
		index++
	}
}

// ForEachBackward executa uma função para cada elemento (do final para o início)
func (d *Deque) ForEachBackward(action func(int, int)) {
	current := d.rear
	index := d.size - 1
	
	for current != nil {
		action(current.data, index)
		current = current.prev
		index--
	}
}

// Filter cria um novo deque com elementos que satisfazem uma condição
func (d *Deque) Filter(predicate func(int) bool) *Deque {
	result := NewDeque()
	
	current := d.front
	for current != nil {
		if predicate(current.data) {
			result.EnqueueRear(current.data)
		}
		current = current.next
	}
	
	return result
}

// Map aplica uma função a todos os elementos e retorna novo deque
func (d *Deque) Map(mapper func(int) int) *Deque {
	result := NewDeque()
	
	current := d.front
	for current != nil {
		result.EnqueueRear(mapper(current.data))
		current = current.next
	}
	
	return result
}

// ============================================================================
// MÉTODOS DE ESTATÍSTICAS
// ============================================================================

// GetStatistics retorna estatísticas do deque
func (d *Deque) GetStatistics() map[string]interface{} {
	if d.IsEmpty() {
		return map[string]interface{}{
			"size":    0,
			"isEmpty": true,
		}
	}
	
	sum := 0
	min := d.front.data
	max := d.front.data
	current := d.front
	
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
	
	average := float64(sum) / float64(d.size)
	
	return map[string]interface{}{
		"size":    d.size,
		"isEmpty": false,
		"sum":     sum,
		"average": average,
		"min":     min,
		"max":     max,
	}
}

// GetMemoryInfo retorna informações sobre uso de memória
func (d *Deque) GetMemoryInfo() map[string]interface{} {
	// Cada nó tem: int (8 bytes) + 2 ponteiros (16 bytes) = 24 bytes
	nodeSize := 24
	totalMemory := d.size * nodeSize
	
	return map[string]interface{}{
		"nodes":          d.size,
		"estimatedBytes": totalMemory,
		"bytesPerNode":   nodeSize,
		"overhead":       "2 ponteiros por nó (prev/next)",
	}
}

// ============================================================================
// MÉTODOS COMPATÍVEIS COM INTERFACE IDEQUE
// ============================================================================

// CloneIDeque implementa IDeque.Clone() retornando IDeque
func (d *Deque) CloneIDeque() IDeque {
	return d.CloneDeque()
}

// EqualsIDeque implementa IDeque.Equals() aceitando IDeque
func (d *Deque) EqualsIDeque(other IDeque) bool {
	if d.Size() != other.Size() {
		return false
	}
	
	slice1 := d.ToSlice()
	slice2 := other.ToSlice()
	
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	
	return true
}

// FilterIDeque implementa IDeque.Filter() retornando IDeque
func (d *Deque) FilterIDeque(predicate func(int) bool) IDeque {
	return d.Filter(predicate)
}

// MapIDeque implementa IDeque.Map() retornando IDeque
func (d *Deque) MapIDeque(mapper func(int) int) IDeque {
	return d.Map(mapper)
}

// ============================================================================
// EXEMPLO DE USO
// ============================================================================

/*
func main() {
	deque := NewDeque()
	
	// Teste de inserções
	fmt.Println("=== Testando Deque ===")
	deque.EnqueueFront(10)
	deque.EnqueueRear(20)
	deque.EnqueueFront(5)
	deque.EnqueueRear(25)
	
	fmt.Printf("Deque: %s\n", deque)
	fmt.Printf("Tamanho: %d\n", deque.Size())
	
	// Teste de acesso
	if front, err := deque.Front(); err == nil {
		fmt.Printf("Primeiro: %d\n", front)
	}
	if rear, err := deque.Rear(); err == nil {
		fmt.Printf("Último: %d\n", rear)
	}
	
	// Teste de remoções
	fmt.Println("\nRemoções:")
	if value, err := deque.DequeueFront(); err == nil {
		fmt.Printf("Removido do início: %d\n", value)
	}
	if value, err := deque.DequeueRear(); err == nil {
		fmt.Printf("Removido do final: %d\n", value)
	}
	
	fmt.Printf("Deque após remoções: %s\n", deque)
	
	// Teste de busca
	fmt.Printf("Contém 10: %t\n", deque.Contains(10))
	fmt.Printf("Índice do 20: %d\n", deque.IndexOf(20))
	
	// Estatísticas
	stats := deque.GetStatistics()
	fmt.Printf("Estatísticas: %+v\n", stats)
}
*/