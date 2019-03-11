# Collected data releases

## Overview
This directory contains time series data collected from public transportation vehicles of the city of Olsztyn (53.773056, 20.476111) - Poland.

## Technical overview
- Sharded per day only (except first two months).
- Timestamps are in traditional POSTGRES "Zulu" format.
- There's always a gap between 1AM and 3AM.
- All columns have data types respective to definitions in `db/1-olsztyn_live.sql`.
    - Header fields naming is original, meaning it was taken from transportation agency system's responses.
    - I may translate it to English later.

## Downloads
### November 2018
[Link](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2018/11.csv.xz)

### December 2018
[Link](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2018/12.csv.xz)

### January 2019
|                                         Tuesday                                        |                                        Wednesday                                       |                                        Thursday                                        |                                         Friday                                         |                                        Saturday                                        |                                         Sunday                                         |                                         Monday                                         |
|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|
|  [1](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/1.csv.xz)  |  [2](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/2.csv.xz)  |  [3](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/3.csv.xz)  |  [4](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/4.csv.xz)  |  [5](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/5.csv.xz)  |  [6](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/6.csv.xz)  |  [7](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/7.csv.xz)  |
|  [8](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/8.csv.xz)  |  [9](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/9.csv.xz)  | [10](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/10.csv.xz) | [11](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/11.csv.xz) | [12](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/12.csv.xz) | [13](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/13.csv.xz) | [14](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/14.csv.xz) |
| [15](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/15.csv.xz) | [16](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/16.csv.xz) | [17](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/17.csv.xz) | [18](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/18.csv.xz) | [19](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/19.csv.xz) | [20](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/20.csv.xz) | [21](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/21.csv.xz) |
| [22](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/22.csv.xz) | [23](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/23.csv.xz) | [24](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/24.csv.xz) | [25](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/25.csv.xz) | [26](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/26.csv.xz) | [27](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/27.csv.xz) | [28](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/28.csv.xz) |
| [29](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/29.csv.xz) | [30](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/30.csv.xz) | [31](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/1/31.csv.xz) |                                                                                        |                                                                                        |                                                                                        |                                                                                        |

### February 2019
|                                         Friday                                         |                                        Saturday                                        |                                         Sunday                                         |                                         Monday                                         |                                         Tuesday                                        |                                        Wednesday                                       |                                        Thursday                                        |
|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|
|  [1](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/1.csv.xz)  |  [2](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/2.csv.xz)  |  [3](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/3.csv.xz)  |  [4](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/4.csv.xz)  |  [5](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/5.csv.xz)  |  [6](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/6.csv.xz)  |  [7](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/7.csv.xz)  |
|  [8](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/8.csv.xz)  |  [9](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/9.csv.xz)  | [10](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/10.csv.xz) | [11](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/11.csv.xz) | [12](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/12.csv.xz) | [13](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/13.csv.xz) | [14](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/14.csv.xz) |
| [15](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/15.csv.xz) | [16](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/16.csv.xz) | [17](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/17.csv.xz) | [18](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/18.csv.xz) | [19](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/19.csv.xz) | [20](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/20.csv.xz) | [21](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/21.csv.xz) |
| [22](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/22.csv.xz) | [23](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/23.csv.xz) | [24](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/24.csv.xz) | [25](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/25.csv.xz) | [26](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/26.csv.xz) | [27](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/27.csv.xz) | [28](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/2/28.csv.xz) |

### March 2019
|                                         Friday                                         |                                        Saturday                                        |                                         Sunday                                         |                                         Monday                                         |                                         Tuesday                                        |                                        Wednesday                                       |                                        Thursday                                        |
|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|:--------------------------------------------------------------------------------------:|
|  [1](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/1.csv.xz)  |  [2](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/2.csv.xz)  |  [3](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/3.csv.xz)  |  [4](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/4.csv.xz)  |  [5](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/5.csv.xz)  |  [6](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/6.csv.xz)  |  [7](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/7.csv.xz)  |
|  [8](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/8.csv.xz)  |  [9](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/9.csv.xz)  | [10](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/10.csv.xz) | [11](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/11.csv.xz) | [12](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/12.csv.xz) | [13](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/13.csv.xz) | [14](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/14.csv.xz) |
| [15](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/15.csv.xz) | [16](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/16.csv.xz) | [17](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/17.csv.xz) | [18](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/18.csv.xz) | [19](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/19.csv.xz) | [20](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/20.csv.xz) | [21](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/21.csv.xz) |
| [22](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/22.csv.xz) | [23](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/23.csv.xz) | [24](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/24.csv.xz) | [25](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/25.csv.xz) | [26](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/26.csv.xz) | [27](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/27.csv.xz) | [28](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/28.csv.xz) |
| [29](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/29.csv.xz) | [30](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/30.csv.xz) | [31](https://s3.eu-central-1.amazonaws.com/olsztynskie-autobusy-data/2019/3/31.csv.xz) |                                                                                        |                                                                                        |                                                                                        |                                                                                        |

## Extraction
- Decompress: `xz -d ./2019/1/1.xz`;
- That's it.

## License
It's public domain data. I believe everything falls under CC0.

***

*Scratches*:
```
docker run --rm -it -v $GOPATH/src/github.com/amwolff/oa/data/2019/3/:/oadata/ -w /oadata/ postgres psql -h 159.89.5.189 -p 5432 -d oadb -U data_service # inside tty: \i dump.psql
```