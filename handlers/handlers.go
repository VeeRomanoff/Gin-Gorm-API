package handlers

import (
	"github.com/VeeRomanoff/gin_gorm_api/helpers"
	"github.com/VeeRomanoff/gin_gorm_api/models"
	"github.com/VeeRomanoff/gin_gorm_api/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// *gin.Context - w http.ResponseWriter, r *http.Request, maybe something else..

func GetAllArticles(ctx *gin.Context) {
	var articles []models.Article
	err := storage.GetAllArticles(articles)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusNotFound, articles)
		return
	}
	helpers.ResponseJSON(ctx, http.StatusOK, articles)
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var article models.Article
	err := storage.GetArticleById(&article, id)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusNotFound, article)
	}
	helpers.ResponseJSON(ctx, http.StatusOK, article)
}

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	ctx.BindJSON(&article) // method to parse json
	err := storage.CreateArticle(&article)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusBadRequest, article)
		return
	}
	helpers.ResponseJSON(ctx, http.StatusCreated, article)
}

func UpdateArticleById(ctx *gin.Context) {
	var article models.Article
	id := ctx.Params.ByName("id")
	err := storage.GetArticleById(&article, id)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusNotFound, article)
		return
	}
	ctx.BindJSON(&article)
	err = storage.UpdateArticleById(&article, id)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusNotFound, article)
	}
	helpers.ResponseJSON(ctx, http.StatusAccepted, article)
}

func DeleteArticleById(ctx *gin.Context) {
	var article models.Article
	id := ctx.Params.ByName("id")
	err := storage.DeleteArticleById(&article, id)
	if err != nil {
		helpers.ResponseJSON(ctx, http.StatusNotFound, article)
		return
	}
	helpers.ResponseJSON(ctx, http.StatusOK, article)
}
