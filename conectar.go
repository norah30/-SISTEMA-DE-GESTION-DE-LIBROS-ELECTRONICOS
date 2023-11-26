// importamos paquetes estandar, externos y definidos
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"clientes"
	"libros"
	"proveedores"

	_ "github.com/denisenkom/go-mssqldb"
)

// definimos la funcion dbConn para establecer la conexion
func dbConn() (conexion *sql.DB) {

	Usuario := "sa"
	contrasenia := "9224"
	Base := "Gestion_libros"
	db, err := sql.Open("sqlserver", "sqlserver://"+Usuario+":"+contrasenia+"@localhost:1433?database="+Base)

	//manejo de error que pueda ocurrir al abrir la conexión a la base de datos
	if err != nil {
		panic(err.Error())
	}
	return db
}

// definimos la funcion main
func main() {
	log.Println("Server conectado: http://localhost:8080") //registro de la conexión del servidor

	db := dbConn() // la conexión a la base de datos
	fmt.Fprintf(os.Stdout, "abrio base de datos")

	//consulta de las bases de datos e iteracion de redultados de clientes
	var cliente clientes.Cliente
	selDB, err := db.Query("SELECT CedulaCli,CodigoLib,Nombres,Apellidos,Correo,Direccion,NumeroTar,Tarjeta,CodigoSeg,Celular,Eliminado FROM clientes")

	r := mux.NewRouter() //creamos un enrutador
	if err != nil {
		log.Fatal(err)

	}
	//iteramos con el bucle for
	for selDB.Next() {
		err = selDB.Scan(&cliente.CedulaCli, &cliente.CodigoLib, &cliente.Nombres, &cliente.Apellidos, &cliente.Correo, &cliente.Direccion, &cliente.NumeroTar, &cliente.Tarjeta, &cliente.CodigoSeg, &cliente.Celular, &cliente.Eliminado)
		if err != nil {
			log.Fatal(err)
		}
	}

	r.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			// Insertar los datos en la base de datos
			err = cliente.InsertarCliente(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Enviar una respuesta al cliente
			fmt.Fprint(w, "Datos insertados con éxito!")
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}).Methods("POST")

	//consulta de las bases de datos e iteracion de redultados de proveedores
	var proveedor proveedores.Proveedor
	selDB, err = db.Query("SELECT Nombres, Apellidos, Usuario, Clave,CedulaPro, Correo, Telefono, EstadoPro FROM proveedores")
	if err != nil {
		log.Fatal(err)
	}

	//iteramos con el bucle for
	for selDB.Next() {
		err = selDB.Scan(&proveedor.Nombres, &proveedor.Apellidos, &proveedor.Usuario, &proveedor.Clave, &proveedor.CedulaPro, &proveedor.Correo, &proveedor.Telefono, &proveedor.EstadoPro)
		if err != nil {
			log.Fatal(err)
		}

		//proveedores
		r.HandleFunc("/proveedores", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {

				// Insertar los datos en la base de datos
				err = proveedor.InsertarProveedor(db)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Enviar una respuesta al proveedor
				fmt.Fprint(w, "Datos insertados con éxito!")
			} else {
				http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			}
		}).Methods("POST")

		//consulta de las bases de datos e iteracion de redultados de libros
		var libro libros.Libro
		selDB, err = db.Query("SELECT CodigoLib, Categoria, Titulo, Precio, Descripcion, Estadolib, CedulaPro FROM libros")
		if err != nil {
			log.Fatal(err)
		}
		//iteramos con el bucle for
		for selDB.Next() {

			err = selDB.Scan(&libro.CodigoLib, &libro.Categoria, &libro.Titulo, &libro.Precio, &libro.Descripcion, &libro.EstadoLib, &libro.CedulaPro)
			if err != nil {
				log.Fatal(err)
			}
		}

		r.HandleFunc("/libros", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {

				// Insertar los datos en la base de datos
				err = libro.InsertarLibro(db)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Enviar una respuesta
				fmt.Fprint(w, "Datos insertados con éxito!")
			} else {
				http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			}
		}).Methods("POST")

		log.Println("Server conectado: http://localhost:8080")
		log.Println("Servidor HTTP iniciado correctamente")
		http.ListenAndServe(":8080", r)
		defer db.Close()
	}
}
