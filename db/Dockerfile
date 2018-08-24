FROM postgres

ENV POSTGRES_USER=data_service
ENV POSTGRES_PASSWORD=data_service
ENV POSTGRES_DB=oadb

COPY 0-create.sql /docker-entrypoint-initdb.d/