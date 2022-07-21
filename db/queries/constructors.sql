-- relatorio de status
SELECT status.status, COUNT(results.statusid) 
  FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN constructors ON constructors.constructorid = results.constructorid
  WHERE constructors.constructorid = (select idoriginal from users where userid = %d)
  GROUP BY status.statusid, constructors.constructorid
  ORDER BY count DESC

-- relatorio de pilotos da escuderia
		SELECT nome_completo,
		COALESCE(vitorias, 0) AS vitorias
		FROM
		(SELECT p.driverid,
			nome_completo
		FROM
		(SELECT DISTINCT driverid
		FROM results
		WHERE constructorid=(select idoriginal from users where userid = %d)) p

		LEFT JOIN
		(SELECT driverid,
				CONCAT(forename, ' ', surname) AS nome_completo
		FROM driver) n 
		ON p.driverid = n.driverid) parcial
		LEFT JOIN
		(SELECT driverid,
			count(*) AS vitorias
		FROM results
		WHERE constructorid=(select idoriginal from users where userid = %d)
		AND POSITION=1
		GROUP BY driverid) AS v
		ON parcial.driverid = v.driverid;

-- consulta de pilotos por nome
	SELECT y.nome_completo,
			y.data_nascimento,
			y.nacionalidade
		FROM
		(SELECT DISTINCT driverid
		FROM results
		WHERE constructorid = (select idoriginal from users where userid = %d)) x
		INNER JOIN
		(SELECT driverid,
			CONCAT(forename, ' ', surname) AS nome_completo,
			dob AS data_nascimento,
			nationality AS nacionalidade
		FROM driver
		WHERE forename='%s') y ON x.driverid = y.driverid;