# Testing Elasticsearch App

Go 言語(gin)+Elasticsearch on docker.

# Features

ElasticSearch を利用したレストラン検索を実施する練習用のリポジトリです。
[ライブドアグルメデータ](https://github.com/livedoor/datasets)を編集し、
ElasticSearch へ登録してクエリの検証を実施します。

DataContet については[こちら](https://github.com/KEMPER0530/elastic-demo)を参照

## Dependency

- Go:v1.19
- ElasticSearch:v8.0.0
- Kibana:v8.0.0

## テストコードの実施

[Qiita の記事](https://qiita.com/y-aka/items/cc01a4a70b2b2ac5beb7)で記載しているテストコード実施のコマンドです。
下記のテストを実施します。

- [Elasticsearch サーバーの動作をモックする方法](https://github.com/KEMPER0530/go-es-testcode/blob/main/src/interfaces/elasticsearch/test/shopRepository_1_test.go)
- [Docker コンテナーで Elasticsearch のインスタンスを実行する方法](https://github.com/KEMPER0530/go-es-testcode/blob/main/src/interfaces/elasticsearch/test/shopRepository_2_test.go)

```
$ make test
```

## Setup（以下は ElasticSearch を Docker で構築し Go 経由で結果を返却するアプリになります）

### ElasticSearch のテストデータ解凍（[Zstandard](https://qiita.com/oioi_tec/items/e66ec93824f694a473c9)で圧縮しています）

```
$ zstd -d /config/elasticsearch/index_settings/test_shop.json.zst
```

### Docker の起動、テストデータ投入（30 分くらいかかります）

```
$ make setup
```

## Usage

### 検索例（jq コマンドを利用することで見やすくなります）

```
$ curl -X GET "http://localhost:8090/v1/findshop?keyword=中華料理&area=東京&name=謝" | jq
```

### kibana の使用

kibana は[こちら](http://localhost:5601)
