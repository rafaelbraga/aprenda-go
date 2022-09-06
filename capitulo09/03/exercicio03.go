// Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
//     - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//     - Do quinto ao último item do slice (incluindo o último item!)
//     - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//     - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//     - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main() {
	array := []int{00, 10, 20, 30, 40, 50, 60, 70, 90, 80}

	fmt.Println(array[:3])
	fmt.Println(array[4:])
	fmt.Println(array[1:7])
	fmt.Println(array[2 : len(array)-1])

}
