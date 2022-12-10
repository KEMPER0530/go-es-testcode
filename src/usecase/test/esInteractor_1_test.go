package usecase

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-es-testcode/src/domain"
	"go-es-testcode/src/usecase"
	mock_repository "go-es-testcode/src/usecase/mock"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_FindShop_MockingServerBehavior(t *testing.T) {

	// 共通利用するstructを設定
	var i usecase.ESInteractor

	// 検索ワードの設定
	keyword := "ラーメン"
	area := "東京"
	name := ""

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
		MockESRepository.EXPECT().FindShop(gomock.Any(), gomock.Any(), gomock.Any()).Return(&apiResult, nil)

		// mock対象メソッドはレシーバーを設定しているのでmock用のレシーバーに差替え
		i.ES = MockESRepository
		// テスト対象のメソッド実行
		fs, _ := i.FindShop(keyword, area, name)
		// テストの実施
		assert.Equal(t, fs.Hits.Hits[0].Source.Id, int64(14018))
		assert.Equal(t, fs.Hits.Hits[0].Source.Name, "テストラーメン")
		assert.Equal(t, fs.Hits.Hits[1].Source.Id, int64(24137))
		assert.Equal(t, fs.Hits.Hits[1].Source.Name, "テストラーメン２")
		assert.Equal(t, fs.Hits.Hits[2].Source.Id, int64(18007))
		assert.Equal(t, fs.Hits.Hits[2].Source.Name, "テストラーメン３")
	})
}
