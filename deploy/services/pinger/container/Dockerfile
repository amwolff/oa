FROM debian:jessie

WORKDIR /app

COPY . /app/pinger

# RUN apt-get update -y && \
#     apt-get install -y ca-certificates && \
#     update-ca-certificates

RUN apt-get update -y && \
    apt-get install -y curl

RUN ["chmod", "+x", "/app/pinger/pinger.sh"]

ENV CLIENT_COOKIE <your-ClientCookie>

ENTRYPOINT ["/app/pinger/pinger.sh"]
