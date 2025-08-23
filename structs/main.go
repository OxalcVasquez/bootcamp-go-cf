package main

import "fmt"

// Definimos una estructura (equivalentea una clase)
type Coche struct {
	Marca string
	Color string
	Motor //Podemos acceder directo a la propiedad Coche.Time
}

type Motor struct {
	Time int
}

// Esta asociado a la structura coche, un metodo para una estructura
func (c Coche) Arrancar() {
	fmt.Printf("El coche %s se esta arrancando \n", c.Marca)
}

// Sin el asterisco no cambia
func (c *Coche) ModificarMarca() {
	c.Marca = "Honda"
}

func main() {
	//creamos una instancia
	miCoche := Coche{Marca: "Toyota", Color: "Rojo"}
	miCoche.Arrancar()
	miCoche.ModificarMarca()

	fmt.Println(miCoche.Marca)
}
