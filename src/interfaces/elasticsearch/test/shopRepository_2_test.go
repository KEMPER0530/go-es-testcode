package elasticsearch

import (
	"bytes"
	"context"
	"fmt"
	v8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	infra "go-es-testcode/src/infrastructure"
	"go-es-testcode/src/interfaces/elasticsearch"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_interfaces_FindShop_RunningServer(t *testing.T) {

	// 検索ワードの設定
	keyword := "ラーメン"
	area := "東京"
	name := ""

	// 共通利用するstructを設定
	var r elasticsearch.SearchRepository

	// 環境変数定義
	os.Setenv("ELASTIC_INDEX_SHOP", "test_shop")
	os.Setenv("MAX_CONNS_PER_HOST", "30")
	os.Setenv("RESPONSE_HEADER_TIMEOUT", "30")
	os.Setenv("TIME_OUT", "10")
	os.Setenv("KEEP_ALIVE", "10")

	// ElasticSearchの立ち上げ
	ctx := context.Background()
	elastic, baseUrl, err := initElastic(ctx)
	if err != nil {
		log.Error("Bulk insert failed.")
	}
	os.Setenv("ELASTIC_SEARCH", baseUrl)
	defer elastic.Terminate(ctx)

	// データ投入
	res, _ := fillElasticWithData(baseUrl)
	if res.StatusCode == 400 {
		log.Error("Bulk insert failed.")
	}

	t.Run("【正常系】FindShopメソッドのテスト(DockerコンテナーでElasticsearchの実際のインスタンスを実行)", func(t *testing.T) {

		// ElasticSearchの接続先を設定
		r.EsHost = baseUrl
		r.EsIndexShop = os.Getenv("ELASTIC_INDEX_SHOP")
		r.EsCon = &infra.ElasticConnection{}

		// time.Sleep(time.Second * 300)
		// テスト対象メソッドの呼び出し
		fs, err := r.FindShop(keyword, area, name)

		// テストの実施
		assert.NotEmpty(t, fs)
		assert.Equal(t, err, nil)
	})
}

// ElasticSearchのコンテナ作成 Port:9200でテスト用のElasticSearchコンテナを立ち上げ
func initElastic(ctx context.Context) (testcontainers.Container, string, error) {
	e, err := startEsContainer("9200", "9300")
	if err != nil {
		log.Error("Could not start ES container: " + err.Error())
		return nil, "", err
	}
	ip, err := e.Host(ctx)
	if err != nil {
		log.Error("Could not get host where the container is exposed: " + err.Error())
		return nil, "", err
	}
	port, err := e.MappedPort(ctx, "9200")
	if err != nil {
		log.Error("Could not retrive the mapped port: " + err.Error())
		return nil, "", err
	}
	baseUrl := fmt.Sprintf("http://%s:%s", ip, port.Port())

	// Clientの作成
	cfg := v8.Config{
		Addresses: []string{
			baseUrl,
		},
	}
	es, _ := v8.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, "", err
	}
	// mapping内容の読み込み
	bytes, err := ioutil.ReadFile("../../../../config/elasticsearch/index_settings/shop.json")
	if err != nil {
		log.Error("Could not read shop.json: " + err.Error())
		return nil, "", err
	}
	mapping := string(bytes)
	// indexの作成
	if err != createIndex(es, mapping) {
		log.Error(err.Error())
		return nil, "", err
	}
	return e, baseUrl, nil
}

func startEsContainer(restPort string, nodesPort string) (testcontainers.Container, error) {
	ctx := context.Background()

	rp := fmt.Sprintf("%s:%s/tcp", restPort, restPort)
	np := fmt.Sprintf("%s:%s/tcp", nodesPort, nodesPort)

	reqes5 := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../../../config/elasticsearch/",
			Dockerfile: "Dockerfile",
		},
		Name:         "es-mock",
		Env:          map[string]string{"discovery.type": "single-node"},
		ExposedPorts: []string{rp, np},
		WaitingFor:   wait.ForLog("started"),
	}
	elastic, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: reqes5,
		Started:          true,
	})

	return elastic, err
}

// createIndex indexを作成します
func createIndex(client *v8.Client, mapping string) error {
	req := esapi.IndicesCreateRequest{
		Index: os.Getenv("ELASTIC_INDEX_SHOP"),
		Body:  strings.NewReader(mapping),
	}

	// コンテナ起動後にスリープを実施。ESが起動していないため
	time.Sleep(time.Second * 30)
	// Perform the request with the client.
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	return nil
}

// データ投入 BulkInsertでデータを投入する
func fillElasticWithData(baseUrl string) (*http.Response, error) {

	b, err := ioutil.ReadFile("../../../../config/elasticsearch/test_data/test_data_2.json")
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", baseUrl+"/_bulk?pretty", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/x-ndjson")
	res, err := client.Do(req)
	if err != nil {
		log.Error("Could not perform a bulk operation")
	}
	defer res.Body.Close()
	log.Info("Bulk-insert:", res.StatusCode)

	return res, err
}
