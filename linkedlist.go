package main

import (
	"fmt"
	"errors"
)

// ============================================================================
// LINKEDLIST - IMPLEMENTAÇÃO BASEADA EM LISTA LIGADA
// ============================================================================

// Node representa um nó na lista ligada
// Cada nó contém um valor e um ponteiro para o próximo nó
type Node struct {
	val  int   // Valor armazenado no nó
	next *Node // Ponteiro para o próximo nó (nil se for o último)
}

// NewNode cria um novo nó com o valor especificado
func NewNode(value int) *Node {
	return &Node{
		val:  value,
		next: nil,
	}
}

// LinkedList implementa uma lista usando nós ligados
// Características:
// - Inserção/remoção no início é rápida O(1)
// - Acesso sequencial O(n)
// - Uso dinâmico de memória (aloca conforme necessário)
// - Não há desperdício de memória
type LinkedList struct {
	head     *Node // Ponteiro para o primeiro nó
	inserted int   // Contador de elementos (para Size() em O(1))
}

// NewLinkedList cria uma nova instância de LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:     nil,
		inserted: 0,
	}
}

// Size retorna o número de elementos na lista
// Complexidade: Θ(1) - Mantemos um contador
func (list *LinkedList) Size() int { // Θ(1)
	return list.inserted
}

// IsEmpty verifica se a lista está vazia
// Complexidade: Θ(1)
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

// Get obtém elemento na posição especificada
// Complexidade: O(n) pior caso, Ω(1) melhor caso
// Pseudocódigo:
// 1. Validar índice
// 2. Percorrer lista do início até a posição desejada
// 3. Retornar valor do nó encontrado
func (list *LinkedList) Get(index int) (int, error) { // O(n), Ω(1)
	if index >= 0 && index < list.inserted {
		aux := list.head
		// Percorre a lista até a posição desejada
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.val, nil
	} else {
		return -1, errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// Set define o valor do elemento na posição especificada
// Complexidade: O(n)
func (list *LinkedList) Set(index int, value int) error {
	if index >= 0 && index < list.inserted {
		aux := list.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		aux.val = value
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// AddFirst adiciona elemento no início da lista
// Complexidade: Θ(1)
// Pseudocódigo:
// 1. Criar novo nó
// 2. Conectar novo nó ao head atual
// 3. Atualizar head para o novo nó
// 4. Incrementar contador
func (list *LinkedList) AddFirst(val int) {
	newNode := NewNode(val)
	newNode.next = list.head
	list.head = newNode
	list.inserted++
}

// Add adiciona elemento no final da lista
// Complexidade: O(n) - precisa percorrer até o final
// Pseudocódigo:
// 1. Criar novo nó com o valor
// 2. Se lista vazia: novo nó vira head
// 3. Senão: percorrer até o último nó e conectar novo nó
// 4. Incrementar contador
func (list *LinkedList) Add(val int) {
	newNode := NewNode(val)
	
	if list.head == nil {
		// Lista vazia - novo nó vira o primeiro
		list.head = newNode
	} else {
		// Percorre até o último nó
		aux := list.head
		for aux.next != nil {
			aux = aux.next
		}
		// Conecta novo nó ao final
		aux.next = newNode
	}
	list.inserted++
}

// AddOnIndex adiciona elemento em posição específica
// Complexidade: O(n) - pode precisar percorrer até a posição
// Pseudocódigo:
// 1. Validar índice
// 2. Se índice 0: inserir no início
// 3. Senão: percorrer até posição anterior e inserir
// 4. Incrementar contador
func (list *LinkedList) AddOnIndex(val int, index int) error {
	if index >= 0 && index <= list.inserted {
		if index == 0 {
			// Inserção no início - O(1)
			list.AddFirst(val)
			return nil
		}
		
		newNode := NewNode(val)
		
		// Percorre até a posição anterior
		aux := list.head
		for i := 0; i < index-1; i++ {
			aux = aux.next
		}
		// Insere novo nó entre aux e aux.next
		newNode.next = aux.next
		aux.next = newNode
		list.inserted++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// RemoveFirst remove o primeiro elemento da lista
// Complexidade: Θ(1)
func (list *LinkedList) RemoveFirst() (int, error) {
	if list.head == nil {
		return 0, errors.New("Lista vazia")
	}
	
	removedValue := list.head.val
	list.head = list.head.next
	list.inserted--
	return removedValue, nil
}

// Remove remove elemento de posição específica
// Complexidade: O(n) - pode precisar percorrer até a posição
// Pseudocódigo:
// 1. Validar índice
// 2. Se índice 0: remover primeiro nó
// 3. Senão: percorrer até posição anterior e reconectar ponteiros
// 4. Decrementar contador
func (list *LinkedList) Remove(index int) error {
	if index >= 0 && index < list.inserted {
		if index == 0 {
			// Remoção do primeiro nó - O(1)
			_, err := list.RemoveFirst()
			return err
		} else {
			// Percorre até o nó anterior ao que será removido
			aux := list.head
			for i := 0; i < index-1; i++ {
				aux = aux.next
			}
			// Remove o nó reconectando os ponteiros
			aux.next = aux.next.next
			list.inserted--
			return nil
		}
	} else {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
}

// RemoveValue remove a primeira ocorrência do valor especificado
// Complexidade: O(n)
func (list *LinkedList) RemoveValue(value int) bool {
	if list.head == nil {
		return false
	}
	
	// Caso especial: remover o primeiro nó
	if list.head.val == value {
		list.RemoveFirst()
		return true
	}
	
	// Procurar o valor nos nós seguintes
	current := list.head
	for current.next != nil {
		if current.next.val == value {
			current.next = current.next.next
			list.inserted--
			return true
		}
		current = current.next
	}
	
	return false
}

// Clear remove todos os elementos da lista
// Complexidade: Θ(1)
func (list *LinkedList) Clear() {
	list.head = nil
	list.inserted = 0
}

// Contains verifica se a lista contém o valor especificado
// Complexidade: O(n)
func (list *LinkedList) Contains(value int) bool {
	current := list.head
	for current != nil {
		if current.val == value {
			return true
		}
		current = current.next
	}
	return false
}

// IndexOf retorna o índice da primeira ocorrência do valor
// Complexidade: O(n)
func (list *LinkedList) IndexOf(value int) int {
	current := list.head
	index := 0
	
	for current != nil {
		if current.val == value {
			return index
		}
		current = current.next
		index++
	}
	
	return -1
}

// ToSlice retorna uma cópia dos elementos como slice
// Complexidade: Θ(n)
func (list *LinkedList) ToSlice() []int {
	result := make([]int, 0, list.inserted)
	current := list.head
	
	for current != nil {
		result = append(result, current.val)
		current = current.next
	}
	
	return result
}

// String retorna uma representação em string da lista
// Complexidade: O(n)
func (list *LinkedList) String() string {
	if list.head == nil {
		return "[]"
	}
	
	result := "["
	current := list.head
	first := true
	
	for current != nil {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%d", current.val)
		current = current.next
		first = false
	}
	
	result += "]"
	return result
}

// Reverse inverte a ordem dos elementos na lista
// Complexidade: O(n)
func (list *LinkedList) Reverse() {
	if list.head == nil || list.head.next == nil {
		return // Lista vazia ou com um elemento
	}
	
	var prev *Node = nil
	current := list.head
	
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	
	list.head = prev
}

// GetMiddle retorna o elemento do meio da lista
// Complexidade: O(n)
// Usa algoritmo "tortoise and hare" (Floyd's algorithm)
func (list *LinkedList) GetMiddle() (int, error) {
	if list.head == nil {
		return 0, errors.New("Lista vazia")
	}
	
	slow := list.head
	fast := list.head
	
	// Fast avança 2 posições, slow avança 1
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	
	return slow.val, nil
}

// HasCycle detecta se há um ciclo na lista
// Complexidade: O(n)
// Usa algoritmo de Floyd (tortoise and hare)
func (list *LinkedList) HasCycle() bool {
	if list.head == nil {
		return false
	}
	
	slow := list.head
	fast := list.head
	
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		
		if slow == fast {
			return true
		}
	}
	
	return false
}

// RemoveDuplicates remove elementos duplicados da lista
// Complexidade: O(n²) - versão simples
func (list *LinkedList) RemoveDuplicates() {
	if list.head == nil {
		return
	}
	
	current := list.head
	
	for current != nil {
		runner := current
		
		// Procura e remove duplicatas do valor atual
		for runner.next != nil {
			if runner.next.val == current.val {
				runner.next = runner.next.next
				list.inserted--
			} else {
				runner = runner.next
			}
		}
		
		current = current.next
	}
}

// AddAll adiciona todos os elementos do slice fornecido no final
// Complexidade: O(n + m) onde n é o tamanho atual e m é o tamanho do slice
func (list *LinkedList) AddAll(elements []int) {
	for _, element := range elements {
		list.Add(element)
	}
}

// AddAllFirst adiciona todos os elementos do slice fornecido no início
// Complexidade: O(m) onde m é o tamanho do slice
func (list *LinkedList) AddAllFirst(elements []int) {
	// Adiciona em ordem reversa para manter a ordem original
	for i := len(elements) - 1; i >= 0; i-- {
		list.AddFirst(elements[i])
	}
}