package controllers

import (
    "go-es-testcode/src/interfaces/elasticsearch"
    "go-es-testcode/src/usecase"
)

type ESController struct {
    Interactor usecase.ESInteractor
}

func NewESController() *ESController {
  return &ESController{
		Interactor: usecase.SearchInteractor{
			Search: &elasticsearch.SearchRepository{
				EsHost:         os.Getenv("ELASTIC_SEARCH"),
				EsIndexShop:    os.Getenv("ELASTIC_INDEX_SHOP"),
				EsCon:          &elasticsearch.SearchRepository{EsCon: ec},
			},
		},
  }
}

func (controller *PCController) Healthcheck(c Context) {
    PC,res := controller.Interactor.Healthcheck()
    _ = PC
    c.JSON(res.Code, res)
}
