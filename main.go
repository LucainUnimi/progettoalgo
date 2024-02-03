package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	g := newGioco()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		operation, operand, _ := strings.Cut(scanner.Text(), " ")
		switch operation {
		case "m":
			names := strings.Split(operand, " ")
			inserisciMattoncino(*g, names[0], names[1], names[2])
		case "s":
			stampaMattoncino(*g, operand)
			fmt.Println()
		case "d":
			listaNomi := operand //per esempio "+ciao -cane -gatto +macchina"
			disponiFila(*g, listaNomi)
		case "S":
			stampaFila(*g, operand)
		case "e":
			eliminaFila(*g, operand)
		case "q":
			return
		}
	}
}
