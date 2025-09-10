# Análise de Código C - Pilha e Avaliação de Afirmações

## Pergunta

Uma pilha é uma estrutura de dados que armazena uma coleção de itens de dados relacionados e que garante o seguinte funcionamento: o último elemento a ser inserido é o primeiro a ser removido. É comum na literatura utilizar os nomes push e pop para as operações de inserção e remoção de um elemento em uma pilha respectivamente. O seguinte trecho de código em linguagem C define uma estrutura de dados pilha utilizando um vetor de inteiros, bem como algumas funções para sua manipulação.

```c
#include <stdlib.h>
#include <stdio.h>

typedef struct {
   int elementos[100];
   int topo;
} pilha;

pilha * cria_pilha() {
   pilha * p = malloc(sizeof(pilha));
   p->topo = -1;
   return p;  // ERRO: deveria ser 'return p;'
}

void push(pilha *p, int elemento) {
   if (p->topo >= 99)
      return;
   p->elementos[++p->topo] = elemento;
}

int pop(pilha *p) {
   int a = p->elementos[p->topo];
   p->topo--;
   return a;
}
```

O programa a seguir utiliza uma pilha:

```c
int main() {
    pilha * p = cria_pilha();
    push(p, 2);
    push(p, 3);
    push(p, 4);
    pop(p);
    push(p, 2);
    int a = pop(p) + pop(p);
    push(p, a);
    a += pop(p);
    printf("%d", a);
    return 0;
}
```

A esse respeito, avalie as afirmações a seguir:

**I.** A complexidade computacional de ambas funções push e pop é O(1).
**II.** O valor exibido pelo programa seria o mesmo caso a instrução `a += pop(p);` fosse trocada por `a += a;`
**III.** Em relação ao vazamento de memória (memory leak), é opcional chamar a função free(p), pois o vetor usado pela pilha é alocado estaticamente.

É correto o que se afirma em:
a) I, apenas.
b) III, apenas.
c) I e II, apenas.
d) II e III apenas.
e) I, II e III.

## Resposta

### Correção do Código

Primeiro, há um erro no código original na função `cria_pilha()`:

```c
// CÓDIGO ORIGINAL (INCORRETO)
pilha * cria_pilha() {
   pilha * p = malloc(sizeof(pilha));
   p->topo = -1;
   return pilha;  // ERRO: deveria ser 'return p;'
}

// CÓDIGO CORRIGIDO
pilha * cria_pilha() {
   pilha * p = malloc(sizeof(pilha));
   p->topo = -1;
   return p;  // CORRETO
}
```

### Rastreamento da Execução

Vamos rastrear passo a passo a execução do programa:

```c
int main() {
    pilha * p = cria_pilha();  // topo = -1, pilha vazia
    
    push(p, 2);               // topo = 0, elementos[0] = 2
    // Estado: [2]
    
    push(p, 3);               // topo = 1, elementos[1] = 3
    // Estado: [2, 3]
    
    push(p, 4);               // topo = 2, elementos[2] = 4
    // Estado: [2, 3, 4]
    
    pop(p);                   // remove 4, topo = 1
    // Estado: [2, 3]
    
    push(p, 2);               // topo = 2, elementos[2] = 2
    // Estado: [2, 3, 2]
    
    int a = pop(p) + pop(p);  // a = 2 + 3 = 5, topo = 0
    // Primeiro pop(): remove 2, retorna 2
    // Segundo pop(): remove 3, retorna 3
    // Estado: [2]
    
    push(p, a);               // push(p, 5), topo = 1
    // Estado: [2, 5]
    
    a += pop(p);              // a = 5 + 5 = 10
    // pop() remove 5, retorna 5
    // Estado: [2]
    
    printf("%d", a);          // Imprime: 10
    return 0;
}
```

**Resultado:** O programa imprime **10**.

### Análise das Afirmações

#### Afirmação I: "A complexidade computacional de ambas funções push e pop é O(1)"

✅ **VERDADEIRA**

**Análise da função push:**
```c
void push(pilha *p, int elemento) {
   if (p->topo >= 99)        // O(1) - comparação
      return;                // O(1) - retorno
   p->elementos[++p->topo] = elemento;  // O(1) - acesso direto ao array
}
```

**Análise da função pop:**
```c
int pop(pilha *p) {
   int a = p->elementos[p->topo];  // O(1) - acesso direto ao array
   p->topo--;                      // O(1) - decremento
   return a;                       // O(1) - retorno
}
```

**Justificativa:**
- Ambas as funções executam apenas operações de **tempo constante**
- **Acesso ao array** por índice é O(1)
- **Incremento/decremento** de variáveis é O(1)
- **Não há loops** ou operações que dependam do tamanho da pilha

#### Afirmação II: "O valor exibido seria o mesmo caso `a += pop(p);` fosse trocada por `a += a;`"

✅ **VERDADEIRA**

**Análise:**

No momento da execução de `a += pop(p);`:
- `a` tem valor **5**
- `pop(p)` retorna **5** (topo da pilha)
- Resultado: `a = 5 + 5 = 10`

Se fosse `a += a;`:
- `a` tem valor **5**
- Resultado: `a = 5 + 5 = 10`

**Ambos produzem o mesmo resultado: 10**

#### Afirmação III: "É opcional chamar free(p), pois o vetor é alocado estaticamente"

❌ **FALSA**

**Análise:**

```c
typedef struct {
   int elementos[100];  // Array DENTRO da struct
   int topo;
} pilha;

pilha * cria_pilha() {
   pilha * p = malloc(sizeof(pilha));  // ALOCAÇÃO DINÂMICA da struct
   p->topo = -1;
   return p;
}
```

**Problemas com a afirmação:**

1. **A struct é alocada dinamicamente**: `malloc(sizeof(pilha))`
2. **O array está dentro da struct**: Quando a struct é alocada dinamicamente, o array também é
3. **Memory leak**: Sem `free(p)`, a memória da struct nunca é liberada
4. **Confusão conceitual**: O array não é "estaticamente alocado" - ele faz parte de uma struct dinamicamente alocada

**Código correto deveria incluir:**
```c
int main() {
    pilha * p = cria_pilha();
    // ... uso da pilha ...
    free(p);  // NECESSÁRIO para evitar memory leak
    return 0;
}
```

### Problemas Adicionais no Código

#### 1. Função pop sem verificação
```c
// CÓDIGO ORIGINAL (PERIGOSO)
int pop(pilha *p) {
   int a = p->elementos[p->topo];  // E se topo == -1?
   p->topo--;
   return a;
}

// CÓDIGO MELHORADO
int pop(pilha *p) {
   if (p->topo < 0) {
       // Pilha vazia - comportamento indefinido
       printf("Erro: tentativa de pop em pilha vazia\n");
       return -1; // ou exit(1)
   }
   int a = p->elementos[p->topo];
   p->topo--;
   return a;
}
```

#### 2. Verificação de ponteiro nulo
```c
void push(pilha *p, int elemento) {
   if (p == NULL) return;  // Verificação de segurança
   if (p->topo >= 99)
      return;
   p->elementos[++p->topo] = elemento;
}
```

### Implementação Melhorada

```c
#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

typedef struct {
   int elementos[100];
   int topo;
} pilha;

pilha * cria_pilha() {
   pilha * p = malloc(sizeof(pilha));
   if (p == NULL) {
       printf("Erro: falha na alocação de memória\n");
       return NULL;
   }
   p->topo = -1;
   return p;
}

bool push(pilha *p, int elemento) {
   if (p == NULL || p->topo >= 99) {
       return false;  // Falha
   }
   p->elementos[++p->topo] = elemento;
   return true;  // Sucesso
}

bool pop(pilha *p, int *elemento) {
   if (p == NULL || p->topo < 0) {
       return false;  // Falha
   }
   *elemento = p->elementos[p->topo--];
   return true;  // Sucesso
}

void destroi_pilha(pilha *p) {
   if (p != NULL) {
       free(p);
   }
}

bool pilha_vazia(pilha *p) {
   return (p == NULL || p->topo < 0);
}

bool pilha_cheia(pilha *p) {
   return (p != NULL && p->topo >= 99);
}
```

### Análise de Complexidade

| Operação | Complexidade | Justificativa |
|----------|--------------|---------------|
| `cria_pilha()` | O(1) | Alocação de tamanho fixo |
| `push()` | O(1) | Acesso direto ao array |
| `pop()` | O(1) | Acesso direto ao array |
| `pilha_vazia()` | O(1) | Comparação simples |
| `pilha_cheia()` | O(1) | Comparação simples |
| `destroi_pilha()` | O(1) | Liberação de memória |

### Comparação com Implementações Dinâmicas

#### Array Fixo (Código dado)
**Vantagens:**
- ✅ Simplicidade de implementação
- ✅ Acesso O(1) garantido
- ✅ Boa localidade de memória

**Desvantagens:**
- ❌ Tamanho limitado (100 elementos)
- ❌ Desperdício de memória se pilha pequena
- ❌ Overflow se exceder capacidade

#### Lista Ligada
**Vantagens:**
- ✅ Tamanho dinâmico
- ✅ Usa apenas memória necessária

**Desvantagens:**
- ❌ Overhead de ponteiros
- ❌ Pior localidade de memória
- ❌ Alocação/desalocação frequente

### Conclusão

**Resposta correta: c) I e II, apenas.**

**Justificativa:**
- ✅ **Afirmação I**: Verdadeira - ambas as funções são O(1)
- ✅ **Afirmação II**: Verdadeira - ambas as instruções produzem o mesmo resultado (10)
- ❌ **Afirmação III**: Falsa - `free(p)` é obrigatório para evitar memory leak

**Pontos importantes:**
1. **Complexidade O(1)** é característica fundamental de pilhas bem implementadas
2. **Rastreamento de execução** é essencial para entender o comportamento
3. **Gerenciamento de memória** em C requer cuidado com malloc/free
4. **Verificações de erro** são importantes para código robusto