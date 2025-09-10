# Aplicativo para Registro de Recordes Olímpicos

## Pergunta

O coordenador geral de um comitê olímpico solicitou a implementação de um aplicativo que permita o registro dos recordes dos atletas à medida que forem sendo quebrados, mantendo a ordem cronológica dos acontecimentos, e possibilitando a leitura dos dados a partir dos mais recentes. Considerando os requisitos do aplicativo, a estrutura de dados mais adequada para a solução a ser implementada é:

**Alternativas:**

a) o deque: tipo especial de lista encadeada, que permite a inserção e a remoção em qualquer das duas extremidades da fila e que deve possuir um nó com a informação (recorde) e dois apontadores, respectivamente, para os nós próximo e anterior.

b) a fila: tipo especial de lista encadeada, tal que o primeiro objeto a ser inserido na fila é o primeiro a ser lido; nesse mecanismo, conhecido como estrutura FIFO (First In- First Out), a inserção e a remoção são feitas em extremidades contrárias e a estrutura deve possuir um nó com a informação (recorde) e um apontador, respectivamente, para o próximo nó.

**c) a pilha: tipo especial de lista encadeada, na qual o último objeto a ser inserido na fila é o primeiro a ser lido; nesse mecanismo, conhecido como estrutura LIFO (Last In- First Out), a inserção e a remoção são feitas na mesma extremidade e a estrutura deve possuir um nó com a informação (recorde) e um apontador para o próximo nó.**

d) a fila invertida: tipo especial de lista encadeada, tal que o primeiro objeto a ser inserido na fila é o primeiro a ser lido; nesse mecanismo, conhecido como estrutura FIFO (First In- First Out), a inserção e a remoção são feitas em extremidades contrárias e a estrutura deve possuir um nó com a informação (recorde) e um apontador, respectivamente, para o nó anterior.

e) lista circular: tipo especial de lista encadeada, na qual o último elemento tem como próximo o primeiro elemento da lista, formando um ciclo, não havendo diferença entre primeiro e último, e a estrutura deve possuir um nó com a informação (recorde) e um apontador, respectivamente, para o próximo nó.

## Resposta

### Análise dos Requisitos

**Requisitos do aplicativo:**

1. ✅ **Registrar recordes** à medida que são quebrados
2. ✅ **Manter ordem cronológica** dos acontecimentos
3. ✅ **Leitura a partir dos mais recentes** (último inserido primeiro)

### Mapeamento para Estruturas de Dados

**Padrão de uso identificado:**

- **Inserção**: Sempre no "final" (novo recorde)
- **Leitura**: Sempre do "mais recente" (último inserido)
- **Comportamento**: **LIFO** (Last In, First Out)

### Análise de Cada Alternativa

#### a) Deque

❌ **Inadequado**

**Problemas:**

- **Complexidade desnecessária**: Permite inserção/remoção em ambas extremidades
- **Não otimizado**: Para este caso, só precisamos de uma extremidade
- **Overhead**: Dois ponteiros por nó (prev/next) sem necessidade

#### b) Fila (FIFO)

❌ **Inadequado**

**Problemas:**

- **Ordem errada**: FIFO lê os **mais antigos** primeiro
- **Requisito violado**: Precisa ler os **mais recentes** primeiro
- **Comportamento oposto**: Primeiro inserido é primeiro lido

#### c) Pilha (LIFO) ✅

✅ **ADEQUADO**

**Vantagens:**

- **Comportamento correto**: LIFO - último inserido é primeiro lido
- **Atende requisitos**: Leitura dos mais recentes primeiro
- **Simplicidade**: Uma extremidade para inserção/remoção
- **Eficiência**: O(1) para inserção e leitura

**Exemplo:**

```
Pilha LIFO:
Inserção: [Recorde1] → [Recorde2] → [Recorde3] ← (topo)
Leitura:  [Recorde3] ← (mais recente sai primeiro) ✅
```

#### d) Fila Invertida

❌ **Inadequado**

**Problemas:**

- **Conceito confuso**: "Fila invertida" não é uma estrutura padrão
- **Ainda FIFO**: Mesmo invertida, mantém comportamento FIFO
- **Complexidade**: Ponteiro para nó anterior adiciona complexidade

#### e) Lista Circular

❌ **Inadequado**

**Problemas:**

- **Sem extremidades definidas**: "Não há diferença entre primeiro e último"
- **Acesso complexo**: Como identificar o "mais recente"?
- **Overhead**: Estrutura circular desnecessária para este caso

### Implementação da Solução

```go
type Recorde struct {
    atleta string
    tempo  float64
    data   time.Time
}

type PilhaRecordes struct {
    recordes []Recorde
    topo     int
}

// Registrar novo recorde
func (p *PilhaRecordes) RegistrarRecorde(r Recorde) {
    p.recordes = append(p.recordes, r)
    p.topo++
    fmt.Printf("Novo recorde registrado: %s - %.2fs\n", r.atleta, r.tempo)
}

// Ler recordes do mais recente
func (p *PilhaRecordes) LerRecordesMaisRecentes() []Recorde {
    resultado := make([]Recorde, 0)

    // Lê do topo (mais recente) para baixo
    for i := len(p.recordes) - 1; i >= 0; i-- {
        resultado = append(resultado, p.recordes[i])
    }

    return resultado
}
```

### Casos de Uso Similares

**Pilha é ideal para:**

- 📚 Histórico de navegação (browser)
- ↩️ Operações de desfazer (undo)
- 📞 Chamadas de função (call stack)
- 📰 Feed de notícias (mais recentes primeiro)
- 🏆 **Registro de recordes** (este caso)

### Conclusão

### Resposta Correta: Alternativa c) Stack

**Justificativa:**

1. **LIFO** atende perfeitamente o requisito de "leitura dos mais recentes"
2. **Ordem cronológica** é preservada naturalmente
3. **Simplicidade** de implementação e uso
4. **Eficiência** para as operações necessárias
5. **Intuitividade** do modelo mental (empilhar recordes)

A pilha é a estrutura de dados **mais adequada** para este aplicativo específico, pois seu comportamento LIFO alinha perfeitamente com os requisitos de registrar cronologicamente e ler a partir dos mais recentes.
