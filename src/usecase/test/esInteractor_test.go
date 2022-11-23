package usecase

import (
	"github.com/golang/mock/gomock"
	"go-es-testcode/src/domain"
	"go-es-testcode/src/usecase"
	mock_repository "go-es-testcode/src/usecase/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
	"encoding/json"
)

func Test_FindShop(t *testing.T) {

	// 共通利用するstructを設定
	var i usecase.ESInteractor

	t.Run("1:FindShopメソッドのテスト(Elasticsearchサーバーの動作をモックするパターン)", func(t *testing.T) {

		// gomockの利用設定
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockESRepository := mock_repository.NewMockESRepository(ctrl)

		// ElasticSearchサーバーの動作を返却するテストデータ
		bytes, err := ioutil.ReadFile("../../../config/elasticsearch/test_data/test_data_1.json")
		if err != nil {
			panic(err)
		}
		m := string(bytes)
		b := ioutil.NopCloser(strings.NewReader(m))
		// レスポンスデータの作成
		var apiResult domain.ShopSearch
		if err := json.NewDecoder(b).Decode(&apiResult); err != nil {
			panic(err)
		}

		// mockで利用するメソッドの返却値を設定する
		// FindShopListメソッドをmock化
		MockESRepository.EXPECT().FindShop(gomock.Any(),gomock.Any(),gomock.Any()).Return(&apiResult, nil)

		// mock対象メソッドはレシーバーを設定しているのでmock用のレシーバーに差替え
		i.ES = MockESRepository
		// テスト対象のメソッド実行
		keyword := "ラーメン"
		area := "東京都"
		name := ""
		fs, _ := i.FindShop(keyword,area,name)
		// テストの実施
		assert.Equal(t, fs.Hits.Hits[0].Source.Id, int64(14018))
		assert.Equal(t, fs.Hits.Hits[0].Source.Name, "テストラーメン")
		assert.Equal(t, fs.Hits.Hits[1].Source.Id, int64(24137))
		assert.Equal(t, fs.Hits.Hits[1].Source.Name, "テストラーメン２")
		assert.Equal(t, fs.Hits.Hits[2].Source.Id, int64(18007))
		assert.Equal(t, fs.Hits.Hits[2].Source.Name, "テストラーメン３")
	})

	t.Run("2:FindShopメソッドのテスト(DockerコンテナーでElasticsearchの実際のインスタンスを実行)", func(t *testing.T) {
	})
}
