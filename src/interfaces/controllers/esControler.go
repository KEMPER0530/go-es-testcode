package controllers

import (
	"go-es-testcode/src/interfaces/elasticsearch"
	"go-es-testcode/src/usecase"
	"os"
	"github.com/gin-gonic/gin"
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
	ES, res := controller.Interactor.FindShop(c.Query("keyword"), c.Query("area"), c.Query("name"))
	c.JSON(res.Code, ES)
}
