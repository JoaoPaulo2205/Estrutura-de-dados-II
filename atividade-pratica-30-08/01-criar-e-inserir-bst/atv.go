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

func (node *No) EmOrdem() {
	if node != nil {
		node.Esquerdo.EmOrdem()
		fmt.Printf("%d ", node.Valor)
		node.Direito.EmOrdem()
	}
}

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

	duplicados := []int{50, 30, 65, 80}
	fmt.Println("\nTentando inserir valores duplicados:", duplicados)
	for _, v := range duplicados {
		arvore.Inserir(v)
	}

	fmt.Print("Árvore após tentativa de duplicados: ")
	arvore.ImprimirEmOrdem()
}
