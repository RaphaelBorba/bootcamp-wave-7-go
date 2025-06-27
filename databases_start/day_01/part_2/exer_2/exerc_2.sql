/*

    - Primery Key de clientes é o ID e o mesmo serve para os planos, eu poderia 
        usar nomes diferentes como por exemplo 'id_client' e 'plan_id'
    - A relação foi de planos 1:n clientes, porque um plano pode ter muitos clientes, mas um cliente pode ter apenas 1 plano.
        Seria válido fazer uma relação de n:m já que um cliente pode ter mais de um endereço.



*/



DROP DATABASE IF EXISTS `empresa_internet`;
CREATE DATABASE `empresa_internet`
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
USE `empresa_internet`;

CREATE TABLE `internt_plans` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `internet_velocity` INT       NOT NULL,
  `price`             FLOAT     NOT NULL,
  `discount`          FLOAT     NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `client` (
  `id`                INT       NOT NULL AUTO_INCREMENT,
  `name`              TEXT      NOT NULL,
  `last_name`         TEXT,
  `birth_date`        DATETIME,
  `region`            TEXT,
  `city`              TEXT,
  `internet_plan_id`  INT,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  KEY `fk_client_plan_idx` (`internet_plan_id`),
  CONSTRAINT `fk_internt_plans_id_client`
    FOREIGN KEY (`internet_plan_id`)
    REFERENCES `internt_plans` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `internt_plans` (`internet_velocity`, `price`, `discount`) VALUES
  (10,  49.90, 0.00),
  (50,  89.90, 5.00),
  (100,129.90, 10.0),
  (200,199.90, 15.0),
  (500,299.90, 20.0);

INSERT INTO `client` (`name`, `last_name`, `birth_date`, `region`, `city`, `internet_plan_id`) VALUES
  ('Alice',    'Silva',    '1985-03-12 00:00:00', 'Norte',        'Manaus',       1),
  ('Bruno',    'Costa',    '1990-07-21 00:00:00', 'Nordeste',     'Recife',       2),
  ('Camila',   'Oliveira', '1978-11-05 00:00:00', 'Sul',          'Porto Alegre', 3),
  ('Daniel',   'Santos',   '1995-01-30 00:00:00', 'Sudeste',      'São Paulo',    4),
  ('Eduarda',  'Pereira',  '1982-05-17 00:00:00', 'Centro-Oeste', 'Goiânia',      1),
  ('Felipe',   'Rodrigues','1988-09-09 00:00:00', 'Sudeste',      'Rio de Janeiro',5),
  ('Gabriela', 'Almeida',  '1992-12-25 00:00:00', 'Nordeste',     'Salvador',     2),
  ('Henrique', 'Ferreira', '1975-04-14 00:00:00', 'Sul',          'Curitiba',     3),
  ('Isabela',  'Gomes',    '1987-08-02 00:00:00', 'Norte',        'Belém',        4),
  ('João',     'Martins',  '1983-02-18 00:00:00', 'Centro-Oeste', 'Brasília',     5);


SELECT * 
  FROM client;

SELECT id, name, last_name, city
  FROM client
 WHERE region = 'Sul';

SELECT id, name, last_name, internet_plan_id
  FROM client
 WHERE internet_plan_id = 3;

SELECT COUNT(*) AS total_clients
  FROM client;

SELECT COUNT(*) AS clients_sudeste
  FROM client
 WHERE region = 'Sudeste';

SELECT * 
  FROM internt_plans;

SELECT id, internet_velocity, price, discount
  FROM internt_plans
 WHERE discount > 0;

SELECT AVG(price) AS avg_plan_price
  FROM internt_plans;

SELECT MAX(internet_velocity) AS top_speed_mb
  FROM internt_plans;

SELECT id, internet_velocity, price
  FROM internt_plans
 ORDER BY price DESC
 LIMIT 3;
