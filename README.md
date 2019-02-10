# eks_demo_app

EKSの勉強会で使う簡単なアプリケーション

## Docker Image

[grandcolline/eks_demo_app](https://hub.docker.com/r/grandcolline/eks_demo_app/)

## RUN

```
$ docker run --rm -p 80:8080 -it grandcolline/eks_demo_app
```

## USAGE

```
GET /hc         : ヘルスチェック
GET /info       : インスタンスIDとバージョンを返す
GET /down       : アプリケーション終了
GET /fibo?n=5   : n番目のフィボナッチ数を返す
```
