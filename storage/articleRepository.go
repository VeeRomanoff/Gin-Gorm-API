package storage

import "github.com/VeeRomanoff/gin_gorm_api/models"

func GetAllArticles(a []models.Article) error {
	if err := DB.Find(a).Error; err != nil {
		return err
	}
	return nil
}

func CreateArticle(a *models.Article) error {
	if err := DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func GetArticleById(a *models.Article, id string) error {
	if err := DB.Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticleById(a *models.Article, id string) error {
	if err := DB.Where("id = ?", id).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func UpdateArticleById(a *models.Article, id string) error {
	if err := DB.Save(&a).Error; err != nil {
		return err
	}
	return nil
}
