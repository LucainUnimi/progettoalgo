package main

type mattoncino struct {
	direzione          bool
	alpha, beta, sigma string
	fila               *fila
}

type fila = *linkedList

type gioco = map[string]*mattoncino

type linkedList struct {
	head *node
	tail *node
}

type node struct {
	next  *node
	prev  *node
	value *mattoncino
}

func newGioco() *gioco {
	m := make(map[string]*mattoncino, 100)
	return &m
}

func newFila() fila {
	return newList()
}

func newList() *linkedList {
	return &linkedList{nil, nil}
}

func newMattoncino(alpha, beta, sigma string) *mattoncino {
	return &mattoncino{alpha: alpha, beta: beta, sigma: sigma, direzione: true, fila: nil}
}

func newNode(m *mattoncino) *node {
	return &node{nil, nil, m}
}

func swapLati(m *mattoncino) {
	m.alpha, m.beta = m.beta, m.alpha
}

func addLast(l *linkedList, m *mattoncino) {
	n := newNode(m)
	if l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}
}
