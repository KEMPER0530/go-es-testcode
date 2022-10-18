# 既存のインデックスの削除
curl -X DELETE "http://localhost:9200/test_shop"

# Index定義の追加
curl -v -H "Content-Type: application/json" \
    -X POST 'http://localhost:9200/test_shop/_doc?pretty' \
    -d @config/elasticsearch/index_settings/shop.json

# データ挿入
elasticdump --input=./config/elasticsearch/index_settings/test_shop.json --output=http://localhost:9200/test_shop

# データの確認
curl -X GET "http://localhost:9200/test_shop/_count"
