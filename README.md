# ecs_demo_app

## Docker Image

[grandcolline/ecs_demo_app](https://hub.docker.com/r/grandcolline/ecs_demo_app/)

## RUN

```
$ docker run --rm -p 80:8080 -it grandcolline/ecs_demo_app
```

## USAGE

```
GET /hc         : ヘルスチェック
GET /info       : インスタンスIDとローカルIPを返す
GET /fibo?n=5   : n番目のフィボナッチ数を返す
```
