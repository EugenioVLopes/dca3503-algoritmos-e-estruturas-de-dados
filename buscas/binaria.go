package buscas

func BuscaBinaria(lista []int, elemento int) int {
	esquerda := 0
	direita := len(lista) - 1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if lista[meio] == elemento {
			return meio // Elemento encontrado
		} else if lista[meio] < elemento {
			esquerda = meio + 1 // Busca na metade direita
		} else {
			direita = meio - 1 // Busca na metade esquerda
		}
	}

	return -1 // Elemento não encontrado
}