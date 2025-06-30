CREATE TEMPORARY TABLE TWD AS
SELECT 
  sr.title as serie, e.title as ep_name, s.`number` as season
FROM episodes AS e
JOIN seasons  AS s  ON e.season_id = s.id
JOIN series   AS sr ON s.serie_id  = sr.id
WHERE sr.title = 'The Walking Dead';

SELECT 
  t.* 
FROM TWD AS t
WHERE season = 1



CREATE INDEX idx_movies_release_date
  ON movies (release_date);

SHOW INDEX
  FROM movies
 WHERE Key_name = 'idx_movies_release_date';

 -- Acredito que a filtragem por data seja um dos select mais utilizado, principalmente quando tratamos de filmes.
 -- Por isso criei um index na coluna release_date dos filmes, para quando uma filtragem por data for feita, o retorno será mais rápido