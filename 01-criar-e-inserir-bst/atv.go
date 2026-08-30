package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

// NovaArvore cria e retorna uma nova árvore binária de busca vazia.
func NovaArvore() *Arvore {
	return &Arvore{Raiz: nil}
}

// NovoNo cria e retorna um novo nó com o valor especificado.
func NovoNo(v int) *No {
	return &No{Valor: v}
}

// inserir insere recursivamente um valor na subárvore enraizada em node, ignorando duplicatas.
func (node *No) inserir(v int) *No {
	if node == nil {
		return NovoNo(v)
	}
	if v < node.Valor {
		node.Esquerdo = node.Esquerdo.inserir(v)
	} else if v > node.Valor {
		node.Direito = node.Direito.inserir(v)
	}
	// Se v == node.Valor, o valor duplicado é ignorado
	return node
}

// Inserir insere um novo valor na árvore binária de busca.
func (a *Arvore) Inserir(v int) {
	a.Raiz = a.Raiz.inserir(v)
}

// EmOrdem percorre a subárvore em ordem simétrica (in-order: Esquerda, Raiz, Direita).
func (node *No) EmOrdem() {
	if node != nil {
		node.Esquerdo.EmOrdem()
		fmt.Printf("%d ", node.Valor)
		node.Direito.EmOrdem()
	}
}

// ImprimirEmOrdem exibe os elementos da árvore em ordem crescente.
func (a *Arvore) ImprimirEmOrdem() {
	if a.Raiz == nil {
		fmt.Println("(árvore vazia)")
		return
	}
	a.Raiz.EmOrdem()
	fmt.Println()
}

func main() {
	arvore := NovaArvore()
	fmt.Print("Árvore inicialmente vazia: ")
	arvore.ImprimirEmOrdem()

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	fmt.Println("\nInserindo valores:", valores)
	for _, v := range valores {
		arvore.Inserir(v)
	}

	fmt.Print("Árvore após inserções (percurso em ordem): ")
	arvore.ImprimirEmOrdem()

	// 3. Teste de inserção de valores duplicados (devem ser ignorados)
	duplicados := []int{50, 30, 65, 80}
	fmt.Println("\nTentando inserir valores duplicados:", duplicados)
	for _, v := range duplicados {
		arvore.Inserir(v)
	}

	fmt.Print("Árvore após tentativa de duplicados: ")
	arvore.ImprimirEmOrdem()
}
