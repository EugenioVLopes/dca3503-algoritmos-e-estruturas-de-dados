# Análise de Código para Automação Industrial

## Pergunta

Estruturas de dados, tais como filas e pilhas, são utilizadas em diversas aplicações para automação industrial por meio de linguagens de programação textuais. O texto estruturado (ST) é uma das opções de linguagem de programação definidas pela norma IEC 61131-3. O trecho de código a seguir foi implementado nesse contexto.

```c
#define MAX 1000

struct eventos {
    char ocorrencia[200];
    char dataHora[50];
};

struct eventos eve[MAX];
int inicio = 0;
int fim = 0;

int processaEvento (struct eventos *recuperado) {
    if(inicio == fim) {
        return -1;
    }
    else {
        inicio++;
        copiaEvento (recuperado, eve[inicio - 1]);
        return 0;
    }
}

int insereEvento (struct eventos *novo) {
    if (fim == MAX) {
        return -1;
    }
    else {
        copiaEvento (eve[fim], novo);
        fim++;
        return 0;
    }
}
```

É correto afirmar que a estrutura de dados e a funcionalidade desse código tratam-se de:

**Alternativas:**
**a) uma fila que processa primeiro os eventos mais antigos.**
b) uma pilha que processa primeiro os eventos mais antigos.
c) uma pilha que processa primeiro os eventos mais recentes.
d) uma pilha que processa os eventos na ordem escolhida pelo operador.
e) uma fila que processa os eventos de acordo com seu respectivo grau de prioridade.

## Resposta

### Análise do Código

**Estrutura implementada:**

- Array `eve[MAX]` para armazenar eventos
- Variável `inicio` para controlar remoção
- Variável `fim` para controlar inserção
- **Comportamento FIFO** (First In, First Out)

### Funcionamento Detalhado

#### Inserção de Eventos (`insereEvento`)

```c
int insereEvento (struct eventos *novo) {
    if (fim == MAX) {
        return -1;  // Fila cheia
    }
    else {
        copiaEvento (eve[fim], novo);  // Insere no final
        fim++;                         // Avança ponteiro fim
        return 0;
    }
}
```

#### Processamento de Eventos (`processaEvento`)

```c
int processaEvento (struct eventos *recuperado) {
    if(inicio == fim) {
        return -1;  // Fila vazia
    }
    else {
        inicio++;                                    // Avança ponteiro início
        copiaEvento (recuperado, eve[inicio - 1]);  // Remove do início
        return 0;
    }
}
```

### Análise das Alternativas

#### a) Fila que processa eventos mais antigos (FIFO) ✅

✅ **CORRETO - Comportamento do código**

**Evidências no código:**

- **Inserção**: `eve[fim]` - adiciona no final (índice `fim`)
- **Remoção**: `eve[inicio - 1]` - remove do início (índice `inicio`)
- **Comportamento FIFO**: Primeiro a entrar, primeiro a sair
- **Ordem cronológica**: Eventos mais antigos são processados primeiro

#### b) Pilha que processa eventos mais antigos

❌ **Contraditório**

**Problema conceitual:**

- **Pilha** é LIFO (Last In, First Out)
- **Não pode** processar "mais antigos" primeiro
- **Contradição** na definição

#### c) Pilha que processa eventos mais recentes (LIFO)

❌ **Não corresponde ao código**

**Problemas:**

- **Código não é pilha**: Usa dois ponteiros (`inicio` e `fim`)
- **Comportamento diferente**: Pilha usaria apenas um ponteiro (topo)
- **Ordem errada**: Código processa antigos primeiro, não recentes

#### d) Pilha com ordem escolhida pelo operador

❌ **Não é comportamento do código**

**Problemas:**

- **Código não é pilha**: Estrutura é claramente uma fila
- **Ordem fixa**: Não há escolha do operador, ordem é FIFO
- **Comportamento determinístico**: Sempre processa do início

#### e) Fila com prioridade

❌ **Não há sistema de prioridades**

**Problemas:**

- **Sem prioridades**: Código não verifica prioridade dos eventos
- **FIFO simples**: Processa na ordem de chegada
- **Estrutura básica**: Não há heap ou ordenação por prioridade

### Casos de Uso para Fila FIFO em Automação

1. **Log de Eventos**

   - Mantém histórico cronológico
   - Facilita auditoria e rastreamento

2. **Processamento Sequencial**

   - Tarefas devem ser executadas em ordem
   - Evita conflitos de dependência

3. **Buffer de Comunicação**

   - Mensagens processadas na ordem de chegada
   - Garante integridade da comunicação

4. **Controle de Produção**
   - Peças processadas na ordem de chegada
   - Mantém fluxo organizado

### Resumo da Análise

**Estrutura identificada:** Fila FIFO (First In, First Out)

**Características principais:**

- **Array estático**: `struct eventos eve[MAX]`
- **Dois ponteiros**: `inicio` (remoção) e `fim` (inserção)
- **Inserção**: No final da fila (`eve[fim]`)
- **Remoção**: Do início da fila (`eve[inicio-1]`)
- **Comportamento**: Primeiro inserido é primeiro processado

**Aplicação em automação:**
Esta implementação é adequada para sistemas que precisam processar eventos na **ordem cronológica** de ocorrência, mantendo a sequência temporal dos acontecimentos.

### Conclusão

**Resposta correta: Alternativa a) uma fila que processa primeiro os eventos mais antigos.**

**Justificativa baseada na análise do código:**

1. **Estrutura de fila**: Código usa dois ponteiros (`inicio` e `fim`) característicos de fila
2. **Comportamento FIFO**: Inserção no final (`eve[fim]`) e remoção do início (`eve[inicio-1]`)
3. **Ordem cronológica**: Eventos mais antigos são processados primeiro
4. **Implementação clara**: Funções `insereEvento` e `processaEvento` demonstram comportamento FIFO
5. **Evidência no código**: `inicio++` e `fim++` mostram movimento unidirecional típico de fila

O código implementa uma **fila FIFO clássica** onde eventos são inseridos no final e removidos do início, garantindo que o **primeiro evento inserido seja o primeiro a ser processado**. Esta é uma implementação padrão para sistemas que precisam manter **ordem cronológica** de processamento.
