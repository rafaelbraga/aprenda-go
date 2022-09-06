// Usando uma literal composta:
//     - Crie uma slice de tipo int
//     - Atribua 10 valores a ela
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.

package main

import "fmt"

func main() {
	array := []int{00, 10, 20, 30, 40, 50, 60, 70, 90, 80}

	for i, num := range array {
		println(i, " ", num)

	}
	fmt.Printf("%T", array)

}
