# Sistema de Gestión de Libros Electronicos
Este proyecto se centra en la creación de un sistema de gestión de libros electrónicos utilizando el lenguaje de programación Go. 
Su principal objetivo es promover la venta de libros digitales, este repositorio contiene el código fuente del sistema.
 
Realizado por: Nora Calle, José Córdova, Omar Llano.

## FindBook

Pensamos en ofrecer grandes oportunidades para nuestros vendedores (autores).
Ofreciéndoles una plataforma que les permita fácilmente publicar y gestionar sus libros.

![](https://github.com/norah30/-SISTEMA-DE-GESTION-DE-LIBROS-ELECTRONICOS/blob/9c47e1ad6529517457d6b450e9933859caf3acf4/imagenes/WhatsApp%20Image%202023-12-21%20at%2023.20.38.jpeg)

## Descripcion del proyecto 
Este proyecto es un sistema realizado en Go que permite gestionar una base de datos de libros electrónicos. Se puede utilizar para realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) en los datos de los libros, clientes y proveedores.

## Tecnologias Utilizadas 
Base de datos MYSQL, 
Lenguaje de programación Go, 
GitHub.  
Plantillas html
Gorilla/mux
Go-sql-driver/mysql

## Diagrama de Funcionalidad del Software.
![](https://github.com/norah30/-SISTEMA-DE-GESTION-DE-LIBROS-ELECTRONICOS/blob/fa5a073a844dc210e7e111113f791f4699324004/sistemas%20de%20venta%20de%20libros%20digitales.jpeg)
## Diagrama de Arquitectura del Software.
![](https://github.com/norah30/-SISTEMA-DE-GESTION-DE-LIBROS-ELECTRONICOS/blob/0fc2f0c244d579dc03cab7970aef38fca1897001/SIS%20(1).jpg)
## Diagrama de Clases 
![](https://github.com/norah30/-SISTEMA-DE-GESTION-DE-LIBROS-ELECTRONICOS/blob/45bcd85b046d5bbcfdbceb5c99a1a2840c25c5ac/Proyecto%20GOLAND%202.jpeg)
 ## Avance del sistema de gestion de libros electronicos 
Definimos estructuras y métodos para manejar entidades como clientes, libros y proveedores. 
Tambien utilizamos una base de datos SQL para almacenar y gestionar los datos.
El código HTML proporciona las páginas de inicio y 404 para la interfaz de usuario. 
Los paquetes base y bdsql se encargan de la conexión a la base de datos y las operaciones CRUD.
El paquete servidor maneja las solicitudes HTTP a diferentes rutas.
![](https://github.com/norah30/-SISTEMA-DE-GESTION-DE-LIBROS-ELECTRONICOS/blob/6c8716addc9e6aa5ef7da74a103272ea724016b0/Inicio_Sistema.png)

## Funcionalidades del Proyecto
Gestión Libros: Permite a los usuarios crear, leer, actualizar y eliminar registros de libros en la base de datos. Esto significa que los usuarios pueden agregar nuevos libros, ver detalles de los libros existentes, modificar la información de los libros existentes y eliminar libros que ya no están en stock.

Gestión Clientes: Similar a la gestión de libros, esta característica permite a los usuarios gestionar los datos de los clientes, incluye la posibilidad de agregar nuevos clientes, ver detalles de los clientes existentes, modificar la información de los clientes existentes y eliminar clientes que ya no son clientes.

Gestión Proveedores: Permite a los proveedores gestionar los datos. 

Interfaz Web: La aplicación proporciona una interfaz web para interactuar con la base de datos. 

Enrutamiento: Utiliza el paquete gorilla/mux para manejar las solicitudes HTTP y dirigirlas a las funciones correspondientes según la ruta y el método HTTP. Esto permite a la aplicación manejar múltiples rutas y métodos HTTP de manera eficiente.

Conexión a MYSQL:  El sistema se conecta a una base datos MySQL utilizando el controlador go-sql-driver/mysql. Esto significa que la aplicación puede almacenar y recuperar datos de la base de datos de manera eficiente.







