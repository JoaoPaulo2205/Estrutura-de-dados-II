package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

func EmOrdem(n *No, f func(int)) {
	if n == nil {
		return
	}
	EmOrdem(n.Esquerdo, f)
	f(n.Valor)
	EmOrdem(n.Direito, f)
}

func PreOrdem(n *No, f func(int)) {
	if n == nil {
		return
	}
	f(n.Valor)
	PreOrdem(n.Esquerdo, f)
	PreOrdem(n.Direito, f)
}

func PosOrdem(n *No, f func(int)) {
	if n == nil {
		return
	}
	PosOrdem(n.Esquerdo, f)
	PosOrdem(n.Direito, f)
	f(n.Valor)
}

func main() {
	n4 := &No{Valor: 4}
	n5 := &No{Valor: 5}
	n2 := &No{Valor: 2, Esquerdo: n4, Direito: n5}
	n3 := &No{Valor: 3}
	raiz := &No{Valor: 1, Esquerdo: n2, Direito: n3}

	fmt.Print("Em-ordem: ")
	EmOrdem(raiz, func(x int) { fmt.Print(x, " ") })
	fmt.Println()

	fmt.Print("Pré-ordem: ")
	PreOrdem(raiz, func(x int) { fmt.Print(x, " ") })
	fmt.Println()

	fmt.Print("Pós-ordem: ")
	PosOrdem(raiz, func(x int) { fmt.Print(x, " ") })
	fmt.Println()
}
