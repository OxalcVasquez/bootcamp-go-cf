package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Println("Antes del swap")
	fmt.Println("x = ", x, "y = ", y)

	swap(&x, &y)

	fmt.Println("Despues del swap")
	fmt.Println("x = ", x, "y = ", y)

}
