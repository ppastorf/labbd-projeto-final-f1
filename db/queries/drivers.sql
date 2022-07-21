-- relatorio de status
SELECT status.status, COUNT(results.statusid)
  FROM results INNER JOIN status ON results.statusid = status.statusid INNER JOIN driver ON driver.driverid = results.driverid
  WHERE driver.driverid = (select idoriginal from users where userid = %d)
  GROUP BY status.statusid, driver.driverid
  ORDER BY count DESC

-- relatorio de vitorias
  SELECT YEAR AS ano, name AS corrida, SUM(POSITION) AS vitoria
    FROM (SELECT POSITION, driverid, raceid FROM results WHERE driverid=(select idoriginal from users where userid = %d) AND POSITION=1) x
    LEFT JOIN (SELECT YEAR, raceid, name FROM races) y ON x.raceid = y.raceid
    GROUP BY ROLLUP(ano, corrida)
    ORDER BY ano, corrida;
