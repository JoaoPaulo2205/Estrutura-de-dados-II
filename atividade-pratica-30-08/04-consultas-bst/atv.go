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

func (node *No) altura() int {
	if node == nil {
		return -1
	}
	esq := node.Esquerdo.altura()
	dir := node.Direito.altura()
	if esq > dir {
		return 1 + esq
	}
	return 1 + dir
}

func (a *Arvore) Altura() int {
	return a.Raiz.altura()
}

func (node *No) contar() int {
	if node == nil {
		return 0
	}
	return 1 + node.Esquerdo.contar() + node.Direito.contar()
}

func (a *Arvore) Contar() int {
	return a.Raiz.contar()
}

func (node *No) contarFolhas() int {
	if node == nil {
		return 0
	}
	if node.Esquerdo == nil && node.Direito == nil {
		return 1
	}
	return node.Esquerdo.contarFolhas() + node.Direito.contarFolhas()
}

func (a *Arvore) ContarFolhas() int {
	return a.Raiz.contarFolhas()
}

func (node *No) minimo() *No {
	if node == nil {
		return nil
	}
	atual := node
	for atual.Esquerdo != nil {
		atual = atual.Esquerdo
	}
	return atual
}

func (a *Arvore) Minimo() *No {
	return a.Raiz.minimo()
}

func (node *No) maximo() *No {
	if node == nil {
		return nil
	}
	atual := node
	for atual.Direito != nil {
		atual = atual.Direito
	}
	return atual
}

func (a *Arvore) Maximo() *No {
	return a.Raiz.maximo()
}

func (node *No) emOrdem() {
	if node != nil {
		node.Esquerdo.emOrdem()
		fmt.Printf("%d ", node.Valor)
		node.Direito.emOrdem()
	}
}

func (a *Arvore) ImprimirEmOrdem() {
	if a.Raiz == nil {
		fmt.Println("(árvore vazia)")
		return
	}
	a.Raiz.emOrdem()
	fmt.Println()
}

func exibirConsultas(a *Arvore, nome string) {
	fmt.Printf("=== %s ===\n", nome)
	fmt.Print("Elementos (em ordem): ")
	a.ImprimirEmOrdem()
	fmt.Printf("Altura: %d\n", a.Altura())
	fmt.Printf("Total de nós: %d\n", a.Contar())
	fmt.Printf("Total de folhas: %d\n", a.ContarFolhas())

	if min := a.Minimo(); min != nil {
		fmt.Printf("Mínimo: %d\n", min.Valor)
	} else {
		fmt.Println("Mínimo: nil (árvore vazia)")
	}

	if max := a.Maximo(); max != nil {
		fmt.Printf("Máximo: %d\n", max.Valor)
	} else {
		fmt.Println("Máximo: nil (árvore vazia)")
	}
	fmt.Println()
}

func main() {
	vazia := NovaArvore()
	exibirConsultas(vazia, "Árvore Vazia")

	umNo := NovaArvore()
	umNo.Inserir(50)
	exibirConsultas(umNo, "Árvore com 1 Nó")

	arvore := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, v := range valores {
		arvore.Inserir(v)
	}
	exibirConsultas(arvore, "Árvore Completa")
}
