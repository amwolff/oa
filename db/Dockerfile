FROM postgres

ENV POSTGRES_USER=data_service
ENV POSTGRES_PASSWORD=data_service
ENV POSTGRES_DB=oadb

COPY 0-create.sql 1-olsztyn_live.sql 2-olsztyn_static.sql /docker-entrypoint-initdb.d/
