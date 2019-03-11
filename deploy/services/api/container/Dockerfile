FROM debian:jessie

WORKDIR /app

COPY . /app/api

# RUN apt-get update -y && \
#     apt-get install -y ca-certificates && \
#     update-ca-certificates

EXPOSE 80

ENTRYPOINT ["/app/api/api", "--config", "/app/api/config.yml"]
