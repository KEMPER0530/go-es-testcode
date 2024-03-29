setup:
	docker-compose up -d &&\
	sleep 60 &&\
	sh ./config/elasticsearch/insert_data.sh

test:
	go test -v ./src/interfaces/...
	go test -v ./src/usecase/...

mockgen:
	mockgen -source=src/usecase/esRepository.go -destination src/usecase/mock/esRepository.go
	mockgen -source=src/interfaces/elasticsearch/elastic.go -destination src/interfaces/elasticsearch/mock/elastic.go
