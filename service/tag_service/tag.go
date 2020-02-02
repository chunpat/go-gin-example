package tag_service

import "github.com/FromChinaBoy/go-gin-example/models"

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) Add() (bool, error) {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() (bool, error) {
	var tag = make(map[string]interface{})
	if t.Name != "" {
		tag["name"] = t.Name
	}
	if t.State != -1 {
		tag["state"] = t.State
	}
	if t.ModifiedBy != "" {
		tag["modified_by"] = t.ModifiedBy
	}

	return models.EditTag(t.ID, tag)
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) Del() (bool, error) {
	return models.DeleteTag(t.ID)
}

func (t *Tag) getMaps() map[string]interface{} {
	var maps map[string]interface{}
	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.Name != "" {
		maps["state"] = t.State
	}

	if t.ModifiedBy != "" {
		maps["modified_by"] = t.ModifiedBy
	}

	return maps
}
