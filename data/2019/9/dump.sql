\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-01' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-02' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 1.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-02' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-03' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 2.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-03' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-04' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 3.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-04' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-05' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 4.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-05' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-06' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 5.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-06' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-07' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 6.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-07' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-08' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 7.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-08' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-09' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 8.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-09' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-10' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 9.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-10' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-11' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 10.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-11' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-12' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 11.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-12' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-13' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 12.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-13' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-14' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 13.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-14' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-15' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 14.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-15' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-16' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 15.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-16' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-17' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 16.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-17' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-18' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 17.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-18' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-19' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 18.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-19' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-20' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 19.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-20' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-21' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 20.csv.xz' DELIMITER ',' CSV HEADER
\COPY (SELECT * FROM olsztyn_live.vehicles WHERE ts BETWEEN (SELECT ts FROM olsztyn_live.vehicles WHERE ts >= '2019-09-21' ORDER BY ts ASC LIMIT 1) AND (SELECT ts FROM olsztyn_live.vehicles WHERE ts < '2019-09-22' ORDER BY ts DESC LIMIT 1) ORDER BY ts ASC) TO PROGRAM 'xz > 21.csv.xz' DELIMITER ',' CSV HEADER
