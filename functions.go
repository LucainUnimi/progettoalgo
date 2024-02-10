package main

import (
	"container/list"
	"fmt"

	"strings"
)

func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if _, isIn := g.mattoncini[sigma]; !isIn && alpha != beta {
		m := newMattoncino(alpha, beta, sigma)
		g.mattoncini[sigma] = m
		if _, isIn := g.forme[alpha]; isIn {
			g.forme[alpha].PushBack(m)
		} else {
			g.forme[alpha] = list.New()
			g.forme[alpha].PushBack(m)
		}
		if _, isIn := g.forme[beta]; isIn {
			g.forme[beta].PushBack(m)
		} else {
			g.forme[beta] = list.New()
			g.forme[beta].PushBack(m)
		}
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
		if m, exist := g.mattoncini[sigma]; exist && ((*m).fila == nil || *((*m).fila) == nil) {
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
			c += len(sottoStringaMassima(f.Value.(*mattoncino).sigma, f.Next().Value.(*mattoncino).sigma))
		}
		fmt.Printf("%d\n", c)
	}
}

// elimina fila sfrutta il garbage collector di go per eliminare
// puntatori inutili  QUESTIONE SU UN CONTROLLO
func eliminaFila(g gioco, sigma string) {
	if m, isIn := g.mattoncini[sigma]; isIn && m.fila != nil {
		*((*m).fila) = nil
	}
}

/*
Crea e posiziona sul tavolo da gioco una fila di lunghezza minima da α a β. Tutti i mattoncini
della fila devono essere presi dalla scatola. Se non è possibile creare alcuna fila da α a β, stampa il
messaggio: non esiste fila da α a β
*/

func disponiFilaMinima(g gioco, alpha, beta string) {
	_, isInAlpha := g.forme[alpha]
	_, isInBeta := g.forme[beta]
	if !(isInAlpha && isInBeta) {
		return
	}

	if alpha == beta {
		adjs := g.forme[alpha].Front()
		pathLen := len(g.mattoncini)
		var pathMin string

		for ; adjs != nil; adjs = adjs.Next() {
			m := adjs.Value.(*mattoncino)
			var s string
			l := len(g.mattoncini)
			visitedArch := make(map[*mattoncino]bool)
			visitedArch[m] = true
			if (*m).alpha == alpha && ((*m).fila == nil || *((*m).fila) == nil) {
				l, s = BFSCamminoMinimo(g, (*m).beta, beta, visitedArch)
			} else if (*m).beta == alpha && ((*m).fila == nil || *((*m).fila) == nil) {
				l, s = BFSCamminoMinimo(g, (*m).alpha, beta, visitedArch)
			}
			if l <= pathLen {
				pathLen = l
				pathMin = s
				if (m.alpha == alpha && (*m).direzione) || (!(m.alpha == alpha) && !((*m).direzione)) {
					pathMin = "+" + (*m).sigma + " " + pathMin
				} else {
					pathMin = "-" + (*m).sigma + " " + pathMin
				}
			}
		}
		if pathLen != 0 {
			disponiFila(g, pathMin)
		} else {
			fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
		}
	} else {
		c, path := BFSCamminoMinimo(g, alpha, beta, make(map[*mattoncino]bool))
		if c == 0 {
			fmt.Printf("non esiste fila da %s a %s\n", alpha, beta)
		} else {
			disponiFila(g, path)
		}
	}

}

func BFSCamminoMinimo(g gioco, alpha, beta string, visitedArch map[*mattoncino]bool) (int, string) {
	visited := make(map[string]string)
	queue := list.New()
	queue.PushBack(alpha)
	visited[alpha] = "Rt%JV+3*tFN3=Lvxj-SG"

	for queue.Len() != 0 {
		curr := queue.Remove(queue.Front()).(string)
		if g.forme[curr] == nil {
			fmt.Println(curr, g.forme[curr])
			return 0, ""
		}
		adjs := g.forme[curr].Front()
		for ; adjs != nil; adjs = adjs.Next() {
			m := adjs.Value.(*mattoncino)
			if m.alpha == curr && visited[m.beta] == "" && !visitedArch[m] && ((*m).fila == nil || *(*m).fila == nil) {
				visited[m.beta] = m.sigma
				queue.PushBack(m.beta)
				if m.beta == beta {
					break
				}
			} else if m.beta == curr && visited[m.alpha] == "" && !visitedArch[m] && ((*m).fila == nil || *(*m).fila == nil) {
				visited[m.alpha] = m.sigma
				queue.PushBack(m.alpha)
				if m.alpha == beta {
					break
				}
			}
			visitedArch[m] = true
		}
	}
	var c int
	if visited[beta] != "" {
		fila := ""
		key := beta
		for key != alpha {
			c++
			s := visited[key]
			m := g.mattoncini[s]
			if m.beta == key {
				if (*m).direzione {
					fila = "+" + s + " " + fila
				} else {
					fila = "-" + s + " " + fila
				}
				key = m.alpha
			} else {
				if (*m).direzione {
					fila = " -" + s + " " + fila
				} else {
					fila = " +" + s + " " + fila
				}
				key = m.beta
			}
		}
		return c, fila
	}
	return c, ""
}
