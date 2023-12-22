package basedatos

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Db es una instancia de sql.DB que representa la conexión a la base de datos.
var Db *sql.DB

// DbConn establece una conexión con la base de datos SQL.
func DbConn() (*sql.DB, error) {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Establecer la conexión con la base de datos MySQL
	conexion, err := sql.Open("mysql", os.Getenv("BD_USER")+":"+os.Getenv("BD_PASSWORD")+"@tcp("+os.Getenv("BD_SERVER")+":"+os.Getenv("BD_PORT")+")/"+os.Getenv("BD_NAME"))
	if err != nil {
		return nil, err
	}

	// Asignar la conexión a la variable global Db
	Db = conexion

	return Db, nil
}

// CerrarConexion cierra la conexión a la base de datos.
func CerrarConexion(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
