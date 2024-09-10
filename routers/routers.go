package routers

import (
	"github.com/VeeRomanoff/gin_gorm_api/handlers"
	"github.com/gin-gonic/gin"
)

// Gin engine. Through gin.Engine the further actions will be done
func SetupRouter() *gin.Engine {
	router := gin.Default() // mux.Router alternative. Default gin router

	// В gin принято группировать ресурсы. Это позволяет нам не прописывать префиксы каждый раз
	apiV1Group := router.Group("/api/v1") // @RequestMapping
	{
		apiV1Group.GET("article", handlers.GetAllArticles)
		apiV1Group.POST("article", handlers.CreateArticle)
		apiV1Group.GET("article/:id", handlers.GetArticleById)
		apiV1Group.PUT("article/:id", handlers.UpdateArticleById)
		apiV1Group.DELETE("article/:id", handlers.DeleteArticleById)
	}

	//... maybe initializing some other groups
	return router
}
