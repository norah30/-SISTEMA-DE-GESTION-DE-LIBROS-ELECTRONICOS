package clientes

//importamos
import (
	"fmt"

	"github.com/norah30/sistema/base"
	"github.com/norah30/sistema/basedatos"
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

// definimos el metodo mostrar clientes, recibe una conexion a la base de datos y devuelve un error si hay un problema.
func (c *Cliente) MostrarCliente(b *base.Base, args ...interface{}) {
	query := "SELECT CedulaCli, CodigoLib, Nombres, Apellidos, Correo, Direccion, NumeroTar, Targeta, CodigoSeg, Celular, Eliminado FROM clientes WHERE CedulaCli= ?;"
	rows, err := b.DB.Query(query, args...)
	if err != nil {
		err = fmt.Errorf("error al consultar cliente: %w", err)
		fmt.Println(err)

	}
	defer rows.Close()

	for rows.Next() {
		var cliente Cliente
		err := rows.Scan(&cliente.CedulaCli, &cliente.CodigoLib, &cliente.Nombres, &cliente.Apellidos, &cliente.Correo, &cliente.Direccion, &cliente.NumeroTar, &cliente.Tarjeta, &cliente.CodigoSeg, &cliente.Celular, &cliente.Eliminado)
		if err != nil {
			err = fmt.Errorf("error al mostrar: %w", err)
			fmt.Println(err)

		}
		fmt.Println(cliente)
	}
	defer basedatos.CerrarConexion(b.DB)
}

// definimos el metodo editar clientes, recibe una conexion a la base de datos y devuelve un error si hay un problema con la actualizacion de la base.
func (c *Cliente) EditarCliente(b *base.Base, nuevo Cliente) error {
	query := "UPDATE clientes SET CedulaCli= ?, CodigoLib= ?, Nombres= ?, Apellidos= ?, Correo= ?, Direccion= ?, NumeroTar= ?, Targeta= ?, CodigoSeg= ?, Celular= ?, Eliminado= ? WHERE CedulaCli= ?; "
	_, err := b.DB.Exec(query, nuevo.CedulaCli, nuevo.CodigoLib, nuevo.Nombres, nuevo.Apellidos, nuevo.Correo, nuevo.Direccion, nuevo.NumeroTar, nuevo.Tarjeta, nuevo.CodigoSeg, nuevo.Celular, nuevo.Eliminado, c.CedulaCli)
	if err != nil {
		return fmt.Errorf("error al editar cliente: %w", err)
	}
	return nil
}

// definimos el metodo eliminar cliente, recibe una conexion a la base de datos y devuelve un error si hay un problema con la eliminacion del dato.
func (c *Cliente) EliminarCliente(b *base.Base) error {
	query := `DELETE FROM clientes WHERE CedulaCli = ?`
	_, err := b.DB.Exec(query, c.CedulaCli)
	if err != nil {
		err = fmt.Errorf("error al eliminar cliente: %w", err)
		fmt.Println(err)
	}
	return nil
}

// definimos el metodo insertar cliente, recibe una conexion a la base de datos y devuelve un error si hay un problema con insertar datos en  la base.
func (c *Cliente) InsertarCliente(b *base.Base) error {
	query := `INSERT INTO clientes (CedulaCli, CodigoLib, Nombres, Apellidos, Correo, Direccion, NumeroTar, Tarjeta, CodigoSeg, Celular, Eliminado) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`
	_, err := b.DB.Exec(query, c.CedulaCli, c.CodigoLib, c.Nombres, c.Apellidos, c.Correo, c.Direccion, c.NumeroTar, c.Tarjeta, c.CodigoSeg, c.Celular)
	if err != nil {
		return fmt.Errorf("error al insertar cliente: %w", err)
	}
	return nil
}
