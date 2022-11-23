package elasticsearch

import (
	v8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/golang/mock/gomock"
	"go-es-testcode/src/infrastructure"
	"go-es-testcode/src/interfaces/elasticsearch"
	mock_elasticsearch "go-es-testcode/src/interfaces/elasticsearch/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_FindShopList(t *testing.T) {

	// 共通利用するstructを設定
	var r elasticsearch.SearchRepository
	r.EsHost = "dummy-host"
	r.EsIndexShop = "dummy-shop"
	r.EsIndexMenu = "dummy-menu"
	r.EsIndexSuggest = "dummy-suggest"
	r.EsCon = &infrastructure.ElasticConnection{} // ←は後でmockに差し替える

	t.Run("1:FindShopメソッドのテスト(Elasticsearchサーバーの動作をモックするパターン)", func(t *testing.T) {

		// gomockの利用設定
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockElastic := mock_elasticsearch.NewMockElastic(ctrl)

		// mockで利用するメソッドの返却値を設定する
		// ConnectElasticメソッドをmock化
		es, _ := v8.NewClient(
			v8.Config{},
		)
		MockElastic.EXPECT().ConnectElastic(r.EsHost).Return(es, nil)

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
		MockElastic.EXPECT().Search(r.EsIndexShop, gomock.Any(), es).Return(&res, nil)

		// mock対象メソッドはレシーバーを設定しているのでmock用のレシーバーに差替え
		r.EsCon = MockElastic
		// テスト対象のメソッド実行
		keyword := "ラーメン"
		area := "東京都"
		name := ""
		fs, _ := r.FindShop(keyword,area,name)
		// テストの実施
		assert.Equal(t, fs.Hits.Hits[0].Source.Id, int64(14018))
		assert.Equal(t, fs.Hits.Hits[0].Source.Name, "テストラーメン")
		assert.Equal(t, fs.Hits.Hits[1].Source.Id, int64(24137))
		assert.Equal(t, fs.Hits.Hits[1].Source.Name, "テストラーメン２")
		assert.Equal(t, fs.Hits.Hits[2].Source.Id, int64(18007))
		assert.Equal(t, fs.Hits.Hits[2].Source.Name, "テストラーメン３")
	})
}
