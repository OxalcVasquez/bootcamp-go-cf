package main

import "fmt"

func main() {

	testCards := []string{
		"4532015112830366",
		"4532015112830367",
	}

	for _, card := range testCards {
		isValid := luhnCheck(card)

		if isValid {
			fmt.Printf("%s es valida \n", card)
		} else {
			fmt.Printf("%s es invalida \n", card)
		}
	}

}

func luhnCheck(cardNumber string) bool {
	sum := 0

	//Recorrir digitos de derecha a izquierda
	for i := (len(cardNumber) - 1); i >= 0; i-- {
		digit := int(cardNumber[i] - '0') //De forma explicita

		//Duplicar cada segundo digitio
		if (len(cardNumber)-i)%2 == 0 {
			digit *= 2
			// Si el resultado > 9, sumar sus digitos
			if digit > 9 {
				digit = digit/10 + digit%10
			}
		}

		sum += digit
	}

	return sum%10 == 0
}
