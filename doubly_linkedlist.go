package main

import (
	"fmt"
	"errors"
)

// ============================================================================
// DOUBLY LINKEDLIST - IMPLEMENTAÇÃO BASEADA EM LISTA DUPLAMENTE LIGADA
// ============================================================================

// DoublyNode representa um nó na lista duplamente ligada
// Cada nó contém um valor e ponteiros para o próximo e anterior nó
type DoublyNode struct {
	data int          // Valor armazenado no nó
	next *DoublyNode  // Ponteiro para o próximo nó
	prev *DoublyNode  // Ponteiro para o nó anterior
}

// NewDoublyNode cria um novo nó com o valor especificado
func NewDoublyNode(value int) *DoublyNode {
	return &DoublyNode{
		data: value,
		next: nil,
		prev: nil,
	}
}

// DoublyLinkedList implementa uma lista usando nós duplamente ligados
// Características:
// - Navegação bidirecional O(1)
// - Inserção/remoção em ambas extremidades O(1)
// - Remoção por referência O(1)
// - Maior uso de memória (ponteiro extra por nó)
type DoublyLinkedList struct {
	head *DoublyNode // Ponteiro para o primeiro nó
	tail *DoublyNode // Ponteiro para o último nó
	size int         // Contador de elementos
}

// NewDoublyLinkedList cria uma nova instância de DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Size retorna o número de elementos na lista
// Complexidade: Θ(1)
func (list *DoublyLinkedList) Size() int {
	return list.size
}

// IsEmpty verifica se a lista está vazia
// Complexidade: Θ(1)
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

// AddFirst adiciona elemento no início da lista
// Complexidade: Θ(1)
func (list *DoublyLinkedList) AddFirst(element int) {
	newNode := NewDoublyNode(element)
	
	if list.head == nil {
		// Lista vazia
		list.head = newNode
		list.tail = newNode
	} else {
		// Conectar novo nó
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
	
	list.size++
}

// AddLast adiciona elemento no final da lista
// Complexidade: Θ(1) - Vantagem sobre LinkedList simples!
func (list *DoublyLinkedList) AddLast(element int) {
	newNode := NewDoublyNode(element)
	
	if list.tail == nil {
		// Lista vazia
		list.head = newNode
		list.tail = newNode
	} else {
		// Conectar novo nó
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	
	list.size++
}

// Add é um alias para AddLast para compatibilidade com interface List
func (list *DoublyLinkedList) Add(element int) {
	list.AddLast(element)
}

// Get obtém elemento na posição especificada
// Complexidade: O(n/2) - Otimizado para escolher direção mais próxima
func (list *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.size {
		return 0, errors.New(fmt.Sprintf("Índice inválido: %d", index))
	}
	
	var current *DoublyNode
	
	// Otimização: escolher direção mais próxima
	if index < list.size/2 {
		// Percorrer do início
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		// Percorrer do final
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}
	
	return current.data, nil
}

// Set define o valor do elemento na posição especificada
// Complexidade: O(n/2)
func (list *DoublyLinkedList) Set(index int, value int) error {
	if index < 0 || index >= list.size {
		return errors.New(fmt.Sprintf("Índice inválido: %d", index))
	}
	
	var current *DoublyNode
	
	// Otimização: escolher direção mais próxima
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}
	
	current.data = value
	return nil
}

// GetNode obtém referência ao nó na posição especificada
// Útil para operações que precisam da referência do nó
// Complexidade: O(n/2)
func (list *DoublyLinkedList) GetNode(index int) (*DoublyNode, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New(fmt.Sprintf("Índice inválido: %d", index))
	}
	
	var current *DoublyNode
	
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}
	
	return current, nil
}

// AddOnIndex adiciona elemento em posição específica
// Complexidade: O(n/2) - Otimizado
func (list *DoublyLinkedList) AddOnIndex(element int, index int) error {
	if index < 0 || index > list.size {
		return errors.New(fmt.Sprintf("Índice inválido: %d", index))
	}
	
	if index == 0 {
		list.AddFirst(element)
		return nil
	}
	
	if index == list.size {
		list.AddLast(element)
		return nil
	}
	
	// Encontrar posição
	var current *DoublyNode
	if index < list.size/2 {
		current = list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = list.tail
		for i := list.size - 1; i > index; i-- {
			current = current.prev
		}
	}
	
	// Inserir antes de current
	newNode := NewDoublyNode(element)
	newNode.next = current
	newNode.prev = current.prev
	current.prev.next = newNode
	current.prev = newNode
	
	list.size++
	return nil
}

// RemoveNode remove nó específico da lista
// Complexidade: Θ(1) - GRANDE VANTAGEM da Doubly LinkedList!
func (list *DoublyLinkedList) RemoveNode(node *DoublyNode) (int, error) {
	if node == nil {
		return 0, errors.New("Nó inválido")
	}
	
	removedData := node.data
	
	// Reconectar nó anterior
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next // Removendo primeiro nó
	}
	
	// Reconectar nó posterior
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev // Removendo último nó
	}
	
	list.size--
	return removedData, nil
}

// RemoveFirst remove o primeiro elemento
// Complexidade: Θ(1)
func (list *DoublyLinkedList) RemoveFirst() (int, error) {
	if list.head == nil {
		return 0, errors.New("Lista vazia")
	}
	
	return list.RemoveNode(list.head)
}

// RemoveLast remove o último elemento
// Complexidade: Θ(1) - Vantagem sobre LinkedList simples!
func (list *DoublyLinkedList) RemoveLast() (int, error) {
	if list.tail == nil {
		return 0, errors.New("Lista vazia")
	}
	
	return list.RemoveNode(list.tail)
}

// Remove remove elemento de posição específica
// Complexidade: O(n/2)
func (list *DoublyLinkedList) Remove(index int) error {
	if index < 0 || index >= list.size {
		return errors.New(fmt.Sprintf("Índice inválido: %d", index))
	}
	
	node, err := list.GetNode(index)
	if err != nil {
		return err
	}
	
	_, err = list.RemoveNode(node)
	return err
}

// RemoveValue remove a primeira ocorrência do valor especificado
// Complexidade: O(n)
func (list *DoublyLinkedList) RemoveValue(value int) bool {
	current := list.head
	
	for current != nil {
		if current.data == value {
			list.RemoveNode(current)
			return true
		}
		current = current.next
	}
	
	return false
}

// Clear remove todos os elementos da lista
// Complexidade: Θ(1)
func (list *DoublyLinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Contains verifica se a lista contém o valor especificado
// Complexidade: O(n/2) - Busca bidirecional otimizada
func (list *DoublyLinkedList) Contains(value int) bool {
	return list.FindNode(value) != nil
}

// IndexOf retorna o índice da primeira ocorrência do valor
// Complexidade: O(n)
func (list *DoublyLinkedList) IndexOf(value int) int {
	current := list.head
	index := 0
	
	for current != nil {
		if current.data == value {
			return index
		}
		current = current.next
		index++
	}
	
	return -1
}

// FindNode encontra o primeiro nó com o valor especificado
// Complexidade: O(n/2) - Busca bidirecional
func (list *DoublyLinkedList) FindNode(value int) *DoublyNode {
	if list.size == 0 {
		return nil
	}
	
	// Busca bidirecional para otimizar
	front := list.head
	back := list.tail
	
	for front != nil && back != nil && front.prev != back {
		if front.data == value {
			return front
		}
		if back.data == value {
			return back
		}
		
		if front == back {
			break
		}
		
		front = front.next
		back = back.prev
	}
	
	// Verificar nó do meio se necessário
	if front != nil && front.data == value {
		return front
	}
	
	return nil
}

// ToSlice retorna uma cópia dos elementos como slice
// Complexidade: Θ(n)
func (list *DoublyLinkedList) ToSlice() []int {
	result := make([]int, 0, list.size)
	current := list.head
	
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	
	return result
}

// ToSliceReverse retorna uma cópia dos elementos em ordem reversa
// Complexidade: Θ(n) - Vantagem da navegação bidirecional
func (list *DoublyLinkedList) ToSliceReverse() []int {
	result := make([]int, 0, list.size)
	current := list.tail
	
	for current != nil {
		result = append(result, current.data)
		current = current.prev
	}
	
	return result
}

// String retorna uma representação em string da lista
// Complexidade: O(n)
func (list *DoublyLinkedList) String() string {
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
		result += fmt.Sprintf("%d", current.data)
		current = current.next
		first = false
	}
	
	result += "]"
	return result
}

// StringReverse retorna representação em string em ordem reversa
// Complexidade: O(n)
func (list *DoublyLinkedList) StringReverse() string {
	if list.tail == nil {
		return "[]"
	}
	
	result := "["
	current := list.tail
	first := true
	
	for current != nil {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%d", current.data)
		current = current.prev
		first = false
	}
	
	result += "]"
	return result
}

// ============================================================================
// ALGORITMOS ESPECIAIS
// ============================================================================

// Reverse inverte a ordem dos elementos na lista
// Complexidade: O(n) - Mais simples que na LinkedList simples
func (list *DoublyLinkedList) Reverse() {
	if list.head == nil {
		return
	}
	
	current := list.head
	
	// Trocar ponteiros prev e next de cada nó
	for current != nil {
		next := current.next
		current.next = current.prev
		current.prev = next
		current = next
	}
	
	// Trocar head e tail
	list.head, list.tail = list.tail, list.head
}

// GetMiddle retorna o elemento do meio da lista
// Complexidade: O(n/2)
func (list *DoublyLinkedList) GetMiddle() (int, error) {
	if list.head == nil {
		return 0, errors.New("Lista vazia")
	}
	
	// Usar navegação bidirecional para encontrar o meio
	front := list.head
	back := list.tail
	
	for front != back && front.next != back {
		front = front.next
		back = back.prev
	}
	
	return front.data, nil
}

// IsPalindrome verifica se a lista é um palíndromo
// Complexidade: O(n/2) - Vantagem da navegação bidirecional
func (list *DoublyLinkedList) IsPalindrome() bool {
	if list.size <= 1 {
		return true
	}
	
	front := list.head
	back := list.tail
	
	for front != back && front.next != back {
		if front.data != back.data {
			return false
		}
		front = front.next
		back = back.prev
	}
	
	return true
}

// RemoveDuplicates remove elementos duplicados da lista
// Complexidade: O(n²) - versão simples
func (list *DoublyLinkedList) RemoveDuplicates() {
	if list.head == nil {
		return
	}
	
	current := list.head
	
	for current != nil {
		runner := current.next
		
		// Procura e remove duplicatas do valor atual
		for runner != nil {
			next := runner.next
			if runner.data == current.data {
				list.RemoveNode(runner)
			}
			runner = next
		}
		
		current = current.next
	}
}

// RotateLeft rotaciona a lista n posições para a esquerda
// Complexidade: O(n)
func (list *DoublyLinkedList) RotateLeft(positions int) {
	if list.size <= 1 || positions <= 0 {
		return
	}
	
	positions = positions % list.size
	if positions == 0 {
		return
	}
	
	// Encontrar novo head
	newHead := list.head
	for i := 0; i < positions; i++ {
		newHead = newHead.next
	}
	
	// Reconectar
	newTail := newHead.prev
	newTail.next = nil
	newHead.prev = nil
	
	list.tail.next = list.head
	list.head.prev = list.tail
	
	list.head = newHead
	list.tail = newTail
}

// RotateRight rotaciona a lista n posições para a direita
// Complexidade: O(n)
func (list *DoublyLinkedList) RotateRight(positions int) {
	if list.size <= 1 || positions <= 0 {
		return
	}
	
	positions = positions % list.size
	if positions == 0 {
		return
	}
	
	// Rotacionar para direita é equivalente a rotacionar para esquerda (size - positions)
	list.RotateLeft(list.size - positions)
}

// AddAll adiciona todos os elementos do slice fornecido no final
// Complexidade: O(m) onde m é o tamanho do slice
func (list *DoublyLinkedList) AddAll(elements []int) {
	for _, element := range elements {
		list.AddLast(element)
	}
}

// AddAllFirst adiciona todos os elementos do slice fornecido no início
// Complexidade: O(m) onde m é o tamanho do slice
func (list *DoublyLinkedList) AddAllFirst(elements []int) {
	// Adiciona em ordem reversa para manter a ordem original
	for i := len(elements) - 1; i >= 0; i-- {
		list.AddFirst(elements[i])
	}
}

// ============================================================================
// ITERADORES
// ============================================================================

// Iterator representa um iterador para DoublyLinkedList
type DoublyIterator struct {
	current *DoublyNode
	list    *DoublyLinkedList
}

// NewIterator cria um novo iterador começando do início
func (list *DoublyLinkedList) NewIterator() *DoublyIterator {
	return &DoublyIterator{
		current: list.head,
		list:    list,
	}
}

// NewReverseIterator cria um novo iterador começando do final
func (list *DoublyLinkedList) NewReverseIterator() *DoublyIterator {
	return &DoublyIterator{
		current: list.tail,
		list:    list,
	}
}

// HasNext verifica se há próximo elemento
func (iter *DoublyIterator) HasNext() bool {
	return iter.current != nil
}

// Next retorna o próximo elemento
func (iter *DoublyIterator) Next() (int, error) {
	if iter.current == nil {
		return 0, errors.New("Não há próximo elemento")
	}
	
	value := iter.current.data
	iter.current = iter.current.next
	return value, nil
}

// HasPrev verifica se há elemento anterior
func (iter *DoublyIterator) HasPrev() bool {
	return iter.current != nil
}

// Prev retorna o elemento anterior
func (iter *DoublyIterator) Prev() (int, error) {
	if iter.current == nil {
		return 0, errors.New("Não há elemento anterior")
	}
	
	value := iter.current.data
	iter.current = iter.current.prev
	return value, nil
}