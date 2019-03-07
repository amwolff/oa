FROM debian:jessie

WORKDIR /app

COPY . /app/dataharvester

# RUN apt-get update -y && \
#     apt-get install -y ca-certificates && \
#     update-ca-certificates

ENTRYPOINT ["/app/dataharvester/dataharvester", "--config", "/app/dataharvester/config.yml"]
