package usecase

import (
	"go-es-testcode/src/domain"
)

type ESInteractor struct {
	ES ESRepository
}

func (interactor *ESInteractor) FindShop(keyword string, area string, name string) (ss *domain.ShopSearch, resultStatus ResultStatus) {
	c, err := interactor.ES.FindShop(keyword, area, name)
	if err != nil {
		return &domain.ShopSearch{}, NewResultStatus(500)
	}

	return c, NewResultStatus(200)
}
