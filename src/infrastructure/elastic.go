package infrastructure

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
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

func (e *ElasticConnection) ConnectElastic(eshost string) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{eshost},
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   e.getEnvAsDuration("TIME_OUT") * time.Second,
				KeepAlive: e.getEnvAsDuration("KEEP_ALIVE") * time.Second,
			}).DialContext,
			MaxConnsPerHost:       e.getEnvAsInt("MAX_CONNS_PER_HOST"),
			ResponseHeaderTimeout: e.getEnvAsDuration("RESPONSE_HEADER_TIMEOUT") * time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating the client: %s", err)
	}

	return client, nil
}

func (e *ElasticConnection) Search(index string, body bytes.Buffer, es *elasticsearch.Client) (*esapi.Response, error) {
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithBody(&body),
		es.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("Error searching: %s", err)
	}

	return res, nil
}

func (e *ElasticConnection) getEnvAsDuration(key string) time.Duration {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Invalid env value for %s, using default", key)
		return 0
	}
	return time.Duration(value)
}

func (e *ElasticConnection) getEnvAsInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Invalid env value for %s, using default", key)
		return 0
	}
	return value
}
