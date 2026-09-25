package main

import "fmt"

type No struct {
	Valor  int
	Altura int
	Esq    *No
	Dir    *No
}

func altura(no *No) int {
	if no == nil {
		return -1
	}
	return no.Altura
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fatorBalanceamento(no *No) int {
	if no == nil {
		return 0
	}
	return altura(no.Esq) - altura(no.Dir)
}

func atualizarAltura(no *No) {
	no.Altura = 1 + max(altura(no.Esq), altura(no.Dir))
}

func rotacaoDireita(y *No) *No {
	x := y.Esq
	y.Esq = x.Dir
	x.Dir = y
	atualizarAltura(y)
	atualizarAltura(x)

	return x
}

func rotacaoEsquerda(x *No) *No {
	y := x.Dir
	x.Dir = y.Esq
	y.Esq = x
	atualizarAltura(x)
	atualizarAltura(y)

	return y
}

func rotacaoLR(no *No) *No {
	no.Esq = rotacaoEsquerda(no.Esq)
	return rotacaoDireita(no)
}

func rotacaoRL(no *No) *No {
	no.Dir = rotacaoDireita(no.Dir)
	return rotacaoEsquerda(no)
}

func balancear(no *No) *No {
	atualizarAltura(no)
	fator := fatorBalanceamento(no)

	if fator > 1 {
		if fatorBalanceamento(no.Esq) < 0 {
			no.Esq = rotacaoEsquerda(no.Esq)
		}
		return rotacaoDireita(no)
	}
	if fator < -1 {
		if fatorBalanceamento(no.Dir) > 0 {
			no.Dir = rotacaoDireita(no.Dir)
		}
		return rotacaoEsquerda(no)
	}

	return no
}

func inserir(no *No, chave int) *No {
	if no == nil {
		novoNo := new(No)
		novoNo.Valor = chave
		novoNo.Altura = 0
		return novoNo
	}
	if chave < no.Valor {
		no.Esq = inserir(no.Esq, chave)
	} else if chave > no.Valor {
		no.Dir = inserir(no.Dir, chave)
	} else {
		return no
	}
	atualizarAltura(no)

	return balancear(no)
}

func inserirBST(no *No, chave int) *No {
	if no == nil {
		novoNo := new(No)
		novoNo.Valor = chave
		novoNo.Altura = 0
		return novoNo
	}

	if chave < no.Valor {
		no.Esq = inserirBST(no.Esq, chave)
	} else if chave > no.Valor {
		no.Dir = inserirBST(no.Dir, chave)
	}

	atualizarAltura(no)

	return no
}

func imprimirArvore(n *No, prefixo string) {
	if n != nil {
		fb := altura(n.Esq) - altura(n.Dir)
		fmt.Printf("%sNó %d (Altura: %d, FB: %+d)\n", prefixo, n.Valor, n.Altura, fb)
		imprimirArvore(n.Esq, prefixo+"  Esq-> ")
		imprimirArvore(n.Dir, prefixo+"  Dir-> ")
	}
}

func main() {
	var raiz *No

	raiz = inserirBST(raiz, 50)
	raiz = inserirBST(raiz, 40)
	raiz = inserirBST(raiz, 30)
	raiz = inserirBST(raiz, 20)
	raiz = inserirBST(raiz, 10)
	fmt.Println("Exercicio 01")

	fmt.Print("------- Arvore Degenerada -------\n")
	imprimirArvore(raiz, "")
	fmt.Print("\n")

	raiz = rotacaoDireita(raiz)

	fmt.Print("------- Arvore Com Rotação a Direita -------\n")
	imprimirArvore(raiz, "")
	fmt.Print("\n")

	fmt.Println("-------------------------------------------------------------")

	var raiz2 *No

	raiz2 = inserirBST(raiz2, 20)
	raiz2 = inserirBST(raiz2, 40)
	raiz2 = inserirBST(raiz2, 30)
	raiz2 = inserirBST(raiz2, 60)
	raiz2 = inserirBST(raiz2, 70)

	fmt.Println("Exercicio 02")
	fmt.Print("------- Arvore Degenerada -------\n")
	imprimirArvore(raiz2, "")
	fmt.Print("\n")

	raiz2 = rotacaoEsquerda(raiz2)

	fmt.Print("------- Arvore Com Rotação a Esquerda -------\n")
	imprimirArvore(raiz2, "")
	fmt.Print("\n")

	fmt.Println("-------------------------------------------------------------")

	var raiz3 *No

	raiz3 = inserirBST(raiz3, 60)
	raiz3 = inserirBST(raiz3, 20)
	raiz3 = inserirBST(raiz3, 10)
	raiz3 = inserirBST(raiz3, 40)
	raiz3 = inserirBST(raiz3, 30)

	fmt.Println("Exercicio 03")
	fmt.Print("------- Arvore Degenerada -------\n")
	imprimirArvore(raiz3, "")
	fmt.Print("\n")

	raiz3 = rotacaoLR(raiz3)

	fmt.Print("------- Arvore Com Rotação Dupla a Direita (LR) -------\n")
	imprimirArvore(raiz3, "")
	fmt.Print("\n")

	fmt.Println("-------------------------------------------------------------")

	var raiz4 *No

	raiz4 = inserirBST(raiz4, 20)
	raiz4 = inserirBST(raiz4, 60)
	raiz4 = inserirBST(raiz4, 40)
	raiz4 = inserirBST(raiz4, 70)
	raiz4 = inserirBST(raiz4, 50)

	fmt.Println("Exercicio 04")
	fmt.Print("------- Arvore Degenerada -------\n")
	imprimirArvore(raiz4, "")
	fmt.Print("\n")

	raiz4 = rotacaoRL(raiz4)

	fmt.Print("------- Arvore Com Rotação Dupla a Esquerda (RL) -------\n")
	imprimirArvore(raiz4, "")
	fmt.Print("\n")
}
