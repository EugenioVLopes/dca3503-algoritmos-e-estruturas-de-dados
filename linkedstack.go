package main

import (
	"fmt"
	"errors"
	"strings"
)

// ============================================================================
// LINKEDSTACK - IMPLEMENTAÇÃO BASEADA EM LISTA LIGADA
// ============================================================================

// StackNode representa um nó na pilha ligada
type StackNode struct {
	data int        // Valor armazenado no nó
	next *StackNode // Ponteiro para o próximo nó
}

// LinkedStack implementa uma pilha usando lista ligada
// Características:
// - Push/Pop são sempre O(1)
// - Não há limite de capacidade (limitado apenas pela memória)
// - Uso dinâmico de memória (aloca conforme necessário)
// - Cada elemento tem overhead de ponteiro
type LinkedStack struct {
	top  *StackNode // Ponteiro para o nó do topo
	size int        // Contador de elementos
}

// NewLinkedStack cria uma nova instância de LinkedStack
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		top:  nil,
		size: 0,
	}
}

// ============================================================================
// IMPLEMENTAÇÃO DA INTERFACE STACK
// ============================================================================

// Push adiciona um elemento no topo da pilha
// Complexidade: O(1)
func (s *LinkedStack) Push(element int) {
	newNode := &StackNode{
		data: element,
		next: s.top,
	}
	s.top = newNode
	s.size++
}

// Pop remove e retorna o elemento do topo da pilha
// Complexidade: O(1)
func (s *LinkedStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("pilha vazia: não é possível fazer pop")
	}
	
	value := s.top.data
	s.top = s.top.next
	s.size--
	
	return value, nil
}

// Peek retorna o elemento do topo sem removê-lo
// Complexidade: O(1)
func (s *LinkedStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("pilha vazia: não há elemento no topo")
	}
	return s.top.data, nil
}

// Size retorna o número de elementos na pilha
// Complexidade: O(1)
func (s *LinkedStack) Size() int {
	return s.size
}

// IsEmpty verifica se a pilha está vazia
// Complexidade: O(1)
func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}

// IsFull verifica se a pilha está cheia
// Para LinkedStack, nunca está "cheia" pois usa alocação dinâmica
// Complexidade: O(1)
func (s *LinkedStack) IsFull() bool {
	return false // Lista ligada nunca está "cheia"
}

// Clear remove todos os elementos da pilha
// Complexidade: O(1) - apenas redefine ponteiros, GC limpa os nós
func (s *LinkedStack) Clear() {
	s.top = nil
	s.size = 0
	// O Garbage Collector do Go automaticamente limpa os nós órfãos
}

// ToSlice converte a pilha para um slice (do topo para a base)
// Complexidade: O(n)
func (s *LinkedStack) ToSlice() []int {
	if s.IsEmpty() {
		return []int{}
	}
	
	result := make([]int, s.size)
	current := s.top
	index := 0
	
	for current != nil {
		result[index] = current.data
		current = current.next
		index++
	}
	
	return result
}

// String retorna uma representação em string da pilha
// Complexidade: O(n)
func (s *LinkedStack) String() string {
	if s.IsEmpty() {
		return "[vazia]"
	}
	
	var builder strings.Builder
	builder.WriteString("[")
	
	current := s.top
	first := true
	
	for current != nil {
		if !first {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("%d", current.data))
		current = current.next
		first = false
	}
	
	builder.WriteString("] ← topo")
	return builder.String()
}

// ============================================================================
// MÉTODOS AUXILIARES ESPECÍFICOS DO LINKEDSTACK
// ============================================================================

// PushAll adiciona múltiplos elementos de uma vez
// Complexidade: O(n) onde n é o número de elementos
func (s *LinkedStack) PushAll(elements []int) {
	for _, element := range elements {
		s.Push(element)
	}
}

// PopMultiple remove e retorna múltiplos elementos do topo
// Complexidade: O(n) onde n é o número de elementos a remover
func (s *LinkedStack) PopMultiple(count int) ([]int, error) {
	if count <= 0 {
		return []int{}, nil
	}
	
	if count > s.size {
		return nil, fmt.Errorf("não há elementos suficientes: solicitado %d, disponível %d", count, s.size)
	}
	
	result := make([]int, count)
	for i := 0; i < count; i++ {
		value, _ := s.Pop()
		result[i] = value
	}
	
	return result, nil
}

// Contains verifica se a pilha contém um elemento específico
// Complexidade: O(n)
func (s *LinkedStack) Contains(element int) bool {
	current := s.top
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

// Search procura um elemento e retorna sua posição a partir do topo
// Retorna -1 se não encontrado
// Complexidade: O(n)
func (s *LinkedStack) Search(element int) int {
	current := s.top
	position := 1
	
	for current != nil {
		if current.data == element {
			return position
		}
		current = current.next
		position++
	}
	
	return -1
}

// Clone cria uma cópia independente da pilha
// Complexidade: O(n)
func (s *LinkedStack) Clone() *LinkedStack {
	newStack := NewLinkedStack()
	
	if s.IsEmpty() {
		return newStack
	}
	
	// Usa uma pilha auxiliar para manter a ordem
	aux := NewLinkedStack()
	
	// Move elementos para auxiliar
	current := s.top
	for current != nil {
		aux.Push(current.data)
		current = current.next
	}
	
	// Move de volta para nova pilha (restaura ordem original)
	for !aux.IsEmpty() {
		value, _ := aux.Pop()
		newStack.Push(value)
	}
	
	return newStack
}

// Equals verifica se duas pilhas são iguais (mesmo conteúdo e ordem)
// Complexidade: O(n)
func (s *LinkedStack) Equals(other *LinkedStack) bool {
	if s.size != other.size {
		return false
	}
	
	current1 := s.top
	current2 := other.top
	
	for current1 != nil && current2 != nil {
		if current1.data != current2.data {
			return false
		}
		current1 = current1.next
		current2 = current2.next
	}
	
	return true
}

// Reverse inverte a ordem dos elementos na pilha
// Complexidade: O(n)
func (s *LinkedStack) Reverse() {
	if s.size <= 1 {
		return
	}
	
	var prev *StackNode = nil
	current := s.top
	
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	
	s.top = prev
}

// ============================================================================
// MÉTODOS DE ANÁLISE E ESTATÍSTICAS
// ============================================================================

// GetStatistics retorna estatísticas da pilha
// Complexidade: O(n)
func (s *LinkedStack) GetStatistics() map[string]interface{} {
	if s.IsEmpty() {
		return map[string]interface{}{
			"size":    0,
			"isEmpty": true,
		}
	}
	
	// Calcula estatísticas
	sum := 0
	min := s.top.data
	max := s.top.data
	current := s.top
	
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
	
	average := float64(sum) / float64(s.size)
	
	return map[string]interface{}{
		"size":    s.size,
		"isEmpty": false,
		"sum":     sum,
		"average": average,
		"min":     min,
		"max":     max,
	}
}

// GetMemoryInfo retorna informações sobre uso de memória
// Estima o uso de memória baseado no número de nós
func (s *LinkedStack) GetMemoryInfo() map[string]interface{} {
	// Cada nó tem: int (8 bytes) + ponteiro (8 bytes) = 16 bytes
	nodeSize := 16
	totalMemory := s.size * nodeSize
	
	return map[string]interface{}{
		"nodes":           s.size,
		"estimatedBytes":  totalMemory,
		"bytesPerNode":    nodeSize,
		"memoryEfficient": s.size > 0, // Sem overhead de capacidade não utilizada
	}
}

// ============================================================================
// MÉTODOS AVANÇADOS
// ============================================================================

// Filter cria uma nova pilha com elementos que satisfazem uma condição
// Mantém a ordem relativa dos elementos
func (s *LinkedStack) Filter(predicate func(int) bool) *LinkedStack {
	result := NewLinkedStack()
	aux := NewLinkedStack()
	
	// Coleta elementos que satisfazem a condição
	current := s.top
	for current != nil {
		if predicate(current.data) {
			aux.Push(current.data)
		}
		current = current.next
	}
	
	// Inverte para manter ordem original
	for !aux.IsEmpty() {
		value, _ := aux.Pop()
		result.Push(value)
	}
	
	return result
}

// Map aplica uma função a todos os elementos e retorna nova pilha
func (s *LinkedStack) Map(mapper func(int) int) *LinkedStack {
	result := NewLinkedStack()
	aux := NewLinkedStack()
	
	// Aplica função e coleta resultados
	current := s.top
	for current != nil {
		aux.Push(mapper(current.data))
		current = current.next
	}
	
	// Inverte para manter ordem original
	for !aux.IsEmpty() {
		value, _ := aux.Pop()
		result.Push(value)
	}
	
	return result
}

// Reduce aplica uma função de redução a todos os elementos
func (s *LinkedStack) Reduce(reducer func(int, int) int, initialValue int) int {
	result := initialValue
	current := s.top
	
	for current != nil {
		result = reducer(result, current.data)
		current = current.next
	}
	
	return result
}

// ForEach executa uma função para cada elemento (do topo para a base)
func (s *LinkedStack) ForEach(action func(int, int)) {
	current := s.top
	index := 0
	
	for current != nil {
		action(current.data, index)
		current = current.next
		index++
	}
}

// ToReversedSlice converte para slice na ordem inversa (da base para o topo)
func (s *LinkedStack) ToReversedSlice() []int {
	if s.IsEmpty() {
		return []int{}
	}
	
	// Primeiro, coleta todos os elementos
	elements := s.ToSlice()
	
	// Inverte o slice
	for i, j := 0, len(elements)-1; i < j; i, j = i+1, j-1 {
		elements[i], elements[j] = elements[j], elements[i]
	}
	
	return elements
}