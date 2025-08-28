package main

import (
	"fmt"
	"errors"
	"strings"
)

// ============================================================================
// ARRAYSTACK - IMPLEMENTAÇÃO BASEADA EM ARRAY DINÂMICO
// ============================================================================

// ArrayStack implementa uma pilha usando um array dinâmico (slice em Go)
// Características:
// - Push/Pop no topo são operações O(1)
// - Acesso direto ao topo O(1)
// - Uso eficiente de memória (elementos contíguos)
// - Redimensionamento automático quando necessário
type ArrayStack struct {
	data     []int // Array interno que armazena os elementos
	top      int   // Índice do elemento no topo (-1 se vazia)
	capacity int   // Capacidade atual do array
}

// NewArrayStack cria uma nova instância de ArrayStack com capacidade inicial
func NewArrayStack(initialCapacity int) *ArrayStack {
	if initialCapacity <= 0 {
		initialCapacity = 10 // Capacidade padrão
	}
	return &ArrayStack{
		data:     make([]int, initialCapacity),
		top:      -1,
		capacity: initialCapacity,
	}
}

// ============================================================================
// IMPLEMENTAÇÃO DA INTERFACE STACK
// ============================================================================

// Push adiciona um elemento no topo da pilha
// Complexidade: O(1) amortizado (pode ser O(n) quando redimensiona)
func (s *ArrayStack) Push(element int) {
	// Verifica se precisa redimensionar
	if s.top+1 >= s.capacity {
		s.resize(s.capacity * 2)
	}
	
	s.top++
	s.data[s.top] = element
}

// Pop remove e retorna o elemento do topo da pilha
// Complexidade: O(1)
func (s *ArrayStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("pilha vazia: não é possível fazer pop")
	}
	
	element := s.data[s.top]
	s.top--
	
	// Redimensiona para baixo se necessário (economiza memória)
	if s.Size() > 0 && s.Size() == s.capacity/4 {
		s.resize(s.capacity / 2)
	}
	
	return element, nil
}

// Peek retorna o elemento do topo sem removê-lo
// Complexidade: O(1)
func (s *ArrayStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("pilha vazia: não há elemento no topo")
	}
	return s.data[s.top], nil
}

// Size retorna o número de elementos na pilha
// Complexidade: O(1)
func (s *ArrayStack) Size() int {
	return s.top + 1
}

// IsEmpty verifica se a pilha está vazia
// Complexidade: O(1)
func (s *ArrayStack) IsEmpty() bool {
	return s.top == -1
}

// IsFull verifica se a pilha está cheia
// Para ArrayStack, nunca está "cheia" pois redimensiona automaticamente
// Complexidade: O(1)
func (s *ArrayStack) IsFull() bool {
	return false // Array dinâmico nunca está "cheio"
}

// Clear remove todos os elementos da pilha
// Complexidade: O(1)
func (s *ArrayStack) Clear() {
	s.top = -1
	// Opcionalmente, pode redimensionar para capacidade inicial
	if s.capacity > 10 {
		s.resize(10)
	}
}

// ToSlice converte a pilha para um slice (do topo para a base)
// Complexidade: O(n)
func (s *ArrayStack) ToSlice() []int {
	if s.IsEmpty() {
		return []int{}
	}
	
	result := make([]int, s.Size())
	for i := 0; i < s.Size(); i++ {
		result[i] = s.data[s.top-i] // Do topo para a base
	}
	return result
}

// String retorna uma representação em string da pilha
// Complexidade: O(n)
func (s *ArrayStack) String() string {
	if s.IsEmpty() {
		return "[vazia]"
	}
	
	var builder strings.Builder
	builder.WriteString("[")
	
	// Mostra do topo para a base
	for i := s.top; i >= 0; i-- {
		if i < s.top {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("%d", s.data[i]))
	}
	
	builder.WriteString("] ← topo")
	return builder.String()
}

// ============================================================================
// MÉTODOS AUXILIARES ESPECÍFICOS DO ARRAYSTACK
// ============================================================================

// Capacity retorna a capacidade atual do array interno
func (s *ArrayStack) Capacity() int {
	return s.capacity
}

// resize redimensiona o array interno para uma nova capacidade
func (s *ArrayStack) resize(newCapacity int) {
	if newCapacity < s.Size() {
		newCapacity = s.Size() // Não pode ser menor que o tamanho atual
	}
	
	newData := make([]int, newCapacity)
	copy(newData, s.data[:s.Size()])
	s.data = newData
	s.capacity = newCapacity
}

// TrimToSize reduz a capacidade para o tamanho atual (economiza memória)
func (s *ArrayStack) TrimToSize() {
	if s.capacity > s.Size() {
		newCapacity := s.Size()
		if newCapacity == 0 {
			newCapacity = 1 // Mantém pelo menos 1 de capacidade
		}
		s.resize(newCapacity)
	}
}

// EnsureCapacity garante que a pilha tenha pelo menos a capacidade especificada
func (s *ArrayStack) EnsureCapacity(minCapacity int) {
	if s.capacity < minCapacity {
		s.resize(minCapacity)
	}
}

// PushAll adiciona múltiplos elementos de uma vez
func (s *ArrayStack) PushAll(elements []int) {
	// Garante capacidade suficiente
	s.EnsureCapacity(s.Size() + len(elements))
	
	for _, element := range elements {
		s.Push(element)
	}
}

// PopMultiple remove e retorna múltiplos elementos do topo
func (s *ArrayStack) PopMultiple(count int) ([]int, error) {
	if count <= 0 {
		return []int{}, nil
	}
	
	if count > s.Size() {
		return nil, fmt.Errorf("não há elementos suficientes: solicitado %d, disponível %d", count, s.Size())
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
func (s *ArrayStack) Contains(element int) bool {
	for i := 0; i <= s.top; i++ {
		if s.data[i] == element {
			return true
		}
	}
	return false
}

// Search procura um elemento e retorna sua posição a partir do topo
// Retorna -1 se não encontrado
// Complexidade: O(n)
func (s *ArrayStack) Search(element int) int {
	for i := s.top; i >= 0; i-- {
		if s.data[i] == element {
			return s.top - i + 1 // Posição a partir do topo (1-indexado)
		}
	}
	return -1
}

// Clone cria uma cópia independente da pilha
func (s *ArrayStack) Clone() *ArrayStack {
	newStack := NewArrayStack(s.capacity)
	newStack.top = s.top
	copy(newStack.data, s.data[:s.Size()])
	return newStack
}

// Equals verifica se duas pilhas são iguais (mesmo conteúdo e ordem)
func (s *ArrayStack) Equals(other *ArrayStack) bool {
	if s.Size() != other.Size() {
		return false
	}
	
	for i := 0; i <= s.top; i++ {
		if s.data[i] != other.data[i] {
			return false
		}
	}
	
	return true
}

// ============================================================================
// MÉTODOS DE ANÁLISE E ESTATÍSTICAS
// ============================================================================

// GetMemoryUsage retorna informações sobre uso de memória
func (s *ArrayStack) GetMemoryUsage() (int, int, float64) {
	usedSlots := s.Size()
	totalSlots := s.capacity
	utilization := float64(usedSlots) / float64(totalSlots) * 100
	return usedSlots, totalSlots, utilization
}

// GetStatistics retorna estatísticas da pilha
func (s *ArrayStack) GetStatistics() map[string]interface{} {
	if s.IsEmpty() {
		return map[string]interface{}{
			"size":         0,
			"capacity":     s.capacity,
			"utilization": 0.0,
			"isEmpty":      true,
		}
	}
	
	// Calcula estatísticas
	sum := 0
	min := s.data[0]
	max := s.data[0]
	
	for i := 0; i <= s.top; i++ {
		value := s.data[i]
		sum += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
	average := float64(sum) / float64(s.Size())
	utilization := float64(s.Size()) / float64(s.capacity) * 100
	
	return map[string]interface{}{
		"size":         s.Size(),
		"capacity":     s.capacity,
		"utilization":  utilization,
		"isEmpty":      false,
		"sum":          sum,
		"average":      average,
		"min":          min,
		"max":          max,
	}
}