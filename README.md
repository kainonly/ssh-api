# ssh api

ssh server api

[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/ssh-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-api)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/kainonly/ssh-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-api)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/kainonly/ssh-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-api)
[![TypeScript](https://img.shields.io/badge/%3C%2F%3E-TypeScript-blue.svg?style=flat-square)](https://github.com/kainonly/ssh-api)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/ssh-api/master/LICENSE)

```shell
docker pull kainonly/ssh-api
```

## Docker Compose

example

```yml
version: '3.7'
services:
  ssh:
    image: kainonly/ssh-api
    restart: always
    volumes: 
      - ./app:/app
    ports:
      - 3000:3000
```

## Api docs

Assume that the underlying request path is `http://localhost:3000`

#### Testing Connect

- url `/testing`
- method `POST`
- body
  - **host** `string`
  - **port** `number`
  - **username** `string`
  - **password** `string` SSH password, default empty
  - **private_key** `string` SSH private key (Base64)
  - **passphrase** `string` private key passphrase

Password connect

```json
{
	"host":"imac",
	"port":22,
	"username":"root",
	"password":"123456"
}
```

Private key connect

```json
{
	"host":"imac",
	"port":22,
	"username":"root",
	"private_key":"LS0tL.......tFWS0tLS0tCg=="
}
```

- response
  - **error** `number` status
  - **msg** `string` Message

```json
{
    "error": 0,
    "msg": "ok"
}
```

#### Put SSH

- url `/put`
- method `POST`
- body
  - **identity** `string` ssh identity code
  - **host** `string`
  - **port** `number`
  - **username** `string`
  - **password** `string` SSH password, default empty
  - **private_key** `string` SSH private key (Base64)
  - **passphrase** `string` private key passphrase

```json
{
    "identity":"test",
	"host":"imac",
	"port":22,
	"username":"root",
	"private_key":"LS0tL.......tFWS0tLS0tCg=="
}
```

- response
  - **error** `number` status
  - **msg** `string` Message

```json
{
    "error": 0,
    "msg": "ok"
}
```

#### Exec SSH

- url `/exec`
- method `POST`
- body
  - **identity** `string` ssh identity code
  - **bash** `string`

```json
{
	"identity":"test",
	"bash":"uptime"
}
```

- response `application/octet-stream`

```
10:38:39 up 12 days, 19:34,  1 user,  load average: 0.11, 0.09, 0.06
```

#### Delete SSH

- url `/delete`
- method `POST`
- body
  - **identity** `string` ssh identity code

```json
{
	"identity":"test"
}
```

- response
  - **error** `number` status
  - **msg** `string` Message

```json
{
    "error": 0,
    "msg": "ok"
}
```

#### Get SSH

- url `/get`
- method `POST`
- body
  - **identity** `string` ssh identity code

```json
{
	"identity":"test"
}
```

- response
  - **error** `number` status
  - data
    - **identity** `string` ssh identity code
    - **host** `string`
    - **port** `number`
    - **username** `string`
    - **connected** `boolean` ssh connected status
    - **tunnels** `array` ssh tunnels set
    - **tunnelsListening** `array` running tunnels

```json
{
    "error": 0,
    "data": {
        "identity": "test",
        "host": "192.168.1.102",
        "port": 22,
        "username": "root",
        "connected": true,
        "tunnels": [],
        "tunnelsListening": []
    }
}
```

#### Get All Identity

- url `/all`
- method `POST`

#### Lists SSH

- url `/lists`
- method `POST`
- body
  - **identity** `array` ssh identity code

```json
{
	"identity":["test"]
}
```

- response
  - **error** `number` status
  - **data** `array`
    - **identity** `string` ssh identity code
    - **host** `string`
    - **port** `number`
    - **username** `string`
    - **connected** `boolean` ssh connected status
    - **tunnels** `array` ssh tunnels set
    - **tunnelsListening** `array` running tunnels

```json
{
    "error": 0,
    "data": [
        {
            "identity": "test",
            "host": "192.168.1.102",
            "port": 22,
            "username": "root",
            "connected": true,
            "tunnels": [
                [
                    "127.0.0.1",
                    27017,
                    "127.0.0.1",
                    27017
                ]
            ],
            "tunnelsListening": [
                true
            ]
        }
    ]
}
```

#### Set Tunnel

- url `/tunnels`
- method `POST`
- body
  - **identity** `string` ssh identity code
  - **tunnels** `array`, tunnel config `[<srcIp>,<srcPort>,<dstIP>,<dstPort>]`

```json
{
	"identity":"test",
	"tunnels":[
		["127.0.0.1",27017,"127.0.0.1",27017]
	]
}
```

- response
  - **error** `number` status
  - **msg** `string` Message

```json
{
    "error": 0,
    "msg": "ok"
}
```
