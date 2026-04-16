package main

import "fmt"

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}
func AgregarProducto(productos []Producto, p Producto) []Producto {
	return append(productos, p)
}
func calcularTotal(productos []Producto) float64 {
	total := 0.0
	for _, p := range productos {
		total += p.Precio * float64(p.Stock)
	}
	return total
}
func buscarProducto(productos []Producto, nombre string) (Producto, bool) {
	for _, p := range productos {
		if p.Nombre == nombre {
			return p, true
		}
	}
	return Producto{}, false
}



func main() {
	productos := []Producto{
		{Nombre: "Producto1", Precio: 100, Stock: 10},
		{Nombre: "Producto2", Precio: 200, Stock: 20},	
		{Nombre: "Producto3", Precio: 300, Stock: 30},
	}

	productos = AgregarProducto(productos, Producto{Nombre: "Producto4", Precio: 400, Stock: 40})
	productoNuevo := Producto{Nombre: "Producto5", Precio: 500, Stock: 50}
	productos = AgregarProducto(productos, productoNuevo)

 fmt.Printf ("El total de inventario es: %v\n", calcularTotal(productos))

	producto, encontrado := buscarProducto(productos, "Producto1")
	if encontrado {
		fmt.Printf("Producto encontrado:%s\n", producto.Nombre)
	} else {
		fmt.Println("producto no encontrado")
	}
	if productoNuevo, encontrado := buscarProducto(productos, "Producto5"); encontrado {
		fmt.Printf("Producto encontrado: %s\n", productoNuevo.Nombre)
	} else {
		fmt.Println("Producto no encontrado")
	}
}	