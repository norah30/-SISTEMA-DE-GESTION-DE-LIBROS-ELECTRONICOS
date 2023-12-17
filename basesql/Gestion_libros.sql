CREATE TABLE [clientes] (
 [CedulaCli] varchar(15) NOT NULL,
 [CodigoLib] int NOT NULL,
 [Nombres] varchar(20) NOT NULL,
 [Apellidos] varchar(20) NOT NULL,
 [Correo] varchar(20) NOT NULL,
 [Direccion] varchar(40) NOT NULL,
 [NumeroTar] varchar(15) NOT NULL,
 [Tarjeta] varchar(15) NOT NULL,
 [CodigoSeg] int NOT NULL,
 [Celular] varchar(10) NOT NULL,
 [Eliminado] bit NOT NULL,
 PRIMARY KEY ([CedulaCli])
)

CREATE TABLE [proveedores] (
 [Nombres] varchar(20) NOT NULL,
 [Apellidos] varchar(20) NOT NULL,
 [Usuario] varchar(15) NOT NULL,
 [Clave] varchar(10) NOT NULL,
 [CedulaPro] varchar(15) NOT NULL,
 [Correo] varchar(20) NOT NULL,
 [Telefono] int NOT NULL,
 [EstadoPro] varchar(1) NOT NULL,
 PRIMARY KEY ([CedulaPro])
)

CREATE TABLE [libros] (
 [CodigoLib] int NOT NULL,
 [Categoria] varchar(10) NOT NULL,
 [Titulo] varchar(20) NOT NULL,
 [Precio] int NOT NULL,
 [Descripcion] varchar(30) NOT NULL,
 [EstadoLib] varchar(1) NOT NULL,
 [CedulaPro] varchar(15) NOT NULL,
 PRIMARY KEY ([CodigoLib]),
 FOREIGN KEY ([CedulaPro]) REFERENCES [proveedores] ([CedulaPro])
)
