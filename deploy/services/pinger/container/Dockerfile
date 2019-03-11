FROM debian:jessie

ENV CLIENT_COOKIE <your-ClientCookie>

WORKDIR /app

COPY . /app/pinger

# RUN apt-get update -y && \
#     apt-get install -y ca-certificates && \
#     update-ca-certificates

RUN apt-get update -y && \
    apt-get install -y curl

RUN ["chmod", "+x", "/app/pinger/pinger.sh"]

ENTRYPOINT ["/app/pinger/pinger.sh"]
