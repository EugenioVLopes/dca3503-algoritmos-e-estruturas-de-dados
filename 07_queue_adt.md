# Queue (Fila) - Tipo Abstrato de Dados

## Introdução

Uma **fila** (queue) é uma estrutura de dados linear que segue o princípio **FIFO** (First In, First Out - "primeiro a entrar, primeiro a sair"). Imagine uma fila de pessoas: a primeira pessoa a chegar é a primeira a ser atendida.

## Características Principais

### Princípio FIFO
- O primeiro elemento inserido é o primeiro a ser removido
- Inserção acontece sempre no final (rear/back)
- Remoção acontece sempre no início (front)
- Acesso restrito às extremidades

### Operações Fundamentais

1. **Enqueue**: Adiciona elemento no final da fila
2. **Dequeue**: Remove e retorna o elemento do início
3. **Front**: Consulta o elemento do início sem removê-lo
4. **Rear**: Consulta o elemento do final sem removê-lo
5. **IsEmpty**: Verifica se a fila está vazia
6. **Size**: Retorna o número de elementos

## Interface Queue

```go
type Queue interface {
    // Operações básicas
    Enqueue(element int)        // Adiciona no final
    Dequeue() (int, error)      // Remove do início
    Front() (int, error)        // Consulta o início
    Rear() (int, error)         // Consulta o final
    
    // Operações de consulta
    Size() int                  // Número de elementos
    IsEmpty() bool             // Verifica se vazia
    IsFull() bool              // Verifica se cheia
    
    // Operações auxiliares
    Clear()                    // Remove todos
    ToSlice() []int            // Converte para slice
    String() string            // Representação textual
}
```

## Implementações

### 1. ArrayQueue (Baseada em Array Circular)

**Características:**
- Usa array circular para otimizar operações
- Evita necessidade de mover elementos
- Redimensionamento automático quando necessário
- Reutiliza espaço liberado por dequeue

**Vantagens:**
- Operações O(1) para enqueue/dequeue
- Uso eficiente de memória (sem fragmentação)
- Cache-friendly (localidade espacial)
- Controle preciso sobre capacidade

**Desvantagens:**
- Redimensionamento ocasional pode ser custoso: O(n)
- Capacidade fixa entre redimensionamentos
- Complexidade adicional do array circular
- Pode desperdiçar memória se mal dimensionada

**Complexidades:**
```
Enqueue:  O(1) amortizado, O(n) no pior caso (redimensionamento)
Dequeue:  O(1)
Front:    O(1)
Rear:     O(1)
IsEmpty:  O(1)
Size:     O(1)
```

**Array Circular:**
```
Índices:  0  1  2  3  4
Array:   [A][B][ ][ ][ ]
         ↑           ↑
       front       rear

Após enqueue(C):
Array:   [A][B][C][ ][ ]
         ↑        ↑
       front    rear

Após dequeue():
Array:   [ ][B][C][ ][ ]
            ↑     ↑
          front rear
```

### 2. LinkedQueue (Baseada em Lista Ligada)

**Características:**
- Usa nós ligados por ponteiros
- Mantém ponteiros para front e rear
- Crescimento dinâmico sem limite predefinido
- Cada elemento tem overhead de ponteiro

**Vantagens:**
- Enqueue/Dequeue sempre O(1) (sem redimensionamento)
- Não há limite de capacidade (exceto memória disponível)
- Não desperdiça memória (aloca exatamente o necessário)
- Flexibilidade total de tamanho
- Implementação mais simples

**Desvantagens:**
- Overhead de memória por ponteiros
- Menor localidade espacial (cache menos eficiente)
- Fragmentação de memória
- Acesso sequencial obrigatório

**Complexidades:**
```
Enqueue:  O(1) sempre
Dequeue:  O(1) sempre
Front:    O(1)
Rear:     O(1)
IsEmpty:  O(1)
Size:     O(1)
```

**Estrutura da Lista:**
```
front → [A|•] → [B|•] → [C|•] → null
                        ↑
                      rear

Após enqueue(D):
front → [A|•] → [B|•] → [C|•] → [D|•] → null
                                ↑
                              rear

Após dequeue():
front → [B|•] → [C|•] → [D|•] → null
                        ↑
                      rear
```

## Comparação das Implementações

| Aspecto | ArrayQueue | LinkedQueue |
|---------|------------|-------------|
| **Complexidade Enqueue** | O(1) amortizado | O(1) sempre |
| **Complexidade Dequeue** | O(1) | O(1) |
| **Uso de Memória** | Mais eficiente | Overhead de ponteiros |
| **Cache Performance** | Melhor | Pior |
| **Limite de Tamanho** | Capacidade atual | Apenas memória total |
| **Redimensionamento** | Automático | Não necessário |
| **Fragmentação** | Baixa | Pode ocorrer |
| **Implementação** | Mais complexa | Mais simples |

## Aplicações Práticas

### 1. Sistemas de Processamento
```go
// Fila de tarefas para processamento
type TaskQueue struct {
    queue Queue
}

func (tq *TaskQueue) AddTask(task int) {
    tq.queue.Enqueue(task)
}

func (tq *TaskQueue) ProcessNext() (int, error) {
    return tq.queue.Dequeue()
}
```

### 2. Algoritmos de Busca (BFS)
```go
// Busca em largura usando fila
func BreadthFirstSearch(graph [][]int, start int) []int {
    queue := NewArrayQueue(len(graph))
    visited := make([]bool, len(graph))
    result := []int{}
    
    queue.Enqueue(start)
    visited[start] = true
    
    for !queue.IsEmpty() {
        current, _ := queue.Dequeue()
        result = append(result, current)
        
        // Adiciona vizinhos não visitados
        for _, neighbor := range graph[current] {
            if !visited[neighbor] {
                queue.Enqueue(neighbor)
                visited[neighbor] = true
            }
        }
    }
    
    return result
}
```

### 3. Geração de Números Binários
```go
// Gera representações binárias de 1 a n
func GenerateBinaryNumbers(n int) []string {
    queue := NewLinkedQueue()
    result := []string{}
    
    queue.Enqueue(1) // Começa com "1"
    
    for i := 1; i <= n; i++ {
        current, _ := queue.Dequeue()
        binary := fmt.Sprintf("%b", current)
        result = append(result, binary)
        
        // Gera próximos: current*2 e current*2+1
        queue.Enqueue(current * 2)
        queue.Enqueue(current*2 + 1)
    }
    
    return result
}
```

### 4. Buffer Circular
```go
// Buffer para streaming de dados
type CircularBuffer struct {
    queue *ArrayQueue
    maxSize int
}

func (cb *CircularBuffer) Add(data int) {
    if cb.queue.Size() >= cb.maxSize {
        cb.queue.Dequeue() // Remove mais antigo
    }
    cb.queue.Enqueue(data)
}
```

### 5. Simulação de Sistemas
```go
// Simulação de fila de atendimento
type ServiceQueue struct {
    customers Queue
    serviceTime map[int]int
}

func (sq *ServiceQueue) ArriveCustomer(id int, serviceTime int) {
    sq.customers.Enqueue(id)
    sq.serviceTime[id] = serviceTime
}

func (sq *ServiceQueue) ServeNext() (int, int, error) {
    if sq.customers.IsEmpty() {
        return 0, 0, errors.New("nenhum cliente na fila")
    }
    
    customer, _ := sq.customers.Dequeue()
    time := sq.serviceTime[customer]
    delete(sq.serviceTime, customer)
    
    return customer, time, nil
}
```

## Algoritmos com Filas

### Primeiro Caractere Não Repetido
```go
func FirstNonRepeatingCharacter(stream string) []rune {
    queue := NewLinkedQueue()
    frequency := make(map[rune]int)
    result := []rune{}
    
    for _, char := range stream {
        frequency[char]++
        
        if frequency[char] == 1 {
            queue.Enqueue(int(char))
        }
        
        // Remove caracteres repetidos da frente
        for !queue.IsEmpty() {
            front, _ := queue.Front()
            if frequency[rune(front)] > 1 {
                queue.Dequeue()
            } else {
                break
            }
        }
        
        // Resultado atual
        if queue.IsEmpty() {
            result = append(result, 0)
        } else {
            front, _ := queue.Front()
            result = append(result, rune(front))
        }
    }
    
    return result
}
```

### Intercalar Fila
```go
// Intercala primeira e segunda metade da fila
func InterleaveQueue(queue Queue) {
    size := queue.Size()
    if size < 2 {
        return
    }
    
    half := size / 2
    aux := NewArrayQueue(half)
    
    // Move primeira metade para auxiliar
    for i := 0; i < half; i++ {
        value, _ := queue.Dequeue()
        aux.Enqueue(value)
    }
    
    // Intercala elementos
    for !aux.IsEmpty() {
        // Da primeira metade
        value1, _ := aux.Dequeue()
        queue.Enqueue(value1)
        
        // Da segunda metade
        if !queue.IsEmpty() {
            value2, _ := queue.Dequeue()
            queue.Enqueue(value2)
        }
    }
}
```

## Considerações de Design

### Quando Usar ArrayQueue
- Tamanho previsível ou com limite superior
- Performance crítica com muitas operações
- Uso intensivo de cache
- Controle rigoroso de memória
- Sistemas embarcados

### Quando Usar LinkedQueue
- Tamanho muito variável e imprevisível
- Sem limite de capacidade conhecido
- Simplicidade de implementação
- Flexibilidade máxima
- Prototipagem rápida

### Tratamento de Erros
- Dequeue/Front/Rear em fila vazia deve retornar erro
- Verificar sempre se operação foi bem-sucedida
- Usar padrão Go de retorno (valor, erro)
- Considerar comportamento em condições limite

## Padrões de Uso

### 1. Processamento em Lote
```go
// Processa elementos em lotes
func ProcessBatch(queue Queue, batchSize int) {
    for !queue.IsEmpty() {
        batch := make([]int, 0, batchSize)
        
        // Coleta lote
        for i := 0; i < batchSize && !queue.IsEmpty(); i++ {
            element, _ := queue.Dequeue()
            batch = append(batch, element)
        }
        
        // Processa lote
        processBatch(batch)
    }
}
```

### 2. Buffer com Timeout
```go
// Buffer que processa após timeout ou quando cheio
type TimedBuffer struct {
    queue Queue
    maxSize int
    timeout time.Duration
    lastFlush time.Time
}

func (tb *TimedBuffer) Add(element int) {
    tb.queue.Enqueue(element)
    
    // Flush se cheio ou timeout
    if tb.queue.Size() >= tb.maxSize || 
       time.Since(tb.lastFlush) > tb.timeout {
        tb.Flush()
    }
}
```

### 3. Fila com Prioridade Simples
```go
// Duas filas para diferentes prioridades
type PriorityQueue struct {
    highPriority Queue
    lowPriority Queue
}

func (pq *PriorityQueue) Enqueue(element int, isHighPriority bool) {
    if isHighPriority {
        pq.highPriority.Enqueue(element)
    } else {
        pq.lowPriority.Enqueue(element)
    }
}

func (pq *PriorityQueue) Dequeue() (int, error) {
    if !pq.highPriority.IsEmpty() {
        return pq.highPriority.Dequeue()
    }
    return pq.lowPriority.Dequeue()
}
```

## Otimizações

### ArrayQueue
- **Capacidade Inicial**: Baseada no uso esperado
- **Fator de Crescimento**: 2x é comum
- **Shrinking**: Reduz quando utilização < 25%
- **Alinhamento**: Considera alinhamento de memória

### LinkedQueue
- **Pool de Nós**: Reutiliza nós para evitar alocações
- **Batch Operations**: EnqueueAll/DequeueMultiple
- **Memory Pooling**: Gerenciamento customizado

## Variações e Extensões

### Deque (Double-Ended Queue)
- Inserção e remoção em ambas as extremidades
- Combina funcionalidades de pilha e fila
- Útil para algoritmos de janela deslizante

### Fila Circular
- Tamanho fixo que sobrescreve elementos antigos
- Útil para buffers de dados contínuos
- Implementação simples com módulo

### Fila Thread-Safe
- Adiciona mutexes para concorrência
- Operações atômicas
- Considerações de deadlock

### Fila Persistente
- Mantém versões anteriores após modificações
- Útil para undo/redo
- Compartilhamento eficiente de estrutura

## Análise de Performance

### Benchmarks Típicos
```
Operação         ArrayQueue    LinkedQueue
1M Enqueues      ~50ms         ~80ms
1M Dequeues      ~30ms         ~40ms
Memória (1M)     ~8MB          ~24MB
Cache Misses     Baixo         Alto
```

### Fatores de Performance
- **Localidade de Cache**: ArrayQueue vence
- **Alocações**: LinkedQueue tem mais overhead
- **Redimensionamento**: Impacto esporádico no ArrayQueue
- **Fragmentação**: Maior no LinkedQueue

## Conclusão

Filas são estruturas fundamentais com aplicações em:
- Sistemas operacionais (escalonamento)
- Redes de computadores (buffers)
- Algoritmos de grafos (BFS)
- Simulações e modelagem
- Processamento de streams

A escolha entre ArrayQueue e LinkedQueue depende dos requisitos específicos:
- **Performance**: ArrayQueue geralmente mais rápida
- **Flexibilidade**: LinkedQueue mais adaptável
- **Memória**: ArrayQueue mais eficiente
- **Simplicidade**: LinkedQueue mais fácil de implementar

Ambas as implementações oferecem operações O(1) para as operações básicas, tornando filas uma escolha excelente para muitos algoritmos e sistemas.