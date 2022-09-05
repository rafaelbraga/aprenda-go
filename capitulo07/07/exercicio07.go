// Utilizando a solução anterior, adicione as opções else if e else.
package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		if i%4 == 0 {
			fmt.Printf("%d é divisivel por 4\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d não é divisivel por 4 mas é por 2\n", i)
		} else {
			fmt.Printf("%d não é divisivel por 4 nem 2\n", i)
		}

	}

}
