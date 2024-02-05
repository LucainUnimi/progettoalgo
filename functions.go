package main

import (
	"fmt"
	"strings"
)

func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if _, isIn := g.mattoncini[sigma]; !isIn && alpha != beta {
		m := newMattoncino(alpha, beta, sigma)
		g.mattoncini[sigma] = m
		g.forme[alpha] = append(g.forme[alpha], m)
		g.forme[beta] = append(g.forme[beta], m)
	}
}

func stampaMattoncino(g gioco, sigma string) {
	if value, isIn := g.mattoncini[sigma]; isIn {
		if value.direzione {
			fmt.Printf("%s: %s, %s", value.sigma, value.alpha, value.beta)
		} else {
			fmt.Printf("%s: %s, %s", value.sigma, value.beta, value.alpha)
		}
	}
}

func stampaFila(g gioco, sigma string) {
	if m, exist := g.mattoncini[sigma]; exist && (*m).fila != nil && (*(*m).fila) != nil {
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
		if m, exist := g.mattoncini[sigma]; exist && ((*m).fila == nil || (*(*m).fila) == nil) {
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

func sottostringaMassima(sigma, tau string) (int, string) {
	n, m := len(sigma), len(tau)
	// Creazione di una matrice per memorizzare i risultati dei sottoproblemi
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Riempimento della matrice dp
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if sigma[i-1] == tau[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Ricostruzione della LCS a partire dalla matrice dp
	lcsLength := dp[n][m]
	lcs := make([]rune, lcsLength)
	for i, j := n, m; i > 0 && j > 0; {
		if sigma[i-1] == tau[j-1] {
			lcs[lcsLength-1] = rune(sigma[i-1]) // Aggiunta del carattere comune
			i--
			j--
			lcsLength--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return len(lcs), string(lcs)
}

func indiceCacofonia(g gioco, sigma string) {
	if m, isIn := g.mattoncini[sigma]; isIn && (*m).fila != nil && *(*m).fila != nil {
		var c int
		for f := m.fila.head; f.next.next != nil; f = f.next {
			c += sottoStringMassima(f.value.sigma, f.next.value.sigma)
		}
	}
}

// elimina fila sfrutta il garbage collector di go per eliminare
// puntatori inutili  QUESTIONE SU UN CONTROLLO
func eliminaFila(g gioco, sigma string) {
	if m, isIn := g.mattoncini[sigma]; isIn && m.fila != nil {
		(*(*m).fila) = nil
	}
}
