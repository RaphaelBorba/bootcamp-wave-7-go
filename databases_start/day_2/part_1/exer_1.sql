CREATE TABLE FUNCIONARIO (
    cod_emp VARCHAR(10),
    nome VARCHAR(50),
    sobrenome VARCHAR(50),
    posto VARCHAR(50),
    data_alta DATE,
    salario DECIMAL(10, 2),
    comissao DECIMAL(10, 2),
    depto_nro VARCHAR(10)
);

CREATE TABLE DEPARTAMENTO (
    depto_nro VARCHAR(10),
    nombre_depto VARCHAR(50),
    localidad VARCHAR(50)
);

/*

Selecciona o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
Mostra os departamentos com mais de cinco empregados.
Mostra o nome, o salário e o nome do departamento dos empregados que têm a mesma posição que "Mito Barchuk".
Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.
Mostra o nome do empregado com o salário mais baixo.
Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
*/

SELECT f.nome, f.posto, d.localidad FROM FUNCIONARIO f INNER JOIN departamento d ON f.depto_nro = d.depto_nro

SELECT 
  d.depto_nro,
  d.nombre_depto,
  d.localidad,
  COUNT(*) AS total_empregados
FROM funcionario AS f
JOIN departamento AS d
  ON f.depto_nro = d.depto_nro
GROUP BY 
  d.depto_nro,
  d.nombre_depto,
  d.localidad
HAVING 
  COUNT(*) > 5;

SELECT 
  f.nome,
  f.salario,
  d.nombre_depto
FROM funcionario AS f
JOIN departamento AS d
  ON f.depto_nro = d.depto_nro
WHERE f.posto = (
  SELECT posto
  FROM funcionario
  WHERE nome = 'Mito'
    AND sobrenome = 'Barchuk'
);

SELECT
    f.*
FROM funcionario AS f
JOIN departamento AS d
  ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Contabilidade'
ORDER BY f.nome;

SELECT nome FROM FUNCIONARIO ORDER BY salario ASC LIMIT 1

SELECT 
  f.*
FROM funcionario AS f
JOIN departamento AS d
  ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Vendas'
ORDER BY f.salario DESC
LIMIT 1;