package buscas

func BuscaSequencial(lista []int, elemento int) int {
	for i := 0; i < len(lista); i++ {
		if lista[i] == elemento {
			return i
		}
	}
	return -1
}