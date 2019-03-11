FROM debian:jessie

WORKDIR /app

COPY . /app/dirserver

# RUN apt-get update -y && \
#     apt-get install -y ca-certificates && \
#     update-ca-certificates

EXPOSE 80

ENTRYPOINT ["/app/dirserver/dirserver", "--config", "/app/dirserver/config.yml"]
