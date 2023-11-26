package clientes

//importamos
import (
	"database/sql"
	"fmt"
)

// definimos la estructura cliente y sus campos
type Cliente struct {
	CedulaCli string
	CodigoLib int
	Nombres   string
	Apellidos string
	Correo    string
	Direccion string
	NumeroTar string
	Tarjeta   string
	CodigoSeg int
	Celular   string
	Eliminado bool
}

// definimos el metodo mostrar clientes
func (c *Cliente) MostrarCliente() {
	fmt.Println("Cédula: ", c.CedulaCli)
	fmt.Println("Nombres: ", c.Nombres)
	fmt.Println("Apellidos: ", c.Apellidos)
	fmt.Println("Dirección: ", c.Direccion)
	fmt.Println("Celular: ", c.Celular)
	fmt.Println("Correo: ", c.Correo)
	fmt.Println("Código de libro: ", c.CodigoLib)
	fmt.Println("Tarjeta: ", c.Tarjeta)
	fmt.Println("Número de tarjeta: ", c.NumeroTar)
	fmt.Println("Código de seguridad: ", c.CodigoSeg)
}

// definimos el metodo editar clientes
func (c *Cliente) EditarCliente(nuevaCedulaCli string, nuevaCodigoLib int, nuevoNombres string, nuevoApellidos string, nuevoCorreo string, nuevaDireccion string, nuevoNumeroTar string, nuevaTarjeta string, nuevoCodigoSeg int, nuevoCelular string) {
	c.CedulaCli = nuevaCedulaCli
	c.CodigoLib = nuevaCodigoLib
	c.Nombres = nuevoNombres
	c.Apellidos = nuevoApellidos
	c.Correo = nuevoCorreo
	c.Direccion = nuevaDireccion
	c.NumeroTar = nuevoNumeroTar
	c.Tarjeta = nuevaTarjeta
	c.CodigoSeg = nuevoCodigoSeg
	c.Celular = nuevoCelular
}

// definimos el metodo eliminar cliente
func (c *Cliente) EliminarCliente() {
	c.Eliminado = true
}

// definimos el metodo insertar cliente
func (c *Cliente) InsertarCliente(db *sql.DB) error {
	query := `INSERT INTO clientes (cedula_cli, codigo_lib, nombres, apellidos, correo, direccion, numero_tar, tarjeta, codigo_seg, celular) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, c.CedulaCli, c.CodigoLib, c.Nombres, c.Apellidos, c.Correo, c.Direccion, c.NumeroTar, c.Tarjeta, c.CodigoSeg, c.Celular)
	return err
}
