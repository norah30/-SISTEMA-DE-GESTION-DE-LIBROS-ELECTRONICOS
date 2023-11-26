package libros

import (
	"database/sql"
	"fmt"
)

// definimos una estructura libro y sus campos
type Libro struct {
	CodigoLib   int
	Categoria   string
	Titulo      string
	Precio      int
	Descripcion string
	EstadoLib   string
	CedulaPro   string
}

// definimos el metodo mostrar libro
func (l *Libro) MostrarLibro() {
	fmt.Println("Código de libro: ", l.CodigoLib)
	fmt.Println("Categoría: ", l.Categoria)
	fmt.Println("Título: ", l.Titulo)
	fmt.Println("Precio: ", l.Precio)
	fmt.Println("Descripción: ", l.Descripcion)
	fmt.Println("Estado: ", l.EstadoLib)
	fmt.Println("Cédula del proveedor: ", l.CedulaPro)
}

// definimos el metodo editar libro
func (l *Libro) EditarLibro(nuevoCodigoLib int, nuevaCategoria string, nuevoTitulo string, nuevoPrecio int, nuevaDescripcion string, nuevoEstado string, nuevaCedulaPro string) {
	l.CodigoLib = nuevoCodigoLib
	l.Categoria = nuevaCategoria
	l.Titulo = nuevoTitulo
	l.Precio = nuevoPrecio
	l.Descripcion = nuevaDescripcion
	l.EstadoLib = nuevoEstado
	l.CedulaPro = nuevaCedulaPro
}

// definimos el metodo eliminar libro
func (l *Libro) EliminarLibro() {
	l.EstadoLib = "Eliminado"
}

// definimos el metodo insertar libro
func (l *Libro) InsertarLibro(db *sql.DB) error {
	query := `INSERT INTO libros (codigo_lib, categoria, titulo, precio, descripcion, estado_lib, cedula_pro) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, l.CodigoLib, l.Categoria, l.Titulo, l.Precio, l.Descripcion, l.EstadoLib, l.CedulaPro)
	return err
}
