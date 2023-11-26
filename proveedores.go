package proveedores

import (
	"database/sql"
	"fmt"
)

// definimos una estructura proveedor y sus campos
type Proveedor struct {
	Nombres   string
	Apellidos string
	Usuario   string
	Clave     string
	CedulaPro string
	Correo    string
	Telefono  int
	EstadoPro string
}

// definimos el metodo mostrar proveedor
func (p *Proveedor) MostrarProveedor() {
	fmt.Println("Nombres: ", p.Nombres)
	fmt.Println("Apellidos: ", p.Apellidos)
	fmt.Println("Usuario: ", p.Usuario)
	fmt.Println("Clave: ", p.Clave)
	fmt.Println("Cédula: ", p.CedulaPro)
	fmt.Println("Correo: ", p.Correo)
	fmt.Println("Teléfono: ", p.Telefono)
	fmt.Println("Estado: ", p.EstadoPro)
}

// definimos el metodo editar proveedor
func (p *Proveedor) EditarProveedor(nuevoNombre string, nuevoApellido string, nuevoUsuario string, nuevaClave string, nuevaCedula string, nuevoCorreo string, nuevoTelefono int, nuevoEstado string) {
	p.Nombres = nuevoNombre
	p.Apellidos = nuevoApellido
	p.Usuario = nuevoUsuario
	p.Clave = nuevaClave
	p.CedulaPro = nuevaCedula
	p.Correo = nuevoCorreo
	p.Telefono = nuevoTelefono
	p.EstadoPro = nuevoEstado
}

// definimos el metodo eliminar proveedor
func (p *Proveedor) EliminarProveedor() {
	p.EstadoPro = "Eliminado"
}

// definimos el metodo insertar proveedor
func (p *Proveedor) InsertarProveedor(db *sql.DB) error {
	query := `INSERT INTO proveedor (nombres, apellidos, usuario, clave, cedula_pro, correo, telefono, estado_pro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, p.Nombres, p.Apellidos, p.Usuario, p.Clave, p.CedulaPro, p.Correo, p.Telefono, p.EstadoPro)
	return err
}
