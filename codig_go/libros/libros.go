package libros

import (
	"fmt"

	"github.com/norah30/sistema/base"
	"github.com/norah30/sistema/basedatos"
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

// definimos el metodo mostrar libro, recibe una conexion a la base de datos y devuelve un error si hay un problema.
func (l *Libro) MostrarLibro(b *base.Base, args ...interface{}) {
	query := "SELECT CodigoLib, Categoria, Titulo, Precio, Descripcion, EstadoLib, CedulaPro FROM libros WHERE CodigoLib = ?;"
	rows, err := b.DB.Query(query, args...)
	if err != nil {
		err = fmt.Errorf("error al consultar libro: %w", err)
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var libro Libro
		err := rows.Scan(&libro.CodigoLib, &libro.Categoria, &libro.Titulo, &libro.Precio, &libro.Descripcion, &libro.EstadoLib, &libro.CedulaPro)
		if err != nil {
			err = fmt.Errorf("error al mostrar: %w", err)
			fmt.Println(err)
		}
		fmt.Println(libro)
	}
	defer basedatos.CerrarConexion(b.DB)
}

// definimos el metodo editar libro, recibe una conexion a la base de datos y devuelve un error si hay un problema con la actualizacion de la base.
func (l *Libro) EditarLibro(b *base.Base, nuevo Libro) error {
	query := "UPDATE libros SET CodigoLib =?, Categoria=?, Titulo=?, Precio=?, Descripcion=?, EstadoLib=?, CedulaPro =? WHERE CodigoLib = ?;"
	_, err := b.DB.Exec(query, nuevo.CodigoLib, nuevo.Categoria, nuevo.Titulo, nuevo.Precio, nuevo.Descripcion, nuevo.EstadoLib, nuevo.CedulaPro, l.CodigoLib)
	if err != nil {
		return fmt.Errorf("error al editar: %w", err)
	}
	return nil
}

// definimos el metodo eliminar libro, recibe una conexion a la base de datos y devuelve un error si hay un problema con la eliminacion del dato.
func (l *Libro) EliminarLibro(b *base.Base) error {
	query := `DELETE FROM libros WHERE CodigoLib = ?`
	_, err := b.DB.Exec(query, l.CodigoLib)
	if err != nil {
		err = fmt.Errorf("error al eliminar libro: %w", err)
		fmt.Println(err)
	}
	return nil
}

// definimos el metodo insertar libro, recibe una conexion a la base de datos y devuelve un error si hay un problema al insertar datos en la base.
func (l *Libro) InsertarLibro(b *base.Base) error {
	query := `INSERT INTO libros (CodigoLib, Categoria, Titulo, Precio, Descripcion, EstadoLib, CedulaPro) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := b.DB.Exec(query, l.CodigoLib, l.Categoria, l.Titulo, l.Precio, l.Descripcion, l.EstadoLib, l.CedulaPro)

	if err != nil {
		return fmt.Errorf("error al insertar libro: %w", err)
	}
	return nil
}
