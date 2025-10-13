package main

import (
	"fmt"
	zlp "zlp/Zalupki"
)

func main() {
	a := zlp.Sum(5, 6)
	fmt.Println("Hello Zalupki, 5 + 6 =", a)

	zlp.Out1()

	zlp.Out2() // Надо активировать сане эту функицю, почему-то не работает
}
