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

func (node *No) remover(v int) *No {
	if node == nil {
		return nil
	}

	if v < node.Valor {
		node.Esquerdo = node.Esquerdo.remover(v)
	} else if v > node.Valor {
		node.Direito = node.Direito.remover(v)
	} else {
		if node.Esquerdo == nil && node.Direito == nil {
			return nil
		}
		if node.Esquerdo == nil {
			return node.Direito
		}
		if node.Direito == nil {
			return node.Esquerdo
		}

		sucessor := node.Direito.minimo()
		node.Valor = sucessor.Valor
		node.Direito = node.Direito.remover(sucessor.Valor)
	}
	return node
}

func (a *Arvore) Remover(v int) {
	a.Raiz = a.Raiz.remover(v)
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

func criarArvoreBase() *Arvore {
	a := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, v := range valores {
		a.Inserir(v)
	}
	return a
}

func main() {
	// Caso 1: Remoção de nó folha (ex: 35)
	fmt.Println("=== Caso 1: Remoção de Nó Folha (35) ===")
	t1 := criarArvoreBase()
	fmt.Print("Antes da remoção:  ")
	t1.ImprimirEmOrdem()
	t1.Remover(35)
	fmt.Print("Depois da remoção: ")
	t1.ImprimirEmOrdem()
	fmt.Println()

	// Caso 2: Remoção de nó com 1 filho (ex: 40, que possui apenas o filho esquerdo 35)
	fmt.Println("=== Caso 2: Remoção de Nó com 1 Filho (40) ===")
	t2 := criarArvoreBase()
	fmt.Print("Antes da remoção:  ")
	t2.ImprimirEmOrdem()
	t2.Remover(40)
	fmt.Print("Depois da remoção: ")
	t2.ImprimirEmOrdem()
	fmt.Println()

	// Caso 3: Remoção de nó com 2 filhos (ex: 30, que possui filhos 20 e 40)
	fmt.Println("=== Caso 3: Remoção de Nó com 2 Filhos (30) ===")
	t3 := criarArvoreBase()
	fmt.Print("Antes da remoção:  ")
	t3.ImprimirEmOrdem()
	t3.Remover(30)
	fmt.Print("Depois da remoção: ")
	t3.ImprimirEmOrdem()
	fmt.Println()

	// Caso 3 (extra): Remoção da raiz com 2 filhos (50)
	fmt.Println("=== Caso 3 (Raiz): Remoção de Nó com 2 Filhos (50) ===")
	t4 := criarArvoreBase()
	fmt.Print("Antes da remoção:  ")
	t4.ImprimirEmOrdem()
	t4.Remover(50)
	fmt.Print("Depois da remoção: ")
	t4.ImprimirEmOrdem()
}