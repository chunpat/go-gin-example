package article_service

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/chunpat/go-gin-example/models"
	"github.com/chunpat/go-gin-example/pkg/export"
	"github.com/chunpat/go-gin-example/pkg/gredis"
	"github.com/chunpat/go-gin-example/pkg/logging"
	"github.com/chunpat/go-gin-example/service/cache_service"
	"strconv"
	"time"
	"fmt"
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
	//var cacheArticle []*models.Article

	//cache := cache_service.Article{
	//	TagID:    a.TagID,
	//	State:    a.State,
	//	PageNum:  a.PageNum,
	//	PageSize: a.PageSize,
	//}
	//key := cache.GetArticlesKey()
	//exist, err := gredis.Exists(key)
	//if err != nil {
	//	//logging redis error
	//	logging.Error("redis error", err)
	//}
	//
	////use redis cache data
	//if exist == true {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info("redis error :", err)
	//	} else {
	//		json.Unmarshal(data, &cacheArticle)
	//		return cacheArticle, nil
	//	}
	//}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, articles, 3600)
	return articles, nil
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
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

//使用excelize流导出
func (a *Article) Export() (string, error) {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		println(err.Error())
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("A1", []interface{}{excelize.Cell{StyleID: styleID, Value: "文章编号"}}); err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("B1", []interface{}{excelize.Cell{StyleID: styleID, Value: "文章tags"}}); err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("C1", []interface{}{excelize.Cell{StyleID: styleID, Value: "文章名"}}); err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("D1", []interface{}{excelize.Cell{StyleID: styleID, Value: "描述"}}); err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("E1", []interface{}{excelize.Cell{StyleID: styleID, Value: "内容"}}); err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("F1", []interface{}{excelize.Cell{StyleID: styleID, Value: "创建时间"}}); err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("G1", []interface{}{excelize.Cell{StyleID: styleID, Value: "封面"}}); err != nil {
		println(err.Error())
	}

	articles, err := a.GetAll()

	for key, a := range articles {
		fmt.Printf("%+v\n", key) // {x:1 y:2}
		fmt.Printf("%+v\n", a) // {x:1 y:2}
		row := make([]interface{}, 7)
		values := []string{
			strconv.Itoa(a.ID),
			a.Tag.Name,
			a.Title,
			a.Desc,
			a.Content,
			strconv.Itoa(a.CreatedOn),
			a.CoverImageUrl,
		}

		for key, value := range values {
			row[key] = value
		}

		cell, _ := excelize.CoordinatesToCellName(1, key + 2)
		if err := streamWriter.SetRow(cell, row); err != nil {
			println(err.Error())
		}
	}

	if err := streamWriter.Flush(); err != nil {
		println(err.Error())
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "artcles-" + time + ".xlsx"

	//mkdir
	fullPath,err := export.GetPwdFullPath(filename)
	if err != nil {
		return "permission dir", err
	}

	if err := file.SaveAs(fullPath); err != nil {
		println(err.Error())
	}
	return filename, nil
}
