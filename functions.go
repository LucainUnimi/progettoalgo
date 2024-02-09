package main

import (
	"fmt"
	"strings"
)

func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if _, isIn := g.mattoncini[sigma]; !isIn && alpha != beta {
		m := newMattoncino(alpha, beta, sigma)
		g.mattoncini[sigma] = m
		g.forme[alpha].PushBack(m)
		g.forme[beta].PushBack(m)
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
		for f := (*m.fila).Front(); f != nil; f = f.Next() {
			m1, _ := f.Value.(*mattoncino)
			fmt.Printf("%s: %s, %s\n", m1.sigma, m1.alpha, m1.beta)
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
			if f.Len() != 0 && (*f.Back().Value.(*mattoncino)).beta != m.alpha {
				return
			}
			f.PushBack(m)
		} else {
			return
		}
	}
	for p := f.Front(); p != nil; p = p.Next() {
		(*p.Value.(*mattoncino)).fila = &f
	}
}

func sottoStringaMassima(sigma, tau string) string {
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

	return string(lcs)
}

/*
   Se non esiste alcun mattoncino di nome σ oppure se il mattoncino di nome σ non appartiene ad alcuna fila, non compie alcuna operazione.
   Altrimenti stampa l’indice di cacofonia della fila cui appartiene il mattoncino di nome σ.
*/
func indiceCacofonia(g gioco, sigma string) {
	if m, isIn := g.mattoncini[sigma]; isIn && (*m).fila != nil && *(*m).fila != nil {
		var c int
		for f := (*m.fila).Front(); f.Next() != nil; f = f.Next() {
			m, _ = f.Value.(*mattoncino)
			c += len(sottoStringaMassima((*m).sigma, (*m).sigma))
		}
		fmt.Printf("%d\n", c)
	}
}

// elimina fila sfrutta il garbage collector di go per eliminare
// puntatori inutili  QUESTIONE SU UN CONTROLLO
func eliminaFila(g gioco, sigma string) {
	if m, isIn := g.mattoncini[sigma]; isIn && m.fila != nil {
		(*(*m).fila) = nil
	}
}

/*
func disponiFilaMinima(g gioco, alpha, beta string) {
	if starts, isIn := g.forme[alpha]
		isIn {

	}
}
*/
