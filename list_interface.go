package main

import "fmt"

// ============================================================================
// INTERFACE LIST - TIPO ABSTRATO DE DADOS
// ============================================================================

// List define o contrato que todas as implementações de lista devem seguir
// Esta é uma abstração que permite polimorfismo - diferentes implementações
// podem ser usadas de forma intercambiável
type List interface {
	// Operações de consulta
	Size() int                    // Retorna o número de elementos na lista
	IsEmpty() bool               // Verifica se a lista está vazia
	Get(index int) (int, error)  // Obtém elemento em uma posição específica
	
	// Operações de modificação
	Add(element int)                          // Adiciona elemento no final da lista
	AddOnIndex(element int, index int) error  // Adiciona elemento em posição específica
	Remove(index int) error                   // Remove elemento de posição específica
	Clear()                                   // Remove todos os elementos
	
	// Operações de busca
	Contains(element int) bool // Verifica se contém elemento
	IndexOf(element int) int   // Encontra posição do elemento
	
	// Operações de conversão
	ToSlice() []int // Converte para slice
	String() string // Representação em string
}

// ============================================================================
// FUNÇÕES UTILITÁRIAS QUE TRABALHAM COM A INTERFACE
// ============================================================================

// PrintList imprime uma lista usando a interface
func PrintList(list List, name string) {
	fmt.Printf("%s: %s (tamanho: %d)\n", name, list.String(), list.Size())
}

// CopyList copia elementos de uma lista para outra
func CopyList(source List, destination List) {
	destination.Clear()
	for i := 0; i < source.Size(); i++ {
		value, _ := source.Get(i)
		destination.Add(value)
	}
}

// ReverseList inverte os elementos de uma lista usando apenas a interface
func ReverseList(list List) {
	size := list.Size()
	for i := 0; i < size/2; i++ {
		// Trocar elementos nas posições i e size-1-i
		left, _ := list.Get(i)
		right, _ := list.Get(size - 1 - i)
		
		list.Set(i, right)
		list.Set(size-1-i, left)
	}
}

// FindMax encontra o maior elemento na lista
func FindMax(list List) (int, error) {
	if list.IsEmpty() {
		return 0, fmt.Errorf("lista vazia")
	}
	
	max, _ := list.Get(0)
	for i := 1; i < list.Size(); i++ {
		value, _ := list.Get(i)
		if value > max {
			max = value
		}
	}
	return max, nil
}

// FindMin encontra o menor elemento na lista
func FindMin(list List) (int, error) {
	if list.IsEmpty() {
		return 0, fmt.Errorf("lista vazia")
	}
	
	min, _ := list.Get(0)
	for i := 1; i < list.Size(); i++ {
		value, _ := list.Get(i)
		if value < min {
			min = value
		}
	}
	return min, nil
}

// Sum calcula a soma de todos os elementos
func Sum(list List) int {
	sum := 0
	for i := 0; i < list.Size(); i++ {
		value, _ := list.Get(i)
		sum += value
	}
	return sum
}

// Average calcula a média dos elementos
func Average(list List) float64 {
	if list.IsEmpty() {
		return 0.0
	}
	return float64(Sum(list)) / float64(list.Size())
}

// RemoveAll remove todas as ocorrências de um valor
func RemoveAll(list List, value int) int {
	removed := 0
	i := 0
	
	for i < list.Size() {
		current, _ := list.Get(i)
		if current == value {
			list.Remove(i)
			removed++
			// Não incrementa i porque os elementos se deslocaram
		} else {
			i++
		}
	}
	
	return removed
}

// IsSorted verifica se a lista está ordenada em ordem crescente
func IsSorted(list List) bool {
	for i := 1; i < list.Size(); i++ {
		current, _ := list.Get(i)
		previous, _ := list.Get(i - 1)
		if current < previous {
			return false
		}
	}
	return true
}

// BubbleSort ordena a lista usando bubble sort
func BubbleSort(list List) {
	n := list.Size()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			current, _ := list.Get(j)
			next, _ := list.Get(j + 1)
			
			if current > next {
				// Trocar elementos
				list.Set(j, next)
				list.Set(j+1, current)
			}
		}
	}
}

// BinarySearch realiza busca binária em uma lista ordenada
// Retorna o índice do elemento ou -1 se não encontrado
func BinarySearch(list List, target int) int {
	left, right := 0, list.Size()-1
	
	for left <= right {
		mid := (left + right) / 2
		midValue, _ := list.Get(mid)
		
		if midValue == target {
			return mid
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

// MergeSorted mescla duas listas ordenadas em uma nova lista ordenada
func MergeSorted(list1, list2 List, result List) {
	result.Clear()
	i, j := 0, 0
	
	// Mesclar enquanto ambas as listas têm elementos
	for i < list1.Size() && j < list2.Size() {
		val1, _ := list1.Get(i)
		val2, _ := list2.Get(j)
		
		if val1 <= val2 {
			result.Add(val1)
			i++
		} else {
			result.Add(val2)
			j++
		}
	}
	
	// Adicionar elementos restantes da lista1
	for i < list1.Size() {
		val, _ := list1.Get(i)
		result.Add(val)
		i++
	}
	
	// Adicionar elementos restantes da lista2
	for j < list2.Size() {
		val, _ := list2.Get(j)
		result.Add(val)
		j++
	}
}

// ============================================================================
// EXTENSÕES PARA ARRAYLIST E LINKEDLIST IMPLEMENTAREM Set()
// ============================================================================

// Set define o valor do elemento na posição especificada (ArrayList)
func (list *ArrayList) Set(index int, value int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = value
		return nil
	}
	return fmt.Errorf("índice inválido: %d", index)
}

// Set define o valor do elemento na posição especificada (LinkedList)
func (list *LinkedList) Set(index int, value int) error {
	if index >= 0 && index < list.inserted {
		current := list.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		current.val = value
		return nil
	}
	return fmt.Errorf("índice inválido: %d", index)
}