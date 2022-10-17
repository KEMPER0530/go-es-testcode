package infrastructure

import (
	"bytes"
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/spf13/viper"
)

type ElasticConnection struct{}

// ElasticSearchへの接続
func (e *ElasticConnection) ConnectElastic(eshost string) (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			eshost,
		},
		Transport: &http.Transport{
			MaxConnsPerHost:       strconv.Atoi(os.Getenv("MAX_CONNS_PER_HOST")),
			ResponseHeaderTimeout: time.Duration(strconv.Atoi(os.Getenv("RESPONSE_HEADER_TIMEOUT"))) * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(strconv.Atoi(os.Getenv("TIME_OUT"))) * time.Second,
				KeepAlive: time.Duration(strconv.Atoi(os.Getenv("KEEP_ALIVE"))) * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
		},
	})
	if err != nil {
		log.Printf("Error creating the client: %s\n", err)
	}
	return es, err
}

func (e *ElasticConnection) Search(index string, body bytes.Buffer, es *elasticsearch.Client) (*esapi.Response, error) {
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithBody(&body),
		es.Search.WithPretty(),
	)
	return res, err
}
