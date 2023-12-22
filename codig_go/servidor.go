package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/norah30/sistema/base"
	"github.com/norah30/sistema/basedatos"
	"github.com/norah30/sistema/clientes"
	"github.com/norah30/sistema/libros"
	"github.com/norah30/sistema/proveedores"
)

var tmpl *template.Template

// definimos una estructura Inicio que contiene un título y un mensaje para la página.
type Inicio struct {
	Titulo  string
	Mensaje string
}

var err error

// handleInicio maneja la petición GET a la ruta raíz ("/") y muestra la página de inicio con el título y el mensaje.
func handleInicio(w http.ResponseWriter, r *http.Request) {
	
	template, err := template.ParseFiles("sistem/inicio.html")

	if err != nil {
		panic(err)
	}else{
		template.Execute(w, nil)
	}
	// datos := Inicio{
	// 	Titulo:  "GESTION DE LIBROS ELECTRONICOS",
	// 	Mensaje: "BIENVENIDOS AL SISTEMA DE GESTION DE LIBROS ELECTRONICOS",
	// }
	// err := tmpl.ExecuteTemplate(w, "inicio.html", datos)
	// if err != nil {
	// 	log.Printf("Error al ejecutar plantilla: %v", err)
	// 	http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
	// }
}

// Pag404 maneja la petición GET a cualquier ruta que no esté definida en el enrutador.
func Pag404(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("404.html"))
	tmpl.Execute(w, nil)
}

/*
	handleClientes maneja las peticiones GET, POST, PUT y DELETE a la ruta "/clientes"

y dependiendo del método de la petición, inserta, edita, elimina o muestra los datos de los clientes.
*/
func handleClientes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var cliente clientes.Cliente
	tmpl, err := template.ParseGlob("clientes.html")
	if err != nil {
		log.Printf("Error al cargar plantilla: %v", err)
		http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
	}
	// Si el método de la petición es POST inserta los datos del cliente en la base de datos.
	if r.Method == "POST" {
		err = cliente.InsertarCliente(base.NewBase(db))
		if err != nil {
			log.Printf("Error al insertar clientes : %v ", err)
			http.Error(w, "Error al preparar la consulta de clientes ", http.StatusInternalServerError)
			return
		}

		// Si no hay errores, envía una respuesta al cliente indicando que los datos se insertaron con éxito.
		fmt.Fprint(w, "Datos insertados con éxito!")
		// Si el método de la petición es PUT edita los datos del cliente en la base de datos
	} else if r.Method == "PUT" {
		err = cliente.EditarCliente(base.NewBase(db), cliente)
		if err != nil {
			log.Printf("Error al editar clientes : %v ", err)
			http.Error(w, "Error al preparar la consulta de clientes ", http.StatusInternalServerError)
			return
		}
		// Si no hay errores envía una respuesta al cliente indicando que los datos se editaron con éxito.
		fmt.Fprint(w, "Datos editados con éxito!")
		// Si el método de la petición es DELETE elimina los datos del cliente en la base de datos
	} else if r.Method == "DELETE" {
		err = cliente.EliminarCliente(base.NewBase(db))
		if err != nil {
			log.Printf("Error al eliminar clientes : %v ", err)
			http.Error(w, "Error al preparar la consulta de clientes ", http.StatusInternalServerError)
			return
		}

		// Si no hay errores, envía una respuesta al cliente indicando que los datos se eliminaron con éxito.
		fmt.Fprint(w, "Datos eliminados con éxito!")
	} else {
		// Si el método de la petición es GET muestra los datos del cliente de la base de datos.
		cliente.MostrarCliente(base.NewBase(db), cliente.CedulaCli)
		if err != nil {
			log.Printf("Error al mostrar clientes: %v", err)
			http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
			return
		}
		// Si no hay errores, ejecuta la plantilla con los datos del cliente y la envía al cliente.
		err = tmpl.Execute(w, cliente)
		if err != nil {
			log.Printf("Error al ejecutar clientes: %v", err)
			http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
			return
		}
	}
}

/*
	handleProveedores maneja las peticiones GET, POST, PUT y DELETE a la ruta "/proveedores"

y dependiendo del método de la petición, inserta, edita, elimina o muestra los datos de los proveedores.
*/
func handleProveedores(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var proveedor proveedores.Proveedor
	tmpl, err := template.ParseGlob("proveedores.html")
	if err != nil {
		log.Printf("Error al cargar plantilla: %v", err)
		http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
	}
	// Si el método de la petición es POST inserta los datos del proveedor en la base de datos.
	if r.Method == "POST" {
		err = proveedor.InsertarProveedor(base.NewBase(db))
		if err != nil {
			log.Printf("Error al insertar proveedor : %v ", err)
			http.Error(w, "Error al preparar la consulta de proveedor ", http.StatusInternalServerError)
			return
		}

		// Si no hay errores, envía una respuesta al cliente indicando que los datos se insertaron con éxito.
		fmt.Fprint(w, "Datos insertados con éxito!")
		// Si el método de la petición es PUT, edita los datos del proveedor en la base de datos.
	} else if r.Method == "PUT" {
		err = proveedor.EditarProveedor(base.NewBase(db), proveedor)
		if err != nil {
			log.Printf("Error al editar proveedores : %v ", err)
			http.Error(w, "Error al preparar la consulta de proveedores ", http.StatusInternalServerError)
			return
		}

		// Si no hay errores, envía una respuesta al cliente indicando que los datos se editaron con éxito.
		fmt.Fprint(w, "Datos editados con éxito!")
		// Si el método de la petición es DELETE eliminara los datos del proveedor en la base de datos.
	} else if r.Method == "DELETE" {
		err = proveedor.EliminarProveedor(base.NewBase(db))
		if err != nil {
			log.Printf("Error al eliminar proveedores : %v ", err)
			http.Error(w, "Error al preparar la consulta de proveedores ", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Datos eliminados con éxito!")
		// Si el método de la petición es GET mostrara los datos del proveedor de la base de datos.
	} else {
		proveedor.MostrarProveedor(base.NewBase(db), proveedor.CedulaPro)
		if err != nil {
			log.Printf("Error al mostrar proveedores: %v", err)
			http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
			return
		}
		// Si no hay errores, intenta ejecutar la plantilla con los datos del proveedor y envia.
		err = tmpl.Execute(w, proveedor)
		if err != nil {
			log.Printf("Error al ejecutar Proveedores: %v", err)
			http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
			return
		}
	}
}

/* handleLibros maneja las peticiones GET, POST, PUT y DELETE a la ruta "/libros"
y dependiendo del método de la petición, inserta, edita, elimina o muestra los datos de los libros  */

func handleLibros(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var libro libros.Libro
	tmpl, err := template.ParseGlob("libros.html")
	if err != nil {
		log.Printf("Error al cargar plantilla: %v", err)
		http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
	}

	// Si el método de la petición es POST inserta los datos del libro en la base de datos.
	if r.Method == "POST" {
		err = libro.InsertarLibro(base.NewBase(db))
		if err != nil {
			log.Printf("Error al insertar libros : %v ", err)
			http.Error(w, "Error al preparar la consulta de libros ", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Datos insertados con éxito!")
		// Si el método de la petición es PUT edita los datos del libro en la base de datos.
	} else if r.Method == "PUT" {
		err = libro.EditarLibro(base.NewBase(db), libro)
		if err != nil {
			log.Printf("Error al editar libros : %v ", err)
			http.Error(w, "Error al preparar la consulta de libros ", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Datos editados con éxito!")
		// Si el método de la petición es DELETE elimina los datos del libro en la base de datos.
	} else if r.Method == "DELETE" {
		err = libro.EliminarLibro(base.NewBase(db))
		if err != nil {
			log.Printf("Error al eliminar libros : %v ", err)
			http.Error(w, "Error al preparar la consulta de libros ", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "Datos eliminados con éxito!")
		// Si el método de la petición es GET mostrara los datos del libro de la base de datos.
	} else {
		libro.MostrarLibro(base.NewBase(db), libro.CodigoLib)
		if err != nil {
			log.Printf("Error al mostrar libros: %v", err)
			http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, libro)
		if err != nil {
			log.Printf("Error al ejecutar libros: %v", err)
			http.Error(w, "Error Interno del servidor", http.StatusInternalServerError)
			return
		}
	}
}



// definimos la funcion main
func main() {
	//registra la conexión del servidor
	log.Println("Server conectado: http://localhost:8080")

	// establece una conexión con la base de datos
	db, err := basedatos.DbConn() // la conexión a la base de datos
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	//verificar si la base de datos se abrio correctamente
	if err = db.Ping(); err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	fmt.Println("La base de datos se abrió correctamente")


	



	//crea un nuevo enrutador
	r := mux.NewRouter()

	// Carga las plantillas html
	tmpl, err = template.ParseGlob("sistem/*.html")
	if err != nil {
		log.Fatalf("Error al cargar plantillas: %v", err)
	}

	// Registra las rutas para el enrutador
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleInicio(w, r)
	}).Methods("GET")

	r.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		handleClientes(w, r, db)
	}).Methods("GET", "POST", "PUT", "DELETE")

	r.HandleFunc("/proveedores", func(w http.ResponseWriter, r *http.Request) {
		handleProveedores(w, r, db)
	}).Methods("GET", "POST", "PUT", "DELETE")

	r.HandleFunc("/libros", func(w http.ResponseWriter, r *http.Request) {
		handleLibros(w, r, db)
	}).Methods("GET", "POST", "PUT", "DELETE")

	//archivos estaticos hacia mux
	s := http.StripPrefix("/estaticos/", http.FileServer(http.Dir("./estaticos/")))
	r.PathPrefix("/estaticos/").Handler(s)

	// Maneja las peticiones no encontradas (error 404)
	r.NotFoundHandler = r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Pag404(w, r)
	}).GetHandler()

	log.Println("Server conectado: http://localhost:8080")
	log.Println("Servidor HTTP iniciado correctamente")

	// Inicia el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", r)

	// Cierra la conexión a la base de datos cuando el programa termina
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error al cerrar la conexión a la base de datos: %v", err)
		}
	}()

}
