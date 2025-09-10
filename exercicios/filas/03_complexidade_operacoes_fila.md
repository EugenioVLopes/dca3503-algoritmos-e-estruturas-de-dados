# Complexidade de Tempo das Operações de Fila

## Pergunta

Considere as estruturas de dados apresentadas na tabela a seguir e responda o desempenho de tempo de pior caso e melhor caso para cada operação listada. Assuma que a fila tem um tamanho máximo (o vetor não precisa duplicar).

| Operação | ArrayQueue |  | LinkedListQueue |  |
|----------|------------|--|-----------------|--|
|          | Pior Caso | Melhor Caso | Pior Caso | Melhor Caso |
| Enqueue(value int) | O(1) | Ω(1) | O(1) | Ω(1) |
| Dequeue() (int, error) | O(1) | Ω(1) | O(1) | Ω(1) |
| Front() (int, error) | O(1) | Ω(1) | O(1) | Ω(1) |
| Size() | O(1) | Ω(1) | O(1) | Ω(1) |

## Resposta

### ArrayQueue (Fila baseada em Array Circular)

#### Enqueue(value int)
- **Pior caso O(1)**: Inserção direta no índice rear
- **Melhor caso Ω(1)**: Inserção direta no índice rear

**Explicação:**
```go
func (q *ArrayQueue) Enqueue(value int) {
    if q.size >= q.capacity {
        return // Fila cheia - O(1)
    }
    
    q.rear = (q.rear + 1) % q.capacity  // O(1) - aritmética modular
    q.values[q.rear] = value            // O(1) - acesso direto ao array
    q.size++                            // O(1) - incremento
}
```

**Por que sempre O(1):**
- **Acesso direto**: Array permite acesso O(1) por índice
- **Sem deslocamento**: Não move elementos existentes
- **Aritmética simples**: Cálculo modular é O(1)
- **Sem redimensionamento**: Tamanho fixo (conforme enunciado)

#### Dequeue() (int, error)
- **Pior caso O(1)**: Remoção direta do índice front
- **Melhor caso Ω(1)**: Remoção direta do índice front

**Explicação:**
```go
func (q *ArrayQueue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia") // O(1)
    }
    
    value := q.values[q.front]              // O(1) - acesso direto
    q.front = (q.front + 1) % q.capacity    // O(1) - aritmética modular
    q.size--                                // O(1) - decremento
    
    return value, nil
}
```

**Por que sempre O(1):**
- **Sem deslocamento**: Não move elementos restantes
- **Ponteiro front**: Apenas atualiza índice de início
- **Acesso direto**: Leitura O(1) do array

#### Front() (int, error)
- **Pior caso O(1)**: Acesso direto ao índice front
- **Melhor caso Ω(1)**: Acesso direto ao índice front

**Explicação:**
```go
func (q *ArrayQueue) Front() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia") // O(1)
    }
    
    return q.values[q.front], nil // O(1) - acesso direto
}
```

#### Size()
- **Pior caso O(1)**: Retorna variável armazenada
- **Melhor caso Ω(1)**: Retorna variável armazenada

**Explicação:**
```go
func (q *ArrayQueue) Size() int {
    return q.size // O(1) - acesso direto à variável
}
```

### LinkedListQueue (Fila baseada em Lista Ligada)

#### Enqueue(value int)
- **Pior caso O(1)**: Inserção no final com ponteiro rear
- **Melhor caso Ω(1)**: Inserção no final com ponteiro rear

**Explicação:**
```go
func (q *LinkedListQueue) Enqueue(value int) {
    newNode := NewQueueNode(value) // O(1) - criação de nó
    
    if q.rear == nil {
        // Fila vazia - O(1)
        q.front = newNode
        q.rear = newNode
    } else {
        // Conecta ao final - O(1)
        q.rear.next = newNode
        q.rear = newNode
    }
    
    q.size++ // O(1)
}
```

**Por que sempre O(1):**
- **Ponteiro rear**: Acesso direto ao último nó
- **Sem percurso**: Não precisa percorrer a lista
- **Alocação**: Criação de nó é O(1)

#### Dequeue() (int, error)
- **Pior caso O(1)**: Remoção do início com ponteiro front
- **Melhor caso Ω(1)**: Remoção do início com ponteiro front

**Explicação:**
```go
func (q *LinkedListQueue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia") // O(1)
    }
    
    value := q.front.data  // O(1) - acesso direto
    q.front = q.front.next // O(1) - atualização de ponteiro
    
    if q.front == nil {
        q.rear = nil // O(1) - fila ficou vazia
    }
    
    q.size-- // O(1)
    return value, nil
}
```

**Por que sempre O(1):**
- **Ponteiro front**: Acesso direto ao primeiro nó
- **Sem percurso**: Não precisa percorrer a lista
- **Desalocação**: Remoção de referência é O(1)

#### Front() (int, error)
- **Pior caso O(1)**: Acesso direto ao nó front
- **Melhor caso Ω(1)**: Acesso direto ao nó front

**Explicação:**
```go
func (q *LinkedListQueue) Front() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia") // O(1)
    }
    
    return q.front.data, nil // O(1) - acesso direto
}
```

#### Size()
- **Pior caso O(1)**: Retorna variável armazenada
- **Melhor caso Ω(1)**: Retorna variável armazenada

**Explicação:**
```go
func (q *LinkedListQueue) Size() int {
    return q.size // O(1) - acesso direto à variável
}
```

### Resumo Comparativo

| Operação | ArrayQueue | LinkedListQueue | Observações |
|----------|------------|-----------------|-------------|
| **Enqueue** | Θ(1) | Θ(1) | Ambas eficientes |
| **Dequeue** | Θ(1) | Θ(1) | Ambas eficientes |
| **Front** | Θ(1) | Θ(1) | Ambas eficientes |
| **Size** | Θ(1) | Θ(1) | Ambas mantêm contador |

### Análise Detalhada

#### Por que todas as operações são O(1)?

**Princípio fundamental das filas:**
- **FIFO**: Inserção no final, remoção no início
- **Ponteiros estratégicos**: Mantém referências para ambas extremidades
- **Sem percurso**: Nunca precisa percorrer elementos intermediários

#### ArrayQueue - Estratégia Circular

```
Array Circular:
┌─────┬─────┬─────┬─────┬─────┐
│  -  │ 20  │ 30  │ 40  │  -  │
└─────┴─────┴─────┴─────┴─────┘
   0     1     2     3     4
         ↑           ↑
      front        rear

Enqueue(50): rear = (3+1) % 5 = 4
Dequeue(): front = (1+1) % 5 = 2
```

**Vantagens:**
- Reutiliza espaço do array
- Evita deslocamento de elementos
- Operações sempre O(1)

#### LinkedListQueue - Ponteiros Duplos

```
Lista Ligada:
front → [10|•] → [20|•] → [30|•] → [40|nil] ← rear

Enqueue(50): rear.next = newNode, rear = newNode
Dequeue(): front = front.next
```

**Vantagens:**
- Tamanho dinâmico
- Sem desperdício de memória
- Operações sempre O(1)

### Comparação com Implementações Ingênuas

#### Fila com Array Simples (Implementação Ruim)
```go
// IMPLEMENTAÇÃO RUIM - NÃO USAR
func (q *BadQueue) Dequeue() (int, error) {
    if q.IsEmpty() {
        return 0, errors.New("fila vazia")
    }
    
    value := q.values[0]
    
    // PROBLEMA: Desloca todos os elementos - O(n)
    for i := 1; i < q.size; i++ {
        q.values[i-1] = q.values[i]
    }
    
    q.size--
    return value, nil
}
```

**Por que é ruim:**
- Dequeue vira O(n) devido ao deslocamento
- Desperdiça processamento
- Não aproveita características da fila

#### Fila com Lista sem Ponteiro Rear (Implementação Ruim)
```go
// IMPLEMENTAÇÃO RUIM - NÃO USAR
func (q *BadLinkedQueue) Enqueue(value int) {
    newNode := NewNode(value)
    
    if q.front == nil {
        q.front = newNode
        return
    }
    
    // PROBLEMA: Percorre até o final - O(n)
    current := q.front
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}
```

**Por que é ruim:**
- Enqueue vira O(n) devido ao percurso
- Não mantém ponteiro rear
- Desperdiça a vantagem da lista ligada

### Fatores que Garantem O(1)

#### ArrayQueue
1. **Array circular**: Evita deslocamento
2. **Índices front/rear**: Acesso direto às extremidades
3. **Aritmética modular**: Cálculo O(1)
4. **Contador size**: Evita percurso para contar

#### LinkedListQueue
1. **Ponteiro front**: Acesso direto ao início
2. **Ponteiro rear**: Acesso direto ao final
3. **Alocação dinâmica**: Criação de nó é O(1)
4. **Contador size**: Evita percurso para contar

### Casos Especiais

#### Fila Vazia
```go
// Todas as operações ainda são O(1)
if q.IsEmpty() {
    return 0, errors.New("fila vazia") // O(1)
}
```

#### Fila com Um Elemento
```go
// ArrayQueue
q.front == q.rear // O(1) - comparação simples

// LinkedListQueue  
q.front == q.rear // O(1) - mesmo nó
```

#### Fila Cheia (ArrayQueue)
```go
if q.size >= q.capacity {
    return // O(1) - verificação simples
}
```

### Análise de Memória

| Aspecto | ArrayQueue | LinkedListQueue |
|---------|------------|------------------|
| **Overhead por elemento** | 0 bytes | 8 bytes (ponteiro) |
| **Memória total** | Fixa | Dinâmica |
| **Fragmentação** | Nenhuma | Possível |
| **Cache locality** | Melhor | Pior |

### Conclusão

**Todas as operações são Θ(1) em ambas implementações** porque:

1. **Design inteligente**: Mantém ponteiros para ambas extremidades
2. **Evita percursos**: Nunca precisa percorrer elementos intermediários
3. **Operações atômicas**: Cada operação faz apenas atualizações simples
4. **Contador size**: Evita contagem manual

A **eficiência O(1)** é uma característica fundamental de implementações corretas de fila, tornando-as ideais para aplicações que requerem alta performance em operações de inserção e remoção.