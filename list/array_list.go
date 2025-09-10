package list

import (
	"fmt"
)

// ============================================================================
// ARRAYLIST - IMPLEMENTAÇÃO BASEADA EM ARRAY DINÂMICO
// ============================================================================

// ArrayList implementa uma lista usando um array dinâmico (slice em Go)
// Características:
// - Acesso aleatório rápido O(1)
// - Inserção/remoção no final é rápida O(1) amortizado
// - Inserção/remoção no meio é lenta O(n)
// - Uso eficiente de memória (elementos contíguos)
type ArrayList struct {
	elements        []int // Array interno que armazena os elementos
	size int              // Contador de elementos inseridos (tamanho lógico)
}

// NewArrayList cria uma nova instância de ArrayList com capacidade inicial
func NewArrayList(initialCapacity int) *ArrayList {
	return &ArrayList{
		elements: make([]int, initialCapacity),
		size:     0,
	}
}

// Init inicializa o ArrayList com capacidade inicial
// Pseudocódigo:
// 1. Criar array interno com tamanho especificado
// 2. Inicializar contador de elementos como 0
func (l *ArrayList) Init(size int) {
	l.elements = make([]int, size) // Aloca memória para 'size' elementos
	l.size = 0          // Inicialmente não há elementos inseridos
}

// Size retorna o número de elementos atualmente na lista
// Complexidade: Θ(1) - Tempo constante
// Pseudocódigo:
// 1. Retornar valor do contador de elementos inseridos
func (list *ArrayList) Size() int { // Θ(1)
	return list.size
}

// IsEmpty verifica se a lista está vazia
// Complexidade: Θ(1)
func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

// Capacity retorna a capacidade atual do array interno
// Complexidade: Θ(1)
func (list *ArrayList) Capacity() int {
	return len(list.elements)
}

// Get obtém o elemento na posição especificada
// Complexidade: Θ(1) - Acesso direto por índice
// Pseudocódigo:
// 1. Verificar se índice é válido (0 <= index < size)
// 2. Se válido: retornar elemento na posição
// 3. Se inválido: retornar erro
func (list *ArrayList) Get(index int) (int, error) { // Θ(1)
	if index >= 0 && index < list.size {
		return list.elements[index], nil // Acesso direto O(1)
	} else {
		return -1, fmt.Errorf("index inválido: %d", index)
	}
}

// Set define o valor do elemento na posição especificada
// Complexidade: Θ(1)
func (list *ArrayList) Set(index int, value int) error {
	if index >= 0 && index < list.size {
		list.elements[index] = value
		return nil
	} else {
		return fmt.Errorf("index inválido: %d", index)
	}
}

// resize redimensiona o array para uma nova capacidade
// Complexidade: Θ(n)
func (list *ArrayList) resize(newCapacity int) {
	if newCapacity < list.size {
		return // Não pode reduzir abaixo do tamanho atual
	}
	
	newElements := make([]int, newCapacity)
	for i := 0; i < list.size; i++ {
		newElements[i] = list.elements[i]
	}
	list.elements = newElements
}

// doubleV dobra a capacidade do array interno quando necessário
// Complexidade: Θ(n) - Precisa copiar todos os elementos
func (list *ArrayList) doubleV() { // Θ(n)
	list.resize(len(list.elements) * 2)
}

// Add adiciona elemento no final da lista
// Complexidade: O(n) pior caso, Ω(1) melhor caso, O(1) amortizado
// Pseudocódigo:
// 1. Se array está cheio: dobrar capacidade
// 2. Inserir elemento na próxima posição disponível
// 3. Incrementar contador de elementos
func (list *ArrayList) Add(value int) { // O(n), Ω(1)
	// Verifica se precisa expandir o array
	if list.size == len(list.elements) {
		list.doubleV() // O(n) apenas quando necessário
	}
	
	// Inserção no final é sempre O(1)
	list.elements[list.size] = value
	list.size++
}

// AddOnIndex adiciona elemento em posição específica
// Complexidade: O(n) pior caso, Ω(1) melhor caso
// Pseudocódigo:
// 1. Validar índice (0 <= index <= inserted)
// 2. Se array cheio: dobrar capacidade
// 3. Deslocar elementos à direita do índice uma posição para frente
// 4. Inserir novo elemento na posição
// 5. Incrementar contador
func (list *ArrayList) AddOnIndex(val int, index int) error { // O(n), Ω(1)
	if index < 0 || index > list.size {
		return fmt.Errorf("index inválido: %d", index)
	}
	
	if list.size == len(list.elements) {
		list.doubleV()
	}

	for i := list.size; i > index; i-- {
		list.elements[i] = list.elements[i-1]
	}
	
	list.elements[index] = val
	list.size++
	return nil
}
// Remove remove elemento de posição específica
// Complexidade: Ω(1) melhor caso, O(n) pior caso
// Pseudocódigo:
// 1. Validar índice (0 <= index < size)
// 2. Deslocar elementos à direita do índice uma posição para esquerda
// 3. Decrementar contador de elementos
func (list *ArrayList) Remove(index int) error { // Ω(1), O(n)
	if index >= 0 && index < list.size {
		// Desloca elementos para a esquerda - O(n) no pior caso
		for i := index; i < list.size-1; i++ {
			list.elements[i] = list.elements[i+1]
		}
		list.size--
		return nil
	} else {
		return fmt.Errorf("index inválido: %d", index)
	}
}

// RemoveValue remove a primeira ocorrência do valor especificado
// Complexidade: O(n)
func (list *ArrayList) RemoveValue(value int) bool {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == value {
			list.Remove(i)
			return true
		}
	}
	return false
}

// Clear remove todos os elementos da lista
// Complexidade: Θ(1)
func (list *ArrayList) Clear() {
	list.size = 0
}

// Contains verifica se a lista contém o valor especificado
// Complexidade: O(n)
func (list *ArrayList) Contains(value int) bool {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == value {
			return true
		}
	}
	return false
}

// IndexOf retorna o índice da primeira ocorrência do valor
// Complexidade: O(n)
func (list *ArrayList) IndexOf(value int) int {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == value {
			return i
		}
	}
	return -1
}

// ToSlice retorna uma cópia dos elementos como slice
// Complexidade: Θ(n)
func (list *ArrayList) ToSlice() []int {
	result := make([]int, list.size)
	for i := 0; i < list.size; i++ {
		result[i] = list.elements[i]
	}
	copy(result, list.elements[:list.size])
	return result
}

// String retorna uma representação em string da lista
// Complexidade: O(n)
func (list *ArrayList) String() string {
	if list.size == 0 {
		return "[]"
	}
	
	result := "["
	for i := 0; i < list.size; i++ {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d", list.elements[i])
	}
	result += "]"
	return result
}

// TrimToSize reduz a capacidade para o tamanho atual
// Complexidade: Θ(n)
func (list *ArrayList) TrimToSize() {
	if list.size < len(list.elements) {
		list.resize(list.size)
	}
}

// EnsureCapacity garante que a lista tenha pelo menos a capacidade especificada
// Complexidade: O(n) se precisar redimensionar, O(1) caso contrário
func (list *ArrayList) EnsureCapacity(minCapacity int) {
	if minCapacity > len(list.elements) {
		newCapacity := len(list.elements)
		for newCapacity < minCapacity {
			newCapacity *= 2
		}
		list.resize(newCapacity)
	}
}
