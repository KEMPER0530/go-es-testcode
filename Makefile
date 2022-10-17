setup:
	docker-compose up -d &&\
	sleep 60 &&\
	sh ./config/elasticsearch/insert_data.sh

test:
	go test -v ./interfaces/...
	go test -v ./usecase/...

mockgen:
	mockgen -source=usecase/search_repository.go -destination usecase/mock/search_repository.go
	mockgen -source=interfaces/elasticsearch/elastic.go -destination interfaces/elasticsearch/mock/elastic.go
