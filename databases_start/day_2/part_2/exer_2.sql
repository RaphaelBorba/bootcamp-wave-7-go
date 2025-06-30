INSERT INTO movies
  (title, rating, awards, release_date, length, genre_id)
VALUES
  ('O Grande Desafio', 8.2, 2, '2025-07-01 00:00:00', 130, NULL);


INSERT INTO genres
  (`created_at`,`name`,`ranking`,`active`)
VALUES
  (NOW(), 'Aventura Épica', 13, 1);

UPDATE movies
   SET genre_id = 13
 WHERE title = 'O Grande Desafio';

UPDATE actors
   SET favorite_movie_id = (
     SELECT id FROM movies WHERE title = 'O Grande Desafio'
   )
 WHERE id = 1;

CREATE TEMPORARY TABLE tmp_movies AS
SELECT * FROM movies;

DELETE FROM tmp_movies
 WHERE awards < 5;

SELECT DISTINCT g.id, g.name
  FROM genres AS g
  JOIN movies AS m ON m.genre_id = g.id;

SELECT a.first_name, a.last_name
  FROM actors AS a
  JOIN movies AS m ON a.favorite_movie_id = m.id
 WHERE m.awards > 3;

CREATE INDEX idx_movies_title
  ON movies (title);


SHOW INDEX
  FROM movies
 WHERE Key_name = 'idx_movies_title';


-- Percebi que, ao criar índices em colunas muito usadas da movies para filtragem (title e release_date) as consultas 
-- ficarão mais rápidas, porque o banco não precisa mais varrer a tabela inteira toda vez que fizer um filtro 
-- ou ordenação nessas colunas.

-- Eu também incluiria um índice em episodes(season_id), já que praticamente todas as consultas puxam episódios
--  por temporada. Com esse índice, buscar só os episódios de uma determinada temporada se torna muito mais eficiente.













