# Olsztyńskie Autobusy
<!-- Infrastructure to pull, retain and represent real-time data collected from Olsztyn public transportation vehicles -->
## About
Olsztyńskie Autobusy combine harvesting vehicles data from systems of Olsztyn public transportation agency, serving current data for web app users and retaining that data.

[![Screenshot](screenshot.png "https://autobusy.olsztyn.pl")](https://autobusy.olsztyn.pl)

Quite simply built (I believe exception might be that data pulling part), it has modular microservice architecture divided into front/back-end.
Main part of this infrastructure is the data harvester service and a Postgres instance.

## Build & Run
In order to start a dockerized instance locally, you will need to obtain session cookie and build services from source.
This walk-through assumes you have Go & Docker installed and configured.

Clone repository:
```
git clone git@github.com:amwolff/oa.git $GOPATH/src/github.com/amwolff/oa
```

Install dependencies:
```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
cd $GOPATH/src/github.com/amwolff/oa
dep ensure
```

Get session cookie:
1. In browser navigate to `sip.zdzit.olsztyn.eu`;
2. Open up inspector (Google Chrome: Control+Shift+I) > Application > Storage > Cookies;
3. Copy value of ASP.NET_SessionId cookie.

Configure services:

In files
```
# $GOPATH/src/github.com/amwolff/oa/deploy/services/dataharvester/example_config.yml
Line 3: ClientCookie: <your-ClientCookie>
# $GOPATH/src/github.com/amwolff/oa/deploy/services/pinger/Dockerfile
Line 16: ENV CLIENT_COOKIE <your-ClientCookie>
```
substitue `<your-ClientCookie>` for copied value of your session cookie.

Build everything:
```
cd $GOPATH/src/github.com/amwolff/oa/deploy
./build_all.sh -d -v 4a032c9
```

Start services:
```
docker network create backend
docker network create frontend
docker run -d --name oa_db --network backend --network-alias oa-postgres-ip-alias --rm amwolff/oa:oadb_4a032c9
docker run -d --name oa_pinger --network backend --rm amwolff/oa:pinger_4a032c9
docker run -d --name oa_dataharvester --network backend --rm amwolff/oa:dataharvester_4a032c9
docker run -d --name oa_api --network backend -p 8080:80 --rm amwolff/oa:api_4a032c9
docker network connect frontend oa_api
docker run -d --name oa_dirserver --network frontend -p 80:80 --rm amwolff/oa:dirserver_4a032c9
```

The app should be now available at `http://localhost`.

### Notes
Started instance is very similar to what I deployed in cloud for "production" use.
Additionally, I've used docker-compose along with Docker in swarm mode (to encapsulate startup instructions) and traefik (for requests routing).

### CI status
[![Build Status](https://travis-ci.com/amwolff/oa.svg?token=8LTVaXtVR2rYts8pwRmn&branch=master)](https://travis-ci.com/amwolff/oa)
