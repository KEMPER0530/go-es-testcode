package controllers

import (
	"github.com/gin-gonic/gin"
	"go-es-testcode/src/interfaces/elasticsearch"
	"go-es-testcode/src/usecase"
	"os"
)

type ESController struct {
	Interactor usecase.ESInteractor
}

func NewESController(ec elasticsearch.Elastic) *ESController {
	return &ESController{
		Interactor: usecase.ESInteractor{
			ES: &elasticsearch.SearchRepository{
				EsHost:      os.Getenv("ELASTIC_SEARCH"),
				EsIndexShop: os.Getenv("ELASTIC_INDEX_SHOP"),
				EsCon:       &elasticsearch.SearchRepository{EsCon: ec},
			},
		},
	}
}

func (controller *ESController) FindShop(c *gin.Context) {
	ss, res := controller.Interactor.FindShop(c.Query("keyword"), c.Query("area"), c.Query("name"))
	c.JSON(res.Code, ss)
}
