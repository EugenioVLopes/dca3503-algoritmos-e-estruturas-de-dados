# Stack (Pilha) - Tipo Abstrato de Dados

## Introdução

Uma **pilha** (stack) é uma estrutura de dados linear que segue o princípio **LIFO** (Last In, First Out - "último a entrar, primeiro a sair").

## Características Principais

### Princípio LIFO
- O último elemento inserido é o primeiro a ser removido
- Acesso restrito: apenas o elemento do topo pode ser acessado diretamente
- Operações principais concentradas em uma extremidade (topo)

### Operações Fundamentais

1. **Push**: Adiciona elemento no topo da pilha
2. **Pop**: Remove e retorna o elemento do topo
3. **Peek/Top**: Consulta o elemento do topo sem removê-lo
4. **IsEmpty**: Verifica se a pilha está vazia
5. **Size**: Retorna o número de elementos

## Interface Stack

```go
type Stack interface {
    // Operações básicas
    Push(element int)           // Adiciona no topo
    Pop() (int, error)          // Remove do topo
    Peek() (int, error)         // Consulta o topo
    
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

### 1. ArrayStack (Baseada em Array)

**Características:**
- Usa array dinâmico (slice) interno
- Redimensionamento automático
- Elementos armazenados contiguamente na memória

**Vantagens:**
- Acesso rápido aos elementos: O(1)
- Uso eficiente de memória (sem overhead de ponteiros)
- Cache-friendly (localidade espacial)
- Push/Pop amortizado O(1)

**Desvantagens:**
- Redimensionamento ocasional pode ser custoso: O(n)
- Pode desperdiçar memória se a capacidade for muito maior que o uso
- Tamanho máximo limitado pela memória contígua disponível

**Complexidades:**
```
Push:     O(1) amortizado, O(n) no pior caso (redimensionamento)
Pop:      O(1)
Peek:     O(1)
IsEmpty:  O(1)
Size:     O(1)
```

### 2. LinkedStack (Baseada em Lista Ligada)

**Características:**
- Usa nós ligados por ponteiros
- Crescimento dinâmico sem limite predefinido
- Cada elemento tem overhead de ponteiro

**Vantagens:**
- Push/Pop sempre O(1) (sem redimensionamento)
- Não há limite de capacidade (exceto memória disponível)
- Não desperdiça memória (aloca exatamente o necessário)
- Flexibilidade total de tamanho

**Desvantagens:**
- Overhead de memória por ponteiros
- Menor localidade espacial (cache menos eficiente)
- Fragmentação de memória
- Acesso sequencial aos elementos

**Complexidades:**
```
Push:     O(1) sempre
Pop:      O(1) sempre
Peek:     O(1)
IsEmpty:  O(1)
Size:     O(1)
```

## Comparação das Implementações

| Aspecto | ArrayStack | LinkedStack |
|---------|------------|-------------|
| **Complexidade Push** | O(1) amortizado | O(1) sempre |
| **Complexidade Pop** | O(1) | O(1) |
| **Uso de Memória** | Mais eficiente | Overhead de ponteiros |
| **Cache Performance** | Melhor | Pior |
| **Limite de Tamanho** | Memória contígua | Apenas memória total |
| **Redimensionamento** | Automático | Não necessário |
| **Fragmentação** | Baixa | Pode ocorrer |

## Aplicações Práticas

### 1. Verificação de Parênteses Balanceados
```go
func IsValidParentheses(s string) bool {
    stack := NewArrayStack(len(s))
    
    for _, char := range s {
        switch char {
        case '(', '[', '{':
            stack.Push(int(char))
        case ')', ']', '}':
            if stack.IsEmpty() {
                return false
            }
            top, _ := stack.Pop()
            if !isMatchingPair(rune(top), char) {
                return false
            }
        }
    }
    
    return stack.IsEmpty()
}
```

### 2. Avaliação de Expressões Pós-fixas
```go
// Exemplo: "3 4 + 2 *" = (3 + 4) * 2 = 14
func EvaluatePostfix(expression []string) (int, error) {
    stack := NewArrayStack(len(expression))
    
    for _, token := range expression {
        switch token {
        case "+":
            b, _ := stack.Pop()
            a, _ := stack.Pop()
            stack.Push(a + b)
        case "-":
            b, _ := stack.Pop()
            a, _ := stack.Pop()
            stack.Push(a - b)
        // ... outros operadores
        default:
            // Número
            var num int
            fmt.Sscanf(token, "%d", &num)
            stack.Push(num)
        }
    }
    
    return stack.Pop()
}
```

### 3. Chamadas de Função (Call Stack)
- Cada chamada de função é empilhada
- Retorno remove a função do topo
- Recursão usa implicitamente uma pilha

### 4. Desfazer/Refazer (Undo/Redo)
- Pilha de ações para desfazer
- Pilha de ações para refazer
- Editores de texto, IDEs

### 5. Navegação em Browsers
- Histórico de páginas visitadas
- Botão "Voltar" usa pilha

## Algoritmos com Pilhas

### Inversão de String
```go
func ReverseString(s string) string {
    stack := NewArrayStack(len(s))
    
    // Push todos os caracteres
    for _, char := range s {
        stack.Push(int(char))
    }
    
    // Pop para formar string invertida
    var result strings.Builder
    for !stack.IsEmpty() {
        char, _ := stack.Pop()
        result.WriteRune(rune(char))
    }
    
    return result.String()
}
```

### Conversão Decimal para Binário
```go
func DecimalToBinary(n int) string {
    if n == 0 {
        return "0"
    }
    
    stack := NewArrayStack(32) // Suficiente para int32
    
    for n > 0 {
        stack.Push(n % 2)
        n /= 2
    }
    
    var result strings.Builder
    for !stack.IsEmpty() {
        digit, _ := stack.Pop()
        result.WriteString(fmt.Sprintf("%d", digit))
    }
    
    return result.String()
}
```

## Considerações de Design

### Quando Usar ArrayStack
- Tamanho previsível ou limitado
- Performance crítica
- Uso intensivo de cache
- Memória limitada

### Quando Usar LinkedStack
- Tamanho muito variável
- Sem limite de capacidade
- Operações sempre O(1) necessárias
- Flexibilidade máxima

### Tratamento de Erros
- Pop/Peek em pilha vazia deve retornar erro
- Verificar sempre se operação foi bem-sucedida
- Usar padrão Go de retorno (valor, erro)

## Padrões de Uso

### 1. Processamento de Sequências
```go
// Processa elementos em ordem reversa
for !stack.IsEmpty() {
    element, err := stack.Pop()
    if err != nil {
        break
    }
    // Processa element
}
```

### 2. Backup e Restauração
```go
// Salva estado atual
backup := stack.Clone()

// Faz operações...

// Restaura se necessário
if errorOccurred {
    stack = backup
}
```

### 3. Iteração Não-Destrutiva
```go
// Usa pilha auxiliar para não modificar original
aux := NewArrayStack(stack.Size())
CopyStack(stack, aux)

for !aux.IsEmpty() {
    element, _ := aux.Pop()
    // Processa element
}
```

## Otimizações

### ArrayStack
- **Capacidade Inicial**: Escolha baseada no uso esperado
- **Fator de Crescimento**: 2x é comum (balança memória vs. realocações)
- **Shrinking**: Reduz capacidade quando utilização < 25%
- **TrimToSize**: Remove capacidade não utilizada

### LinkedStack
- **Pool de Nós**: Reutiliza nós para evitar alocações
- **Batch Operations**: PushAll/PopMultiple para eficiência
- **Memory Pooling**: Gerenciamento customizado de memória

## Variações e Extensões

### Stack com Máximo
- Mantém track do elemento máximo
- Push/Pop/GetMax todos O(1)
- Usa pilha auxiliar para máximos

### Stack Thread-Safe
- Adiciona mutexes para concorrência
- Operações atômicas
- Considerações de performance

### Stack Limitada
- Capacidade máxima fixa
- IsFull() retorna true quando cheia
- Push pode falhar se cheia

## Conclusão

Pilhas são estruturas fundamentais em ciência da computação, com aplicações em:
- Parsing e compilação
- Gerenciamento de memória
- Algoritmos de busca (DFS)
- Sistemas operacionais
- Aplicações interativas

A escolha entre ArrayStack e LinkedStack depende dos requisitos específicos de performance, memória e flexibilidade da aplicação.