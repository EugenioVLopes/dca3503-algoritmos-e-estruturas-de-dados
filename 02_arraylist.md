# ArrayList - Lista Baseada em Array Dinâmico

## Visão Geral

O **ArrayList** é uma implementação da Lista ADT baseada em um **array dinâmico**. É uma das estruturas de dados mais utilizadas devido ao seu excelente desempenho para acesso aleatório e simplicidade de implementação.

### Características Principais:
- 🎯 **Acesso direto**: O(1) por índice
- 📈 **Redimensionamento automático**: Cresce conforme necessário
- 💾 **Localidade de memória**: Elementos contíguos
- ⚡ **Cache-friendly**: Excelente performance de iteração

---

## Estrutura Interna

### Componentes:

```go
type ArrayList struct {
    data     []int  // Array interno que armazena os elementos
    size     int    // Número atual de elementos
    capacity int    // Capacidade total do array
}
```

### Representação Visual:

```
Capacidade: 8
Tamanho: 5

Índices:  [0] [1] [2] [3] [4] [5] [6] [7]
Array:    [10][20][30][40][50][ ][ ][ ]
           ↑                   ↑
        Elementos           Espaço livre
```

---

## Algoritmos e Implementação

### 1. Inicialização

**Propósito**: Criar ArrayList com capacidade inicial

**Pseudocódigo**:
```
ALGORITMO Init(capacidade_inicial)
INÍCIO
    data ← criar_array(capacidade_inicial)
    size ← 0
    capacity ← capacidade_inicial
FIM
```

**Implementação Go**:
```go
func NewArrayList(initialCapacity int) *ArrayList {
    return &ArrayList{
        data:     make([]int, initialCapacity),
        size:     0,
        capacity: initialCapacity,
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
func (list *ArrayList) Size() int {
    return list.size
}
```

**Complexidade**: Θ(1)
**Explicação**: Acesso direto ao campo size.

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
    
    RETORNAR data[index]
FIM
```

**Implementação Go**:
```go
func (list *ArrayList) Get(index int) (int, error) {
    if index < 0 || index >= list.size {
        return 0, fmt.Errorf("índice inválido: %d", index)
    }
    return list.data[index], nil
}
```

**Complexidade**: Θ(1)
**Explicação**: Acesso direto por índice em array.

---

### 4. Resize() - Redimensionamento

**Propósito**: Aumentar capacidade quando necessário

**Pseudocódigo**:
```
ALGORITMO Resize(nova_capacidade)
INÍCIO
    novo_array ← criar_array(nova_capacidade)
    
    PARA i DE 0 ATÉ size-1 FAÇA
        novo_array[i] ← data[i]
    FIM_PARA
    
    data ← novo_array
    capacity ← nova_capacidade
FIM
```

**Implementação Go**:
```go
func (list *ArrayList) resize(newCapacity int) {
    newData := make([]int, newCapacity)
    copy(newData, list.data[:list.size])
    list.data = newData
    list.capacity = newCapacity
}
```

**Complexidade**: Θ(n)
**Explicação**: Precisa copiar todos os elementos existentes.

---

### 5. Add(element) - Inserção no Final

**Propósito**: Adicionar elemento no final da lista

**Pseudocódigo**:
```
ALGORITMO Add(element)
INÍCIO
    SE size = capacity ENTÃO
        Resize(capacity * 2)  // Dobra a capacidade
    FIM_SE
    
    data[size] ← element
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *ArrayList) Add(element int) {
    if list.size == list.capacity {
        list.resize(list.capacity * 2)
    }
    
    list.data[list.size] = element
    list.size++
}
```

**Complexidade**:
- **Melhor caso**: Θ(1) - quando há espaço
- **Pior caso**: O(n) - quando precisa redimensionar
- **Amortizado**: O(1) - redimensionamento é raro

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
    
    SE size = capacity ENTÃO
        Resize(capacity * 2)
    FIM_SE
    
    // Deslocar elementos para direita
    PARA i DE size ATÉ index+1 FAÇA
        data[i] ← data[i-1]
    FIM_PARA
    
    data[index] ← element
    size ← size + 1
FIM
```

**Implementação Go**:
```go
func (list *ArrayList) AddOnIndex(element int, index int) error {
    if index < 0 || index > list.size {
        return fmt.Errorf("índice inválido: %d", index)
    }
    
    if list.size == list.capacity {
        list.resize(list.capacity * 2)
    }
    
    // Desloca elementos para direita
    for i := list.size; i > index; i-- {
        list.data[i] = list.data[i-1]
    }
    
    list.data[index] = element
    list.size++
    return nil
}
```

**Complexidade**:
- **Melhor caso**: Θ(1) - inserção no final
- **Pior caso**: O(n) - inserção no início
- **Caso médio**: O(n/2) = O(n)

**Visualização**:
```
Antes:  [10][20][30][40][ ][ ]
Inserir 25 no índice 2

1. Deslocar: [10][20][ ][30][40][ ]
2. Inserir:  [10][20][25][30][40][ ]
```

---

### 7. Remove(index) - Remoção por Posição

**Propósito**: Remover elemento de posição específica

**Pseudocódigo**:
```
ALGORITMO Remove(index)
INÍCIO
    SE index < 0 OU index >= size ENTÃO
        RETORNAR erro
    FIM_SE
    
    elemento_removido ← data[index]
    
    // Deslocar elementos para esquerda
    PARA i DE index ATÉ size-2 FAÇA
        data[i] ← data[i+1]
    FIM_PARA
    
    size ← size - 1
    RETORNAR elemento_removido
FIM
```

**Implementação Go**:
```go
func (list *ArrayList) Remove(index int) (int, error) {
    if index < 0 || index >= list.size {
        return 0, fmt.Errorf("índice inválido: %d", index)
    }
    
    removedElement := list.data[index]
    
    // Desloca elementos para esquerda
    for i := index; i < list.size-1; i++ {
        list.data[i] = list.data[i+1]
    }
    
    list.size--
    return removedElement, nil
}
```

**Complexidade**:
- **Melhor caso**: Θ(1) - remoção do final
- **Pior caso**: O(n) - remoção do início
- **Caso médio**: O(n/2) = O(n)

---

## Análise de Complexidade Detalhada

### Resumo das Operações:

| Operação | Melhor Caso | Caso Médio | Pior Caso | Espaço |
|----------|-------------|------------|-----------|--------|
| Get(index) | Θ(1) | Θ(1) | Θ(1) | O(1) |
| Add(final) | Θ(1) | Θ(1) | O(n)* | O(1) |
| AddOnIndex | Θ(1) | O(n) | O(n) | O(1) |
| Remove | Θ(1) | O(n) | O(n) | O(1) |
| Size | Θ(1) | Θ(1) | Θ(1) | O(1) |

*O(n) apenas quando redimensiona (raro)

### Análise Amortizada do Add():

**Sequência de inserções com redimensionamento**:
```
Capacidade inicial: 4

Inserção 1-4: O(1) cada = 4 operações
Inserção 5: O(4) para redimensionar + O(1) = 5 operações
Inserção 6-8: O(1) cada = 3 operações
Inserção 9: O(8) para redimensionar + O(1) = 9 operações

Total para 9 inserções: 4 + 5 + 3 + 9 = 21 operações
Média: 21/9 ≈ 2.33 = O(1) amortizado
```

---

## Estratégias de Redimensionamento

### 1. **Dobrar Capacidade** (Padrão)
```go
newCapacity = oldCapacity * 2
```
**Vantagens**: Simples, O(1) amortizado
**Desvantagens**: Pode desperdiçar até 50% da memória

### 2. **Fator 1.5**
```go
newCapacity = oldCapacity + oldCapacity/2
```
**Vantagens**: Menos desperdício de memória
**Desvantagens**: Mais redimensionamentos

### 3. **Crescimento Linear**
```go
newCapacity = oldCapacity + BLOCK_SIZE
```
**Vantagens**: Controle preciso da memória
**Desvantagens**: Não é O(1) amortizado

### 4. **Shrinking (Encolhimento)**
```go
if size < capacity/4 {
    resize(capacity/2)
}
```
**Objetivo**: Liberar memória não utilizada

---

## Vantagens do ArrayList

### ✅ **Performance**
- **Acesso O(1)**: Direto por índice
- **Cache-friendly**: Elementos contíguos na memória
- **Iteração rápida**: Localidade de referência excelente
- **Baixo overhead**: Apenas alguns campos extras

### ✅ **Simplicidade**
- **Implementação direta**: Baseado em conceitos simples
- **Debugging fácil**: Estado visível e previsível
- **Menos ponteiros**: Reduz chance de erros

### ✅ **Flexibilidade**
- **Redimensionamento automático**: Cresce conforme necessário
- **Acesso bidirecional**: Frente para trás e vice-versa
- **Suporte a algoritmos**: Ordenação, busca binária, etc.

---

## Desvantagens do ArrayList

### ❌ **Inserção/Remoção no Meio**
- **O(n) complexidade**: Precisa deslocar elementos
- **Operação custosa**: Especialmente para listas grandes

### ❌ **Uso de Memória**
- **Desperdício**: Pode ter espaços não utilizados
- **Redimensionamento**: Picos temporários de uso
- **Fragmentação**: Em sistemas com pouca memória

### ❌ **Inserções Frequentes**
- **Redimensionamentos**: Podem causar pausas
- **Cópia de dados**: Operação custosa

---

## Casos de Uso Ideais

### 🎯 **Use ArrayList quando:**

1. **Acesso frequente por índice**
   ```go
   // Exemplo: Acessar elemento aleatório
   element := list.Get(randomIndex)
   ```

2. **Iteração sobre todos elementos**
   ```go
   // Exemplo: Somar todos elementos
   for i := 0; i < list.Size(); i++ {
       sum += list.Get(i)
   }
   ```

3. **Inserções principalmente no final**
   ```go
   // Exemplo: Log de eventos
   eventLog.Add(newEvent)
   ```

4. **Busca binária** (em lista ordenada)
   ```go
   // Exemplo: Buscar em lista ordenada
   index := binarySearch(sortedList, target)
   ```

5. **Algoritmos que precisam de acesso aleatório**
   ```go
   // Exemplo: Algoritmo de ordenação
   quickSort(list, 0, list.Size()-1)
   ```

---

## Otimizações Avançadas

### 1. **Capacidade Inicial Inteligente**
```go
func NewArrayListWithHint(expectedSize int) *ArrayList {
    initialCapacity := nextPowerOfTwo(expectedSize)
    return NewArrayList(initialCapacity)
}
```

### 2. **Lazy Shrinking**
```go
func (list *ArrayList) Remove(index int) {
    // ... remoção normal ...
    
    // Só encolhe se muito vazio
    if list.size < list.capacity/4 && list.capacity > MINIMUM_CAPACITY {
        list.resize(list.capacity / 2)
    }
}
```

### 3. **Batch Operations**
```go
func (list *ArrayList) AddAll(elements []int) {
    requiredCapacity := list.size + len(elements)
    if requiredCapacity > list.capacity {
        list.resize(max(requiredCapacity, list.capacity*2))
    }
    
    copy(list.data[list.size:], elements)
    list.size += len(elements)
}
```

---

## Variações do ArrayList

### 1. **Circular ArrayList**
- **Uso**: Quando inserções/remoções são frequentes no início
- **Técnica**: Usar índices head e tail

### 2. **Segmented ArrayList**
- **Uso**: Para listas muito grandes
- **Técnica**: Dividir em blocos menores

### 3. **Copy-on-Write ArrayList**
- **Uso**: Quando leitura é muito mais frequente que escrita
- **Técnica**: Compartilhar array até modificação

---

## Exercícios Práticos

### 1. **Implementação Básica**
Implemente um ArrayList com todas as operações básicas.

### 2. **Otimizações**
a) Adicione shrinking automático
b) Implemente AddAll() eficiente
c) Crie construtor com capacidade inicial inteligente

### 3. **Análise**
a) Meça performance de inserção vs LinkedList
b) Compare uso de memória com diferentes fatores de crescimento
c) Analise cache misses em diferentes padrões de acesso

### 4. **Aplicações**
a) Implemente um buffer circular usando ArrayList
b) Crie um sistema de undo/redo
c) Desenvolva um algoritmo de ordenação in-place

---

## Comparação com Outras Estruturas

| Aspecto | ArrayList | LinkedList | Array Estático |
|---------|-----------|------------|----------------|
| Acesso por índice | O(1) | O(n) | O(1) |
| Inserção final | O(1)* | O(n) | O(1) |
| Inserção início | O(n) | O(1) | O(n) |
| Remoção final | O(1) | O(n) | O(1) |
| Remoção início | O(n) | O(1) | O(n) |
| Uso de memória | Médio | Alto | Baixo |
| Cache performance | Excelente | Ruim | Excelente |
| Flexibilidade | Alta | Alta | Baixa |

*Amortizado

---

## Resumo

O **ArrayList** é uma estrutura fundamental que oferece:

### 🎯 **Pontos Fortes**
- Acesso rápido por índice O(1)
- Excelente performance de cache
- Implementação simples e robusta
- Redimensionamento automático

### ⚠️ **Limitações**
- Inserção/remoção no meio é custosa O(n)
- Pode desperdiçar memória
- Redimensionamentos podem causar pausas

### 📚 **Quando Usar**
- Acesso frequente por índice
- Iteração sobre elementos
- Inserções principalmente no final
- Algoritmos que precisam de acesso aleatório

O ArrayList é a escolha padrão para a maioria dos casos de uso de listas, sendo superado apenas quando inserções/remoções no meio são muito frequentes.