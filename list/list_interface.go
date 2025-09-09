package list

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
		
		list.Remove(i)
		list.AddOnIndex(right, i)
		
		list.Remove(size - 1 - i)
		list.AddOnIndex(left, size - 1 - i)
	}
}

// SortList ordena os elementos de uma lista usando o algoritmo de seleção
func SortList(list List) {
	size := list.Size()
	for i := 0; i < size-1; i++ {
		minIndex := i
		minValue, _ := list.Get(i)
		
		for j := i+1; j < size; j++ {
			currentValue, _ := list.Get(j)
			if currentValue < minValue {
				minIndex = j
				minValue = currentValue
			}
		}
		
		if minIndex != i {
			// Trocar elementos
			currentValue, _ := list.Get(i)
			
			list.Remove(i)
			list.AddOnIndex(minValue, i)
			
			list.Remove(minIndex)
			list.AddOnIndex(currentValue, minIndex)
		}
	}
}