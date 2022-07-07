1. 

SELECT status.status, COUNT(RESULTS.statusid)
	FROM results INNER JOIN STATUS ON results.statusid = status.statusid
	GROUP BY status.statusid
	ORDER BY count DESC

2.

SELECT c.name AS nome_da_cidade,
       iatacode AS codigo_iata,
       a.name AS nome_aeroporto,
       city AS cidade_aeroporto,
       TYPE AS tipo_aeroporto,
               earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) AS distancia
FROM
  (SELECT name,
          lat, long
   FROM geocities15k
   WHERE name=<INSERIR NOME DA CIDADE ESCOLHIDA>) c -- ATENÇÂO !!!!!
LEFT JOIN
  (SELECT iatacode,
          name,
          city,
          TYPE,
          latdeg,
          longdeg
   FROM airports
   WHERE isocountry='BR'
     AND (TYPE='medium_airport'
          OR TYPE='large_airport')) a ON earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) <= 100000;