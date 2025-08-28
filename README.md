# DCA3503 - Algoritmos e Estruturas de Dados

## 📚 Sobre este Diretório

Este diretório contém material de estudo organizado para a disciplina **DCA3503 - Algoritmos e Estruturas de Dados**, focando em estruturas de dados fundamentais como **Listas**, **Pilhas** e **Filas**.

## 🗂️ Organização dos Arquivos

### 📖 **Documentação Teórica**

1. **[01_lista_adt.md](01_lista_adt.md)** - Conceitos Fundamentais

   - O que é um Tipo Abstrato de Dados (ADT)
   - Interface da Lista
   - Invariantes e propriedades
   - Análise de complexidade geral

2. **[02_arraylist.md](02_arraylist.md)** - ArrayList Detalhado

   - Estrutura interna baseada em array dinâmico
   - Algoritmos e pseudocódigo de cada operação
   - Análise de complexidade específica
   - Estratégias de redimensionamento
   - Vantagens, desvantagens e casos de uso

3. **[03_linkedlist.md](03_linkedlist.md)** - LinkedList Detalhado

   - Estrutura baseada em nós ligados
   - Algoritmos e pseudocódigo de cada operação
   - Variações (simples, dupla, circular)
   - Algoritmos especiais (reverse, detectar ciclo, etc.)
   - Vantagens, desvantagens e casos de uso

4. **[04_comparacao_estruturas.md](04_comparacao_estruturas.md)** - Análise Comparativa
   - Comparação detalhada de complexidades
   - Análise de uso de memória
   - Performance de cache
   - Benchmarks práticos
   - Árvore de decisão para escolha
   - Implementações híbridas

5. **[06_stack_adt.md](06_stack_adt.md)** - Pilhas (Stacks)
   - Conceitos fundamentais de pilhas (LIFO)
   - Interface Stack e operações básicas
   - Comparação entre ArrayStack e LinkedStack
   - Aplicações práticas (parênteses balanceados, expressões pós-fixas)
   - Algoritmos e padrões de uso

6. **[07_queue_adt.md](07_queue_adt.md)** - Filas (Queues)
   - Conceitos fundamentais de filas (FIFO)
   - Interface Queue e operações básicas
   - Comparação entre ArrayQueue (circular) e LinkedQueue
   - Aplicações práticas (BFS, geração de números binários, sistemas)
   - Algoritmos clássicos e padrões de uso

### 💻 **Implementações em Go**

#### **Listas**

6. **[list_interface.go](list_interface.go)** - Interface e Utilitários

   - Definição da interface `List`
   - Funções utilitárias que trabalham com a interface
   - Algoritmos genéricos (busca, ordenação, etc.)

7. **[arraylist.go](arraylist.go)** - Implementação ArrayList

   - Implementação completa da estrutura ArrayList
   - Todos os métodos com comentários detalhados
   - Operações otimizadas (AddAll, TrimToSize, etc.)

8. **[linkedlist.go](linkedlist.go)** - Implementação LinkedList

   - Implementação completa da estrutura LinkedList
   - Algoritmos especiais (Reverse, GetMiddle, etc.)
   - Detecção de ciclos e remoção de duplicatas

#### **Pilhas**

9. **[stack_interface.go](stack_interface.go)** - Interface Stack e Utilitários

   - Definição da interface `Stack`
   - Funções utilitárias para manipulação de pilhas
   - Algoritmos clássicos (parênteses balanceados, expressões pós-fixas)
   - Operações avançadas (inversão, busca, estatísticas)

10. **[arraystack.go](arraystack.go)** - Implementação ArrayStack

    - Pilha baseada em array dinâmico
    - Redimensionamento automático
    - Operações O(1) amortizadas
    - Métodos de análise e estatísticas

11. **[linkedstack.go](linkedstack.go)** - Implementação LinkedStack

     - Pilha baseada em lista ligada
     - Operações sempre O(1)
     - Métodos funcionais (Map, Filter, Reduce)
     - Flexibilidade total de tamanho

#### **Filas**

12. **[queue_interface.go](queue_interface.go)** - Interface Queue e Utilitários

    - Definição da interface `Queue`
    - Funções utilitárias para manipulação de filas
    - Algoritmos clássicos (BFS, geração binária, caractere não repetido)
    - Operações avançadas (rotação, intercalação, estatísticas)

13. **[arrayqueue.go](arrayqueue.go)** - Implementação ArrayQueue

    - Fila baseada em array circular
    - Operações O(1) para enqueue/dequeue
    - Redimensionamento automático
    - Uso eficiente de espaço (reutiliza posições)

14. **[linkedqueue.go](linkedqueue.go)** - Implementação LinkedQueue

    - Fila baseada em lista ligada
    - Operações sempre O(1)
    - Métodos funcionais (Map, Filter, Reduce, Partition)
    - Flexibilidade total de tamanho

15. **[main.go](main.go)** - Demonstrações e Testes
   - Exemplos práticos de uso de listas, pilhas e filas
   - Comparações de performance entre implementações
   - Demonstração da interface polimórfica
   - Algoritmos clássicos usando as estruturas

### 📁 **Arquivos Legados**

9. **[list_completo.go](list_completo.go)** - Versão Original

   - Implementação original com ambas estruturas
   - Mantido para referência histórica

10. **[estruturas_dados_explicacao.md](estruturas_dados_explicacao.md)** - Versão Original
    - Documentação original consolidada
    - Mantido para referência

## 🚀 **Como Executar**

### **Pré-requisitos:**

#### **Instalação do Go:**

1. **Windows:**
   - Baixe o instalador em: https://golang.org/dl/
   - Execute o instalador e siga as instruções
   - Reinicie o terminal após a instalação

2. **Verificar instalação:**
   ```bash
   go version
   ```

### **Executando o código:**

```bash
# Navegar para o diretório
cd "c:\Users\Mateus\Downloads\dca3503-algoritmos-e-estruturas-de-dados"

# Executar todas as demonstrações (listas, pilhas e filas)
go run *.go

# Ou executar arquivos específicos
go run main.go list_interface.go arraylist.go linkedlist.go stack_interface.go arraystack.go linkedstack.go queue_interface.go arrayqueue.go linkedqueue.go
```

### **Executando demonstrações específicas:**

```bash
# Apenas listas
go run main.go list_interface.go arraylist.go linkedlist.go

# Apenas pilhas
go run main.go stack_interface.go arraystack.go linkedstack.go

# Apenas filas
go run main.go queue_interface.go arrayqueue.go linkedqueue.go
```

## 📊 **Resumo de Complexidades**

### **Listas**

| Operação            | ArrayList | LinkedList | Melhor Para |
| ------------------- | --------- | ---------- | ----------- |
| **Acesso Get(i)**   | O(1)      | O(n)       | ArrayList   |
| **Inserção final**  | O(1)\*    | O(n)       | ArrayList   |
| **Inserção início** | O(n)      | O(1)       | LinkedList  |
| **Remoção final**   | O(1)      | O(n)       | ArrayList   |
| **Remoção início**  | O(n)      | O(1)       | LinkedList  |
| **Busca**           | O(n)      | O(n)       | Empate      |
| **Uso memória**     | Médio     | Alto       | ArrayList   |

### **Pilhas**

| Operação        | ArrayStack    | LinkedStack | Melhor Para   |
| --------------- | ------------- | ----------- | ------------- |
| **Push**        | O(1)\*        | O(1)        | LinkedStack   |
| **Pop**         | O(1)          | O(1)        | Empate        |
| **Peek**        | O(1)          | O(1)        | Empate        |
| **Size**        | O(1)          | O(1)        | Empate        |
| **IsEmpty**     | O(1)          | O(1)        | Empate        |
| **Uso memória** | Mais eficiente| Overhead ptr | ArrayStack    |
| **Cache**       | Melhor        | Pior        | ArrayStack    |
| **Flexibilidade**| Limitada     | Total       | LinkedStack   |

*\* O(1) amortizado para ArrayStack devido ao redimensionamento*

### **Filas**

| Operação        | ArrayQueue    | LinkedQueue | Melhor Para   |
| --------------- | ------------- | ----------- | ------------- |
| **Enqueue**     | O(1)\*        | O(1)        | LinkedQueue   |
| **Dequeue**     | O(1)          | O(1)        | Empate        |
| **Front**       | O(1)          | O(1)        | Empate        |
| **Rear**        | O(1)          | O(1)        | Empate        |
| **Size**        | O(1)          | O(1)        | Empate        |
| **IsEmpty**     | O(1)          | O(1)        | Empate        |
| **Uso memória** | Mais eficiente| Overhead ptr | ArrayQueue    |
| **Cache**       | Melhor        | Pior        | ArrayQueue    |
| **Flexibilidade**| Limitada     | Total       | LinkedQueue   |

*\* O(1) amortizado para ArrayQueue devido ao redimensionamento*

\*Amortizado

## 🎓 **Conceitos Importantes**

### **Análise Amortizada**

- ArrayList: O(1) amortizado para inserção no final
- Redimensionamento: custoso individualmente, mas raro

### **Localidade de Cache**

- ArrayList: excelente (elementos contíguos)
- LinkedList: ruim (nós espalhados na memória)

### **Trade-offs Fundamentais**

- **Tempo vs Espaço**: ArrayList pode desperdiçar memória
- **Flexibilidade vs Performance**: LinkedList mais flexível, ArrayList mais rápido

## 📚 **Referências e Leituras Complementares**

### **Livros:**

- "Introduction to Algorithms" - Cormen, Leiserson, Rivest, Stein
- "Data Structures and Algorithms in Java" - Goodrich, Tamassia
- "Algorithms" - Robert Sedgewick

### **Recursos Online:**

- [Visualgo](https://visualgo.net/) - Visualização de algoritmos
- [Big-O Cheat Sheet](https://www.bigocheatsheet.com/)
- [Go Documentation](https://golang.org/doc/)
