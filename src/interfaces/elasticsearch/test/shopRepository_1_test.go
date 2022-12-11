package elasticsearch

import (
	"fmt"
	v8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go-es-testcode/src/infrastructure"
	"go-es-testcode/src/interfaces/elasticsearch"
	mock_elasticsearch "go-es-testcode/src/interfaces/elasticsearch/mock"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_interfaces_FindShop_MockingServerBehavior(t *testing.T) {

	// テスト対象のメソッド実行
	keyword := "ラーメン"
	area := "東京都"
	name := ""

	// 共通利用するstructを設定
	var r elasticsearch.SearchRepository
	var mockElastic *mock_elasticsearch.MockElastic
	var es *v8.Client

	// gomockの利用設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockElastic = mock_elasticsearch.NewMockElastic(ctrl)

	// ElasticSearchの接続先を設定
	r.EsHost = "dummy-host"
	r.EsIndexShop = "dummy-shop"
	r.EsCon = &infrastructure.ElasticConnection{} // ←は後でmockに差し替える

	es, _ = v8.NewClient(v8.Config{})

	t.Run("【正常系】interfaces:FindShopメソッドのテスト(Elasticsearchサーバーの動作をモックするパターン)", func(t *testing.T) {

		// mockで利用するメソッドの返却値を設定する
		// ConnectElasticメソッドをmock化
		mockElastic.EXPECT().ConnectElastic(r.EsHost).Return(es, nil)

		// テストデータ読み込み
		bytes, err := ioutil.ReadFile("../../../../config/elasticsearch/test_data/test_data_1.json")
		if err != nil {
			panic(err)
		}

		// mockで利用するメソッドの返却値を設定する
		var res esapi.Response
		res.StatusCode = 200
		m := string(bytes)
		res.Body = ioutil.NopCloser(strings.NewReader(m))
		// Searchメソッドをmock化
		mockElastic.EXPECT().Search(r.EsIndexShop, gomock.Any(), es).Return(&res, nil)

		// mock対象メソッドはレシーバーを設定しているのでmock用のレシーバーに差替え
		r.EsCon = mockElastic

		fs, err := r.FindShop(keyword, area, name)

		// テストの実施
		assert.Equal(t, fs.Hits.Hits[0].Source.Id, int64(14018))
		assert.Equal(t, fs.Hits.Hits[0].Source.Name, "テストラーメン")
		assert.Equal(t, fs.Hits.Hits[1].Source.Id, int64(24137))
		assert.Equal(t, fs.Hits.Hits[1].Source.Name, "テストラーメン２")
		assert.Equal(t, fs.Hits.Hits[2].Source.Id, int64(18007))
		assert.Equal(t, fs.Hits.Hits[2].Source.Name, "テストラーメン３")
		assert.Equal(t, err, nil)
	})

	t.Run("【異常系】interfaces:FindShopメソッドのテスト(Elasticsearchサーバーの動作をモックするパターン)", func(t *testing.T) {

		// mockで利用するメソッドの返却値を設定する
		// ConnectElasticメソッドをmock化
		mockElastic.EXPECT().ConnectElastic(r.EsHost).Return(es, nil)

		// mockで利用するメソッドの返却値を設定する
		var res esapi.Response
		res.StatusCode = 200

		// Searchメソッドをmock化
		mockErr := errors.New(fmt.Sprintf("Error: %s", "errors.New"))
		mockElastic.EXPECT().Search(r.EsIndexShop, gomock.Any(), es).Return(&res, mockErr)

		// mock対象メソッドはレシーバーを設定しているのでmock用のレシーバーに差替え
		r.EsCon = mockElastic

		_, err := r.FindShop(keyword, area, name)

		// テストの実施
		assert.Equal(t, err, mockErr)
	})
}
