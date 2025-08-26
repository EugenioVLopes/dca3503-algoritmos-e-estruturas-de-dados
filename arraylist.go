package main

import (
	"fmt"
	"errors"
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
	v        []int // Array interno que armazena os elementos
	inserted int   // Contador de elementos inseridos (tamanho lógico)
}

// NewArrayList cria uma nova instância de ArrayList com capacidade inicial
func NewArrayList(initialCapacity int) *ArrayList {
	return &ArrayList{
		v:        make([]int, initialCapacity),
		inserted: 0,
	}
}

// Init inicializa o ArrayList com capacidade inicial
// Pseudocódigo:
// 1. Criar array interno com tamanho especificado
// 2. Inicializar contador de elementos como 0
func (l *ArrayList) Init(size int) {
	l.v = make([]int, size) // Aloca memória para 'size' elementos
	l.inserted = 0          // Inicialmente não há elementos inseridos
}

// Size retorna o número de elementos atualmente na lista
// Complexidade: Θ(1) - Tempo constante
// Pseudocódigo:
// 1. Retornar valor do contador de elementos inseridos
func (list *ArrayList) Size() int { // Θ(1)
	return list.inserted
}

// IsEmpty verifica se a lista está vazia
// Complexidade: Θ(1)
func (list *ArrayList) IsEmpty() bool {
	return list.inserted == 0
}

// Capacity retorna a capacidade atual do array interno
// Complexidade: Θ(1)
func (list *ArrayList) Capacity() int {
	return len(list.v)
}

// Get obtém o elemento na posição especificada
// Complexidade: Θ(1) - Acesso direto por índice
// Pseudocódigo:
// 1. Verificar se índice é válido (0 <= index < inserted)
// 2. Se válido: retornar elemento na posição
// 3. Se inválido: retornar erro
func (list *ArrayList) Get(index int) (int, error) { // Θ(1)
	// Validação de bounds - evita acesso a memória inválida
	if index >= 0 && index < list.inserted {
		return list.v[index], nil // Acesso direto O(1)
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// Set define o valor do elemento na posição especificada
// Complexidade: Θ(1)
func (list *ArrayList) Set(index int, value int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = value
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// doubleV dobra a capacidade do array interno quando necessário
// Complexidade: Θ(n) - Precisa copiar todos os elementos
// Pseudocódigo:
// 1. Criar novo array com dobro do tamanho atual
// 2. Copiar todos os elementos do array antigo para o novo
// 3. Substituir array interno pelo novo
func (list *ArrayList) doubleV() { // Θ(n)
	// Estratégia de crescimento: dobrar o tamanho
	// Isso garante O(1) amortizado para inserções
	newV := make([]int, len(list.v)*2)
	
	// Copia elementos do array antigo para o novo
	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
	}
	
	list.v = newV // Substitui o array interno
}

// resize redimensiona o array para uma nova capacidade
// Complexidade: Θ(n)
func (list *ArrayList) resize(newCapacity int) {
	if newCapacity < list.inserted {
		return // Não pode reduzir abaixo do tamanho atual
	}
	
	newV := make([]int, newCapacity)
	copy(newV, list.v[:list.inserted])
	list.v = newV
}

// Add adiciona elemento no final da lista
// Complexidade: O(n) pior caso, Ω(1) melhor caso, O(1) amortizado
// Pseudocódigo:
// 1. Se array está cheio: dobrar capacidade
// 2. Inserir elemento na próxima posição disponível
// 3. Incrementar contador de elementos
func (list *ArrayList) Add(val int) { // O(n), Ω(1)
	// Verifica se precisa expandir o array
	if list.inserted == len(list.v) {
		list.doubleV() // O(n) apenas quando necessário
	}
	
	// Inserção no final é sempre O(1)
	list.v[list.inserted] = val
	list.inserted++
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
	// Validação: pode inserir no final (index == inserted)
	if index >= 0 && index <= list.inserted {
		// Expande array se necessário
		if list.inserted == len(list.v) {
			list.doubleV()
		}
		
		// Desloca elementos para a direita - O(n) no pior caso
		for i := list.inserted; i > index; i-- {
			list.v[i] = list.v[i-1]
		}
		
		// Insere o novo elemento
		list.v[index] = val
		list.inserted++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// Remove remove elemento de posição específica
// Complexidade: Ω(1) melhor caso, O(n) pior caso
// Pseudocódigo:
// 1. Validar índice (0 <= index < inserted)
// 2. Deslocar elementos à direita do índice uma posição para esquerda
// 3. Decrementar contador de elementos
func (list *ArrayList) Remove(index int) error { // Ω(1), O(n)
	if index >= 0 && index < list.inserted {
		// Desloca elementos para a esquerda - O(n) no pior caso
		for i := index; i < list.inserted-1; i++ {
			list.v[i] = list.v[i+1]
		}
		list.inserted--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// RemoveValue remove a primeira ocorrência do valor especificado
// Complexidade: O(n)
func (list *ArrayList) RemoveValue(value int) bool {
	for i := 0; i < list.inserted; i++ {
		if list.v[i] == value {
			list.Remove(i)
			return true
		}
	}
	return false
}

// Clear remove todos os elementos da lista
// Complexidade: Θ(1)
func (list *ArrayList) Clear() {
	list.inserted = 0
}

// Contains verifica se a lista contém o valor especificado
// Complexidade: O(n)
func (list *ArrayList) Contains(value int) bool {
	for i := 0; i < list.inserted; i++ {
		if list.v[i] == value {
			return true
		}
	}
	return false
}

// IndexOf retorna o índice da primeira ocorrência do valor
// Complexidade: O(n)
func (list *ArrayList) IndexOf(value int) int {
	for i := 0; i < list.inserted; i++ {
		if list.v[i] == value {
			return i
		}
	}
	return -1
}

// ToSlice retorna uma cópia dos elementos como slice
// Complexidade: Θ(n)
func (list *ArrayList) ToSlice() []int {
	result := make([]int, list.inserted)
	copy(result, list.v[:list.inserted])
	return result
}

// String retorna uma representação em string da lista
// Complexidade: O(n)
func (list *ArrayList) String() string {
	if list.inserted == 0 {
		return "[]"
	}
	
	result := "["
	for i := 0; i < list.inserted; i++ {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d", list.v[i])
	}
	result += "]"
	return result
}

// TrimToSize reduz a capacidade para o tamanho atual
// Complexidade: Θ(n)
func (list *ArrayList) TrimToSize() {
	if list.inserted < len(list.v) {
		list.resize(list.inserted)
	}
}

// EnsureCapacity garante que a lista tenha pelo menos a capacidade especificada
// Complexidade: O(n) se precisar redimensionar, O(1) caso contrário
func (list *ArrayList) EnsureCapacity(minCapacity int) {
	if minCapacity > len(list.v) {
		newCapacity := len(list.v)
		for newCapacity < minCapacity {
			newCapacity *= 2
		}
		list.resize(newCapacity)
	}
}

// AddAll adiciona todos os elementos do slice fornecido
// Complexidade: O(n + m) onde n é o tamanho atual e m é o tamanho do slice
func (list *ArrayList) AddAll(elements []int) {
	requiredCapacity := list.inserted + len(elements)
	if requiredCapacity > len(list.v) {
		list.EnsureCapacity(requiredCapacity)
	}
	
	copy(list.v[list.inserted:], elements)
	list.inserted += len(elements)
}