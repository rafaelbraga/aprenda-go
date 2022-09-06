// Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array
// - Utilizando format printing, demonstre o tipo do array

package main

import "fmt"

func main() {
	array := [5]int{00, 10, 20, 30, 40}

	for i, num := range array {
		println(i, " ", num)

	}
	fmt.Printf("%T", array)

}
