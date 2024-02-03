package main

import (
	"fmt"
	"strings"
)

func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if _, isIn := g[sigma]; !isIn && alpha != beta {
		g[sigma] = newMattoncino(alpha, beta, sigma)
	}
}

func stampaMattoncino(g gioco, sigma string) {
	if value, isIn := g[sigma]; isIn {
		if value.direzione {
			fmt.Printf("%s: %s, %s", value.sigma, value.alpha, value.beta)
		} else {
			fmt.Printf("%s: %s, %s", value.sigma, value.beta, value.alpha)
		}
	}
}

func stampaFila(g gioco, sigma string) {
	if m, exist := g[sigma]; exist && (*m).fila != nil && (*(*m).fila) != nil {
		fmt.Println("(")
		for f := (*m.fila).head; f != nil; f = f.next {
			fmt.Printf("\t%s: %s, %s\n", f.value.sigma, f.value.alpha, f.value.beta)
		}
		fmt.Println(")")
	}
}

func disponiFila(g gioco, listaNomi string) {
	var f fila
	f = newFila()
	for _, s := range strings.Fields(listaNomi) {
		sigma := s[1:]
		if m, exist := g[sigma]; exist && ((*m).fila == nil || (*(*m).fila) == nil) {
			if !(s[0] == '+') != !m.direzione {
				m.direzione = false
				swapLati(m)
			}
			if f.tail != nil && f.tail.value.beta != m.alpha {
				return
			}
			addLast(f, m)
		} else {
			return
		}
	}
	for p := f.head; p != nil; p = p.next {
		p.value.fila = &f
	}
}

// elimina fila sfrutta il garbage collector di go per eliminare
// puntatori inutili
func eliminaFila(g gioco, sigma string) {
	if m, isIn := g[sigma]; isIn && m.fila != nil {
		(*(*m).fila) = nil
	}
}
