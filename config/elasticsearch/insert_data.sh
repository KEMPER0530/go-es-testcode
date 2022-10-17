# 既存のインデックスの削除
curl -X DELETE "http://localhost:9200/test_shop"

# Index定義の追加
cd config/elasticsearch/index_settings && curl -H "Content-Type: application/json" -XPOST 'http://localhost:9200/test_shop/_doc?pretty' -d @shop.json

# データ挿入
elasticdump --input=test_shop.json --output=http://localhost:9200/test_shop

# データの確認
curl -X GET "http://localhost:9200/test_shop/_count"
