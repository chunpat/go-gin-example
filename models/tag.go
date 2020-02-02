package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageOffset int, pageSize int, maps interface{}) ([]Tag, error) {
	var tags []Tag
	err := db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func AddTag(name string, state int, createdBy string) (bool, error) {
	err := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error

	if err != nil {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? and deleted_on = ?", id, 0).First(&tag).Error
	if err != nil {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func DeleteTag(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func EditTag(id int, data interface{}) (bool, error) {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 软删除转硬删除tags
func CleanAllTag() (bool, error) {
	// Unscoped 让model不走回调
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
