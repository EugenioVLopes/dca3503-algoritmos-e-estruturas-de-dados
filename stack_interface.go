package main

import "fmt"

// ============================================================================
// INTERFACE STACK - TIPO ABSTRATO DE DADOS
// ============================================================================

// Stack define o contrato que todas as implementações de pilha devem seguir
// Uma pilha é uma estrutura de dados LIFO (Last In, First Out)
// O último elemento inserido é o primeiro a ser removido
type Stack interface {
	// Operações básicas de pilha
	Push(element int)           // Adiciona elemento no topo da pilha
	Pop() (int, error)          // Remove e retorna elemento do topo
	Peek() (int, error)         // Retorna elemento do topo sem remover
	
	// Operações de consulta
	Size() int                  // Retorna número de elementos na pilha
	IsEmpty() bool             // Verifica se a pilha está vazia
	IsFull() bool              // Verifica se a pilha está cheia (para implementações com limite)
	
	// Operações auxiliares
	Clear()                    // Remove todos os elementos
	ToSlice() []int            // Converte para slice (do topo para a base)
	String() string            // Representação em string
}

// ============================================================================
// FUNÇÕES UTILITÁRIAS QUE TRABALHAM COM A INTERFACE
// ============================================================================

// PrintStack imprime uma pilha usando a interface
func PrintStack(stack Stack, name string) {
	fmt.Printf("%s: %s (tamanho: %d)\n", name, stack.String(), stack.Size())
}

// CopyStack copia elementos de uma pilha para outra
// Mantém a ordem original (usa pilha auxiliar)
func CopyStack(source Stack, destination Stack) {
	destination.Clear()
	
	// Pilha auxiliar para manter a ordem
	aux := NewArrayStack(source.Size())
	
	// Move elementos da origem para auxiliar
	for !source.IsEmpty() {
		value, _ := source.Pop()
		aux.Push(value)
	}
	
	// Move elementos da auxiliar para destino e restaura origem
	for !aux.IsEmpty() {
		value, _ := aux.Pop()
		source.Push(value)
		destination.Push(value)
	}
}

// ReverseStack inverte a ordem dos elementos na pilha
func ReverseStack(stack Stack) {
	if stack.IsEmpty() {
		return
	}
	
	// Remove elemento do topo
	value, _ := stack.Pop()
	
	// Recursivamente inverte o resto
	ReverseStack(stack)
	
	// Insere elemento no fundo da pilha
	insertAtBottom(stack, value)
}

// insertAtBottom insere elemento no fundo da pilha
func insertAtBottom(stack Stack, value int) {
	if stack.IsEmpty() {
		stack.Push(value)
		return
	}
	
	// Remove elemento do topo
	top, _ := stack.Pop()
	
	// Recursivamente insere no fundo
	insertAtBottom(stack, value)
	
	// Restaura elemento do topo
	stack.Push(top)
}

// FindInStack procura um elemento na pilha sem alterar sua estrutura
func FindInStack(stack Stack, target int) bool {
	if stack.IsEmpty() {
		return false
	}
	
	// Remove elemento do topo
	value, _ := stack.Pop()
	
	// Verifica se é o elemento procurado
	if value == target {
		stack.Push(value) // Restaura
		return true
	}
	
	// Procura recursivamente no resto
	found := FindInStack(stack, target)
	
	// Restaura elemento
	stack.Push(value)
	
	return found
}

// StackSum calcula a soma de todos os elementos na pilha
func StackSum(stack Stack) int {
	if stack.IsEmpty() {
		return 0
	}
	
	// Remove elemento do topo
	value, _ := stack.Pop()
	
	// Soma recursivamente
	sum := value + StackSum(stack)
	
	// Restaura elemento
	stack.Push(value)
	
	return sum
}

// StackMax encontra o maior elemento na pilha
func StackMax(stack Stack) (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("pilha vazia")
	}
	
	// Caso base: apenas um elemento
	value, _ := stack.Pop()
	if stack.IsEmpty() {
		stack.Push(value)
		return value, nil
	}
	
	// Encontra máximo no resto
	maxRest, _ := StackMax(stack)
	
	// Restaura elemento
	stack.Push(value)
	
	// Retorna o maior
	if value > maxRest {
		return value, nil
	}
	return maxRest, nil
}

// IsValidParentheses verifica se uma sequência de parênteses está balanceada
// Exemplo de aplicação prática de pilhas
func IsValidParentheses(s string) bool {
	stack := NewArrayStack(len(s))
	
	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(int(char))
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}
			top, _ := stack.Pop()
			if !isMatchingPair(rune(top), char) {
				return false
			}
		}
	}
	
	return stack.IsEmpty()
}

// isMatchingPair verifica se dois caracteres formam um par válido
func isMatchingPair(open, close rune) bool {
	return (open == '(' && close == ')') ||
		   (open == '[' && close == ']') ||
		   (open == '{' && close == '}')
}

// EvaluatePostfix avalia uma expressão em notação pós-fixa
// Exemplo: "3 4 + 2 *" = (3 + 4) * 2 = 14
func EvaluatePostfix(expression []string) (int, error) {
	stack := NewArrayStack(len(expression))
	
	for _, token := range expression {
		switch token {
		case "+":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("expressão inválida")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a + b)
		case "-":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("expressão inválida")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a - b)
		case "*":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("expressão inválida")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a * b)
		case "/":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("expressão inválida")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			if b == 0 {
				return 0, fmt.Errorf("divisão por zero")
			}
			stack.Push(a / b)
		default:
			// Assume que é um número
			var num int
			_, err := fmt.Sscanf(token, "%d", &num)
			if err != nil {
				return 0, fmt.Errorf("token inválido: %s", token)
			}
			stack.Push(num)
		}
	}
	
	if stack.Size() != 1 {
		return 0, fmt.Errorf("expressão inválida")
	}
	
	result, _ := stack.Pop()
	return result, nil
}