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
			inserisciMattoncino(g, names[0], names[1], names[2])
		case "s":
			stampaMattoncino(g, operand)
			fmt.Println()
		case "d":
			listaNomi := operand
			disponiFila(g, listaNomi)
		case "S":
			stampaFila(g, operand)
		case "e":
			eliminaFila(g, operand)
		case "i":
			indiceCacofonia(g, operand)
		case "M":
			names := strings.Split(operand, " ")
			fmt.Println(string(sottoSeqMassima([]rune(names[0]), []rune(names[1]))))
		case "f":
			names := strings.Split(operand, " ")
			disponiFilaMinima(g, names[0], names[1])
		case "c":
			name, shapes, _ := strings.Cut(operand, " ")
			costo(g, name, shapes)
		case "q":
			return
		}
	}
}
