# Vantagens e Desvantagens da Lista Duplamente Ligada

## Pergunta

Cite uma vantagem e uma desvantagem da Lista Duplamente Ligada em relação à Lista Ligada.

## Resposta

### Vantagem

**Inserção no final em O(1)** devido ao ponteiro para o último nó, enquanto a lista ligada simples precisa percorrer toda a lista em O(n).

**Explicação:** A DoublyLinkedList mantém uma referência direta para o último nó (`tail`), permitindo inserções no final sem necessidade de percorrer a lista. A LinkedList simples precisa navegar desde o `head` até encontrar o último nó.

**Outras vantagens:**

- Navegação bidirecional (frente e trás)
- Remoção de nó conhecido em O(1)
- Melhor performance para operações no final da lista

### Desvantagem

**Maior consumo de memória** devido aos dois ponteiros por nó (anterior e próximo), enquanto a lista ligada simples usa apenas um ponteiro.

**Explicação:** Cada nó na DoublyLinkedList precisa armazenar:

- `value`: o valor do elemento
- `next`: ponteiro para o próximo nó
- `prev`: ponteiro para o nó anterior

Enquanto na LinkedList simples cada nó armazena apenas:

- `value`: o valor do elemento
- `next`: ponteiro para o próximo nó

**Overhead de memória:** ~50% mais memória por nó devido ao ponteiro adicional.

## Resumo Comparativo

| Aspecto                       | LinkedList    | DoublyLinkedList |
| ----------------------------- | ------------- | ---------------- |
| Inserção no final             | O(n)          | O(1) ✅          |
| Navegação                     | Unidirecional | Bidirecional ✅  |
| Memória por nó                | 1 ponteiro    | 2 ponteiros ❌   |
| Remoção de nó conhecido       | O(n)          | O(1) ✅          |
| Complexidade de implementação | Menor         | Maior            |

## Quando usar DoublyLinkedList?

- Inserções/remoções frequentes no final
- Necessidade de navegação bidirecional
- Operações que requerem acesso ao nó anterior
- Quando a performance é mais importante que o uso de memória
