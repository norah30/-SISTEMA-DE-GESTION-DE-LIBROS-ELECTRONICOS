package base

import (
	"database/sql"
	"fmt"
)

// definimos una estructura Base que contiene la conexion con la base de datos sql
type Base struct {
	DB *sql.DB
}

// NewBase crea una una nueva instancia de Base
func NewBase(db *sql.DB) *Base {
	return &Base{DB: db}
}

// MostrarCliente ejecuta la consulta SQL para mostrar los datos de los clientes.
func (b *Base) MostrarCliente(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := b.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al mostrar cliente: %w", err)
	}
	return rows, nil
}

// MostrarProveedor ejecuta la consulta SQL para mostrar los datos de los proveedores.
func (b *Base) MostrarProveedor(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := b.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al mostrar proveedor: %w", err)
	}
	return rows, nil
}

// MostrarLibro ejecuta la consulta SQL para mostrar los datos de los libros.
func (b *Base) MostrarLibro(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := b.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al mostrar cliente: %w", err)
	}
	return rows, nil
}

// EditarCliente ejecuta la consulta SQL para editar los datos de los clientes.
func (b *Base) EditarCliente(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al eliminar: %w", err)
	}
	return nil
}

// EditarProveedor ejecuta la consulta SQL para editar los datos de los proveedores
func (b *Base) EditarProveedor(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al eliminar: %w", err)
	}
	return nil
}

// EditarLibro ejecuta la consulta SQL para editar los datos de los libros.
func (b *Base) EditarLibro(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al eliminar: %w", err)
	}
	return nil
}

// EliminarCliente ejecuta la consulta SQL para eliminar los datos de los clientes.
func (b *Base) EliminarCliente(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al eliminar: %w", err)
	}
	return nil
}

// EliminarProveedor ejecuta la consulta SQL para eliminar los datos de los proveedores.
func (b *Base) EliminarProveedor(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al eliminar: %w", err)
	}
	return nil
}

// EliminarLibro ejecuta la consulta SQL para eliminar los datos de los libros.
func (b *Base) EliminarLibro(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al eliminar: %w", err)
	}
	return nil
}

// InsertarCliente ejecuta la consulta SQL para insertar los datos de los clientes.
func (b *Base) InsertarCliente(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al insertar: %w", err)
	}
	return nil
}

// InsertarProveedor ejecuta la consulta SQL para insertar los datos de los proveedores.
func (b *Base) InsertarProveedor(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al insertar: %w", err)
	}
	return nil
}

// InsertarLibro ejecuta la consulta SQL para insertar los datos de los libros.
func (b *Base) InsertarLibro(query string, args ...interface{}) error {
	_, err := b.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error al insertar: %w", err)
	}
	return nil
}
