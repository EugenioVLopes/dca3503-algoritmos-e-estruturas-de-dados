# Doubly LinkedList - Lista Duplamente Ligada

## Visão Geral

A **Doubly LinkedList** (Lista Duplamente Ligada) é uma evolução da LinkedList simples que adiciona um ponteiro adicional em cada nó, permitindo navegação bidirecional. Esta estrutura oferece maior flexibilidade para operações que requerem acesso tanto para frente quanto para trás.

### Características Principais:
- 🔄 **Navegação bidirecional**: Pode percorrer em ambas as direções
- ⚡ **Inserção/remoção O(1)**: Com referência ao nó
- 🎯 **Acesso ao final eficiente**: Com ponteiro tail
- 💾 **Maior uso de memória**: Ponteiro extra por nó

---

## Estrutura Interna

### Componentes:

#### DoublyNode (Nó Duplo):
```go
type DoublyNode struct {
    data int           // Valor armazenado
    next *DoublyNode   // Ponteiro para próximo nó
    prev *DoublyNode   // Ponteiro para nó anterior
}
```

#### DoublyLinkedList:
```go
type DoublyLinkedList struct {
    head *DoublyNode   // Ponteiro para primeiro nó
    tail *DoublyNode   // Ponteiro para último nó
    size int           // Número de elementos
}
```

### Representação Visual:

```
DoublyLinkedList:

head → [nil|10|•] ⇄ [•|20|•] ⇄ [•|30|•] ⇄ [•|40|nil] ← tail
        ↑                                      ↑
      Node1                                  Node4
      
size: 4

Legenda:
[prev|data|next] - estrutura do nó
⇄ - conexão bidirecional
• - ponteiro válido
nil - ponteiro nulo
```

### Comparação Visual com Singly LinkedList:

```
Singly LinkedList:
head → [10|•] → [20|•] → [30|•] → [40|nil]
       (apenas uma direção)

Doubly LinkedList:
head → [nil|10|•] ⇄ [•|20|•] ⇄ [•|30|•] ⇄ [•|40|nil] ← tail
       (navegação bidirecional)
```

---

## Algoritmos e Implementação

### 1. Inicialização

**Propósito**: Criar DoublyLinkedList vazia

**Pseudocódigo**:
```
ALGORITMO Init()
INÍCIO
    head ← NULL
    tail ← NULL
    size ← 0
FIM
```

**Implementação Go**:
```go
func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{
        head: nil,
        tail: nil,
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
    RETORNAR size
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) Size() int {
    return list.size
}
```

**Complexidade**: Θ(1)

---

### 3. AddFirst(element) - Inserção no Início

**Propósito**: Adicionar elemento no início da lista

**Pseudocódigo**:
```
ALGORITMO AddFirst(element)
INÍCIO
    novo_no ← criar_no(element)
    novo_no.next ← head
    novo_no.prev ← NULL
    
    SE head ≠ NULL ENTÃO
        head.prev ← novo_no
    SENÃO
        tail ← novo_no  // Lista estava vazia
    FIM_SE
    
    head ← novo_no
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) AddFirst(element int) {
    newNode := &DoublyNode{
        data: element,
        next: list.head,
        prev: nil,
    }
    
    if list.head != nil {
        list.head.prev = newNode
    } else {
        list.tail = newNode // Lista estava vazia
    }
    
    list.head = newNode
    list.size++
}
```

**Complexidade**: Θ(1)

**Visualização**:
```
Antes:  head → [nil|20|•] ⇄ [•|30|nil] ← tail
Inserir 10 no início:

1. Criar nó: [nil|10|nil]
2. Conectar: [nil|10|•] → [nil|20|•]
3. Conectar volta: [nil|10|•] ← [•|20|•]
4. Atualizar head: head → [nil|10|•] ⇄ [•|20|•] ⇄ [•|30|nil] ← tail
```

---

### 4. AddLast(element) - Inserção no Final

**Propósito**: Adicionar elemento no final da lista

**Pseudocódigo**:
```
ALGORITMO AddLast(element)
INÍCIO
    novo_no ← criar_no(element)
    novo_no.next ← NULL
    novo_no.prev ← tail
    
    SE tail ≠ NULL ENTÃO
        tail.next ← novo_no
    SENÃO
        head ← novo_no  // Lista estava vazia
    FIM_SE
    
    tail ← novo_no
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) AddLast(element int) {
    newNode := &DoublyNode{
        data: element,
        next: nil,
        prev: list.tail,
    }
    
    if list.tail != nil {
        list.tail.next = newNode
    } else {
        list.head = newNode // Lista estava vazia
    }
    
    list.tail = newNode
    list.size++
}
```

**Complexidade**: Θ(1)
**Vantagem**: Diferente da LinkedList simples, não precisa percorrer até o final!

---

### 5. Get(index) - Acesso por Índice

**Propósito**: Obter elemento em posição específica

**Pseudocódigo Otimizado**:
```
ALGORITMO Get(index)
INÍCIO
    SE index < 0 OU index >= size ENTÃO
        RETORNAR erro
    FIM_SE
    
    // Otimização: escolher direção mais próxima
    SE index < size/2 ENTÃO
        // Percorrer do início
        atual ← head
        PARA i DE 0 ATÉ index-1 FAÇA
            atual ← atual.next
        FIM_PARA
    SENÃO
        // Percorrer do final
        atual ← tail
        PARA i DE size-1 ATÉ index+1 FAÇA
            atual ← atual.prev
        FIM_PARA
    FIM_SE
    
    RETORNAR atual.data
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) Get(index int) (int, error) {
    if index < 0 || index >= list.size {
        return 0, fmt.Errorf("índice inválido: %d", index)
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
```

**Complexidade**: O(n/2) = O(n), mas 2x mais rápido que LinkedList simples!

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
    
    SE index = size ENTÃO
        AddLast(element)
        RETORNAR
    FIM_SE
    
    // Encontrar posição (otimizado)
    SE index < size/2 ENTÃO
        atual ← head
        PARA i DE 0 ATÉ index-1 FAÇA
            atual ← atual.next
        FIM_PARA
    SENÃO
        atual ← tail
        PARA i DE size-1 ATÉ index FAÇA
            atual ← atual.prev
        FIM_PARA
    FIM_SE
    
    // Inserir antes de 'atual'
    novo_no ← criar_no(element)
    novo_no.next ← atual
    novo_no.prev ← atual.prev
    atual.prev.next ← novo_no
    atual.prev ← novo_no
    
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) AddOnIndex(element int, index int) error {
    if index < 0 || index > list.size {
        return fmt.Errorf("índice inválido: %d", index)
    }
    
    if index == 0 {
        list.AddFirst(element)
        return nil
    }
    
    if index == list.size {
        list.AddLast(element)
        return nil
    }
    
    var current *DoublyNode
    
    // Encontrar posição otimizada
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
    newNode := &DoublyNode{
        data: element,
        next: current,
        prev: current.prev,
    }
    
    current.prev.next = newNode
    current.prev = newNode
    list.size++
    
    return nil
}
```

**Complexidade**: O(n/2) = O(n), mas otimizado!

---

### 7. RemoveNode(node) - Remoção por Referência

**Propósito**: Remover nó específico (vantagem única da Doubly LinkedList)

**Pseudocódigo**:
```
ALGORITMO RemoveNode(node)
INÍCIO
    SE node = NULL ENTÃO
        RETORNAR erro
    FIM_SE
    
    // Reconectar nó anterior
    SE node.prev ≠ NULL ENTÃO
        node.prev.next ← node.next
    SENÃO
        head ← node.next  // Removendo primeiro nó
    FIM_SE
    
    // Reconectar nó posterior
    SE node.next ≠ NULL ENTÃO
        node.next.prev ← node.prev
    SENÃO
        tail ← node.prev  // Removendo último nó
    FIM_SE
    
    size ← size - 1
    RETORNAR node.data
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) RemoveNode(node *DoublyNode) (int, error) {
    if node == nil {
        return 0, fmt.Errorf("nó inválido")
    }
    
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
    return node.data, nil
}
```

**Complexidade**: Θ(1) - Esta é a grande vantagem!
**Uso**: Ideal para implementar LRU Cache, onde você mantém referências aos nós.

---

### 8. RemoveFirst() - Remoção do Início

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
    
    SE head ≠ NULL ENTÃO
        head.prev ← NULL
    SENÃO
        tail ← NULL  // Lista ficou vazia
    FIM_SE
    
    size ← size - 1
    RETORNAR elemento_removido
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) RemoveFirst() (int, error) {
    if list.head == nil {
        return 0, fmt.Errorf("lista vazia")
    }
    
    removedData := list.head.data
    list.head = list.head.next
    
    if list.head != nil {
        list.head.prev = nil
    } else {
        list.tail = nil // Lista ficou vazia
    }
    
    list.size--
    return removedData, nil
}
```

**Complexidade**: Θ(1)

---

### 9. RemoveLast() - Remoção do Final

**Propósito**: Remover último elemento

**Pseudocódigo**:
```
ALGORITMO RemoveLast()
INÍCIO
    SE tail = NULL ENTÃO
        RETORNAR erro
    FIM_SE
    
    elemento_removido ← tail.data
    tail ← tail.prev
    
    SE tail ≠ NULL ENTÃO
        tail.next ← NULL
    SENÃO
        head ← NULL  // Lista ficou vazia
    FIM_SE
    
    size ← size - 1
    RETORNAR elemento_removido
FIM
```

**Implementação Go**:
```go
func (list *DoublyLinkedList) RemoveLast() (int, error) {
    if list.tail == nil {
        return 0, fmt.Errorf("lista vazia")
    }
    
    removedData := list.tail.data
    list.tail = list.tail.prev
    
    if list.tail != nil {
        list.tail.next = nil
    } else {
        list.head = nil // Lista ficou vazia
    }
    
    list.size--
    return removedData, nil
}
```

**Complexidade**: Θ(1)
**Vantagem**: LinkedList simples seria O(n) para esta operação!

---

## Análise de Complexidade Detalhada

### Resumo das Operações:

| Operação | Doubly LinkedList | Singly LinkedList | ArrayList |
|----------|-------------------|-------------------|----------|
| **Acesso** |
| Get(index) | O(n/2) | O(n) | O(1) |
| **Inserção** |
| AddFirst | Θ(1) | Θ(1) | O(n) |
| AddLast | Θ(1) | O(n) | O(1)* |
| AddOnIndex | O(n/2) | O(n) | O(n) |
| **Remoção** |
| RemoveFirst | Θ(1) | Θ(1) | O(n) |
| RemoveLast | Θ(1) | O(n) | O(1) |
| RemoveNode | Θ(1) | O(n) | N/A |
| Remove(index) | O(n/2) | O(n) | O(n) |
| **Memória** |
| Por elemento | 16 bytes | 12 bytes | 4 bytes |
| Overhead | Alto | Médio | Baixo |

*Amortizado

### Uso de Memória Detalhado:

**Por nó (64-bit)**:
```
DoublyNode = data + next + prev
           = 4 + 8 + 8 = 20 bytes
           
Overhead = 20/4 = 500% comparado ao valor puro
```

**Comparação de memória para 1000 elementos**:
- **ArrayList**: ~4KB (sem desperdício) a ~8KB (com capacidade extra)
- **LinkedList**: ~12KB
- **DoublyLinkedList**: ~20KB

---

## Vantagens da Doubly LinkedList

### ✅ **Navegação Bidirecional**
- **Percorrer para trás**: Útil para undo/redo, histórico
- **Acesso otimizado**: Escolhe direção mais próxima
- **Algoritmos mais eficientes**: Merge sort, etc.

### ✅ **Operações no Final O(1)**
- **AddLast()**: Não precisa percorrer
- **RemoveLast()**: Acesso direto via tail
- **Implementação de Deque**: Eficiente em ambas extremidades

### ✅ **Remoção por Referência O(1)**
- **RemoveNode()**: Única estrutura com esta capacidade
- **LRU Cache**: Implementação eficiente
- **Event listeners**: Remoção rápida de callbacks

### ✅ **Flexibilidade Máxima**
- **Inserção em qualquer posição**: Com referência ao nó
- **Algoritmos complexos**: Mais fáceis de implementar
- **Estruturas avançadas**: Base para outras implementações

---

## Desvantagens da Doubly LinkedList

### ❌ **Alto Uso de Memória**
- **Ponteiro extra**: 8 bytes por nó (64-bit)
- **500% overhead**: Comparado ao valor puro
- **Cache performance**: Ainda pior que ArrayList

### ❌ **Complexidade de Implementação**
- **Mais ponteiros**: Maior chance de erros
- **Casos especiais**: Lista vazia, um elemento
- **Manutenção**: Sempre atualizar prev e next

### ❌ **Acesso Sequencial**
- **Ainda O(n)**: Para acesso por índice
- **Sem acesso aleatório**: Como ArrayList
- **Algoritmos de busca**: Menos eficientes

---

## Casos de Uso Ideais

### 🎯 **Use Doubly LinkedList quando:**

#### 1. **Navegação Bidirecional Frequente**
```go
// Exemplo: Editor de texto com cursor
type TextEditor struct {
    lines *DoublyLinkedList
    currentLine *DoublyNode
}

func (editor *TextEditor) MoveCursorUp() {
    if editor.currentLine.prev != nil {
        editor.currentLine = editor.currentLine.prev
    }
}

func (editor *TextEditor) MoveCursorDown() {
    if editor.currentLine.next != nil {
        editor.currentLine = editor.currentLine.next
    }
}
```

#### 2. **LRU Cache Implementation**
```go
type LRUCache struct {
    capacity int
    cache    map[int]*DoublyNode
    list     *DoublyLinkedList
}

func (lru *LRUCache) Get(key int) int {
    if node, exists := lru.cache[key]; exists {
        // Move para início (O(1)!)
        lru.list.RemoveNode(node)
        lru.list.AddFirst(node.data)
        return node.data
    }
    return -1
}
```

#### 3. **Undo/Redo System**
```go
type UndoRedoManager struct {
    actions *DoublyLinkedList
    current *DoublyNode
}

func (manager *UndoRedoManager) Undo() {
    if manager.current.prev != nil {
        manager.current.prev.data.Undo()
        manager.current = manager.current.prev
    }
}

func (manager *UndoRedoManager) Redo() {
    if manager.current.next != nil {
        manager.current = manager.current.next
        manager.current.data.Execute()
    }
}
```

#### 4. **Music Playlist com Navegação**
```go
type MusicPlayer struct {
    playlist *DoublyLinkedList
    current  *DoublyNode
}

func (player *MusicPlayer) NextSong() {
    if player.current.next != nil {
        player.current = player.current.next
        player.PlaySong(player.current.data)
    }
}

func (player *MusicPlayer) PreviousSong() {
    if player.current.prev != nil {
        player.current = player.current.prev
        player.PlaySong(player.current.data)
    }
}
```

#### 5. **Browser History**
```go
type BrowserHistory struct {
    pages   *DoublyLinkedList
    current *DoublyNode
}

func (browser *BrowserHistory) Back() {
    if browser.current.prev != nil {
        browser.current = browser.current.prev
        browser.LoadPage(browser.current.data)
    }
}

func (browser *BrowserHistory) Forward() {
    if browser.current.next != nil {
        browser.current = browser.current.next
        browser.LoadPage(browser.current.data)
    }
}
```

---

## Algoritmos Especiais

### 1. **Reverse() - Inversão da Lista**

```go
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
```

**Complexidade**: O(n)
**Vantagem**: Mais simples que na LinkedList simples!

### 2. **FindFromMiddle() - Busca Otimizada**

```go
func (list *DoublyLinkedList) FindFromMiddle(value int) *DoublyNode {
    if list.size == 0 {
        return nil
    }
    
    // Buscar simultaneamente do início e do fim
    front := list.head
    back := list.tail
    
    for front != back && front.prev != back {
        if front.data == value {
            return front
        }
        if back.data == value {
            return back
        }
        
        front = front.next
        back = back.prev
    }
    
    // Verificar nó do meio se necessário
    if front.data == value {
        return front
    }
    
    return nil
}
```

**Complexidade**: O(n/2) - 2x mais rápido!

### 3. **MergeSorted() - Mesclar Listas Ordenadas**

```go
func MergeSortedDoubly(list1, list2 *DoublyLinkedList) *DoublyLinkedList {
    result := NewDoublyLinkedList()
    
    current1 := list1.head
    current2 := list2.head
    
    for current1 != nil && current2 != nil {
        if current1.data <= current2.data {
            result.AddLast(current1.data)
            current1 = current1.next
        } else {
            result.AddLast(current2.data)
            current2 = current2.next
        }
    }
    
    // Adicionar elementos restantes
    for current1 != nil {
        result.AddLast(current1.data)
        current1 = current1.next
    }
    
    for current2 != nil {
        result.AddLast(current2.data)
        current2 = current2.next
    }
    
    return result
}
```

---

## Implementações Avançadas

### 1. **Circular Doubly LinkedList**

```go
type CircularDoublyLinkedList struct {
    head *DoublyNode
    size int
}

// No circular, tail.next = head e head.prev = tail
func (list *CircularDoublyLinkedList) AddFirst(element int) {
    newNode := &DoublyNode{data: element}
    
    if list.head == nil {
        newNode.next = newNode
        newNode.prev = newNode
        list.head = newNode
    } else {
        tail := list.head.prev
        newNode.next = list.head
        newNode.prev = tail
        list.head.prev = newNode
        tail.next = newNode
        list.head = newNode
    }
    
    list.size++
}
```

### 2. **Thread-Safe Doubly LinkedList**

```go
type SafeDoublyLinkedList struct {
    list *DoublyLinkedList
    mutex sync.RWMutex
}

func (safe *SafeDoublyLinkedList) AddFirst(element int) {
    safe.mutex.Lock()
    defer safe.mutex.Unlock()
    safe.list.AddFirst(element)
}

func (safe *SafeDoublyLinkedList) Get(index int) (int, error) {
    safe.mutex.RLock()
    defer safe.mutex.RUnlock()
    return safe.list.Get(index)
}
```

---

## Comparação Completa

### Performance Benchmark (1M elementos):

| Operação | ArrayList | LinkedList | DoublyLinkedList |
|----------|-----------|------------|------------------|
| **Inserção final** | 2ms | 800ms | 3ms |
| **Inserção início** | 500ms | 1ms | 1ms |
| **Acesso meio** | 0.001ms | 250ms | 125ms |
| **Remoção final** | 0.001ms | 800ms | 0.001ms |
| **Remoção início** | 500ms | 0.001ms | 0.001ms |
| **Iteração completa** | 2ms | 50ms | 80ms |

### Uso de Memória (1M elementos):

| Estrutura | Memória Total | Overhead |
|-----------|---------------|----------|
| **ArrayList** | ~4-8MB | 0-100% |
| **LinkedList** | ~12MB | 200% |
| **DoublyLinkedList** | ~20MB | 400% |

---

## Exercícios Práticos

### 1. **Implementação Básica**
a) Implemente DoublyLinkedList completa
b) Adicione método ToString() para debug
c) Implemente Iterator bidirecional

### 2. **Algoritmos Avançados**
a) Implemente MergeSort para DoublyLinkedList
b) Crie algoritmo de busca bidirecional
c) Implemente rotação da lista
d) Crie método para detectar palíndromos

### 3. **Aplicações Práticas**
a) Implemente LRU Cache completo
b) Crie sistema de undo/redo
c) Desenvolva editor de texto simples
d) Implemente playlist de música

### 4. **Otimizações**
a) Implemente versão circular
b) Crie versão thread-safe
c) Adicione pool de nós
d) Implemente lazy deletion

---

## Resumo

A **Doubly LinkedList** é uma estrutura poderosa que oferece:

### 🎯 **Pontos Fortes**
- Navegação bidirecional eficiente
- Operações O(1) em ambas extremidades
- Remoção por referência O(1)
- Base ideal para estruturas complexas

### ⚠️ **Limitações**
- Alto uso de memória (400% overhead)
- Complexidade de implementação
- Acesso sequencial para índices
- Cache performance ruim

### 📚 **Quando Usar**
- Navegação bidirecional frequente
- LRU Cache e estruturas similares
- Undo/Redo systems
- Quando você mantém referências aos nós
- Implementação de Deque eficiente

### 🔄 **Comparação Rápida**
- **vs ArrayList**: Mais flexível, mas usa mais memória
- **vs LinkedList**: Mais funcionalidades, mas maior overhead
- **vs Deque**: Implementação base para Deque eficiente

A Doubly LinkedList é a escolha ideal quando você precisa de máxima flexibilidade de navegação e não se importa com o overhead de memória adicional!