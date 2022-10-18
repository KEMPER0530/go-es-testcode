package usecase

import (
	"go-es-testcode/src/domain"
)

type ESRepository interface {
	FindShop(string, string, string) (*domain.ShopSearch, error)
}
