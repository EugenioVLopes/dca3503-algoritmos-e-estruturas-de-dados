# Implementação de ArrayQueue e LinkedListQueue

## Pergunta

Na linguagem GoLang, use a interface IQueue definida abaixo e programe as seguintes estruturas de dados: ArrayQueue, LinkedListQueue.

```go
type IQueue interface {
    Enqueue(value int)
    Dequeue() (int, error)
    Front() (int, error)
    IsEmpty() bool
    Size() int
}
```

## Resposta

### Interface IQueue

```go
package main

import (
    "errors"
    "fmt"
)

// IQueue define a interface para implementações de fila
type IQueue interface {
    Enqueue(value int)        // Adiciona elemento no final da fila
    Dequeue() (int, error)    // Remove e retorna elemento do início da fila
    Front() (int, error)      // Retorna elemento do início sem remover
    IsEmpty() bool            // Verifica se a fila está vazia
    Size() int                // Retorna número de elementos na fila
}
```

### 1. ArrayQueue (Fila baseada em Array)

#### Implementação com Array Circular

```go
// ArrayQueue implementa uma fila usando array circular
type ArrayQueue struct {
    values   []int // Array para armazenar elementos
    front    int   // Índice do primeiro elemento
    rear     int   // Índice do último elemento
    size     int   // Número atual de elementos
    capacity int   // Capacidade máxima da fila
}

// NewArrayQueue cria uma nova instância de ArrayQueue
func NewArrayQueue(capacity int) *ArrayQueue {
    return &ArrayQueue{
        values:   make([]int, capacity),
        front:    0,
        rear:     -1,
        size:     0,
        capacity: capacity,
    }
}

// Enqueue adiciona elemento no final da fila
// Complexidade: O(1)
func (q *ArrayQueue) Enqueue(value int) {
    if q.size >= q.capacity {
        fmt.Printf("Erro: Fila cheia, não é possível adicionar %d\n", value)
        return
    }

    // Incrementa rear de forma circular
    q.rear = (q.rear + 1) % q.capacity
    q.values[q.rear] = value
    q.size++
}

// Dequeue remove e retorna elemento do início da fila
// Complexidade: O(1)
func (q *ArrayQueue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia")
    }

    value := q.values[q.front]
    // Incrementa front de forma circular
    q.front = (q.front + 1) % q.capacity
    q.size--

    return value, nil
}

// Front retorna elemento do início sem remover
// Complexidade: O(1)
func (q *ArrayQueue) Front() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia")
    }

    return q.values[q.front], nil
}

// IsEmpty verifica se a fila está vazia
// Complexidade: O(1)
func (q *ArrayQueue) IsEmpty() bool {
    return q.size == 0
}

// Size retorna número de elementos na fila
// Complexidade: O(1)
func (q *ArrayQueue) Size() int {
    return q.size
}

// IsFull verifica se a fila está cheia
func (q *ArrayQueue) IsFull() bool {
    return q.size == q.capacity
}

// String retorna representação em string da fila
func (q *ArrayQueue) String() string {
    if q.IsEmpty() {
        return "[]"
    }

    result := "["
    for i := 0; i < q.size; i++ {
        index := (q.front + i) % q.capacity
        if i > 0 {
            result += ", "
        }
        result += fmt.Sprintf("%d", q.values[index])
    }
    result += "]"

    return result
}
```

### 2. LinkedListQueue (Fila baseada em Lista Ligada)

#### Implementação com Nós Ligados

```go
// QueueNode representa um nó na fila ligada
type QueueNode struct {
    data int        // Valor armazenado no nó
    next *QueueNode // Ponteiro para o próximo nó
}

// NewQueueNode cria um novo nó
func NewQueueNode(value int) *QueueNode {
    return &QueueNode{
        data: value,
        next: nil,
    }
}

// LinkedListQueue implementa uma fila usando lista ligada
type LinkedListQueue struct {
    front *QueueNode // Ponteiro para o primeiro nó (início da fila)
    rear  *QueueNode // Ponteiro para o último nó (final da fila)
    size  int        // Número de elementos na fila
}

// NewLinkedListQueue cria uma nova instância de LinkedListQueue
func NewLinkedListQueue() *LinkedListQueue {
    return &LinkedListQueue{
        front: nil,
        rear:  nil,
        size:  0,
    }
}

// Enqueue adiciona elemento no final da fila
// Complexidade: O(1)
func (q *LinkedListQueue) Enqueue(value int) {
    newNode := NewQueueNode(value)

    if q.rear == nil {
        // Fila vazia - novo nó é tanto front quanto rear
        q.front = newNode
        q.rear = newNode
    } else {
        // Conecta novo nó ao final e atualiza rear
        q.rear.next = newNode
        q.rear = newNode
    }

    q.size++
}

// Dequeue remove e retorna elemento do início da fila
// Complexidade: O(1)
func (q *LinkedListQueue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia")
    }

    value := q.front.data
    q.front = q.front.next

    // Se a fila ficou vazia, rear também deve ser nil
    if q.front == nil {
        q.rear = nil
    }

    q.size--
    return value, nil
}

// Front retorna elemento do início sem remover
// Complexidade: O(1)
func (q *LinkedListQueue) Front() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia")
    }

    return q.front.data, nil
}

// IsEmpty verifica se a fila está vazia
// Complexidade: O(1)
func (q *LinkedListQueue) IsEmpty() bool {
    return q.front == nil
}

// Size retorna número de elementos na fila
// Complexidade: O(1)
func (q *LinkedListQueue) Size() int {
    return q.size
}

// String retorna representação em string da fila
func (q *LinkedListQueue) String() string {
    if q.IsEmpty() {
        return "[]"
    }

    result := "["
    current := q.front
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
```

### 3. Exemplo de Uso

```go
func main() {
    fmt.Println("=== Testando ArrayQueue ===")
    testQueue(NewArrayQueue(5))

    fmt.Println("\n=== Testando LinkedListQueue ===")
    testQueue(NewLinkedListQueue())
}

func testQueue(queue IQueue) {
    // Teste de inserção
    fmt.Println("Inserindo elementos: 10, 20, 30")
    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    fmt.Printf("Tamanho: %d\n", queue.Size())
    fmt.Printf("Fila: %s\n", queue)

    // Teste de Front
    if front, err := queue.Front(); err == nil {
        fmt.Printf("Primeiro elemento: %d\n", front)
    }

    // Teste de remoção
    fmt.Println("\nRemovendo elementos:")
    for !queue.IsEmpty() {
        if value, err := queue.Dequeue(); err == nil {
            fmt.Printf("Removido: %d\n", value)
            fmt.Printf("Tamanho: %d\n", queue.Size())
        }
    }

    // Teste de fila vazia
    fmt.Println("\nTestando fila vazia:")
    if _, err := queue.Dequeue(); err != nil {
        fmt.Printf("Erro esperado: %s\n", err)
    }

    if _, err := queue.Front(); err != nil {
        fmt.Printf("Erro esperado: %s\n", err)
    }
}
```

### 4. Comparação entre Implementações

| Aspecto           | ArrayQueue                 | LinkedListQueue                |
| ----------------- | -------------------------- | ------------------------------ |
| **Memória**       | Fixa (pré-alocada)         | Dinâmica (conforme necessário) |
| **Overhead**      | Menor                      | Maior (ponteiros)              |
| **Capacidade**    | Limitada                   | Ilimitada (memória disponível) |
| **Cache**         | Melhor localidade          | Pior localidade                |
| **Implementação** | Mais complexa (circular)   | Mais simples                   |
| **Desperdício**   | Possível (se subutilizada) | Nenhum                         |

### 5. Análise de Complexidade

#### ArrayQueue

| Operação | Complexidade | Justificativa                 |
| -------- | ------------ | ----------------------------- |
| Enqueue  | O(1)         | Acesso direto ao índice rear  |
| Dequeue  | O(1)         | Acesso direto ao índice front |
| Front    | O(1)         | Acesso direto ao índice front |
| IsEmpty  | O(1)         | Verificação de variável size  |
| Size     | O(1)         | Retorno de variável size      |

#### LinkedListQueue

| Operação | Complexidade | Justificativa                        |
| -------- | ------------ | ------------------------------------ |
| Enqueue  | O(1)         | Inserção no final com ponteiro rear  |
| Dequeue  | O(1)         | Remoção do início com ponteiro front |
| Front    | O(1)         | Acesso direto ao nó front            |
| IsEmpty  | O(1)         | Verificação se front é nil           |
| Size     | O(1)         | Retorno de variável size             |

### 6. Vantagens e Desvantagens

#### ArrayQueue

**Vantagens:**
✅ Melhor performance (cache-friendly)
✅ Menor uso de memória por elemento
✅ Acesso direto por índice
✅ Sem fragmentação de memória

**Desvantagens:**
❌ Tamanho fixo (limitado)
❌ Pode desperdiçar memória
❌ Implementação mais complexa (circular)
❌ Overflow se exceder capacidade

#### LinkedListQueue

**Vantagens:**
✅ Tamanho dinâmico
✅ Usa apenas memória necessária
✅ Implementação mais simples
✅ Sem limite de capacidade

**Desvantagens:**
❌ Overhead de ponteiros
❌ Pior localidade de memória
❌ Alocação/desalocação frequente
❌ Possível fragmentação

### 7. Quando Usar Cada Implementação

**Use ArrayQueue quando:**

- Tamanho máximo é conhecido
- Performance é crítica
- Memória é limitada
- Acesso frequente aos elementos

**Use LinkedListQueue quando:**

- Tamanho varia muito
- Não há limite de capacidade
- Simplicidade é prioridade
- Memória é abundante

### 8. Implementação Alternativa: ArrayQueue Dinâmico

```go
// ArrayQueueDynamic com redimensionamento automático
type ArrayQueueDynamic struct {
    values []int
    front  int
    rear   int
    size   int
}

func (q *ArrayQueueDynamic) Enqueue(value int) {
    if q.size == len(q.values) {
        q.resize() // Dobra o tamanho quando necessário
    }

    q.rear = (q.rear + 1) % len(q.values)
    q.values[q.rear] = value
    q.size++
}

func (q *ArrayQueueDynamic) resize() {
    newCapacity := len(q.values) * 2
    if newCapacity == 0 {
        newCapacity = 1
    }

    newValues := make([]int, newCapacity)

    // Copia elementos na ordem correta
    for i := 0; i < q.size; i++ {
        index := (q.front + i) % len(q.values)
        newValues[i] = q.values[index]
    }

    q.values = newValues
    q.front = 0
    q.rear = q.size - 1
}
```

### Conclusão

Ambas as implementações seguem a interface `IQueue` e oferecem operações O(1) para todas as operações básicas. A escolha entre elas depende dos requisitos específicos:

- **ArrayQueue**: Melhor para performance e uso controlado de memória
- **LinkedListQueue**: Melhor para flexibilidade e simplicidade

A implementação com interface permite trocar facilmente entre as duas, demonstrando o poder da abstração em programação orientada a objetos.
