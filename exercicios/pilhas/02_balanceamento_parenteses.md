# Função para Detectar Parênteses Balanceados

## Pergunta

Escreva uma função que detecta se uma certa combinação de parênteses está balanceada. 

**Dica 1:** usar uma pilha.
**Dica 2:** pensar nos casos de sucesso e casos de falha antes da implementação.

```go
func balparenteses(par string) bool
```

## Resposta

### Análise do Problema

#### Casos de Sucesso (Balanceados)
- `"()"` - Par simples
- `"()()"` - Múltiplos pares
- `"(())"` - Aninhamento
- `"((()))"` - Múltiplos níveis
- `"()(())"` - Combinação de pares e aninhamento
- `""` - String vazia (convenção: balanceada)

#### Casos de Falha (Desbalanceados)
- `"("` - Parêntese aberto sem fechamento
- `")"` - Parêntese fechado sem abertura
- `")("` - Ordem incorreta
- `"(()"` - Falta fechamento
- `"())"` - Fechamento extra
- `"((())"` - Falta um fechamento

### Estratégia com Pilha

1. **Percorrer** cada caractere da string
2. **Parêntese aberto `(`**: empilhar (push)
3. **Parêntese fechado `)`**: 
   - Se pilha vazia → desbalanceado
   - Senão → desempilhar (pop)
4. **Final**: pilha deve estar vazia

### Implementação

```go
package main

import "fmt"

func balparenteses(par string) bool {
    // Pilha simples usando slice
    pilha := make([]rune, 0)
    
    for _, char := range par {
        switch char {
        case '(':
            // Parêntese aberto: empilha
            pilha = append(pilha, char)
            
        case ')':
            // Parêntese fechado: verifica se há correspondente
            if len(pilha) == 0 {
                // Não há parêntese aberto correspondente
                return false
            }
            // Remove o parêntese aberto correspondente
            pilha = pilha[:len(pilha)-1]
            
        default:
            // Ignora outros caracteres (se houver)
            continue
        }
    }
    
    // Balanceado se e somente se a pilha estiver vazia
    return len(pilha) == 0
}
```

### Implementação Alternativa (Mais Robusta)

```go
func balparentesesRobusta(par string) bool {
    contador := 0
    
    for _, char := range par {
        switch char {
        case '(':
            contador++
        case ')':
            contador--
            // Se contador fica negativo, há ')' sem '(' correspondente
            if contador < 0 {
                return false
            }
        }
    }
    
    // Balanceado se contador é zero
    return contador == 0
}
```

### Implementação com Pilha Explícita

```go
type Stack struct {
    items []rune
}

func (s *Stack) Push(item rune) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func balparentesesComPilha(par string) bool {
    stack := &Stack{}
    
    for _, char := range par {
        if char == '(' {
            stack.Push(char)
        } else if char == ')' {
            if stack.IsEmpty() {
                return false
            }
            stack.Pop()
        }
    }
    
    return stack.IsEmpty()
}
```

### Extensão para Múltiplos Tipos de Parênteses

```go
func balanceamentoCompleto(s string) bool {
    stack := make([]rune, 0)
    
    // Mapa de correspondências
    pares := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    for _, char := range s {
        switch char {
        case '(', '[', '{':
            // Caracteres de abertura: empilha
            stack = append(stack, char)
            
        case ')', ']', '}':
            // Caracteres de fechamento: verifica correspondência
            if len(stack) == 0 {
                return false
            }
            
            // Verifica se o topo da pilha corresponde
            topo := stack[len(stack)-1]
            if topo != pares[char] {
                return false
            }
            
            // Remove o caractere correspondente
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}
```

### Testes

```go
func main() {
    // Casos de teste
    testes := []struct {
        input    string
        esperado bool
    }{
        {"", true},           // String vazia
        {"()", true},         // Par simples
        {"(())", true},       // Aninhado
        {"()()", true},       // Múltiplos pares
        {"((()))", true},     // Múltiplos níveis
        {"()(())", true},     // Combinado
        
        {"(", false},         // Aberto sem fechar
        {")", false},         // Fechado sem abrir
        {")(", false},        // Ordem errada
        {"(()", false},       // Falta fechamento
        {"())", false},       // Fechamento extra
        {"((())", false},     // Falta um fechamento
    }
    
    fmt.Println("Testando função balparenteses:")
    for i, teste := range testes {
        resultado := balparenteses(teste.input)
        status := "✅"
        if resultado != teste.esperado {
            status = "❌"
        }
        fmt.Printf("Teste %d: \"%s\" → %v %s\n", 
                   i+1, teste.input, resultado, status)
    }
}
```

### Análise de Complexidade

#### Complexidade de Tempo
- **O(n)** onde n é o comprimento da string
- Cada caractere é processado exatamente uma vez
- Operações de pilha (push/pop) são O(1)

#### Complexidade de Espaço
- **O(n)** no pior caso
- Pior caso: string como "((((((" requer pilha de tamanho n/2
- Melhor caso: string balanceada simples como "()()" usa espaço constante

### Rastreamento de Execução

**Exemplo:** `balparenteses("(())")`

```
String: "(())"
Índice: 0123

Passo 1: char='(' → pilha=['(']
Passo 2: char='(' → pilha=['(', '(']
Passo 3: char=')' → pilha=['('] (pop)
Passo 4: char=')' → pilha=[] (pop)

Resultado: pilha vazia → true (balanceado)
```

**Exemplo:** `balparenteses("())")`

```
String: "())"
Índice: 012

Passo 1: char='(' → pilha=['(']
Passo 2: char=')' → pilha=[] (pop)
Passo 3: char=')' → pilha vazia, mas tentativa de pop → false

Resultado: false (desbalanceado)
```

### Variações do Problema

#### 1. Contar Níveis de Aninhamento
```go
func nivelMaximoAninhamento(par string) int {
    nivel := 0
    maxNivel := 0
    
    for _, char := range par {
        if char == '(' {
            nivel++
            if nivel > maxNivel {
                maxNivel = nivel
            }
        } else if char == ')' {
            nivel--
        }
    }
    
    return maxNivel
}
```

#### 2. Encontrar Posição do Erro
```go
func encontrarErro(par string) int {
    pilha := make([]int, 0) // Armazena índices
    
    for i, char := range par {
        if char == '(' {
            pilha = append(pilha, i)
        } else if char == ')' {
            if len(pilha) == 0 {
                return i // Posição do ')' sem '(' correspondente
            }
            pilha = pilha[:len(pilha)-1]
        }
    }
    
    if len(pilha) > 0 {
        return pilha[0] // Posição do primeiro '(' sem fechamento
    }
    
    return -1 // Balanceado
}
```

### Conclusão

A função `balparenteses` demonstra um uso clássico de pilha para resolver problemas de **correspondência e balanceamento**. A pilha é ideal porque:

1. **LIFO** corresponde naturalmente ao aninhamento de parênteses
2. **Eficiência** O(1) para operações básicas
3. **Simplicidade** de implementação e compreensão
4. **Extensibilidade** para múltiplos tipos de delimitadores

Este padrão é fundamental em:
- **Parsers** de linguagens de programação
- **Validação** de expressões matemáticas
- **Análise sintática** de código
- **Verificação** de estruturas aninhadas