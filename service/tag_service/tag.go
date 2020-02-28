package tag_service

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"strconv"
	"time"

	"github.com/chunpat/go-gin-example/models"
	"github.com/chunpat/go-gin-example/pkg/export"
	"github.com/chunpat/go-gin-example/pkg/gredis"
	"github.com/chunpat/go-gin-example/pkg/logging"
	"github.com/chunpat/go-gin-example/service/cache_service"
	"github.com/tealeg/xlsx"
)

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
	return models.ExistTagByName(t.Name, t.ID)
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

func (t *Tag) GetAll() ([]*models.Tag, error) {
	var cacheTags []*models.Tag
	cacheService := cache_service.Tag{
		Name:     t.Name,
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}

	key := cacheService.GetTagsKey()
	exist, err := gredis.Exists(key)
	print("key:", exist, err)
	if err != nil {
		print("error")
		//logging redis error
		logging.Error("redis error", err)
	}

	//use redis cache data
	if exist == true {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info("redis error :", err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) Del() (bool, error) {
	return models.DeleteTag(t.ID)
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State != -1 {
		maps["state"] = t.State
	}

	if t.ModifiedBy != "" {
		maps["modified_by"] = t.ModifiedBy
	}
	maps["deleted_on"] = 0
	return maps
}

//使用标准库导出
func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx"

	//mkdir
	fullPath,err := export.GetPwdFullPath(filename)
	if err != nil {
		return "permission dir", err
	}

	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}

	return filename, nil
}

//使用excelize 导入
func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows,err := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}
