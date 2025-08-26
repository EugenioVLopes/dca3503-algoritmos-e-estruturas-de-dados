# DCA3503 - Algoritmos e Estruturas de Dados

## 📚 Sobre este Diretório

Este diretório contém material de estudo organizado para a disciplina **DCA3503 - Algoritmos e Estruturas de Dados**, focando especificamente em **Listas** como estrutura de dados fundamental.

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

### 💻 **Implementações em Go**

5. **[list_interface.go](list_interface.go)** - Interface e Utilitários

   - Definição da interface `List`
   - Funções utilitárias que trabalham com a interface
   - Algoritmos genéricos (busca, ordenação, etc.)

6. **[arraylist.go](arraylist.go)** - Implementação ArrayList

   - Implementação completa da estrutura ArrayList
   - Todos os métodos com comentários detalhados
   - Operações otimizadas (AddAll, TrimToSize, etc.)

7. **[linkedlist.go](linkedlist.go)** - Implementação LinkedList

   - Implementação completa da estrutura LinkedList
   - Algoritmos especiais (Reverse, GetMiddle, etc.)
   - Detecção de ciclos e remoção de duplicatas

8. **[main.go](main.go)** - Demonstrações e Testes
   - Exemplos práticos de uso
   - Comparações de performance
   - Demonstração da interface polimórfica
   - Algoritmos usando as estruturas

### 📁 **Arquivos Legados**

9. **[list_completo.go](list_completo.go)** - Versão Original

   - Implementação original com ambas estruturas
   - Mantido para referência histórica

10. **[estruturas_dados_explicacao.md](estruturas_dados_explicacao.md)** - Versão Original
    - Documentação original consolidada
    - Mantido para referência

## 🚀 **Como Executar**

### **Pré-requisitos:**

- Go 1.19 ou superior instalado
- Terminal/Command Prompt

### **Executando o código:**

```bash
# Navegar para o diretório
cd "c:\Users\Mateus\Downloads\dca3503-algoritmos-e-estruturas-de-dados"

# Executar demonstrações
go run *.go

# Ou executar arquivos específicos
go run main.go arraylist.go linkedlist.go list_interface.go
```

### **Testando implementações:**

```bash
# Testar apenas ArrayList
go run arraylist.go list_interface.go -test arraylist

# Testar apenas LinkedList
go run linkedlist.go list_interface.go -test linkedlist
```

## 📊 **Resumo de Complexidades**

| Operação            | ArrayList | LinkedList | Melhor Para |
| ------------------- | --------- | ---------- | ----------- |
| **Acesso Get(i)**   | O(1)      | O(n)       | ArrayList   |
| **Inserção final**  | O(1)\*    | O(n)       | ArrayList   |
| **Inserção início** | O(n)      | O(1)       | LinkedList  |
| **Remoção final**   | O(1)      | O(n)       | ArrayList   |
| **Remoção início**  | O(n)      | O(1)       | LinkedList  |
| **Busca**           | O(n)      | O(n)       | Empate      |
| **Uso memória**     | Médio     | Alto       | ArrayList   |

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
