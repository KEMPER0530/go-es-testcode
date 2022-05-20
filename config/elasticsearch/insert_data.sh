# 既存のインデックスの削除
curl -X DELETE "http://localhost:9200/test-shop"

# Index定義の追加
curl -v -H 'Content-Type: application/json' \
    -d "$(node ./config/elastic/index_settings/shop.js)" \
    -X PUT \
		"http://localhost:9200/test-shop"


# データ挿入
elasticdump --input=./config/elastic/test-shop.json --output=http://localhost:9200/test-shop

# データの確認
curl -X GET "http://localhost:9200/test-shop/_count"
