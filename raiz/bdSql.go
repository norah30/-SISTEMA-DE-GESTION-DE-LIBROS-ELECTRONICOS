package basedatos

import (
	"database/sql"
	"fmt"
)

// DbConn establece una conexión con la base de datos SQL.
func DbConn() (*sql.DB, error) {
	Usuario := os.Getenv("DB_USER")
	contrasenia := os.Getenv("DB_PASSWORD")
	Base := os.Getenv("DB_NAME")
	db, err := sql.Open("sqlserver", "sqlserver://"+Usuario+":"+contrasenia+"@localhost:1433?database="+Base)

	// Si ocurre un error al abrir la conexión a la base de datos, se devuelve nil y el error.
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}
	return db, nil
}

// CerrarConexion cierra la conexión con la base de datos SQL.
func CerrarConexion(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Errorf("error al cerrar conexion: %w", err)
	}
}
