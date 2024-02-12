package main

import "container/list"

type mattoncino struct {
	direzione          bool
	alpha, beta, sigma string
	fila               *fila
}

type fila = *list.List

type giocoStruct struct {
	mattoncini map[string]*mattoncino
	forme      map[string]*list.List
}

type gioco = *giocoStruct

func newGioco() gioco {
	return &giocoStruct{make(map[string]*mattoncino), make(map[string]*list.List)}
}

func newFila() fila {
	return list.New()
}

func newMattoncino(alpha, beta, sigma string) *mattoncino {
	return &mattoncino{alpha: alpha, beta: beta, sigma: sigma, direzione: true, fila: nil}
}

func swapLati(m *mattoncino) {
	m.alpha, m.beta = m.beta, m.alpha
}
