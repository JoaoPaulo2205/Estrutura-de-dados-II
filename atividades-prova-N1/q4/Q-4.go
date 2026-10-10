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
		n.Esquerdo = n.Esquerdo.Inserir(v) // aqui reatribui certinho
	} else if v > n.Valor {
		n.Direito = n.Direito.Inserir(v) // aqui também
	}
	return n
}

func (t *Arvore) Inserir(v int) {
	t.Raiz = t.Raiz.Inserir(v)
}

func Altura(n *No) int {
	if n == nil {
		return -1
	}
	e := Altura(n.Esquerdo)
	d := Altura(n.Direito)
	if e > d {
		return 1 + e
	}
	return 1 + d
}

func main() {
	t := &Arvore{}
	for _, v := range []int{10, 20, 30, 40, 50} {
		t.Inserir(v)
	}
	fmt.Println(Altura(t.Raiz))
}
