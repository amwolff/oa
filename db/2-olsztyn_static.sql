\connect oadb

CREATE SCHEMA IF NOT EXISTS olsztyn_static;
GRANT USAGE ON SCHEMA olsztyn_static TO data_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA olsztyn_static TO data_service;
ALTER DEFAULT PRIVILEGES IN SCHEMA olsztyn_static GRANT ALL ON TABLES TO data_service;

CREATE TABLE olsztyn_static.routes (
  id           SERIAL                   NOT NULL PRIMARY KEY,
  ts           TIMESTAMP WITH TIME ZONE NOT NULL,

  number       VARCHAR(3),
  description  VARCHAR,
  description2 VARCHAR,
  variant      INTEGER,
  transport    VARCHAR(1),
  direction    VARCHAR(1)
);

CREATE TABLE olsztyn_static.stops (
  id           SERIAL                   NOT NULL PRIMARY KEY,
  ts           TIMESTAMP WITH TIME ZONE NOT NULL,

  number       VARCHAR(7),
  name         VARCHAR,
  street_name  VARCHAR,
  latitude     FLOAT(20),
  longitude    FLOAT(20)
);

