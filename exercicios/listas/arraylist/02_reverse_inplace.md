# Função Reverse In-place para ArrayList

## Pergunta

Escreva uma função in-place para inverter a ordem de um ArrayList.

```go
func (list *ArrayList) Reverse()

type ArrayList struct {
    values   []int
    inserted int
}
```

## Resposta

### Implementação

```go
func (list *ArrayList) Reverse() {
    for i := 0; i < list.inserted/2; i++ {
        j := list.inserted - 1 - i
        list.values[i], list.values[j] = list.values[j], list.values[i]
    }
}
```

### Explicação Detalhada

**Algoritmo:** A função percorre apenas **metade do array** (até `inserted/2`) e troca cada elemento da posição `i` com seu correspondente da posição `j = inserted - 1 - i`.

**Por que apenas metade?**
- Se percorrermos o array inteiro, cada elemento seria trocado duas vezes
- Resultado: o array voltaria ao estado original
- Percorrendo apenas metade, cada par é trocado exatamente uma vez

### Exemplo de Execução

**Array inicial:** `[1, 2, 3, 4, 5]` (inserted = 5)

| Iteração | i | j | Troca | Array após troca |
|----------|---|---|-------|------------------|
| 1 | 0 | 4 | 1 ↔ 5 | `[5, 2, 3, 4, 1]` |
| 2 | 1 | 3 | 2 ↔ 4 | `[5, 4, 3, 2, 1]` |

**Loop para:** i = 2, pois 2 ≥ 5/2 (2.5)

**Array final:** `[5, 4, 3, 2, 1]` ✅

### Análise de Complexidade

- **Tempo:** O(n/2) = **O(n)**
  - Percorre metade dos elementos
  - Cada troca é O(1)

- **Espaço:** **O(1)** (in-place)
  - Usa apenas variáveis auxiliares `i` e `j`
  - Não aloca memória adicional
  - Modifica o array original

### Casos Especiais

**Array vazio:** `inserted = 0`
- Loop não executa (0 < 0/2 é falso)
- Array permanece vazio ✅

**Array com 1 elemento:** `inserted = 1`
- Loop não executa (0 < 1/2 é falso)
- Array permanece inalterado ✅

**Array com 2 elementos:** `inserted = 2`
- Loop executa 1 vez (0 < 2/2)
- Troca posições 0 e 1 ✅

### Vantagens da Implementação

✅ **Eficiência de memória:** O(1) espaço
✅ **Performance:** O(n) tempo, ótimo para inversão
✅ **Simplicidade:** Código limpo e fácil de entender
✅ **Robustez:** Funciona para todos os tamanhos de array
✅ **In-place:** Não requer array auxiliar

### Comparação com Outras Abordagens

| Abordagem | Tempo | Espaço | In-place |
|-----------|-------|--------|-----------|
| Troca de extremidades | O(n) | O(1) | ✅ |
| Array auxiliar | O(n) | O(n) | ❌ |
| Recursão | O(n) | O(n) | ❌ |
| Stack | O(n) | O(n) | ❌ |

### Teste da Função

```go
func main() {
    list := &ArrayList{
        values:   []int{1, 2, 3, 4, 5, 0, 0, 0},
        inserted: 5,
    }
    
    fmt.Println("Antes:", list.values[:list.inserted]) // [1 2 3 4 5]
    list.Reverse()
    fmt.Println("Depois:", list.values[:list.inserted]) // [5 4 3 2 1]
}
```