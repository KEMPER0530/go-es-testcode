package elasticsearch

import (
	"bytes"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type SearchRepository struct {
	EsHost      string
	EsIndexShop string
	EsCon       Elastic
}

func (s *SearchRepository) ConnectElastic(eshost string) (*elasticsearch.Client, error) {
	return s.EsCon.ConnectElastic(eshost)
}

func (s *SearchRepository) Search(index string, body bytes.Buffer, es *elasticsearch.Client) (*esapi.Response, error) {
	return s.EsCon.Search(index, body, es)
}
