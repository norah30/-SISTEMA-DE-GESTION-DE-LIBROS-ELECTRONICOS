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
// definimos el metodo mostrar proveedor, recibe una conexion a la base de datos y devuelve un error si hay un problema .
func (p *Proveedor) MostrarProveedor(b *base.Base, args ...interface{}) {
	query := "SELECT Nombres, Apellidos, Usuario, Clave, CedulaPro, Correo, Telefono, EstadoPro FROM proveedores  WHERE CedulaPro= ?;"
	rows, err := b.DB.Query(query, args...)
	if err != nil {
		fmt.Errorf("error al consultar proveedor: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var proveedor Proveedor
		err := rows.Scan(&proveedor.Nombres, &proveedor.Apellidos, &proveedor.Usuario, &proveedor.Clave, &proveedor.CedulaPro, &proveedor.Correo, &proveedor.Telefono, &proveedor.EstadoPro)
		if err != nil {
			fmt.Errorf("error al mostrar: %w", err)
		}
		fmt.Println(proveedor)
	}
	defer basedatos.CerrarConexion(b.DB)
}

// definimos el metodo editar proveedor, recibe una conexion a la base de datos y devuelve un error si hay un problema con la actualizacion de la base.
func (p *Proveedor) EditarProveedor(b *base.Base, nuevo Proveedor) error {
	query := "UPDATE proveedores SET Nombres =? , Apellidos=?, Usuario=?, Clave=?, CedulaPro=?, Correo=?, Telefono=?, EstadoPro=? WHERE CedulaPro= ?;"
	_, err := b.DB.Exec(query, nuevo.Nombres, nuevo.Apellidos, nuevo.Usuario, nuevo.Clave, nuevo.CedulaPro, nuevo.Correo, nuevo.Telefono, nuevo.EstadoPro, p.CedulaPro)
	if err != nil {
		return fmt.Errorf("error al editar: %w", err)
	}
	return nil
}

// definimos el metodo eliminar proveedor, recibe una conexion a la base de datos y devuelve un error si hay un problema con la eliminacion  del dato.
func (p *Proveedor) EliminarProveedor(b *base.Base) error {
	query := `DELETE FROM proveedores WHERE CedulaPro = ?`
	_, err := b.DB.Exec(query, p.CedulaPro)
	if err != nil {
		fmt.Errorf("error al eliminar proveedor: %w", err)
	}
	return nil
}

// definimos el metodo insertar proveedor, recibe una conexion a la base de datos y devuelve un error si hay un problema al insertar datos en la base.
func (p *Proveedor) InsertarProveedor(b *base.Base) error {
	query := `INSERT INTO proveedores (Nombres, Apellidos, Usuario, Clave, CedulaPro, Correo, Telefono, EstadoPro) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := b.DB.Exec(query, p.Nombres, p.Apellidos, p.Usuario, p.Clave, p.CedulaPro, p.Correo, p.Telefono, p.EstadoPro)

	if err != nil {
		return fmt.Errorf("error al insertar proveedor: %w", err)
	}
	return nil
}


