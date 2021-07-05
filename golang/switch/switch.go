package main

import "fmt"

func diaDaSemana(numero uint8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta (Sextou)"
	case 7:
		return "Sabado"

	default:
		return "Numero invalido"
	}

}

func main() {

	// SWITCH

	dia := diaDaSemana(6)

	fmt.Println(dia)
}
