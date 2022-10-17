package elasticsearch

import (
	"bytes"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Elastic interface {
	ConnectElastic(string) (*elasticsearch.Client, error)
	Search(string, bytes.Buffer, *elasticsearch.Client) (*esapi.Response, error)
}
