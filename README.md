# ssh api

ssh server api

[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/ssh-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-api)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/kainonly/ssh-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-api)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/kainonly/ssh-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/ssh-api)
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
      - ./data:/app/data
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

- response
  - **error** `number` status
  - **data** `string` Message

```json
{
    "error": 0,
    "data": " 09:42:22 up 9 days, 23:43,  1 user,  load average: 0.26, 0.22, 0.19\n"
}
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
    - **connected** `string` ssh connected client version
    - **tunnels** `array` ssh tunnels set

```json
{
    "error": 0,
    "data": {
        "identity": "test",
        "host": "imac",
        "port": 22,
        "username": "root",
        "connected": "SSH-2.0-Go",
        "tunnels": [
            {
                "src_ip": "127.0.0.1",
                "src_port": 5601,
                "dst_ip": "127.0.0.1",
                "dst_port": 5601
            }
        ]
    }
}
```

#### Get All Identity

- url `/all`
- method `POST`

```json
{
    "error": 0,
    "data": [
        "test"
    ]
}
```

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
    - **connected** `string` ssh connected client version
    - **tunnels** `array` ssh tunnels set

```json
{
    "error": 0,
    "data": [
        {
            "identity": "test",
            "host": "imac",
            "port": 22,
            "username": "root",
            "connected": "SSH-2.0-Go",
            "tunnels": [
                {
                "src_ip": "127.0.0.1",
                "src_port": 5601,
                "dst_ip": "127.0.0.1",
                "dst_port": 5601
                }
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
  - **tunnels** `array` tunnels options
    - **src_ip** `string` origin ip
    - **src_port** `int` origin port
    - **dst_ip** `string` target ip
    - **dst_port** `int` target port

```json
{
    "identity":"test",
    "tunnels":[
        {
            "src_ip":"127.0.0.1",
            "src_port":3306,
            "dst_ip":"127.0.0.1",
            "dst_port":3306
        },
        {
            "src_ip":"127.0.0.1",
            "src_port":9200,
            "dst_ip":"127.0.0.1",
            "dst_port":9200
        }
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
