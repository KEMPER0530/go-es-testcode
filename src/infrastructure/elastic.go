package infrastructure

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ElasticConnection struct{}

// ElasticSearchへの接続
func (e *ElasticConnection) ConnectElastic(eshost string) (*elasticsearch.Client, error) {
	_MaxConnsPerHost, _ := strconv.Atoi(os.Getenv("MAX_CONNS_PER_HOST"))
	_ResponseHeaderTimeout, _ := strconv.Atoi(os.Getenv("RESPONSE_HEADER_TIMEOUT"))
	_Timeout, _ := strconv.Atoi(os.Getenv("TIME_OUT"))
	_KeepAlive, _ := strconv.Atoi(os.Getenv("KEEP_ALIVE"))
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			eshost,
		},
		Transport: &http.Transport{
			MaxConnsPerHost:       _MaxConnsPerHost,
			ResponseHeaderTimeout: time.Duration(_ResponseHeaderTimeout) * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(_Timeout) * time.Second,
				KeepAlive: time.Duration(_KeepAlive) * time.Second,
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
