
CREATE DATABASE IF NOT EXISTS library_db
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
USE library_db;

CREATE TABLE `AUTOR` (
  `idAutor`     INT            NOT NULL AUTO_INCREMENT,
  `Nombre`      VARCHAR(100)   NOT NULL,
  `Nacionalidad`VARCHAR(100),
  PRIMARY KEY (`idAutor`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `LIBRO` (
  `idLibro`     INT            NOT NULL AUTO_INCREMENT,
  `Titulo`      VARCHAR(255)   NOT NULL,
  `Editorial`   VARCHAR(255),
  `Area`        VARCHAR(100),
  PRIMARY KEY (`idLibro`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `ESTUDIANTE` (
  `idLector`    INT            NOT NULL AUTO_INCREMENT,
  `Nombre`      VARCHAR(100)   NOT NULL,
  `Apellido`    VARCHAR(100),
  `Direccion`   VARCHAR(255),
  `Carrera`     VARCHAR(100),
  `Edad`        INT,
  PRIMARY KEY (`idLector`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `LIBROAUTOR` (
  `idAutor`     INT            NOT NULL,
  `idLibro`     INT            NOT NULL,
  PRIMARY KEY (`idAutor`,`idLibro`),
  CONSTRAINT `fk_la_autor`
    FOREIGN KEY (`idAutor`) REFERENCES `AUTOR`   (`idAutor`)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_la_libro`
    FOREIGN KEY (`idLibro`) REFERENCES `LIBRO`   (`idLibro`)
    ON UPDATE NO ACTION ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `PRESTAMO` (
  `idPrestamo`      INT            NOT NULL AUTO_INCREMENT,
  `idLector`        INT            NOT NULL,
  `idLibro`         INT            NOT NULL,
  `FechaPrestamo`   DATE           NOT NULL,
  `FechaDevolucion` DATE,
  `Devuelto`        BOOLEAN        NOT NULL DEFAULT FALSE,
  PRIMARY KEY (`idPrestamo`),
  KEY `idx_prestamo_lector` (`idLector`),
  KEY `idx_prestamo_libro`  (`idLibro`),
  CONSTRAINT `fk_pr_lector`
    FOREIGN KEY (`idLector`) REFERENCES `ESTUDIANTE` (`idLector`)
    ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_pr_libro`
    FOREIGN KEY (`idLibro`)  REFERENCES `LIBRO`      (`idLibro`)
    ON UPDATE NO ACTION ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `AUTOR` (`idAutor`, `Nombre`, `Nacionalidad`) VALUES
 (1, 'Gabriel García Márquez', 'Colombiana'),
 (2, 'Isabel Allende',         'Chilena'),
 (3, 'J.K. Rowling',           'Britânica'),
 (4, 'Haruki Murakami',        'Japonesa'),
 (5, 'Chimamanda Ngozi Adichie','Nigeriana');

INSERT INTO `LIBRO` (`idLibro`, `Titulo`,                       `Editorial`,               `Area`) VALUES
 (1, 'Cem Anos de Solidão',       'Editorial Sudamericana',   'Literatura'),
 (2, 'A Casa dos Espíritos',      'Plaza & Janés',            'Literatura'),
 (3, 'Harry Potter e a Pedra Filosofal','Bloomsbury',          'Fantasia'),
 (4, 'Kafka à Beira-Mar',         'Shinchosha',               'Ficção'),
 (5, 'Americanah',                'Knopf',                    'Romance');

INSERT INTO `ESTUDIANTE` (`idLector`, `Nombre`,  `Apellido`,     `Direccion`,                  `Carrera`,                 `Edad`) VALUES
 (1, 'Carlos',  'Silva',    'Rua das Flores, 123',       'Engenharia',               21),
 (2, 'Mariana', 'Rodriguez','Avenida Brasil, 456',       'Medicina',                 22),
 (3, 'João',    'Pereira',  'Praça Central, 789',        'Direito',                  20),
 (4, 'Ana',     'Souza',    'Rua XV de Novembro, 10',    'Ciências da Computação',   23),
 (5, 'Luís',    'Fernandes','Travessa do Comércio, 5',  'Administração',            24);

INSERT INTO `LIBROAUTOR` (`idAutor`, `idLibro`) VALUES
 (1, 1),
 (2, 2),
 (3, 3),
 (4, 4),
 (5, 5);

INSERT INTO `PRESTAMO` (`idPrestamo`, `idLector`, `idLibro`, `FechaPrestamo`, `FechaDevolucion`, `Devuelto`) VALUES
 (1, 1, 1, '2025-06-01', '2025-06-15', TRUE),
 (2, 2, 2, '2025-06-02', '2025-06-12', TRUE),
 (3, 3, 3, '2025-06-03', NULL,          FALSE),
 (4, 4, 4, '2025-06-04', '2025-06-20', TRUE),
 (5, 5, 5, '2025-06-05', NULL,          FALSE);


SELECT * FROM AUTOR a 

SELECT e.Nombre , e.Edad FROM ESTUDIANTE e 

SELECT * FROM ESTUDIANTE e WHERE e.Carrera = 'Ciências da Computação'

SELECT * FROM AUTOR a WHERE a.Nacionalidad = 'Francesa' OR a.Nacionalidad = 'Italiana'

SELECT * FROM LIBRO l WHERE l.Area != 'Internet'

SELECT * FROM LIBRO l WHERE l.Editorial = 'Salamandra'

SELECT * FROM ESTUDIANTE e WHERE e.Edad > (SELECT AVG(e.Edad) FROM ESTUDIANTE e)

SELECT * FROM ESTUDIANTE e WHERE e.Apellido LIKE 'G%';

SELECT a.Nombre, a.idAutor FROM LIBROAUTOR l INNER JOIN AUTOR a ON l.idAutor = a.idAutor
	INNER JOIN LIBRO l2 ON l.idLibro  = l2.idLibro
	WHERE l2.Titulo = 'O Universo: Guia de Viagem'
	
SELECT l2.Titulo FROM PRESTAMO p INNER JOIN ESTUDIANTE e ON p.idLector = e.idLector
	INNER JOIN LIBRO l2 ON p.idLibro  = l2.idLibro
	WHERE e.Nombre = 'Filippo' AND e.Apellido = 'Galli'
		

SELECT Min(e.edad) FROM ESTUDIANTE e 

SELECT l2.Titulo FROM LIBROAUTOR l INNER JOIN AUTOR a ON l.idAutor = a.idAutor
	INNER JOIN LIBRO l2 ON l.idLibro  = l2.idLibro
	WHERE a.Nombre = 'J.K. Rowling'

SELECT l2.Titulo FROM PRESTAMO p 
	INNER JOIN LIBRO l2 ON p.idLibro  = l2.idLibro
	WHERE p.FechaPrestamo = '2021-07-16'