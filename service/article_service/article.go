package article_service

import (
	"encoding/json"

	"github.com/FromChinaBoy/go-gin-example/models"
	"github.com/FromChinaBoy/go-gin-example/pkg/gredis"
	"github.com/FromChinaBoy/go-gin-example/pkg/logging"
	"github.com/FromChinaBoy/go-gin-example/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	exist, err := gredis.Exists(key)
	if err != nil {
		//logging redis error
		logging.Error("redis error", err)
	}

	//use redis cache data
	if exist == true {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var cacheArticle []*models.Article

	cache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	exist, err := gredis.Exists(key)
	if err != nil {
		//logging redis error
		logging.Error("redis error", err)
	}

	//use redis cache data
	if exist == true {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info("redis error :", err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 3600)
	return articles, nil
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps)
}

func (a *Article) Add() (bool, error) {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"state":           a.State,
		"cover_image_url": a.CoverImageUrl,
	}
	return models.AddArticle(article)
}

func (a *Article) Edit() (bool, error) {
	data := make(map[string]interface{})
	if a.TagID > 0 {
		data["tag_id"] = a.TagID
	}
	if a.Title != "" {
		data["title"] = a.Title
	}
	if a.Desc != "" {
		data["desc"] = a.Desc
	}
	if a.Content != "" {
		data["content"] = a.Content
	}
	if a.CoverImageUrl != "" {
		data["cover_image_url"] = a.CoverImageUrl
	}

	data["modified_by"] = a.ModifiedBy
	return models.EditArticle(a.ID, data)
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Del() (bool, error) {
	return models.DeleteArticle(a.ID)
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
