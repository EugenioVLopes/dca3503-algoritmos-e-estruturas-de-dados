# LinkedList - Lista Ligada

## Visão Geral

A **LinkedList** é uma implementação da Lista ADT baseada em **nós conectados por ponteiros**. É uma estrutura dinâmica que oferece flexibilidade máxima para inserções e remoções, especialmente no início da lista.

### Características Principais:
- 🔗 **Estrutura dinâmica**: Nós alocados conforme necessário
- ⚡ **Inserção/remoção eficiente**: O(1) no início
- 💾 **Uso exato de memória**: Sem desperdício
- 🎯 **Flexibilidade máxima**: Tamanho limitado apenas pela memória

---

## Estrutura Interna

### Componentes:

#### Node (Nó):
```go
type Node struct {
    data int    // Valor armazenado
    next *Node  // Ponteiro para próximo nó
}
```

#### LinkedList:
```go
type LinkedList struct {
    head *Node  // Ponteiro para primeiro nó
    size int    // Número de elementos (opcional, para Size() O(1))
}
```

### Representação Visual:

```
LinkedList: head → [10|•] → [20|•] → [30|•] → [40|nil]
                    ↑        ↑        ↑        ↑
                  Node1    Node2    Node3    Node4
                  
size: 4
```

### Variações:

#### 1. **Singly Linked List** (Simples)
```
[data|next] → [data|next] → [data|nil]
```

#### 2. **Doubly Linked List** (Dupla)
```
nil ← [prev|data|next] ⇄ [prev|data|next] ⇄ [prev|data|nil]
```

#### 3. **Circular Linked List** (Circular)
```
[data|next] → [data|next] → [data|next] ↗
     ↑                                  ↙
     └──────────────────────────────────
```

---

## Algoritmos e Implementação

### 1. Inicialização

**Propósito**: Criar LinkedList vazia

**Pseudocódigo**:
```
ALGORITMO Init()
INÍCIO
    head ← NULL
    size ← 0
FIM
```

**Implementação Go**:
```go
func NewLinkedList() *LinkedList {
    return &LinkedList{
        head: nil,
        size: 0,
    }
}
```

**Complexidade**: Θ(1)

---

### 2. Size() - Obter Tamanho

**Propósito**: Retornar número de elementos

**Pseudocódigo**:
```
ALGORITMO Size()
INÍCIO
    RETORNAR size  // Se mantemos contador
    
    // OU percorrer lista (se não mantemos contador)
    contador ← 0
    atual ← head
    ENQUANTO atual ≠ NULL FAÇA
        contador ← contador + 1
        atual ← atual.next
    FIM_ENQUANTO
    RETORNAR contador
FIM
```

**Implementação Go**:
```go
// Versão O(1) - com contador
func (list *LinkedList) Size() int {
    return list.size
}

// Versão O(n) - sem contador
func (list *LinkedList) SizeByTraversal() int {
    count := 0
    current := list.head
    for current != nil {
        count++
        current = current.next
    }
    return count
}
```

**Complexidade**: 
- Com contador: Θ(1)
- Sem contador: Θ(n)

---

### 3. Get(index) - Acesso por Índice

**Propósito**: Obter elemento em posição específica

**Pseudocódigo**:
```
ALGORITMO Get(index)
INÍCIO
    SE index < 0 OU index >= size ENTÃO
        RETORNAR erro
    FIM_SE
    
    atual ← head
    PARA i DE 0 ATÉ index-1 FAÇA
        atual ← atual.next
    FIM_PARA
    
    RETORNAR atual.data
FIM
```

**Implementação Go**:
```go
func (list *LinkedList) Get(index int) (int, error) {
    if index < 0 || index >= list.size {
        return 0, fmt.Errorf("índice inválido: %d", index)
    }
    
    current := list.head
    for i := 0; i < index; i++ {
        current = current.next
    }
    
    return current.data, nil
}
```

**Complexidade**: 
- **Melhor caso**: Θ(1) - primeiro elemento
- **Pior caso**: O(n) - último elemento
- **Caso médio**: O(n/2) = O(n)

**Visualização**:
```
Get(2): head → [10] → [20] → [30] → [40]
                       ↑       ↑
                    i=0,1     i=2 (encontrado)
```

---

### 4. AddFirst(element) - Inserção no Início

**Propósito**: Adicionar elemento no início da lista

**Pseudocódigo**:
```
ALGORITMO AddFirst(element)
INÍCIO
    novo_no ← criar_no(element)
    novo_no.next ← head
    head ← novo_no
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *LinkedList) AddFirst(element int) {
    newNode := &Node{
        data: element,
        next: list.head,
    }
    list.head = newNode
    list.size++
}
```

**Complexidade**: Θ(1)

**Visualização**:
```
Antes:  head → [20] → [30] → [40]
Inserir 10 no início:

1. Criar nó: [10|•]
2. Conectar:  [10] → [20] → [30] → [40]
3. Atualizar head: head → [10] → [20] → [30] → [40]
```

---

### 5. Add(element) - Inserção no Final

**Propósito**: Adicionar elemento no final da lista

**Pseudocódigo**:
```
ALGORITMO Add(element)
INÍCIO
    novo_no ← criar_no(element)
    
    SE head = NULL ENTÃO
        head ← novo_no
    SENÃO
        atual ← head
        ENQUANTO atual.next ≠ NULL FAÇA
            atual ← atual.next
        FIM_ENQUANTO
        atual.next ← novo_no
    FIM_SE
    
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *LinkedList) Add(element int) {
    newNode := &Node{
        data: element,
        next: nil,
    }
    
    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    
    list.size++
}
```

**Complexidade**: O(n)
**Explicação**: Precisa percorrer até o último nó.

**Otimização com Tail Pointer**:
```go
type LinkedList struct {
    head *Node
    tail *Node  // Ponteiro para último nó
    size int
}

func (list *LinkedList) Add(element int) {
    newNode := &Node{data: element, next: nil}
    
    if list.head == nil {
        list.head = newNode
        list.tail = newNode
    } else {
        list.tail.next = newNode
        list.tail = newNode
    }
    
    list.size++
}
```
**Complexidade com tail**: Θ(1)

---

### 6. AddOnIndex(element, index) - Inserção em Posição

**Propósito**: Inserir elemento em posição específica

**Pseudocódigo**:
```
ALGORITMO AddOnIndex(element, index)
INÍCIO
    SE index < 0 OU index > size ENTÃO
        RETORNAR erro
    FIM_SE
    
    SE index = 0 ENTÃO
        AddFirst(element)
        RETORNAR
    FIM_SE
    
    novo_no ← criar_no(element)
    atual ← head
    
    // Ir até posição anterior
    PARA i DE 0 ATÉ index-2 FAÇA
        atual ← atual.next
    FIM_PARA
    
    novo_no.next ← atual.next
    atual.next ← novo_no
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *LinkedList) AddOnIndex(element int, index int) error {
    if index < 0 || index > list.size {
        return fmt.Errorf("índice inválido: %d", index)
    }
    
    if index == 0 {
        list.AddFirst(element)
        return nil
    }
    
    newNode := &Node{
        data: element,
        next: nil,
    }
    
    current := list.head
    for i := 0; i < index-1; i++ {
        current = current.next
    }
    
    newNode.next = current.next
    current.next = newNode
    list.size++
    
    return nil
}
```

**Complexidade**:
- **Melhor caso**: Θ(1) - inserção no início
- **Pior caso**: O(n) - inserção no final
- **Caso médio**: O(n/2) = O(n)

**Visualização**:
```
Inserir 25 no índice 2:
Antes:  [10] → [20] → [30] → [40]
                ↑
            index-1=1

1. Criar nó: [25]
2. Conectar: [25] → [30]
3. Religar: [20] → [25] → [30]

Depois: [10] → [20] → [25] → [30] → [40]
```

---

### 7. RemoveFirst() - Remoção do Início

**Propósito**: Remover primeiro elemento

**Pseudocódigo**:
```
ALGORITMO RemoveFirst()
INÍCIO
    SE head = NULL ENTÃO
        RETORNAR erro
    FIM_SE
    
    elemento_removido ← head.data
    head ← head.next
    size ← size - 1
    RETORNAR elemento_removido
FIM
```

**Implementação Go**:
```go
func (list *LinkedList) RemoveFirst() (int, error) {
    if list.head == nil {
        return 0, fmt.Errorf("lista vazia")
    }
    
    removedData := list.head.data
    list.head = list.head.next
    list.size--
    
    return removedData, nil
}
```

**Complexidade**: Θ(1)

---

### 8. Remove(index) - Remoção por Posição

**Propósito**: Remover elemento de posição específica

**Pseudocódigo**:
```
ALGORITMO Remove(index)
INÍCIO
    SE index < 0 OU index >= size ENTÃO
        RETORNAR erro
    FIM_SE
    
    SE index = 0 ENTÃO
        RETORNAR RemoveFirst()
    FIM_SE
    
    atual ← head
    // Ir até posição anterior
    PARA i DE 0 ATÉ index-2 FAÇA
        atual ← atual.next
    FIM_PARA
    
    elemento_removido ← atual.next.data
    atual.next ← atual.next.next
    size ← size - 1
    RETORNAR elemento_removido
FIM
```

**Implementação Go**:
```go
func (list *LinkedList) Remove(index int) (int, error) {
    if index < 0 || index >= list.size {
        return 0, fmt.Errorf("índice inválido: %d", index)
    }
    
    if index == 0 {
        return list.RemoveFirst()
    }
    
    current := list.head
    for i := 0; i < index-1; i++ {
        current = current.next
    }
    
    removedData := current.next.data
    current.next = current.next.next
    list.size--
    
    return removedData, nil
}
```

**Complexidade**:
- **Melhor caso**: Θ(1) - remoção do início
- **Pior caso**: O(n) - remoção do final
- **Caso médio**: O(n/2) = O(n)

---

## Análise de Complexidade Detalhada

### Resumo das Operações:

| Operação | Melhor Caso | Caso Médio | Pior Caso | Espaço |
|----------|-------------|------------|-----------|--------|
| Get(index) | Θ(1) | O(n) | O(n) | O(1) |
| AddFirst | Θ(1) | Θ(1) | Θ(1) | O(1) |
| Add(final) | O(n) | O(n) | O(n) | O(1) |
| Add(final)* | Θ(1) | Θ(1) | Θ(1) | O(1) |
| AddOnIndex | Θ(1) | O(n) | O(n) | O(1) |
| RemoveFirst | Θ(1) | Θ(1) | Θ(1) | O(1) |
| Remove | Θ(1) | O(n) | O(n) | O(1) |
| Size | Θ(1)** | Θ(1)** | Θ(1)** | O(1) |

*Com tail pointer  
**Com contador mantido

### Uso de Memória:

**Por elemento**:
```
Node = data + next_pointer
     = 4 bytes + 8 bytes (64-bit)
     = 12 bytes por elemento
     
Overhead = 12/4 = 300% comparado ao valor puro
```

**LinkedList total**:
```
Memória = (12 bytes × n elementos) + overhead da estrutura
        = 12n + constante
```

---

## Vantagens da LinkedList

### ✅ **Flexibilidade Dinâmica**
- **Tamanho ilimitado**: Limitado apenas pela memória disponível
- **Sem redimensionamento**: Não precisa realocar arrays
- **Crescimento incremental**: Aloca apenas o necessário

### ✅ **Inserção/Remoção Eficiente**
- **Início O(1)**: Operações no head são instantâneas
- **Sem deslocamento**: Não precisa mover outros elementos
- **Operações locais**: Afeta apenas ponteiros adjacentes

### ✅ **Uso Exato de Memória**
- **Sem desperdício**: Aloca exatamente o que precisa
- **Liberação imediata**: Memória liberada na remoção
- **Fragmentação controlada**: Nós pequenos e independentes

---

## Desvantagens da LinkedList

### ❌ **Acesso Sequencial**
- **Sem acesso direto**: Precisa percorrer desde o início
- **O(n) para acesso**: Não há "saltos" para posições
- **Cache misses**: Nós espalhados na memória

### ❌ **Overhead de Memória**
- **Ponteiros extras**: 8 bytes por nó (64-bit)
- **300% overhead**: Comparado ao valor puro
- **Fragmentação**: Nós podem estar espalhados

### ❌ **Complexidade de Implementação**
- **Gerenciamento de ponteiros**: Propenso a erros
- **Casos especiais**: Lista vazia, um elemento, etc.
- **Memory leaks**: Em linguagens sem GC

---

## Casos de Uso Ideais

### 🎯 **Use LinkedList quando:**

1. **Inserções/remoções frequentes no início**
   ```go
   // Exemplo: Stack (pilha)
   stack.AddFirst(element)  // Push
   stack.RemoveFirst()      // Pop
   ```

2. **Tamanho muito variável**
   ```go
   // Exemplo: Lista de tarefas dinâmica
   taskList.AddFirst(urgentTask)
   taskList.Remove(completedTaskIndex)
   ```

3. **Inserções/remoções no meio frequentes**
   ```go
   // Exemplo: Editor de texto (lista de linhas)
   editor.AddOnIndex(newLine, cursorPosition)
   ```

4. **Memória limitada**
   ```go
   // Exemplo: Sistema embarcado
   // LinkedList usa apenas memória necessária
   ```

5. **Implementação de outras estruturas**
   ```go
   // Exemplo: Hash table com chaining
   type HashTable struct {
       buckets []*LinkedList
   }
   ```

---

## Otimizações e Variações

### 1. **Doubly Linked List**

```go
type DoublyNode struct {
    data int
    next *DoublyNode
    prev *DoublyNode
}

type DoublyLinkedList struct {
    head *DoublyNode
    tail *DoublyNode
    size int
}
```

**Vantagens**:
- Navegação bidirecional
- Remoção O(1) com referência ao nó
- Inserção antes de um nó O(1)

**Desvantagens**:
- Mais memória (ponteiro extra)
- Mais complexidade de implementação

### 2. **Circular Linked List**

```go
// Último nó aponta para o primeiro
last.next = head
```

**Vantagens**:
- Percorrer lista infinitamente
- Implementação de round-robin
- Sem ponteiros nil

### 3. **LinkedList com Tail Pointer**

```go
type LinkedList struct {
    head *Node
    tail *Node  // Ponteiro para último nó
    size int
}
```

**Vantagens**:
- Add() no final O(1)
- RemoveLast() mais eficiente
- Melhor para filas (FIFO)

### 4. **Skip List**

```go
type SkipNode struct {
    data int
    forward []*SkipNode  // Array de ponteiros para diferentes níveis
}
```

**Vantagens**:
- Busca O(log n) probabilística
- Inserção/remoção O(log n)
- Alternativa a árvores balanceadas

---

## Implementações Especializadas

### 1. **Stack usando LinkedList**

```go
type Stack struct {
    list *LinkedList
}

func (s *Stack) Push(element int) {
    s.list.AddFirst(element)
}

func (s *Stack) Pop() (int, error) {
    return s.list.RemoveFirst()
}

func (s *Stack) Peek() (int, error) {
    return s.list.Get(0)
}
```

### 2. **Queue usando LinkedList**

```go
type Queue struct {
    list *LinkedList  // Com tail pointer
}

func (q *Queue) Enqueue(element int) {
    q.list.Add(element)  // Adiciona no final
}

func (q *Queue) Dequeue() (int, error) {
    return q.list.RemoveFirst()  // Remove do início
}
```

### 3. **LRU Cache usando Doubly LinkedList**

```go
type LRUCache struct {
    capacity int
    cache    map[int]*DoublyNode
    list     *DoublyLinkedList
}

func (lru *LRUCache) Get(key int) int {
    if node, exists := lru.cache[key]; exists {
        lru.moveToHead(node)  // O(1) com doubly linked
        return node.value
    }
    return -1
}
```

---

## Debugging e Problemas Comuns

### 1. **Memory Leaks**
```go
// PROBLEMA: Não limpar referências
func (list *LinkedList) Remove(index int) {
    // ... encontrar nó ...
    current.next = current.next.next
    // FALTOU: current.next.next = nil (em linguagens sem GC)
}

// SOLUÇÃO: Limpar referências
toRemove := current.next
current.next = toRemove.next
toRemove.next = nil  // Evita vazamentos
```

### 2. **Ponteiros Nulos**
```go
// PROBLEMA: Não verificar nil
func (list *LinkedList) Get(index int) int {
    current := list.head
    for i := 0; i < index; i++ {
        current = current.next  // Pode ser nil!
    }
    return current.data
}

// SOLUÇÃO: Sempre verificar
if current == nil {
    panic("índice fora dos limites")
}
```

### 3. **Loops Infinitos**
```go
// PROBLEMA: Lista circular acidental
node1.next = node2
node2.next = node1  // Cria loop!

// SOLUÇÃO: Algoritmo de detecção (Floyd's)
func (list *LinkedList) HasCycle() bool {
    slow, fast := list.head, list.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if slow == fast {
            return true
        }
    }
    return false
}
```

---

## Exercícios Práticos

### 1. **Implementação Básica**
a) Implemente LinkedList com todas operações
b) Adicione método ToString() para debug
c) Implemente Iterator para percorrer a lista

### 2. **Algoritmos Clássicos**
a) Reverse() - inverter lista
b) FindMiddle() - encontrar elemento do meio
c) RemoveDuplicates() - remover duplicatas
d) Merge() - mesclar duas listas ordenadas

### 3. **Otimizações**
a) Implemente Doubly LinkedList
b) Adicione tail pointer
c) Crie versão thread-safe
d) Implemente lazy deletion

### 4. **Aplicações**
a) Implemente Stack e Queue
b) Crie um LRU Cache
c) Desenvolva um editor de texto simples
d) Implemente undo/redo system

---

## Comparação com ArrayList

| Aspecto | LinkedList | ArrayList |
|---------|------------|----------|
| Acesso por índice | O(n) | O(1) |
| Inserção início | O(1) | O(n) |
| Inserção final | O(n) ou O(1)* | O(1)** |
| Inserção meio | O(n) | O(n) |
| Remoção início | O(1) | O(n) |
| Remoção final | O(n) ou O(1)* | O(1) |
| Remoção meio | O(n) | O(n) |
| Uso memória | Alto overhead | Pode desperdiçar |
| Cache performance | Ruim | Excelente |
| Implementação | Complexa | Simples |

*Com tail pointer  
**Amortizado

---

## Resumo

A **LinkedList** é uma estrutura fundamental que oferece:

### 🎯 **Pontos Fortes**
- Inserção/remoção O(1) no início
- Uso exato de memória
- Tamanho dinâmico ilimitado
- Base para outras estruturas

### ⚠️ **Limitações**
- Acesso sequencial O(n)
- Alto overhead de memória
- Cache performance ruim
- Complexidade de implementação

### 📚 **Quando Usar**
- Inserções/remoções frequentes no início
- Tamanho muito variável
- Implementação de Stack/Queue
- Memória limitada (sem desperdício)

A LinkedList é ideal quando a flexibilidade de inserção/remoção é mais importante que o acesso rápido por índice.