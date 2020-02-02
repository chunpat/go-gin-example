package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	CoverImageUrl string `json:"cover_image_url"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? and deleted_on = ?", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Article{}).Where(maps).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetArticles(pageOffset int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageOffset).Limit(pageSize).Find(&articles).Error

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticle(id int, data interface{}) (bool, error) {
	err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddArticle(data map[string]interface{}) (bool, error) {
	err := db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteArticle(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(Article{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 软删除转硬删除Article
func CleanAllArticle() (bool, error) {
	// Unscoped 让model不走回调
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
