# Classificação de Estruturas de Dados Java

## Pergunta

A biblioteca de coleções da linguagem Java disponibiliza implementações de propósito geral para estruturas de dados elementares, como listas, filas e pilhas. Considere as seguintes definições de classes que representam implementações de estruturas de dados disponíveis na biblioteca da linguagem:

• **Classe A**: os objetos são organizados em uma ordem linear e podem ser inseridos somente no início ou no final dessa sequência;
• **Classe B**: os objetos são organizados em uma ordem linear determinada por uma referência ao próximo objeto;
• **Classe C**: os objetos são removidos na ordem oposta em que foram inseridos;
• **Classe D**: os objetos são inseridos e removidos respeitando a seguinte regra: o elemento a ser removido é sempre aquele que foi inserido primeiro.

Nesse contexto, assinale a alternativa que representa, respectivamente, as estruturas de dados implementadas pelas classes A, B, C e D.

**Alternativas:**

1. Lista circular, lista simplesmente ligada, pilha e fila.
2. **Deque, lista simplesmente ligada, pilha e fila.**
3. Lista duplamente ligada, lista simplesmente ligada, fila e pilha.
4. Pilha, fila, deque e lista simplesmente encadeada.
5. Deque, pilha, lista ligada e fila.

## Resposta

### Análise de Cada Classe

#### Classe A: "Inserção somente no início ou no final"

**Características:**

- Ordem linear ✅
- Inserção **apenas** nas extremidades (início OU final)
- Não permite inserção no meio

### Estrutura correspondente: DEQUE (Double-Ended Queue)

**Justificativa:**

- **Deque** permite inserção/remoção em **ambas as extremidades**
- **Lista comum** permite inserção em qualquer posição
- **Pilha** permite apenas uma extremidade
- **Fila** permite inserção em uma extremidade e remoção na outra

#### Classe B: "Ordem linear determinada por referência ao próximo objeto"

**Características:**

- Ordem linear ✅
- Organização baseada em **referências/ponteiros**
- Cada elemento aponta para o próximo

### Estrutura correspondente: LISTA SIMPLESMENTE LIGADA

**Justificativa:**

- **LinkedList** usa nós com referência `next`
- **ArrayList** usa índices, não referências
- **Pilha/Fila** são abstrações, não especificam implementação

#### Classe C: "Removidos na ordem oposta em que foram inseridos"

**Características:**

- **Último a entrar, primeiro a sair**
- Ordem **inversa** de inserção
- Comportamento **LIFO** (Last In, First Out)

### Estrutura correspondente: PILHA (Stack)

**Justificativa:**

- **Pilha**: LIFO - último inserido é o primeiro removido
- **Fila**: FIFO - primeiro inserido é o primeiro removido
- **Deque**: Pode ser LIFO ou FIFO dependendo do uso

#### Classe D: "Elemento removido é sempre o que foi inserido primeiro"

**Características:**

- **Primeiro a entrar, primeiro a sair**
- Ordem **igual** à de inserção
- Comportamento **FIFO** (First In, First Out)

### Estrutura correspondente: FILA (Queue)

**Justificativa:**

- **Fila**: FIFO - primeiro inserido é o primeiro removido
- **Pilha**: LIFO - último inserido é o primeiro removido

### Resumo da Análise

| Classe | Descrição                         | Estrutura                     | Comportamento |
| ------ | --------------------------------- | ----------------------------- | ------------- |
| **A**  | Inserção só nas extremidades      | **Deque**                     | Acesso duplo  |
| **B**  | Ordem por referência ao próximo   | **Lista Simplesmente Ligada** | Ponteiros     |
| **C**  | Remoção em ordem oposta           | **Pilha**                     | LIFO          |
| **D**  | Remove sempre o primeiro inserido | **Fila**                      | FIFO          |

### Resposta Correta

**Alternativa 2: Deque, lista simplesmente ligada, pilha e fila.**

### Análise das Outras Alternativas

#### Alternativa 1: "Lista circular, lista simplesmente ligada, pilha e fila"

❌ **Incorreta**

- **Classe A** não é lista circular
- Lista circular permite inserção em qualquer posição
- Classe A especifica "somente início ou final"

#### Alternativa 3: "Lista duplamente ligada, lista simplesmente ligada, fila e pilha"

❌ **Incorreta**

- **Classe A** não é lista duplamente ligada
- Lista duplamente ligada permite inserção em qualquer posição
- **Classes C e D** estão trocadas (fila ≠ LIFO, pilha ≠ FIFO)

#### Alternativa 4: "Pilha, fila, deque e lista simplesmente encadeada"

❌ **Incorreta**

- **Classe A** não é pilha (pilha só permite uma extremidade)
- **Classe B** não é fila
- **Classes A, B, C, D** todas incorretas

#### Alternativa 5: "Deque, pilha, lista ligada e fila"

❌ **Incorreta**

- **Classe B** não é pilha
- **Classe C** não é lista ligada
- Pilha tem comportamento LIFO, não "referência ao próximo"

### Conceitos Importantes

#### Deque vs Lista Simplesmente Ligada

- **Deque**: Acesso **restrito** às extremidades (início e fim)
- **Lista**: Acesso **livre** a qualquer posição

#### LIFO vs FIFO

- **LIFO** (Pilha): Último entra, primeiro sai
- **FIFO** (Fila): Primeiro entra, primeiro sai

### Exemplos Práticos em Java

```java
// Classe A - Deque
Deque<String> deque = new ArrayDeque<>();
deque.addFirst("início");  // ✅ Permitido
deque.addLast("final");    // ✅ Permitido
// deque.add(1, "meio");   // ❌ Não existe no Deque

// Classe B - LinkedList
LinkedList<String> lista = new LinkedList<>();
lista.add("qualquer posição");  // ✅ Baseado em referências

// Classe C - Stack (Pilha)
Stack<String> pilha = new Stack<>();
pilha.push("A"); pilha.push("B"); pilha.push("C");
String ultimo = pilha.pop();  // "C" - ordem oposta

// Classe D - Queue (Fila)
Queue<String> fila = new LinkedList<>();
fila.offer("A"); fila.offer("B"); fila.offer("C");
String primeiro = fila.poll();  // "A" - primeiro inserido
```

### Conclusão

A análise correta identifica:

- **A = Deque**: Inserção restrita às extremidades
- **B = Lista Simplesmente Ligada**: Organização por referências
- **C = Pilha**: Comportamento LIFO
- **D = Fila**: Comportamento FIFO

### Resposta: Alternativa 2
