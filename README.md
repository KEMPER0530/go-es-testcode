# Elasticsearch

Elasticsearch on docker.

# Features

ElasticSearch を利用したレストラン検索を実施する練習用のリポジトリです。
[ライブドアグルメデータ](https://github.com/livedoor/datasets)を編集し、
ElasticSearch へ登録してクエリの検証を実施します。

DataContet については[こちら](https://github.com/KEMPER0530/elastic-demo)を参照

## Dependency

- Go:v1.19
- ElasticSearch:v8.0.0
- Kibana:v8.0.0

# Setup

### Docker の起動、テストデータ投入(30 分くらいかかります)

```
$ make setup
```

# Usage

### 検索例

```
$ curl -X GET http://localhost:8090/v1/findshop?keyword=中華料理&area=東京&name=謝
```

### kibana の使用例

kibana は[こちら](http://localhost:9200)
