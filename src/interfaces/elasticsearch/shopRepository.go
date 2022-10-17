package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"

	"log"
	"time"
)

func (r *SearchRepository) FindShopList(keyword *wrappers.StringValue) (*domain.ShopSearch, error) {

	var buf bytes.Buffer

	e, err := r.EsCon.ConnectElastic(r.EsHost)
	if err := json.NewEncoder(&buf).Encode(r.SearchConditionShop(keyword.Value)); err != nil {
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
		return nil, err
	}

	return &apiResult, nil
}

func (repo *SearchRepository) SearchConditionShop(keyword string) *domain.ShopSearchRequest {
	var req domain.ShopSearchRequest

	req.From = 0
	req.Size = 10

	slot, _ := domain.GetCurrentSlot(time.Now())

	req.Source = []interface{}{
		"shop_id",
		"name",
		"brand_name",
		"settings",
		fmt.Sprintf("delivery.waiting_time.%s", slot),
		"location",
		"distance",
		"metadata",
		"point_redeem_rate",
		"services",
		"rating",
		"average_budget",
		"delivery_fee",
		"delivery.immediately_acceptable_time",
		"delivery.reserve_acceptable_time",
		"delivery.fee",
		"payment_methods",
		"food_genres",
		"labels",
		"aggregations",
	}

	repo.applyKeyword(keyword, &req)
	repo.applyDisplayFlag(&req)

	return &req
}

// キーワード検索
func (repo *SearchRepository) applyKeyword(keyword string, req *domain.ShopSearchRequest) {
	if keyword != "" {
		req.Query.Bool.Must = append(req.Query.Bool.Must, domain.CombinedFields{
			CombinedFields: domain.CombinedFieldsValue{
				Query: keyword,
				Fields: []string{
					"name^1.5",
					"brand.name^2",
					"search_text.menu^1",
				},
			},
		})
	}
}

// 検索掲出フラグ
func (repo *SearchRepository) applyDisplayFlag(req *domain.ShopSearchRequest) {
	req.Query.Bool.Filter = append(req.Query.Bool.Filter, domain.Term{
		Term: map[string]domain.TermValue{
			"settings.tags": {
				Value: "present",
			},
		},
	})
}
