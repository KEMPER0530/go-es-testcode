package infrastructure

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-es-testcode/src/interfaces/controllers"
)

type Router struct {
	ginEngine *gin.Engine
	port      string
}

func NewRouter() *Router {
	router := &Router{
		ginEngine: gin.Default(),
		port:      ":" + os.Getenv("PORT"),
	}
	router.setRouting()
	return router
}

func (r *Router) setRouting() {
	r.setCORS()

	// コントローラーの設定
	esController := controllers.NewESController(&ElasticConnection{})
	// ElasticSearchにアクセスして接続確認を行う
	r.ginEngine.GET("/v1/findshop", esController.FindShop)
}

func Run(r *Router) {
	r.ginEngine.Run(r.port)
}

// Cross-Origin Resource Sharing (CORS) 設定
func (r *Router) setCORS() {
	r.ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Accept", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Cache-Control", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))
}
