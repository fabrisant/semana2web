package main

import "fmt"

type Cliente struct {
	ID       int
	Nombre   string
	Carrera  string
	Saldo    float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func ListarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("No hay clientes registrados")
		return
	}

	for _, c := range clientes {
		fmt.Println("-------------")
		fmt.Println("ID:", c.ID)
		fmt.Println("Nombre:", c.Nombre)
		fmt.Println("Carrera:", c.Carrera)
		fmt.Println("Saldo:", c.Saldo)
	}
}

func main() {
	fmt.Println("==== Cafeteria LUIS ====")

	// listas para clientes, productos y pedidos
	clientes := []Cliente{}
	productos := []Producto{}
	pedidos := []Pedido{}

	clientes = append(clientes, Cliente{1, "Valentina", "Medicina", 6.0})
	clientes = append(clientes, Cliente{2, "Camila", "Arquitectura", 5.10})
	clientes = append(clientes, Cliente{3, "Sofia", "Odontologia", 9.0})


	productos = append(productos, Producto{1, "Gaseosa", 1.55, 6, "bebida"})
	productos = append(productos, Producto{2, "Cafe", 1.0, 13, "bebida"})
	productos = append(productos, Producto{3, "Pastel", 2.0, 13, "snack"})
	productos = append(productos, Producto{4, "Ensalada", 3.6, 9, "almuerzo"})
	productos = append(productos, Producto{5, "Sopa", 2.5, 8, "almuerzo"})

	// lista clientes registrados
	ListarClientes(clientes)
	fmt.Println("Cantidad de productos:", len(productos))
	fmt.Println("Cantidad de pedidos:", len(pedidos))
}

func AgregarCliente(clientes []Cliente, c Cliente) []Cliente {
	clientes = append(clientes, c)
	return clientes
}
func buscarCliente(clientes []Cliente, id int) (Cliente, bool) {
	for _, c := range clientes {
		if c.ID == id {
			return c, true
		}
	}
	return Cliente{}, false
}
func listarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("No hay clientes registrados")
		return		
	}	
	for _, c := range clientes {
		fmt.Println("-------------")
		fmt.Println("ID:", c.ID)					
		fmt.Println("Nombre:", c.Nombre)
		fmt.Println("Carrera:", c.Carrera)
		fmt.Println("Saldo:", c.Saldo)
	}
}	