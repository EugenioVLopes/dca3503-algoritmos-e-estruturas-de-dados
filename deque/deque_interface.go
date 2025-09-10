package deque

// ============================================================================
// DEQUE INTERFACE - DEFINIÇÃO DA INTERFACE PARA DEQUES
// ============================================================================

// IDeque define a interface para implementações de deque (double-ended queue)
// Um deque permite inserção e remoção eficiente em ambas extremidades
type IDeque interface {
	// Operações de inserção
	EnqueueFront(value int)  // Adiciona elemento no início
	EnqueueRear(value int)   // Adiciona elemento no final
	
	// Operações de remoção
	DequeueFront() (int, error)  // Remove e retorna elemento do início
	DequeueRear() (int, error)   // Remove e retorna elemento do final
	
	// Operações de acesso
	Front() (int, error)     // Retorna elemento do início sem remover
	Rear() (int, error)      // Retorna elemento do final sem remover
	
	// Operações de estado
	IsEmpty() bool           // Verifica se o deque está vazio
	Size() int               // Retorna o número de elementos
	
	// Operações auxiliares
	Clear()                  // Remove todos os elementos
	ToSlice() []int          // Converte para slice
	String() string          // Representação em string
}

// ============================================================================
// OPERAÇÕES COMUNS PARA DEQUES
// ============================================================================

// DequeOperations fornece operações comuns que podem ser implementadas
// por diferentes tipos de deque
type DequeOperations interface {
	// Operações de busca
	Contains(element int) bool
	IndexOf(element int) int
	
	// Operações funcionais
	Filter(predicate func(int) bool) IDeque
	Map(mapper func(int) int) IDeque
	Reduce(reducer func(int, int) int, initialValue int) int
	ForEach(action func(int, int))
	
	// Operações de transformação
	Reverse()
	Rotate(k int)
}

// DequeComparison fornece operações de comparação e clonagem
// Nota: Clone e Equals são específicos por tipo para melhor type safety
type DequeComparison interface {
	// Operações de cópia compatíveis com interface
	CloneIDeque() IDeque
	EqualsIDeque(other IDeque) bool
}

// ============================================================================
// TIPOS DE DEQUE DISPONÍVEIS
// ============================================================================

// DequeType representa os diferentes tipos de implementação de deque
type DequeType int

const (
	// ArrayDequeType representa deque baseado em array circular
	ArrayDequeType DequeType = iota
	
	// LinkedListDequeType representa deque baseado em lista ligada
	LinkedListDequeType
	
	// DoublyLinkedDequeType representa deque baseado em lista duplamente ligada
	DoublyLinkedDequeType
)

// String retorna a representação em string do tipo de deque
func (dt DequeType) String() string {
	switch dt {
	case ArrayDequeType:
		return "ArrayDeque"
	case LinkedListDequeType:
		return "LinkedListDeque"
	case DoublyLinkedDequeType:
		return "DoublyLinkedDeque"
	default:
		return "Unknown"
	}
}

// GetFactory retorna uma instância do factory
func GetFactory() *DequeFactory {
	return &DequeFactory{}
}

// ============================================================================
// FUNÇÕES UTILITÁRIAS
// ============================================================================

// CompareDeques compara dois deques e retorna true se são iguais
func CompareDeques(d1, d2 IDeque) bool {
	if d1.Size() != d2.Size() {
		return false
	}
	
	slice1 := d1.ToSlice()
	slice2 := d2.ToSlice()
	
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	
	return true
}

// MergeDeques combina dois deques em um novo deque
func MergeDeques(d1, d2 IDeque, dequeType DequeType) IDeque {
	factory := GetFactory()
	result := factory.NewDeque(dequeType)
	
	// Adiciona elementos do primeiro deque
	slice1 := d1.ToSlice()
	for _, element := range slice1 {
		result.EnqueueRear(element)
	}
	
	// Adiciona elementos do segundo deque
	slice2 := d2.ToSlice()
	for _, element := range slice2 {
		result.EnqueueRear(element)
	}
	
	return result
}

// ReverseDeque cria um novo deque com elementos em ordem reversa
func ReverseDeque(d IDeque, dequeType DequeType) IDeque {
	factory := GetFactory()
	result := factory.NewDeque(dequeType)
	
	slice := d.ToSlice()
	for i := len(slice) - 1; i >= 0; i-- {
		result.EnqueueRear(slice[i])
	}
	
	return result
}

// ============================================================================
// FACTORY PARA CRIAÇÃO DE DEQUES
// ============================================================================

// DequeFactory fornece métodos para criar diferentes tipos de deque
type DequeFactory struct{}

// NewDeque cria um novo deque do tipo especificado
func (f *DequeFactory) NewDeque(dequeType DequeType, capacity ...int) IDeque {
	switch dequeType {
	case ArrayDequeType:
		if len(capacity) > 0 {
			return NewArrayDeque(capacity[0])
		}
		return NewArrayDeque(10) // Capacidade padrão
		
	case LinkedListDequeType:
		return NewLinkedListDeque()
		
	case DoublyLinkedDequeType:
		return NewDeque() // Usa NewDeque() para deque duplamente ligado
		
	default:
		return NewLinkedListDeque() // Padrão
	}
}
