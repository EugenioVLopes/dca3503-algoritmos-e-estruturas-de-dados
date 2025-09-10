# Complexidade de Tempo das Estruturas de Dados

## Pergunta

Considere as estruturas de dados apresentadas na tabela e responda o desempenho de tempo de pior caso e melhor caso para cada operação listada.

## Resposta

### ArrayList

- **Add(value int)**: Pior caso O(n), Melhor caso Ω(1)
- **AddOnIndex(value int, index int)**: Pior caso O(n), Melhor caso Ω(1)
- **RemoveOnIndex(index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Get(index int)**: Pior caso O(1), Melhor caso Ω(1)
- **Set(value int, index int)**: Pior caso O(1), Melhor caso Ω(1)
- **Size()**: Pior caso O(1), Melhor caso Ω(1)

### LinkedList

- **Add(value int)**: Pior caso O(n), Melhor caso Ω(n)
- **AddOnIndex(value int, index int)**: Pior caso O(n), Melhor caso Ω(1)
- **RemoveOnIndex(index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Get(index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Set(value int, index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Size()**: Pior caso O(1), Melhor caso Ω(1)

### DoublyLinkedList

- **Add(value int)**: Pior caso O(1), Melhor caso Ω(1)
- **AddOnIndex(value int, index int)**: Pior caso O(n), Melhor caso Ω(1)
- **RemoveOnIndex(index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Get(index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Set(value int, index int)**: Pior caso O(n), Melhor caso Ω(1)
- **Size()**: Pior caso O(1), Melhor caso Ω(1)

## Explicação

### Notação Big O e Omega

- **O(n)** - Notação Big O: representa o **pior caso** (limite superior)
- **Ω(n)** - Notação Omega: representa o **melhor caso** (limite inferior)
- **n** - representa o tamanho da estrutura de dados

### Por que essas complexidades?

#### ArrayList (Array Dinâmico)

- **Vantagens**: Acesso direto por índice O(1), operações no final são rápidas
- **Desvantagens**: Inserções/remoções no meio requerem deslocamento de elementos
- **Add**: Melhor caso Ω(1) quando há espaço; Pior caso O(n) quando precisa redimensionar
- **Get/Set**: Sempre O(1) devido ao acesso direto por índice

#### LinkedList (Lista Ligada Simples)

- **Vantagens**: Inserções/remoções no início são O(1)
- **Desvantagens**: Acesso sequencial, sem acesso direto por índice
- **Add**: Sempre O(n) pois precisa percorrer até o final
- **Get/Set**: Pior caso O(n) quando o elemento está no final

#### DoublyLinkedList (Lista Duplamente Ligada)

- **Vantagens**: Inserção no final é O(1) (mantém referência para o último nó)
- **Vantagens**: Navegação bidirecional otimiza algumas operações
- **Add**: O(1) pois mantém ponteiro para o último elemento
- **Get/Set**: Pode otimizar buscando do início ou fim, mas ainda O(n) no pior caso

## Resumo Objetivo

**ArrayList** oferece acesso direto O(1) mas inserções custam O(n). **LinkedList** tem inserções no início O(1) mas acesso por índice O(n). **DoublyLinkedList** otimiza inserções no final para O(1) mantendo acesso O(n). A escolha depende do padrão de uso: ArrayList para acesso frequente por índice, LinkedList para inserções no início, DoublyLinkedList para inserções no final e navegação bidirecional.
