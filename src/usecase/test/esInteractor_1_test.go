package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-es-testcode/src/domain"
	"go-es-testcode/src/usecase"
	mock_repository "go-es-testcode/src/usecase/mock"
	"io/ioutil"
	"testing"
)

func Test_usecase_FindShop_MockingServerBehavior(t *testing.T) {
	// 検索ワードの設定
	keyword := "ラーメン"
	area := "東京"
	name := ""

	// 共通利用するstructを設定
	var i usecase.ESInteractor

	// gomockの利用設定
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// MockRepositryの作成
	mockESRepository := mock_repository.NewMockESRepository(ctrl)
	i.ES = mockESRepository

	t.Run("【正常系】FindShopメソッドの正常系テスト(Elasticsearchサーバーの動作をモックするパターン)", func(t *testing.T) {
		// ElasticSearchサーバーの動作を返却するレスポンスデータを利用する
		apiResult, err := loadTestData("../../../config/elasticsearch/test_data/test_data_1.json")
		if err != nil {
			t.Fatal(err)
		}

		// c, err := interactor.ES.FindShop(keyword, area, name)メソッドのMock化 正常系
		mockESRepository.EXPECT().FindShop(gomock.Any(), gomock.Any(), gomock.Any()).Return(apiResult, nil)

		// テスト対象のメソッド実行
		fs, status := i.FindShop(keyword, area, name)

		// テストの実施
		// メソッドのステータス、ElasticSearchの返却値の確認を実施
		assert.Equal(t, fs.Hits.Hits[0].Source.Id, int64(14018))
		assert.Equal(t, fs.Hits.Hits[0].Source.Name, "テストラーメン")
		assert.Equal(t, fs.Hits.Hits[1].Source.Id, int64(24137))
		assert.Equal(t, fs.Hits.Hits[1].Source.Name, "テストラーメン２")
		assert.Equal(t, fs.Hits.Hits[2].Source.Id, int64(18007))
		assert.Equal(t, fs.Hits.Hits[2].Source.Name, "テストラーメン３")
		assert.Equal(t, status.Code, domain.NewResultStatus(200).Code)
	})

	t.Run("【異常系】FindShopメソッドのテスト(Elasticsearchサーバーの動作をモックするパターン)", func(t *testing.T) {
		// c, err := interactor.ES.FindShop(keyword, area, name)メソッドのMock化 エラー系
		mockErr := fmt.Errorf("Error: %s", "errors.New")
		mockESRepository.EXPECT().FindShop(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, mockErr)

		// テスト対象のメソッド実行
		_, status := i.FindShop(keyword, area, name)

		// テストの実施
		// メソッドのステータスの確認を実施
		assert.Equal(t, status.Code, domain.NewResultStatus(500).Code)
	})
}

func loadTestData(path string) (*domain.ShopSearch, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var apiResult domain.ShopSearch
	if err := json.Unmarshal(bytes, &apiResult); err != nil {
		return nil, err
	}

	return &apiResult, nil
}
