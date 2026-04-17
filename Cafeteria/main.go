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
//funciones para clientes  
func ListarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("No hay clientes registrados")
		return
	}

	for _, c := range clientes {
		fmt.Println("ID:", c.ID, "| Nombre:", c.Nombre, "| Carrera:", c.Carrera, "| Saldo:", c.Saldo)
	}
}
func AgregarCliente(clientes []Cliente, c Cliente) []Cliente {
	clientes = append(clientes, c)
	return clientes
}

func BuscarCliente(clientes []Cliente, id int) *Cliente {
	for i := range clientes {
		if clientes[i].ID == id {
			return &clientes[i]
		}
	}
	return nil
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	nuevo := []Cliente{}

	for _, c := range clientes {
		if c.ID != id {
			nuevo = append(nuevo, c)
		}
	}

	return nuevo
}

//aplicacin de funciones para produtos 
func agregarProducto(productos []Producto, p Producto) []Producto {
	return append(productos, p)
}



func main() {
	fmt.Println("==== Cafeteria LUIS ====")

	clientes := []Cliente{}
	productos := []Producto{}
	pedidos := []Pedido{}

	// 🔹 DATOS INICIALES
	clientes = append(clientes, Cliente{1, "Valentina", "Medicina", 6.0})
	clientes = append(clientes, Cliente{2, "Camila", "Arquitectura", 5.10})
	clientes = append(clientes, Cliente{3, "Sofia", "Odontologia", 9.0})

	productos = append(productos, Producto{1, "Gaseosa", 1.55, 6, "bebida"})
	productos = append(productos, Producto{2, "Cafe", 1.0, 13, "bebida"})
	productos = append(productos, Producto{3, "Pastel", 2.0, 13, "snack"})
	productos = append(productos, Producto{4, "Ensalada", 3.6, 9, "almuerzo"})
	productos = append(productos, Producto{5, "Sopa", 2.5, 8, "almuerzo"})


	fmt.Println("\nClientes iniciales:")
	ListarClientes(clientes)

	//aplicamos las funciones
	// Agregar un nuevo cliente
	clientes = AgregarCliente(clientes, Cliente{4, "Mario", "Software", 10.0})
	fmt.Println("\nDespués de agregar cliente:")
	ListarClientes(clientes)

	//buscamos un cliente
	c := BuscarCliente(clientes, 3)
	if c != nil {
		fmt.Println("\nCliente encontrado:", c.Nombre)
	} else {
		fmt.Println("\nCliente no encontrado")
	}

	//elimiamos un cliente
	clientes = EliminarCliente(clientes, 1)
	fmt.Println("\nDespués de eliminar cliente:")
	ListarClientes(clientes)

	
	fmt.Println("\nCantidad de productos:", len(productos))
	fmt.Println("Cantidad de pedidos:", len(pedidos))
}
