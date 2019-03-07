# Collected data releases
## Overview
This directory contains time series data collected from public transportation vehicles of the city of Olsztyn (53.773056, 20.476111) - Poland.
## Technical overview
- Sharded per day only (except first two months).
- There's always a gap between 1AM and 3AM.
- All columns have data types respective to definitions in `db/1-olsztyn_live.sql`.
    - Header fields naming is original, meaning it was taken from transportation agency system's responses.
    - I may translate it to English later.
- Timestamps are in traditional POSTGRES "Zulu" format.
## Downloads
<!-- Link to storage -->
## Extraction
- Decompress: `xz -d ./2019/1/1.xz`;
- That's it.
## License
It's public domain data. I believe everything falls under CC0.

***

*Scratches*:
```
docker run --rm -it -v $GOPATH/src/github.com/amwolff/oa/data/2019/2/:/oadata/ -w /oadata/ postgres psql -h 159.89.5.189 -p 5432 -d oadb -U data_service # inside tty: \i dump.psql
```
