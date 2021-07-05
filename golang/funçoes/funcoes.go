package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subs := n1 - n2
	return sum, subs
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}

	result := f("testando variavel com função")
	fmt.Println(result)

	resultado1, resultado2 := calculosMatematicos(100, 90)
	fmt.Println(resultado1, resultado2)

	_, resultadoSub := calculosMatematicos(50, 80)

	resultadoSum, _ := calculosMatematicos(50, 80)
	fmt.Println(resultadoSub)
	fmt.Println(resultadoSum)

	fmt.Println(calculosMatematicos(12, 14))

}
