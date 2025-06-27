
SELECT * from movies_db.movies m 

SELECT first_name, last_name, rating FROM movies_db.actors a 

SELECT title From movies_db.series s 

SELECT first_name, last_name, rating FROM movies_db.actors a WHERE a.rating > 7.5

SELECT m.title, m.rating, m.awards  from movies_db.movies m WHERE m.rating >7.5

SELECT m.title, m.rating  FROM movies_db.movies m ORDER BY m.rating

SELECT m.title FROM movies_db.movies m LIMIT 3

SELECT * FROM movies_db.movies m ORDER BY m.rating DESC LIMIT 5

SELECT * FROM movies_db.actors a LIMIT 10

SELECT m.title, m.rating FROM movies_db.movies m WHERE m.title LIKE '%Toy Story%'

SELECT * FROM movies_db.actors a WHERE a.first_name = 'Sam'

SELECT m.title, m.release_date FROM movies_db.movies m WHERE m.release_date < '2008-01-01' AND m.release_date > '2004-01-01'

SELECT m.title, m.release_date FROM movies_db.movies m WHERE m.release_date < '2009-01-01' AND m.release_date > '1988-01-01' AND m.rating > 3 AND m.awards > 1

