\connect oadb

CREATE SCHEMA IF NOT EXISTS olsztyn_live;
GRANT USAGE ON SCHEMA olsztyn_live TO data_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA olsztyn_live TO data_service;
ALTER DEFAULT PRIVILEGES IN SCHEMA olsztyn_live GRANT ALL ON TABLES TO data_service;

CREATE TABLE olsztyn_live.vehicles (
  id                  SERIAL                   NOT NULL PRIMARY KEY,
  ts                  TIMESTAMP WITH TIME ZONE NOT NULL,

  nr_radia            INTEGER,
  nb                  INTEGER,
  numer_lini          VARCHAR(3),
  war_trasy           VARCHAR(1),
  kierunek            VARCHAR(1),
  id_kursu            INTEGER,
  lp_przyst           INTEGER,
  droga_plan          INTEGER,
  droga_wyko          INTEGER,
  dlugosc             NUMERIC(7, 5),
  szerokosc           NUMERIC(7, 5),
  prev_dlugosc        NUMERIC(7, 5),
  prev_szerokosc      NUMERIC(7, 5),
  odchylenie          INTEGER,
  odchylenie_str      VARCHAR(9),
  stan                INTEGER,
  plan_godz_rozp      VARCHAR(5),
  nast_id_kursu       INTEGER,
  nast_plan_godz_rozp VARCHAR(5),
  nast_num_lini       VARCHAR(3),
  nast_war_trasy      VARCHAR(1),
  nast_kierunek       VARCHAR(1),
  ile_sek_do_odjazdu  INTEGER,
  typ_pojazdu         VARCHAR(1),
  transport           VARCHAR(1),
  cechy               VARCHAR(3),
  opis_tabl           VARCHAR,
  nast_opis_tabl      VARCHAR,
  wektor              NUMERIC(5, 2)
);

CREATE INDEX ON olsztyn_live.vehicles (ts DESC);
