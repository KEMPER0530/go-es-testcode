# 店舗検索

ElasticSearch

## Dependency

- Go:v1.17.5
- ElasticSearch:v8.0.0
- Kibana:v8.0.0

## Directory

```
.
├── Makefile
├── README.md
├── common                            # FOの共通処理
├── config                            # 環境変数の管理ファイル、Elastic, Go の設定ファイルなど共通設定ファイル
│   ├── config_development.yaml
│   ├── elastic
│   │   ├── Dockerfile
│   │   └── insert_data.sh
│   └── go
│       └── Dockerfile
│
├── docker-compose.yml
├── domain　　　　　　　　　　　　　　　　　# [Entities] Entitity、他
├── errors                            # 未実装
├── go.mod
├── go.sum
├── infrastructure                    # [Devices] gRPCサービスの起動、ElasticSearchとの接続
├── inferfaces                        # [Interface Adapters]
│   ├── elasticsearch                   # Presenter: ElasticSearchへ投げる
│   └── service.go                      # Controller protoで定義しているメソッドを記載
├── main.go
├── proto
└── usecase                           # [usecase] ビジネスロジック
```