# Função Reverse In-place para DoublyLinkedList

## Pergunta

Escreva uma função in-place para inverter a ordem de uma DoublyLinkedList.

```go
func (list *DoublyLinkedList) Reverse()

type DoublyLinkedList struct {
    head *Node2P
    tail *Node2P
    size int
}

type Node2P struct {
    prev  *Node2P
    value int
    next  *Node2P
}
```

## Resposta

### Implementação

```go
func (list *DoublyLinkedList) Reverse() {
    if list.head == nil || list.head.next == nil {
        return // Lista vazia ou com apenas um elemento
    }

    current := list.head

    // Percorre a lista trocando prev e next de cada nó
    for current != nil {
        // Troca os ponteiros prev e next do nó atual
        current.prev, current.next = current.next, current.prev
        current = current.prev // Move para o próximo (que agora está em prev)
    }

    // Troca head e tail
    list.head, list.tail = list.tail, list.head
}
```

### Explicação Detalhada

**Algoritmo:** A função percorre a lista **trocando os ponteiros `prev` e `next`** de cada nó usando atribuição múltipla, depois troca `head` e `tail` da lista.

**Estratégia:**

1. **Trocar ponteiros:** Para cada nó, `prev` e `next` são invertidos
2. **Navegação especial:** Após a troca, usa `current.prev` para avançar
3. **Atualizar extremidades:** Troca `head` e `tail` no final

**Por que `current.prev`?**
Após trocar `prev` e `next`, o que era `next` agora está em `prev`, então usamos `current.prev` para continuar a navegação.

### Exemplo de Execução

**Lista inicial:** `1 ⇄ 2 ⇄ 3 ⇄ 4`

```
head → [nil|1|→] ⇄ [←|2|→] ⇄ [←|3|→] ⇄ [←|4|nil] ← tail
```

| Iteração | current | Ação            | Estado após troca |
| -------- | ------- | --------------- | ----------------- | --- | --------- | --- | ------- | --- | --------- | --- | ----- |
| **1**    | nó 1    | Troca prev↔next | `[→               | 1   | nil] ⇄ [← | 2   | →] ⇄ [← | 3   | →] ⇄ [←   | 4   | nil]` |
| **2**    | nó 2    | Troca prev↔next | `[→               | 1   | nil] ⇄ [→ | 2   | ←] ⇄ [← | 3   | →] ⇄ [←   | 4   | nil]` |
| **3**    | nó 3    | Troca prev↔next | `[→               | 1   | nil] ⇄ [→ | 2   | ←] ⇄ [→ | 3   | ←] ⇄ [←   | 4   | nil]` |
| **4**    | nó 4    | Troca prev↔next | `[→               | 1   | nil] ⇄ [→ | 2   | ←] ⇄ [→ | 3   | ←] ⇄ [nil | 4   | ←]`   |

**Após trocar head e tail:**

```
head → [nil|4|←] ⇄ [→|3|←] ⇄ [→|2|←] ⇄ [→|1|nil] ← tail
```

**Lista final:** `4 ⇄ 3 ⇄ 2 ⇄ 1` ✅

### Análise de Complexidade

- **Tempo:** **O(n)**

  - Percorre cada nó exatamente uma vez
  - Cada troca de ponteiros é O(1)
  - Troca de head/tail é O(1)

- **Espaço:** **O(1)** (in-place)
  - Usa apenas a variável `current`
  - Não aloca novos nós
  - Reutiliza a estrutura existente

### Casos Especiais

**Lista vazia:** `head = nil, tail = nil`

```go
if list.head == nil { return } // Retorna imediatamente
```

**Lista com 1 elemento:** `head = tail = nó`

```go
if list.head.next == nil { return } // Já está "invertida"
```

**Lista com 2 elementos:** `1 ⇄ 2`

- Iteração 1: Troca ponteiros do nó 1
- Iteração 2: Troca ponteiros do nó 2
- Troca head/tail
- Resultado: `2 ⇄ 1` ✅

### Visualização Detalhada

```
Estado inicial:
head → [nil|1|ptr2] ⇄ [ptr1|2|ptr3] ⇄ [ptr2|3|nil] ← tail

Iteração 1 (current = nó 1):
Antes:  [nil|1|ptr2]
Troca:  prev=nil, next=ptr2 → prev=ptr2, next=nil
Depois: [ptr2|1|nil]
current = current.prev = ptr2 (nó 2)

Iteração 2 (current = nó 2):
Antes:  [ptr1|2|ptr3]
Troca:  prev=ptr1, next=ptr3 → prev=ptr3, next=ptr1
Depois: [ptr3|2|ptr1]
current = current.prev = ptr3 (nó 3)

Iteração 3 (current = nó 3):
Antes:  [ptr2|3|nil]
Troca:  prev=ptr2, next=nil → prev=nil, next=ptr2
Depois: [nil|3|ptr2]
current = current.prev = nil (fim do loop)

Troca head/tail:
head = antigo tail (nó 3)
tail = antigo head (nó 1)

Resultado final:
head → [nil|3|ptr2] ⇄ [ptr3|2|ptr1] ⇄ [ptr2|1|nil] ← tail
       3 ⇄ 2 ⇄ 1
```

### Vantagens da Implementação

✅ **Eficiência:** O(n) tempo, O(1) espaço
✅ **Elegância:** Usa atribuição múltipla do Go
✅ **In-place:** Reutiliza nós existentes
✅ **Simplicidade:** Menos variáveis que LinkedList simples
✅ **Robustez:** Mantém integridade bidirecional

### Comparação: DoublyLinkedList vs LinkedList

| Aspecto              | LinkedList              | DoublyLinkedList |
| -------------------- | ----------------------- | ---------------- |
| Variáveis auxiliares | 3 (prev, current, next) | 1 (current)      |
| Operações por nó     | 4 atribuições           | 1 troca múltipla |
| Atualização final    | Apenas head             | Head e tail      |
| Complexidade código  | Média                   | Baixa            |
| Navegação            | Unidirecional           | Bidirecional     |

### Vantagem da DoublyLinkedList

A implementação é **mais simples** que a LinkedList porque:

- Cada nó já tem ponteiros bidirecionais
- Não precisa "salvar" o próximo nó (está em `prev` após a troca)
- A atribuição múltipla torna o código mais limpo
- Manter `tail` permite navegação reversa eficiente

### Teste da Função

```go
func main() {
    // Criar lista: 1 ⇄ 2 ⇄ 3 ⇄ 4
    list := &DoublyLinkedList{}

    node1 := &Node2P{nil, 1, nil}
    node2 := &Node2P{nil, 2, nil}
    node3 := &Node2P{nil, 3, nil}
    node4 := &Node2P{nil, 4, nil}

    // Conectar nós
    node1.next, node2.prev = node2, node1
    node2.next, node3.prev = node3, node2
    node3.next, node4.prev = node4, node3

    list.head, list.tail = node1, node4
    list.size = 4

    fmt.Println("Antes:")  // 1 ⇄ 2 ⇄ 3 ⇄ 4
    printList(list)

    list.Reverse()

    fmt.Println("Depois:") // 4 ⇄ 3 ⇄ 2 ⇄ 1
    printList(list)
}
```
