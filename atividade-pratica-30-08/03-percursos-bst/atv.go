package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovaArvore() *Arvore {
	return &Arvore{Raiz: nil}
}

func NovoNo(v int) *No {
	return &No{Valor: v}
}

func (node *No) inserir(v int) *No {
	if node == nil {
		return NovoNo(v)
	}
	if v < node.Valor {
		node.Esquerdo = node.Esquerdo.inserir(v)
	} else if v > node.Valor {
		node.Direito = node.Direito.inserir(v)
	}
	return node
}

func (a *Arvore) Inserir(v int) {
	a.Raiz = a.Raiz.inserir(v)
}

func (node *No) preOrdem() {
	if node != nil {
		fmt.Printf("%d ", node.Valor)
		node.Esquerdo.preOrdem()
		node.Direito.preOrdem()
	}
}

func (a *Arvore) PreOrdem() {
	a.Raiz.preOrdem()
	fmt.Println()
}

func (node *No) emOrdem() {
	if node != nil {
		node.Esquerdo.emOrdem()
		fmt.Printf("%d ", node.Valor)
		node.Direito.emOrdem()
	}
}

func (a *Arvore) EmOrdem() {
	a.Raiz.emOrdem()
	fmt.Println()
}

func (node *No) posOrdem() {
	if node != nil {
		node.Esquerdo.posOrdem()
		node.Direito.posOrdem()
		fmt.Printf("%d ", node.Valor)
	}
}

func (a *Arvore) PosOrdem() {
	a.Raiz.posOrdem()
	fmt.Println()
}

func (a *Arvore) EmLargura() {
	if a.Raiz == nil {
		fmt.Println()
		return
	}

	fila := []*No{a.Raiz}
	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		fmt.Printf("%d ", atual.Valor)

		if atual.Esquerdo != nil {
			fila = append(fila, atual.Esquerdo)
		}
		if atual.Direito != nil {
			fila = append(fila, atual.Direito)
		}
	}
	fmt.Println()
}

/*
Por que o percurso EmOrdem produz valores crescentes em uma BST:
Pela definição de uma Árvore Binária de Busca (BST), para qualquer nó N:
- Todos os nós da subárvore esquerda têm valores estritamente menores que N.Valor.
- Todos os nós da subárvore direita têm valores estritamente maiores que N.Valor.

O percurso EmOrdem visita os nós na ordem: (1) Subárvore Esquerda -> (2) Raiz -> (3) Subárvore Direita.
Aplicando essa regra recursivamente em cada nó, primeiro são visitados todos os valores menores que a raiz atual, em seguida a própria raiz, e por último todos os valores maiores. Consequentemente, todos os elementos são processados em ordem estritamente crescente.
*/

func main() {
	arvore := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, v := range valores {
		arvore.Inserir(v)
	}

	fmt.Print("Pré-Ordem:  ")
	arvore.PreOrdem()

	fmt.Print("Em-Ordem:   ")
	arvore.EmOrdem()

	fmt.Print("Pós-Ordem:  ")
	arvore.PosOrdem()

	fmt.Print("Em-Largura: ")
	arvore.EmLargura()
}
