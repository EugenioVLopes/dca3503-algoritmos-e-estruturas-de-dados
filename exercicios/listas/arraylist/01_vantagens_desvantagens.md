# Vantagens e Desvantagens do ArrayList

## Pergunta

Cite uma vantagem e uma desvantagem do ArrayList em relação à Lista Ligada.

## Resposta

### Vantagem

**Acesso direto por índice em O(1)**, enquanto a lista ligada precisa percorrer sequencialmente em O(n).

**Explicação:** O ArrayList armazena elementos em posições contíguas de memória, permitindo calcular diretamente o endereço de qualquer elemento usando a fórmula: `endereço_base + (índice × tamanho_elemento)`. Isso resulta em acesso instantâneo a qualquer posição.

### Desvantagem

**Inserções e remoções no meio custam O(n)** devido ao deslocamento de elementos, enquanto na lista ligada são O(1) quando se tem a referência do nó.

**Explicação:** Quando inserimos ou removemos um elemento no meio do ArrayList, todos os elementos posteriores precisam ser deslocados uma posição para frente ou para trás. Em uma lista ligada, basta ajustar os ponteiros dos nós adjacentes.

## Resumo Comparativo

| Operação          | ArrayList      | LinkedList                 |
| ----------------- | -------------- | -------------------------- |
| Acesso por índice | O(1) ✅        | O(n) ❌                    |
| Inserção no meio  | O(n) ❌        | O(1) ✅                    |
| Uso de memória    | Menor overhead | Maior overhead (ponteiros) |
| Cache locality    | Melhor         | Pior                       |

## Quando usar ArrayList?

- Acesso frequente por índice
- Poucas inserções/remoções no meio
- Operações principalmente no final da lista
- Quando a performance de acesso é crítica
