-- relatorio de status
SELECT status.status, COUNT(results.statusid)
  FROM results INNER JOIN status ON results.statusid = status.statusid
  GROUP BY status.statusid
  ORDER BY count DESC

-- relatorio de aeroportos
SELECT  c.name AS nome_da_cidade,
    iatacode AS codigo_iata,
    a.name AS nome_aeroporto,
    city AS cidade_aeroporto,
    TYPE AS tipo_aeroporto,
    earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) AS distancia
  FROM (SELECT name, lat, long FROM geocities15k WHERE name='%s') c
  LEFT JOIN (
    SELECT iatacode, name, city, TYPE, latdeg, longdeg
    FROM airports
    WHERE isocountry='BR' AND (TYPE='medium_airport' OR TYPE='large_airport')
  ) a ON earth_distance(ll_to_earth(a.latdeg, a.longdeg), ll_to_earth(c.lat, c.long)) <= 100000;

-- criar escuderia
INSERT INTO public.Constructors (ConstructorId, ConstructorRef, Name, Nationality, Url)
  VALUES((SELECT MAX(ConstructorId) FROM Constructors) + 1, '%s', '%s', '%s',' %s');

-- criar piloto
    INSERT INTO public.Driver (DriverId, Driverref, Number, Code, Forename, Surname, Dob, Nationality)
      VALUES((SELECT MAX(DriverId) FROM Driver) + 1, '%s', '%s', '%s', '%s', '%s', '%s', '%s');

  
