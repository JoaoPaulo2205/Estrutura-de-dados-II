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

func (node *No) buscar(v int) *No {
	if node == nil || node.Valor == v {
		return node
	}
	if v < node.Valor {
		return node.Esquerdo.buscar(v)
	}
	return node.Direito.buscar(v)
}

func (a *Arvore) Buscar(v int) *No {
	return a.Raiz.buscar(v)
}

func (a *Arvore) BuscarIter(v int) *No {
	atual := a.Raiz
	for atual != nil {
		if v == atual.Valor {
			return atual
		} else if v < atual.Valor {
			atual = atual.Esquerdo
		} else {
			atual = atual.Direito
		}
	}
	return nil
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

func testarBusca(a *Arvore, v int) {
	fmt.Printf("\n--- Testando busca pelo valor %d ---\n", v)

	resRec := a.Buscar(v)
	if resRec != nil {
		fmt.Printf("[Recursivo] Valor %d ENCONTRADO no nó (%p) com valor %d\n", v, resRec, resRec.Valor)
	} else {
		fmt.Printf("[Recursivo] Valor %d NÃO ENCONTRADO na árvore\n", v)
	}

	resIter := a.BuscarIter(v)
	if resIter != nil {
		fmt.Printf("[Iterativo] Valor %d ENCONTRADO no nó (%p) com valor %d\n", v, resIter, resIter.Valor)
	} else {
		fmt.Printf("[Iterativo] Valor %d NÃO ENCONTRADO na árvore\n", v)
	}
}

func main() {
	arvore := NovaArvore()

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	fmt.Println("Inserindo valores na árvore:", valores)
	for _, v := range valores {
		arvore.Inserir(v)
	}

	fmt.Print("Árvore em ordem: ")
	arvore.ImprimirEmOrdem()

	testarBusca(arvore, 65)
	testarBusca(arvore, 45)
}
