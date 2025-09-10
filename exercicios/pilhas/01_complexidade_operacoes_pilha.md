# Complexidade de Tempo das Operações de Pilha

## Pergunta

Considere as estruturas de dados apresentadas na tabela a seguir e responda o desempenho de tempo de pior caso e melhor caso para cada operação listada.

| Operação            | ArrayStack |             | LinkedListStack |             |
| ------------------- | ---------- | ----------- | --------------- | ----------- |
|                     | Pior Caso  | Melhor Caso | Pior Caso       | Melhor Caso |
| Push(value int)     | O(n)       | Ω(1)        | O(1)            | Ω(1)        |
| Pop() (int, error)  | O(1)       | Ω(1)        | O(1)            | Ω(1)        |
| Peek() (int, error) | O(1)       | Ω(1)        | O(1)            | Ω(1)        |
| Size()              | O(1)       | Ω(1)        | O(1)            | Ω(1)        |

## Resposta

### ArrayStack (Pilha baseada em Array)

#### Push(value int)

- **Pior caso O(n)**: Quando o array está cheio e precisa ser redimensionado
- **Melhor caso Ω(1)**: Quando há espaço disponível no array

**Explicação:**

```go
func (s *ArrayStack) Push(value int) {
    if s.size >= len(s.items) {
        // Pior caso: precisa redimensionar O(n)
        newItems := make([]int, len(s.items)*2)
        copy(newItems, s.items)  // Copia todos os elementos
        s.items = newItems
    }
    // Melhor caso: inserção direta O(1)
    s.items[s.size] = value
    s.size++
}
```

#### Pop() (int, error)

- **Pior caso O(1)**: Sempre constante
- **Melhor caso Ω(1)**: Sempre constante

**Explicação:**

```go
func (s *ArrayStack) Pop() (int, error) {
    if s.size == 0 {
        return 0, errors.New("pilha vazia")
    }
    s.size--
    return s.items[s.size], nil  // Acesso direto O(1)
}
```

#### Peek() (int, error)

- **Pior caso O(1)**: Acesso direto ao topo
- **Melhor caso Ω(1)**: Acesso direto ao topo

**Explicação:**

```go
func (s *ArrayStack) Peek() (int, error) {
    if s.size == 0 {
        return 0, errors.New("pilha vazia")
    }
    return s.items[s.size-1], nil  // Acesso direto O(1)
}
```

#### Size()

- **Pior caso O(1)**: Retorna variável armazenada
- **Melhor caso Ω(1)**: Retorna variável armazenada

**Explicação:**

```go
func (s *ArrayStack) Size() int {
    return s.size  // Acesso direto à variável O(1)
}
```

### LinkedListStack (Pilha baseada em Lista Ligada)

#### Push(value int)

- **Pior caso O(1)**: Inserção no início da lista
- **Melhor caso Ω(1)**: Inserção no início da lista

**Explicação:**

```go
func (s *LinkedListStack) Push(value int) {
    newNode := &Node{value: value, next: s.head}
    s.head = newNode  // Sempre O(1)
    s.size++
}
```

#### Pop() (int, error)

- **Pior caso O(1)**: Remoção do início da lista
- **Melhor caso Ω(1)**: Remoção do início da lista

**Explicação:**

```go
func (s *LinkedListStack) Pop() (int, error) {
    if s.head == nil {
        return 0, errors.New("pilha vazia")
    }
    value := s.head.value
    s.head = s.head.next  // Sempre O(1)
    s.size--
    return value, nil
}
```

#### Peek() (int, error)

- **Pior caso O(1)**: Acesso ao primeiro nó
- **Melhor caso Ω(1)**: Acesso ao primeiro nó

**Explicação:**

```go
func (s *LinkedListStack) Peek() (int, error) {
    if s.head == nil {
        return 0, errors.New("pilha vazia")
    }
    return s.head.value, nil  // Acesso direto O(1)
}
```

#### Size()

- **Pior caso O(1)**: Retorna variável armazenada
- **Melhor caso Ω(1)**: Retorna variável armazenada

**Explicação:**

```go
func (s *LinkedListStack) Size() int {
    return s.size  // Acesso direto à variável O(1)
}
```

### Comparação Resumida

| Operação    | ArrayStack           | LinkedListStack | Vantagem        |
| ----------- | -------------------- | --------------- | --------------- |
| **Push**    | O(n) amortizado O(1) | O(1) sempre     | LinkedListStack |
| **Pop**     | O(1)                 | O(1)            | Empate          |
| **Peek**    | O(1)                 | O(1)            | Empate          |
| **Size**    | O(1)                 | O(1)            | Empate          |
| **Memória** | Melhor localidade    | Mais overhead   | ArrayStack      |

### Análise Detalhada

#### ArrayStack

**Vantagens:**

- ✅ Melhor localidade de memória (cache-friendly)
- ✅ Menor overhead de memória por elemento
- ✅ Operações geralmente mais rápidas na prática

**Desvantagens:**

- ❌ Push pode ser O(n) quando precisa redimensionar
- ❌ Pode desperdiçar memória se a pilha encolher
- ❌ Tamanho máximo pode ser limitado

#### LinkedListStack

**Vantagens:**

- ✅ Push sempre O(1) garantido
- ✅ Usa apenas a memória necessária
- ✅ Sem limite de tamanho (exceto memória disponível)

**Desvantagens:**

- ❌ Overhead de ponteiros (8 bytes por nó em 64-bit)
- ❌ Pior localidade de memória
- ❌ Alocação/desalocação frequente de nós

### Complexidade Amortizada

**ArrayStack Push:**

- Embora o pior caso seja O(n), a **análise amortizada** mostra que é O(1)
- Redimensionamentos são raros e o custo é "distribuído" entre muitas operações
- Na prática, a maioria dos pushes são O(1)

**Exemplo de análise amortizada:**

```
Sequência de 8 pushes em array inicial de tamanho 1:
Push 1: O(1) - sem redimensionamento
Push 2: O(2) - redimensiona para 2
Push 3: O(3) - redimensiona para 4
Push 4: O(1) - sem redimensionamento
Push 5: O(5) - redimensiona para 8
Push 6: O(1) - sem redimensionamento
Push 7: O(1) - sem redimensionamento
Push 8: O(1) - sem redimensionamento

Custo total: 1+2+3+1+5+1+1+1 = 15
Custo amortizado: 15/8 = 1.875 ≈ O(1)
```

### Quando Usar Cada Implementação

**Use ArrayStack quando:**

- Performance é crítica
- Tamanho da pilha é relativamente previsível
- Memória é limitada
- Acesso frequente aos elementos

**Use LinkedListStack quando:**

- Tamanho da pilha varia muito
- Garantia de O(1) para push é essencial
- Memória é abundante
- Simplicidade de implementação é prioridade

### Implementações Exemplo

#### ArrayStack

```go
type ArrayStack struct {
    items []int
    size  int
}

func NewArrayStack() *ArrayStack {
    return &ArrayStack{
        items: make([]int, 1),
        size:  0,
    }
}
```

#### LinkedListStack

```go
type Node struct {
    value int
    next  *Node
}

type LinkedListStack struct {
    head *Node
    size int
}

func NewLinkedListStack() *LinkedListStack {
    return &LinkedListStack{head: nil, size: 0}
}
```
