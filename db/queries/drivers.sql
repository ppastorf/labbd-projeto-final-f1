5.

SELECT YEAR AS ano,
               name AS corrida,
               SUM(POSITION) AS vitoria
FROM
  (SELECT POSITION,
          driverid,
          raceid
   FROM results
   WHERE driverid=<INSERIR_ID_COMO_PARAMETRO>
     AND POSITION=1) x
LEFT JOIN
  (SELECT YEAR,
          raceid,
          name
   FROM races) y ON x.raceid = y.raceid
GROUP BY ROLLUP(ano, corrida)
ORDER BY ano,
         corrida;

6.

SELECT status.status, COUNT(results.statusid)
	FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN driver ON driver.driverid = results.driverid
	WHERE driver.driverid = 1
	GROUP BY status.statusid, driver.driverid
	ORDER BY count DESC