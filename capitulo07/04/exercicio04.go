// Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

package main

func main() {
	bornYear := 1991
	todayYear := 2022

	for {
		println(bornYear)
		bornYear++
		if bornYear == todayYear {
			break
		}
	}
}
