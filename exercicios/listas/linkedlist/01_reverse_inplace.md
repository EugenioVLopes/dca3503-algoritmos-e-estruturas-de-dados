# Função Reverse In-place para LinkedList

## Pergunta

Escreva uma função in-place para inverter a ordem de uma LinkedList.

```go
func (list *LinkedList) Reverse()

type LinkedList struct {
    head *Node
    size int
}

type Node struct {
    value int
    next  *Node
}
```

## Resposta

### Implementação

```go
func (list *LinkedList) Reverse() {
    if list.head == nil || list.head.next == nil {
        return // Lista vazia ou com apenas um elemento
    }

    var prev *Node = nil
    current := list.head

    for current != nil {
        next := current.next    // Salva o próximo nó
        current.next = prev     // Inverte o ponteiro
        prev = current          // Move prev para frente
        current = next          // Move current para frente
    }

    list.head = prev // Atualiza o head para o último nó processado
}
```

### Explicação Detalhada

**Algoritmo:** A função usa **três ponteiros** (`prev`, `current`, `next`) para inverter os ponteiros `next` de cada nó, percorrendo a lista uma vez e atualizando o `head` para o antigo último nó.

**Estratégia:**

1. **Salvar referência:** `next = current.next` (evita perder o resto da lista)
2. **Inverter ponteiro:** `current.next = prev` (aponta para o nó anterior)
3. **Avançar ponteiros:** Move `prev` e `current` para a próxima iteração
4. **Atualizar head:** No final, `prev` aponta para o novo primeiro nó

### Exemplo de Execução

**Lista inicial:** `1 → 2 → 3 → 4 → nil`

| Iteração    | prev | current | next | Ação          | Estado                      |
| ----------- | ---- | ------- | ---- | ------------- | --------------------------- |
| **Inicial** | nil  | 1       | -    | -             | `nil ← 1 → 2 → 3 → 4 → nil` |
| **1**       | nil  | 1       | 2    | Inverte 1→nil | `nil ← 1   2 → 3 → 4 → nil` |
| **2**       | 1    | 2       | 3    | Inverte 2→1   | `nil ← 1 ← 2   3 → 4 → nil` |
| **3**       | 2    | 3       | 4    | Inverte 3→2   | `nil ← 1 ← 2 ← 3   4 → nil` |
| **4**       | 3    | 4       | nil  | Inverte 4→3   | `nil ← 1 ← 2 ← 3 ← 4`       |
| **Final**   | 4    | nil     | -    | head = prev   | `4 → 3 → 2 → 1 → nil`       |

**Lista final:** `4 → 3 → 2 → 1 → nil` ✅

### Análise de Complexidade

- **Tempo:** **O(n)**

  - Percorre cada nó exatamente uma vez
  - Cada operação por nó é O(1)

- **Espaço:** **O(1)** (in-place)
  - Usa apenas 3 ponteiros auxiliares
  - Não aloca novos nós
  - Reutiliza a estrutura existente

### Casos Especiais

**Lista vazia:** `head = nil`

```go
if list.head == nil { return } // Retorna imediatamente
```

**Lista com 1 elemento:** `head → valor → nil`

```go
if list.head.next == nil { return } // Já está "invertida"
```

**Lista com 2 elementos:** `1 → 2 → nil`

- Iteração 1: `nil ← 1   2 → nil`
- Iteração 2: `nil ← 1 ← 2`
- Resultado: `2 → 1 → nil` ✅

### Visualização Passo a Passo

```
Estado inicial: [head] → 1 → 2 → 3 → nil
                prev=nil, current=1, next=?

Passo 1: next = current.next (next = 2)
         current.next = prev (1 → nil)
         prev = current (prev = 1)
         current = next (current = 2)

Estado: nil ← 1    [current] → 2 → 3 → nil
        prev=1, current=2, next=?

Passo 2: next = current.next (next = 3)
         current.next = prev (2 → 1)
         prev = current (prev = 2)
         current = next (current = 3)

Estado: nil ← 1 ← 2    [current] → 3 → nil
        prev=2, current=3, next=?

Passo 3: next = current.next (next = nil)
         current.next = prev (3 → 2)
         prev = current (prev = 3)
         current = next (current = nil)

Estado final: nil ← 1 ← 2 ← 3    current=nil
              prev=3

Atualizar head: list.head = prev (head = 3)
Resultado: [head] → 3 → 2 → 1 → nil
```

### Vantagens da Implementação

✅ **Eficiência de memória:** O(1) espaço
✅ **Performance:** O(n) tempo, ótimo para inversão
✅ **In-place:** Reutiliza nós existentes
✅ **Robustez:** Trata todos os casos especiais
✅ **Simplicidade:** Algoritmo clássico e bem estabelecido

### Comparação com Outras Abordagens

| Abordagem             | Tempo | Espaço | In-place | Complexidade |
| --------------------- | ----- | ------ | -------- | ------------ |
| Inversão de ponteiros | O(n)  | O(1)   | ✅       | Baixa        |
| Stack auxiliar        | O(n)  | O(n)   | ❌       | Média        |
| Recursão              | O(n)  | O(n)   | ✅       | Alta         |
| Array auxiliar        | O(n)  | O(n)   | ❌       | Baixa        |

### Teste da Função

```go
func main() {
    // Criar lista: 1 → 2 → 3 → 4 → nil
    list := &LinkedList{}
    list.head = &Node{1, &Node{2, &Node{3, &Node{4, nil}}}}
    list.size = 4

    fmt.Println("Antes:")  // 1 → 2 → 3 → 4 → nil
    printList(list)

    list.Reverse()

    fmt.Println("Depois:") // 4 → 3 → 2 → 1 → nil
    printList(list)
}
```
