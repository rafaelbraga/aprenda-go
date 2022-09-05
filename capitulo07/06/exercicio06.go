// Crie um programa que demonstre o funcionamento da declaração if.
package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		if i%4 == 0 {
			fmt.Printf("%d é divisivel por 4\n", i)

		}

	}

}
