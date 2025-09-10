# Função Size() para ArrayQueue Circular

## Pergunta

Escreva uma função que retorne a quantidade de elementos inseridos em uma Fila implementada com vetor. Escreva a função Size() considerando que o struct ArrayQueue não contém a variável size, como apresentado na tabela a seguir. Lembre-se que os índices de front e rear inicialmente assumem o valor -1, e que o ArrayQueue tem um caráter circular.

```go
func (queue *ArrayQueue) Size()

type ArrayQueue struct {
    values []int
    front  int
    rear   int
}
```

## Resposta

### Implementação da Função Size()

```go
func (queue *ArrayQueue) Size() int {
    if queue.front == -1 {
        // Fila vazia
        return 0
    }
    
    if queue.rear >= queue.front {
        // Caso normal: rear está à frente ou igual ao front
        return queue.rear - queue.front + 1
    } else {
        // Caso circular: rear "deu a volta" no array
        return (len(queue.values) - queue.front) + (queue.rear + 1)
    }
}
```

### Análise Detalhada

#### Casos Possíveis

**1. Fila Vazia**
```
front = -1, rear = -1
Size() = 0

Array: [ ][ ][ ][ ][ ]
        0  1  2  3  4
```

**2. Caso Normal (rear >= front)**
```
front = 1, rear = 3
Size() = 3 - 1 + 1 = 3

Array: [ ][A][B][C][ ]
        0  1  2  3  4
           ↑     ↑
        front  rear
```

**3. Caso Circular (rear < front)**
```
front = 3, rear = 1
Size() = (5 - 3) + (1 + 1) = 2 + 2 = 4

Array: [D][E][ ][A][B]
        0  1  2  3  4
           ↑     ↑
        rear   front
```

### Explicação da Lógica

#### Condição 1: Fila Vazia
```go
if queue.front == -1 {
    return 0
}
```

**Justificativa:**
- Quando `front == -1`, a fila está vazia
- Esta é a condição inicial e também quando todos elementos são removidos
- `rear` também será -1 neste caso

#### Condição 2: Caso Normal
```go
if queue.rear >= queue.front {
    return queue.rear - queue.front + 1
}
```

**Justificativa:**
- Elementos estão em sequência contígua
- Fórmula: `rear - front + 1`
- O `+1` é necessário porque ambos os índices são inclusivos

**Exemplo:**
```
front = 2, rear = 4
Elementos nas posições: 2, 3, 4
Quantidade: 4 - 2 + 1 = 3 ✓
```

#### Condição 3: Caso Circular
```go
else {
    return (len(queue.values) - queue.front) + (queue.rear + 1)
}
```

**Justificativa:**
- Elementos estão em duas partes: final + início do array
- **Parte 1**: `len(queue.values) - queue.front` (do front até o final)
- **Parte 2**: `queue.rear + 1` (do início até rear)

**Exemplo:**
```
Array tamanho 5: [C][D][ ][A][B]
                  0  1  2  3  4
                     ↑     ↑
                  rear=1  front=3

Parte 1 (final): posições 3, 4 → 5 - 3 = 2 elementos
Parte 2 (início): posições 0, 1 → 1 + 1 = 2 elementos
Total: 2 + 2 = 4 elementos ✓
```

### Implementação Completa com Operações

```go
package main

import (
    "errors"
    "fmt"
)

type ArrayQueue struct {
    values []int
    front  int
    rear   int
}

// NewArrayQueue cria uma nova fila circular
func NewArrayQueue(capacity int) *ArrayQueue {
    return &ArrayQueue{
        values: make([]int, capacity),
        front:  -1,
        rear:   -1,
    }
}

// Size retorna o número de elementos na fila
func (queue *ArrayQueue) Size() int {
    if queue.front == -1 {
        // Fila vazia
        return 0
    }
    
    if queue.rear >= queue.front {
        // Caso normal: rear está à frente ou igual ao front
        return queue.rear - queue.front + 1
    } else {
        // Caso circular: rear "deu a volta" no array
        return (len(queue.values) - queue.front) + (queue.rear + 1)
    }
}

// IsEmpty verifica se a fila está vazia
func (queue *ArrayQueue) IsEmpty() bool {
    return queue.front == -1
}

// IsFull verifica se a fila está cheia
func (queue *ArrayQueue) IsFull() bool {
    return queue.Size() == len(queue.values)
}

// Enqueue adiciona elemento na fila
func (queue *ArrayQueue) Enqueue(value int) error {
    if queue.IsFull() {
        return errors.New("fila cheia")
    }
    
    if queue.IsEmpty() {
        // Primeiro elemento
        queue.front = 0
        queue.rear = 0
    } else {
        // Incrementa rear de forma circular
        queue.rear = (queue.rear + 1) % len(queue.values)
    }
    
    queue.values[queue.rear] = value
    return nil
}

// Dequeue remove elemento da fila
func (queue *ArrayQueue) Dequeue() (int, error) {
    if queue.IsEmpty() {
        return 0, errors.New("fila vazia")
    }
    
    value := queue.values[queue.front]
    
    if queue.front == queue.rear {
        // Último elemento - fila fica vazia
        queue.front = -1
        queue.rear = -1
    } else {
        // Incrementa front de forma circular
        queue.front = (queue.front + 1) % len(queue.values)
    }
    
    return value, nil
}

// Front retorna o primeiro elemento sem remover
func (queue *ArrayQueue) Front() (int, error) {
    if queue.IsEmpty() {
        return 0, errors.New("fila vazia")
    }
    
    return queue.values[queue.front], nil
}

// String retorna representação da fila
func (queue *ArrayQueue) String() string {
    if queue.IsEmpty() {
        return "[]"
    }
    
    result := "["
    size := queue.Size()
    
    for i := 0; i < size; i++ {
        index := (queue.front + i) % len(queue.values)
        if i > 0 {
            result += ", "
        }
        result += fmt.Sprintf("%d", queue.values[index])
    }
    
    result += "]"
    return result
}
```

### Exemplo de Uso e Teste

```go
func main() {
    queue := NewArrayQueue(5)
    
    fmt.Printf("Fila inicial: %s, Size: %d\n", queue, queue.Size())
    
    // Teste 1: Inserções normais
    fmt.Println("\n=== Inserções Normais ===")
    queue.Enqueue(10)
    fmt.Printf("Após Enqueue(10): %s, Size: %d\n", queue, queue.Size())
    
    queue.Enqueue(20)
    queue.Enqueue(30)
    fmt.Printf("Após Enqueue(20,30): %s, Size: %d\n", queue, queue.Size())
    
    // Teste 2: Remoções e inserções (criando caso circular)
    fmt.Println("\n=== Criando Caso Circular ===")
    queue.Dequeue() // Remove 10
    queue.Dequeue() // Remove 20
    fmt.Printf("Após 2 Dequeues: %s, Size: %d\n", queue, queue.Size())
    
    queue.Enqueue(40)
    queue.Enqueue(50)
    queue.Enqueue(60)
    fmt.Printf("Após Enqueue(40,50,60): %s, Size: %d\n", queue, queue.Size())
    
    // Teste 3: Estado dos índices
    fmt.Printf("\nEstado interno: front=%d, rear=%d\n", queue.front, queue.rear)
    fmt.Printf("Array interno: %v\n", queue.values)
    
    // Teste 4: Esvaziar fila
    fmt.Println("\n=== Esvaziando Fila ===")
    for !queue.IsEmpty() {
        value, _ := queue.Dequeue()
        fmt.Printf("Dequeue: %d, Size: %d\n", value, queue.Size())
    }
}
```

### Saída Esperada

```
Fila inicial: [], Size: 0

=== Inserções Normais ===
Após Enqueue(10): [10], Size: 1
Após Enqueue(20,30): [10, 20, 30], Size: 3

=== Criando Caso Circular ===
Após 2 Dequeues: [30], Size: 1
Após Enqueue(40,50,60): [30, 40, 50, 60], Size: 4

Estado interno: front=2, rear=0
Array interno: [60 0 30 40 50]

=== Esvaziando Fila ===
Dequeue: 30, Size: 3
Dequeue: 40, Size: 2
Dequeue: 50, Size: 1
Dequeue: 60, Size: 0
```

### Casos de Teste Específicos

#### Teste 1: Fila Vazia
```go
queue := NewArrayQueue(3)
fmt.Printf("Size: %d\n", queue.Size()) // Output: 0
```

#### Teste 2: Um Elemento
```go
queue.Enqueue(100)
fmt.Printf("Size: %d\n", queue.Size()) // Output: 1
// front=0, rear=0 → 0-0+1 = 1
```

#### Teste 3: Caso Normal
```go
// front=0, rear=2
queue.Enqueue(10)
queue.Enqueue(20)
queue.Enqueue(30)
fmt.Printf("Size: %d\n", queue.Size()) // Output: 3
// 2-0+1 = 3
```

#### Teste 4: Caso Circular
```go
// Após algumas operações: front=2, rear=0
fmt.Printf("Size: %d\n", queue.Size()) // Output: 2
// (3-2) + (0+1) = 1 + 1 = 2
```

### Análise de Complexidade

**Complexidade de Tempo:** O(1)
- Apenas operações aritméticas simples
- Não há loops ou recursão
- Acesso direto às variáveis

**Complexidade de Espaço:** O(1)
- Não usa memória adicional
- Apenas variáveis locais

### Vantagens da Implementação

✅ **Eficiência**: O(1) para calcular tamanho
✅ **Simplicidade**: Lógica clara e direta
✅ **Robustez**: Trata todos os casos possíveis
✅ **Sem overhead**: Não precisa de variável size extra

### Desvantagens

❌ **Cálculo repetitivo**: Size() precisa calcular a cada chamada
❌ **Complexidade lógica**: Mais complexo que manter contador
❌ **Propenso a erros**: Lógica circular pode confundir

### Comparação: Com vs Sem Variável Size

| Aspecto | Com size | Sem size |
|---------|----------|----------|
| **Memória** | +4 bytes | 0 bytes extra |
| **Size()** | O(1) direto | O(1) calculado |
| **Enqueue/Dequeue** | +1 operação | 0 operações extras |
| **Complexidade** | Menor | Maior |
| **Manutenção** | Mais fácil | Mais difícil |

### Conclusão

A implementação da função `Size()` sem variável auxiliar é **possível e eficiente** (O(1)), mas requer cuidado com a lógica circular. A fórmula considera dois casos:

1. **Normal**: `rear - front + 1`
2. **Circular**: `(capacity - front) + (rear + 1)`

Embora funcional, **manter uma variável size** é geralmente preferível por simplicidade e clareza do código, especialmente em sistemas de produção onde a manutenibilidade é crucial.