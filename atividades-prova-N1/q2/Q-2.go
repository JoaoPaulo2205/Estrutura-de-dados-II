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
		n.Esquerdo.Inserir(v) // bug: não reatribui o retorno
	} else if v > n.Valor {
		n.Direito.Inserir(v) // bug: não reatribui o retorno
	}
	return n
}

func (t *Arvore) Inserir(v int) {
	t.Raiz = t.Raiz.Inserir(v)
}

func (n *No) EmOrdem(f func(int)) {
	if n == nil {
		return
	}
	n.Esquerdo.EmOrdem(f)
	f(n.Valor)
	n.Direito.EmOrdem(f)
}

func main() {
	t := &Arvore{}
	for _, v := range []int{50, 30, 70, 20, 40} {
		t.Inserir(v)
	}
	t.Raiz.EmOrdem(func(x int) { fmt.Print(x, " ") })
}
