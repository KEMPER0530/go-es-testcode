insertElasticSearch:
	docker compose up -d es01 es02 kibana &&\
	sleep 60 &&\
	sh ./config/elastic/insert_data.sh

test:
	go test -v ./interfaces/...
	go test -v ./usecase/...

mockgen:
	mockgen -source=usecase/search_repository.go -destination usecase/mock/search_repository.go
	mockgen -source=interfaces/elasticsearch/elastic.go -destination interfaces/elasticsearch/mock/elastic.go
