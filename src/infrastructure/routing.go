package infrastructure

import (
	// Gin
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"

	// コントローラー
	controllers "go-es-testcode/src/interfaces/controllers"
)

type Routing struct {
    Gin *gin.Engine
    Port string
}

func NewRouting() *Routing {
  r := &Routing{
      Gin: gin.Default(),
      Port: ":" + os.Getenv("PORT"),
  }
  r.setRouting()
	return r
}

func (r *Routing) setRouting() {

	// 本番設定の場合
	if os.Getenv("GO_ENV") == "production" {
		os.Setenv("GIN_MODE", "release")
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS設定
	r.Gin.Use(setCors())

	// コントローラーの設定
	ESController := controllers.NewESController()
	// ElasticSearchにアクセスして接続確認を行う
	r.Gin.GET("/v1/healthCheck", func (c *gin.Context) { ESController.HealthCheck(c,d) })

}

func Run(r *Routing) {
    r.Gin.Run(r.Port)
}

// Cross-Origin Resource Sharing (CORS) is a mechanism
// that uses additional HTTP headers to let a
// user agent gain permission to access selected resources from a server
// on a different origin /(domain) than the site currently in use.
// CORS for All origins, allowing:
// - PUT and PATCH methods
// - Origin header
// - Credentials share
// - Preflight requests cached for 1 hours
func setCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Accept", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Cache-Control", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	})
}
