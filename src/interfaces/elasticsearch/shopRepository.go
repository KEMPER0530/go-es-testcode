package elasticsearch

import (
	"bytes"
	"encoding/json"
	"go-es-testcode/src/domain"
	"log"
)

func (r *SearchRepository) FindShop(keyword string, area string, name string) (*domain.ShopSearch, error) {
	var buf bytes.Buffer
	e, err := r.EsCon.ConnectElastic(r.EsHost)
	if err := json.NewEncoder(&buf).Encode(r.SearchConditionShop(keyword, area, name)); err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := r.EsCon.Search(r.EsIndexShop, buf, e)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("failed to update with elastic search. Not ok. %s\n", res.Status())
	}
	defer res.Body.Close()
	var apiResult domain.ShopSearch
	if err := json.NewDecoder(res.Body).Decode(&apiResult); err != nil {
		log.Println(err)
		return nil, err
	}
	return &apiResult, nil
}

func (repo *SearchRepository) SearchConditionShop(keyword string, area string, name string) *domain.ShopSearchRequest {
	var req domain.ShopSearchRequest

	req.From = 0
	req.Size = 5

	req.Source = []interface{}{
		"id",
		"name",
		"property",
		"alphabet",
		"name_kana",
		"pref_id",
		"area_id",
		"station_id1",
		"station_time1",
		"station_distance1",
		"station_id2",
		"station_time2",
		"station_distance2",
		"station_id3",
		"station_time3",
		"station_distance3",
		"category_id1",
		"category_id2",
		"category_id3",
		"category_id4",
		"category_id5",
		"zip",
		"address",
		"north_latitude",
		"east_longitude",
		"description",
		"purpose",
		"open_morning",
		"open_lunch",
		"open_late",
		"photo_count",
		"special_count",
		"menu_count",
		"fan_count",
		"access_count",
		"created_on",
		"modified_on",
		"closed",
		"area_name",
		"pref_name",
		"pref",
		"location",
		"stas",
		"cates",
		"kuchikomi",
	}

	repo.applyKeyword(keyword, &req)
	repo.applyArea(area, &req)
	repo.applyShopname(name, &req)

	return &req
}

// キーワード検索
func (repo *SearchRepository) applyKeyword(keyword string, req *domain.ShopSearchRequest) {
	if keyword != "" {
		req.Query.Bool.Must = append(req.Query.Bool.Must, domain.CombinedFields{
			CombinedFields: domain.CombinedFieldsValue{
				Query: keyword,
				Fields: []string{
					"cates",
				},
			},
		})
	}
}

// エリア検索
func (repo *SearchRepository) applyArea(area string, req *domain.ShopSearchRequest) {
	if area != "" {
		req.Query.Bool.Must = append(req.Query.Bool.Must, domain.CombinedFields{
			CombinedFields: domain.CombinedFieldsValue{
				Query: area,
				Fields: []string{
					"area_name^3",
					"pref_name^2",
					"stas",
					"address",
				},
			},
		})
	}
}

// 名前検索
func (repo *SearchRepository) applyShopname(shopname string, req *domain.ShopSearchRequest) {
	if shopname != "" {
		req.Query.Bool.Must = append(req.Query.Bool.Must, domain.CombinedFields{
			CombinedFields: domain.CombinedFieldsValue{
				Query: shopname,
				Fields: []string{
					"name",
				},
			},
		})
	}
}
