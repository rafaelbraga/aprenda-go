// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.
package main

import (
	"fmt"
)

const (
	_ = iota + 2023
	b
	c
	d
	e
)

func main() {
	fmt.Println(e, b, c, d)

}
