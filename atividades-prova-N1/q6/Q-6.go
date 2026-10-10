package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func (n *No) Inserir(v int) *No {
	if n == nil {
		return &No{Valor: v}
	}
	if v < n.Valor {
		n.Esquerdo = n.Esquerdo.Inserir(v)
	} else if v > n.Valor {
		n.Direito = n.Direito.Inserir(v)
	}
	return n
}

func (t *Arvore) Inserir(v int) {
	t.Raiz = t.Raiz.Inserir(v)
}

func (n *No) Minimo() *No {
	if n == nil || n.Esquerdo == nil {
		return n
	}
	return n.Esquerdo.Minimo()
}

func (n *No) Remover(v int) *No {
	if n == nil {
		return nil
	}
	if v < n.Valor {
		n.Esquerdo = n.Esquerdo.Remover(v)
	} else if v > n.Valor {
		n.Direito = n.Direito.Remover(v)
	} else {
		if n.Esquerdo == nil {
			return n.Direito
		}
		if n.Direito == nil {
			return n.Esquerdo
		}
		suc := n.Direito.Minimo()
		n.Valor = suc.Valor
		n.Direito = n.Direito.Remover(suc.Valor)
	}
	return n
}

func (n *No) PreOrdem(f func(int)) {
	if n == nil {
		return
	}
	f(n.Valor)
	n.Esquerdo.PreOrdem(f)
	n.Direito.PreOrdem(f)
}

func main() {
	t := &Arvore{}
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80, 35, 45} {
		t.Inserir(v)
	}
	t.Raiz = t.Raiz.Remover(30)
	t.Raiz.PreOrdem(func(x int) { fmt.Print(x, " ") })
}
