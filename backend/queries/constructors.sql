3.

SELECT nome_completo,
       COALESCE(vitorias, 0) AS vitorias
FROM
  (SELECT p.driverid,
          nome_completo
   FROM
     (SELECT DISTINCT driverid
      FROM results
      WHERE constructorid=9) p --consulta pilotos distintos que correram pela escuderia

   LEFT JOIN
     (SELECT driverid,
             CONCAT(forename, ' ', surname) AS nome_completo
      FROM driver) n --obtém nome completo
 ON p.driverid = n.driverid) parcial
LEFT JOIN
  (SELECT driverid,
          count(*) AS vitorias
   FROM results
   WHERE constructorid=9
     AND POSITION=1
   GROUP BY driverid) AS v --obtém vitorias por aquela escuderia
ON parcial.driverid = v.driverid;

4.

SELECT status.status, COUNT(results.statusid)
	FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN constructors ON constructors.constructorid = results.constructorid
	WHERE constructors.constructorid = 1
	GROUP BY status.statusid, constructors.CONSTRUCTORID
	ORDER BY count DESC